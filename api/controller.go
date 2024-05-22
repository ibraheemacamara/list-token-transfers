package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ibraheemacamara/list-token-transfers/config"
	"github.com/ibraheemacamara/list-token-transfers/types"
	"github.com/ibraheemacamara/list-token-transfers/utils"
	log "github.com/sirupsen/logrus"
)

type RpcProviderClientInterface interface {
	SendRequest(ctx context.Context, body types.JSONRPCRequest) types.JSONRPCResponse
}

type Controller struct {
	rpcProviderClient RpcProviderClientInterface
}

func NewController(rpcProviderClient RpcProviderClientInterface) *Controller {
	return &Controller{
		rpcProviderClient: rpcProviderClient,
	}
}

// Request
// curl -X POST --data '{"jsonrpc":"2.0","method":"chiliz_transfersList","params":[{"addresses":["0x407d73d8a49eeb85d32cf465507dd71d507100c1"], "fromBlock":"latest"}],"id":1}
func (c *Controller) MainHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Debugf("Received request %v", ctx.Request)

		//Decode body request
		var req types.JSONRPCRequest
		rawBody, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			log.Errorf("Could not read request body: %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}

		if err := json.Unmarshal(rawBody, &req); err != nil {
			log.Errorf("JSON unmarshal error: %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
			return
		}

		cfg := config.GetConfig()
		allowedMathod := cfg.AllowedMethods()
		if !utils.Contains(allowedMathod, req.Method) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "method not found"})
			return
		}

		var transfersListParams types.TransfersListParams
		err = utils.ConvertInterfaceToStruct(req.Params[0], &transfersListParams)
		if err != nil {
			log.Errorf("Invalid params: %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
			return
		}

		//validate fromBlock
		latestBlockNumberString, jsonRPCError := latesBlock(ctx, c.rpcProviderClient)
		if jsonRPCError != nil {
			log.Errorf("error getting the latest block %v", jsonRPCError.Message)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": types.InternalErrorMessage})
			return
		}
		latestBlockNumber, err := strconv.ParseInt(latestBlockNumberString[2:], 16, 64)
		if err != nil {
			log.Errorf("latestBlock is not a valid hexadecimal number. %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": types.InternalErrorMessage})
			return
		}

		jsonResponse := getLogs(ctx, c.rpcProviderClient, latestBlockNumber, transfersListParams)

		var logs []types.Log
		err = utils.ConvertInterfaceToStruct(jsonResponse.Result, &logs)
		if err != nil {
			log.Errorf("failed to convert logs result to struct. %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": types.InternalErrorMessage})
			return
		}

		filtredLogs := filterLogByAddressList(logs, transfersListParams.Addresses)

		transferList := buildTransferResponse(filtredLogs)

		ctx.JSON(http.StatusOK, transferList)
	}
}
