package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/ibraheemacamara/list-token-transfers/types"
	log "github.com/sirupsen/logrus"
)

type HttpInterface interface {
	Do(req *http.Request) (*http.Response, error)
}

type RpcProviderClient struct {
	Client HttpInterface
	Url    string
}

func Init(u string, timeout time.Duration) *RpcProviderClient {
	return &RpcProviderClient{
		Client: &http.Client{
			Timeout: timeout,
		},
		Url: u,
	}

}

func (h *RpcProviderClient) SendRequest(ctx context.Context, body types.JSONRPCRequest) types.JSONRPCResponse {

	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Errorf("failed to marshal body request: body %v, error: %v", body, err)
		return types.JSONRPCResponse{

			Error: &types.JSONRPCError{
				Code:    types.InternalErrorCode,
				Message: types.InternalErrorMessage,
			},
		}
	}

	log.Debugf("Sending request to rpc provider. url: %v, request id: %v , jsonrpc: %v , method: %v , params: %v", h.Url, body.ID, body.Jsonrpc, body.Method, body.Params)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, h.Url, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Error(err.Error())
		return types.JSONRPCResponse{

			Error: &types.JSONRPCError{
				Code:    types.InternalErrorCode,
				Message: types.InternalErrorMessage,
			},
		}
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := h.Client.Do(req)

	if err != nil {
		log.Errorf("failed to send request to rpc provider. response: %v, err: %v", resp, err)
		return types.JSONRPCResponse{
			Jsonrpc: "2.0",
			ID:      body.ID,
			Error: &types.JSONRPCError{
				Code:    types.InternalErrorCode,
				Message: types.InternalErrorMessage,
			},
		}
	}
	if resp.StatusCode != http.StatusOK {
		log.Errorf("failed to send request to rpc provider. responseCode: %v, err: %v", resp.StatusCode, err)
		return types.JSONRPCResponse{
			Jsonrpc: "2.0",
			ID:      body.ID,
			Error: &types.JSONRPCError{
				Code:    types.InternalErrorCode,
				Message: types.InternalErrorMessage,
			},
		}
	}
	defer resp.Body.Close()

	var result types.JSONRPCResponse
	resultBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.JSONRPCResponse{

			Error: &types.JSONRPCError{
				Code:    types.InternalErrorCode,
				Message: types.InternalErrorMessage,
			},
		}
	}

	json.Unmarshal(resultBytes, &result)
	return result
}
