package tests

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	testassert "github.com/stretchr/testify/assert"
)

// Need to set url to your JSON RPC node
const url = "https://api.trongrid.io/jsonrpc"

func TestGetBlockNumber(t *testing.T) {
	assert := testassert.New(t)

	ctx := context.Background()

	rpcClient, err := ethclient.Dial(url)
	assert.NoError(err, "should connect to RPC without error")
	defer rpcClient.Close()

	blockNumber, err := rpcClient.BlockNumber(ctx)
	assert.NoError(err, "should get block number without error")
	assert.Greater(blockNumber, uint64(0), "block number should be greater than 0")

	t.Logf("Last block number: %d", blockNumber)
}

func TestGetBlockByNumber(t *testing.T) {
	assert := testassert.New(t)

	ctx := context.Background()

	rpcClient, err := ethclient.Dial(url)
	assert.NoError(err, "should connect to RPC without error")
	defer rpcClient.Close()

	address := common.HexToAddress("0x41a614f803b6fd780986a42c78ec9c7f77e6ded13c") // USDT contract in hex (note the `0x41` Tron prefix)
	t.Logf("address: %s", address.String())

	code, err := rpcClient.CodeAt(ctx, address, nil)
	assert.NoError(err, "should get code without error")
	assert.NotNil(code, "code should not be nil")

	t.Logf("code: %+v", code)
}

// Don't work
// func TestGetBlockByNumber(t *testing.T) {
//	assert := testassert.New(t)
//
//	rpcClient, err := ethclient.Dial(url)
//	assert.NoError(err, "should connect to RPC without error")
//	defer rpcClient.Close()
//
//	//blockNumber := big.NewInt(16348182)
//	hash := common.HexToHash("0x2f4efd102173a000eb2e1d4bb6e5474e48da7863ced0f925da8657c9dfca2746")
//	t.Logf("blockHash: %s", hash.String())
//
//	block, _, err := rpcClient.TransactionByHash(context.Background(), hash)
//	assert.NoError(err, "should get block by number without error")
//	assert.NotNil(block, "block should not be nil")
//
//	t.Logf("block: %+v", block)
// }
