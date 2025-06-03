package jsonrpc

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kattana-io/tron-objects-api/pkg/client/jsonrpc"

	"go.uber.org/zap"
)

type API struct {
	url       string
	log       *zap.Logger
	rpcclient *jsonrpc.JSONRPCClient
	ethclient *ethclient.Client
}

func NewAPI(
	url string,
	log *zap.Logger,
	rpccli *jsonrpc.JSONRPCClient,
	ethcli *ethclient.Client,
) *API {
	return &API{
		url:       url,
		log:       log,
		rpcclient: rpccli,
		ethclient: ethcli,
	}
}
