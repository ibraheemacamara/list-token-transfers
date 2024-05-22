package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ibraheemacamara/list-token-transfers/config"
	"github.com/ibraheemacamara/list-token-transfers/types"
	"github.com/ibraheemacamara/list-token-transfers/utils"
	"github.com/stretchr/testify/assert"
)

func MockConfig() config.Config {
	return config.Config{
		AllowedMethods: []string{"evm_transferHistory"},
	}
}

func MockContext(body interface{}) (*gin.Context, *gin.Engine, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, engine := gin.CreateTestContext(w)

	c.Request = &http.Request{
		Header: make(http.Header),
	}

	MockJsonPost(c, body)

	return c, engine, w
}

func MockJsonPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

func TestController_ShouldGetResponse(t *testing.T) {

	//Init a server ()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer server.Close()

	rpcProviderInterface := new(MockRpcProviderClient)
	rpcProviderInterface.RpcProviderClient = utils.Init("myprovider.com", time.Second*4)
	rpcProviderInterface.RpcProviderClient.Client = server.Client()
	controller := NewController(rpcProviderInterface)
	config := MockConfig()
	controller.SetConfig(config)

	//given a request body
	params := types.TransfersListParams{
		ContractAddress: "",
		Addresses:       []string{"0x073b55a2bd0c7b5df3a3453f8cfbd510787d13706"},
		LastBlocks:      2,
	}
	req := types.JSONRPCRequest{
		Jsonrpc: "2.0",
		ID:      1,
		Method:  "evm_transferHistory",
		Params:  []interface{}{params},
	}

	//Mock context
	ctx, engine, recoder := MockContext(req)

	var err error
	ctx.Request, err = http.NewRequest(http.MethodPost, "/", ctx.Request.Body)
	assert.Nil(t, err)

	engine.POST("/", controller.MainHandler())

	engine.ServeHTTP(recoder, ctx.Request)

	assert.EqualValues(t, 200, recoder.Code)
	fmt.Println(recoder.Body)
	var response []types.TransferResponse
	err = json.Unmarshal(recoder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, 1, len(response))
}
