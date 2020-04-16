package rangomorg

import (
	"github.com/ybbus/jsonrpc"
)

const ENDPOINT = "https://api.random.org/json-rpc/2/invoke"

type Rangom struct {
	ApiKey    string
	RPCClient jsonrpc.RPCClient
}

func New(apiKey string) Rangom {
	return Rangom{ApiKey: apiKey, RPCClient: jsonrpc.NewClient(ENDPOINT)}
}



