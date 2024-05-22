package types

const (
	InternalErrorMessage = "internal error"
	InternalErrorCode    = -32604
	ServerErrorCode      = -32000
	Pending              = "pending"
	Latest               = "latest"
	Safe                 = "safe"
	Finalized            = "finalized"
	Earliest             = "earliest"
)

// JSONRPCRequest defines the structure of an incoming JSON-RPC request.
type JSONRPCRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      interface{}   `json:"id"`
}

// JSONRPCResponse defines the structure of a JSON-RPC response.
type JSONRPCResponse struct {
	Jsonrpc string        `json:"jsonrpc"`
	ID      interface{}   `json:"id"`
	Result  interface{}   `json:"result,omitempty"`
	Error   *JSONRPCError `json:"error,omitempty"`
}

// JSONRPCError defines the structure of an error in a JSON-RPC response.
type JSONRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// HTTPError defines the structure of an error in a HTTP response.
type HTTPError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type TransfersListParams struct {
	ContractAddress string   `json:"contractAddress"`
	Addresses       []string `json:"addresses"`
	LastBlocks      int64    `json:"lastBlocks"`
}

type TransferResponse struct {
	From            string  `json:"from"`
	To              string  `json:"to"`
	Value           float64 `json:"value"`
	BlockHash       string  `json:"blockHash"`
	BlockNumber     string  `json:"blockNumber"`
	TransactionHash string  `json:"transactionHash"`
}

type EthGetLogsParams struct {
	Address   string   `json:"address"`
	FromBlock string   `json:"fromBlock"`
	Topics    []string `json:"topics"`
}

type Log struct {
	Address          string   `json:"address"`
	BlockHash        string   `json:"blockHash"`
	BlockNumber      string   `json:"blockNumber"`
	Data             string   `json:"data"`
	LogIndex         string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
	Topics           []string `json:"topics"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
}
