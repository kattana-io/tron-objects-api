package jsonrpc

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kattana-io/tron-objects-api/pkg/client/jsonrpc"
	"github.com/kattana-io/tron-objects-api/pkg/types"
	testassert "github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

const (
	rpcURL = "https://api.trongrid.io/jsonrpc"

	// usdtAddress = "0xa614f803b6fd780986a42c78ec9c7f77e6ded13c" // USDT contract in hex
	usdtAddrBase58 = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t" // USDT contract (on Tron)
)

func TestGetContractInfo(t *testing.T) {
	assert := testassert.New(t)

	ctx := context.Background()
	logger := zap.NewNop()

	ethcli, err := ethclient.DialContext(ctx, rpcURL)
	assert.NoError(err)
	rpccli := jsonrpc.NewJSONRPCClient(rpcURL)

	api := NewAPI(rpcURL, logger, rpccli, ethcli)

	usdtHexAddr := types.NewFromBase58(usdtAddrBase58).ToGoEthAddr()
	decimals, err := api.GetTRC20Decimal(ctx, usdtHexAddr)
	assert.NoError(err, "GetTRC20Decimal should not return error")

	t.Log("decimals", decimals)
}
