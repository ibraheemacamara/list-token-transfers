package api

import (
	"context"
	"testing"

	"github.com/ibraheemacamara/list-token-transfers/types"
	"github.com/ibraheemacamara/list-token-transfers/utils"
	"github.com/stretchr/testify/assert"
)

type MockRpcProviderClient struct {
	*utils.RpcProviderClient
}

var expectedLastBlock = "0x12ff6ec"
var expectedLogs = []types.Log{
	{
		Address: "contractAddress1",
		Topics:  []string{"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef", "0x00000000000000000000000073b55a2bd0c7b5df3a3453f8cfbd510787d13706", "0x0000000000000000000000006564ffa0e5f57628ea21d1591e9320e7906e77c5"},
		Data:    "0x0000000000000000000000000000000000000000000000000000000000000008",
	},
	{
		Address: "contractAddress2",
		Topics:  []string{"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef", "0x00000000000000000000000073b55a2bd0c7b5df3a3453f8cfbd510787d13500", "0x0000000000000000000000006564ffa0e5f57628ea21d1591e9320e7906e88c5"},
		Data:    "0x0000000000000000000000000000000000000000000000000000000000000008",
	},
}

func (m *MockRpcProviderClient) SendRequest(ctx context.Context, body types.JSONRPCRequest) types.JSONRPCResponse {
	if body.Method == "eth_blockNumber" {
		return types.JSONRPCResponse{
			Result: expectedLastBlock,
		}
	} else {

		return types.JSONRPCResponse{
			Result: expectedLogs,
		}
	}
}

func TestFilterLogByAddressList(t *testing.T) {
	expectedLog := types.Log{
		Address: "contractAddress1",
		Topics:  []string{"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef", "0x00000000000000000000000073b55a2bd0c7b5df3a3453f8cfbd510787d13706", "0x0000000000000000000000006564ffa0e5f57628ea21d1591e9320e7906e77c5"},
	}
	logs := []types.Log{
		expectedLog,
		{
			Address: "contractAddress2",
			Topics:  []string{"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef", "0x00000000000000000000000073b55a2bd0c7b5df3a3453f8cfbd510787d13500", "0x0000000000000000000000006564ffa0e5f57628ea21d1591e9320e7906e88c5"},
		},
	}

	addresses := []string{"0x073b55a2bd0c7b5df3a3453f8cfbd510787d13706"}

	filtred := filterLogByAddressList(logs, addresses)

	assert.EqualValues(t, 1, len(filtred))
	assert.Equal(t, expectedLog, filtred[0])
}

func TestGetLastBlock(t *testing.T) {
	rpcProviderInterface := new(MockRpcProviderClient)

	response, err := latesBlock(context.Background(), rpcProviderInterface)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLastBlock, response)
}
func TestGetLogs(t *testing.T) {
	rpcProviderInterface := new(MockRpcProviderClient)

	logsResponse := getLogs(context.Background(), rpcProviderInterface, 1, types.TransfersListParams{})
	assert.Nil(t, logsResponse.Error)
	assert.NotNil(t, logsResponse.Result)

	var logs []types.Log
	err := utils.ConvertInterfaceToStruct(logsResponse.Result, &logs)
	assert.NoError(t, err)

	assert.EqualValues(t, len(expectedLogs), len(logs))
	assert.Equal(t, expectedLogs, logs)
}
