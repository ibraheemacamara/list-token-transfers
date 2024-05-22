package api

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ibraheemacamara/list-token-transfers/config"
	"github.com/ibraheemacamara/list-token-transfers/types"
	"github.com/ibraheemacamara/list-token-transfers/utils"
)

func getLogs(ctx context.Context, client RpcProviderClientInterface, lastestBlock int64, reqParam types.TransfersListParams) types.JSONRPCResponse {
	cfg := config.GetConfig()
	topic_0 := cfg.Erc20TransferTopic()

	fromBlock := fmt.Sprintf("0x%x", lastestBlock-reqParam.LastBlocks)
	params := []interface{}{
		types.EthGetLogsParams{
			Address:   reqParam.ContractAddress,
			FromBlock: fromBlock,
			Topics:    []string{topic_0},
		},
	}
	req := types.JSONRPCRequest{
		ID:      1,
		Jsonrpc: "2.0",
		Method:  "eth_getLogs",
		Params:  params,
	}

	jsonResponse := client.SendRequest(ctx, req)

	return jsonResponse
}

func latesBlock(ctx context.Context, client RpcProviderClientInterface) (string, *types.JSONRPCError) {
	req := types.JSONRPCRequest{
		ID:      1,
		Jsonrpc: "2.0",
		Method:  "eth_blockNumber",
	}

	jsonResponse := client.SendRequest(ctx, req)

	if jsonResponse.Error != nil {
		return "", jsonResponse.Error
	}

	return fmt.Sprintf("%v", jsonResponse.Result), nil
}

// index topic 0 => trensfer event
// index topic 1 => from
// index topic 2 => to
func filterLogByAddressList(logs []types.Log, addresses []string) []types.Log {
	fmt.Println(addresses)
	if len(addresses) < 1 {
		return logs
	}
	filtredLogs := []types.Log{}

	for i := 0; i < len(logs); i++ {
		from := fmt.Sprintf("0x%v", utils.TrimLeftZeroes(logs[i].Topics[1][2:]))
		to := fmt.Sprintf("0x%v", utils.TrimLeftZeroes(logs[i].Topics[2][2:]))
		if utils.Contains(addresses, from) || utils.Contains(addresses, to) {
			filtredLogs = append(filtredLogs, logs[i])
		}
	}

	return filtredLogs
}

func buildTransferResponse(logs []types.Log) []types.TransferResponse {
	txResponse := []types.TransferResponse{}

	for i := 0; i < len(logs); i++ {
		value, _ := strconv.ParseInt(utils.TrimLeftZeroes(logs[i].Data)[2:], 16, 32)
		txResponse = append(txResponse, types.TransferResponse{
			From:            fmt.Sprintf("0x%v", utils.TrimLeftZeroes(logs[i].Topics[1][2:])),
			To:              fmt.Sprintf("0x%v", utils.TrimLeftZeroes(logs[i].Topics[2][2:])),
			BlockHash:       logs[i].BlockHash,
			BlockNumber:     logs[i].BlockNumber,
			TransactionHash: logs[i].TransactionHash,
			Value:           float64(value),
		})
	}

	return txResponse
}
