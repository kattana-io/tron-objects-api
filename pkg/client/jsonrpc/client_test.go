package jsonrpc

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/kattana-io/tron-objects-api/pkg/models"
	"github.com/kattana-io/tron-objects-api/pkg/types"
	testassert "github.com/stretchr/testify/assert"
)

const (
	rpcURL = "https://api.trongrid.io/jsonrpc"

	usdtAddrBase58 = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t" // USDT contract (on Tron)
	jmAddrBase58   = "TVHH59uHVpHzLDMFFpUgCx2dNAQqCzPhcR" // USDT contract (on Tron)
	sunAddrBase58  = "TSSMHYeV2uE9qYH95DqyoCuNCzEL1NvU3S" // USDT contract (on Tron)

	jmUsdtJustMoneyPairAddrBase58 = "THTWV7R3U7XQsHWQt8YHgsqirvY9QttB7u" // JM/USDT (on Tron)

	sunswapV1PairAddrBase58      = "TDNbPAZh1cWvdJDnJ5M56eCwbxFm7b1x8V"
	sunswapV1TestTokenAddrBase58 = "TUqSwiouiFb2M7vBKfLdtn9gLzjSBSnomA" //nolint:gosec

	// sunswapV2FactoryAddrBase58 = "TKWJdrQkqHisa1X8HUdHEfREvTzw4pMAaY"
)

func Test_GetCode(t *testing.T) {
	assert := testassert.New(t)

	ctx := context.Background()
	cli := NewJSONRPCClient(rpcURL)

	usdtHexAddr := types.NewFromBase58(usdtAddrBase58).ToGoEthAddr()

	tag := "latest"
	runtimeCode, err := cli.GetCode(ctx, usdtHexAddr, tag)
	assert.NoError(err, "should not error on getCode")
	assert.NotEmpty(runtimeCode, "runtime code should not be empty")

	t.Logf("runtimeCode: %+v", runtimeCode)
}

func Test_GetBlockByNum(t *testing.T) {
	assert := testassert.New(t)

	ctx := context.Background()
	cli := NewJSONRPCClient(rpcURL)

	blockNum := int64(72735774)
	blockNumHex := fmt.Sprintf("0x%x", blockNum)

	block, err := cli.GetBlockByNum(ctx, blockNum, true)
	assert.NoError(err, "should not error on valid block number")
	assert.NotNil(block, "block should not be nil")
	assert.Equal(blockNumHex, block.Number, "block number should match expected hex")
	assert.NotEmpty(block.Hash, "block hash should not be empty")
	assert.NotEmpty(block.Timestamp, "timestamp should not be empty")

	t.Logf("block: %+v", block)
}

func Test_GetTransactionByHash(t *testing.T) {
	assert := testassert.New(t)

	ctx := context.Background()
	cli := NewJSONRPCClient(rpcURL)

	hash := common.HexToHash("0x7f97e9cba680a6280c79e0c5d8cf4380955a41c2ffb3fb29489f1015be67ab09")

	tx, err := cli.GetTransactionByHash(ctx, hash)
	assert.NoError(err, "should not error on valid transaction hash")
	assert.NotNil(tx, "transaction should not be nil")
	assert.Equal(hash.String(), tx.Hash, "transaction hash should match input")
	assert.NotEmpty(tx.BlockNumber, "block number should not be empty")

	t.Logf("tx: %+v", tx)
}

func Test_GetTransactionByBlockNumAndIndex(t *testing.T) {
	assert := testassert.New(t)

	ctx := context.Background()
	cli := NewJSONRPCClient(rpcURL)

	blockNum := int64(72735774)
	index := int64(0)

	tx, err := cli.GetTransactionByBlockNumAndIndex(ctx, blockNum, index)
	assert.NoError(err, "should not error on valid block number and index")
	assert.NotNil(tx, "transaction should not be nil")
	assert.NotEmpty(tx.Hash, "transaction hash should not be empty")
	assert.NotEmpty(tx.BlockNumber, "block number should not be empty")

	t.Logf("tx: %+v", tx)
}

func Test_GetTransactionReceipt(t *testing.T) {
	assert := testassert.New(t)

	ctx := context.Background()
	cli := NewJSONRPCClient(rpcURL)

	hash := common.HexToHash("0x5b8ef93cded9ffe7f250e5c5bbbee6a935ad25d80415a6fcd847bb6f6e3b4076")

	tx, err := cli.GetTransactionReceipt(ctx, hash)
	assert.NoError(err, "GetTransactionReceipt should not error")
	assert.NotNil(tx, "transaction receipt should not be nil")
	assert.NotEmpty(tx.TransactionHash, "transaction hash should not be empty")
	assert.NotEmpty(tx.BlockNumber, "block number should not be empty")

	t.Logf("transaction receipt: %+v", tx)
}

func Test_GetLogs(t *testing.T) {
	assert := testassert.New(t)

	ctx := context.Background()
	cli := NewJSONRPCClient(rpcURL)

	hash := common.HexToHash("0x000000000456c76cc8bb8cdf2d7f6f8a49ee5c58cc58c511f7976211e7b08ff3")

	logs, err := cli.GetLogs(ctx, &models.GetLogsRequest{BlockHash: hash})
	assert.NoError(err, "GetLogs should not error")
	assert.NotNil(logs, "logs should not be nil")

	t.Logf("logs receipt: %+v", logs)
}

func Test_GetTRC20Decimals(t *testing.T) {
	assert := testassert.New(t)

	ctx := context.Background()
	cli := NewJSONRPCClient(rpcURL)

	usdtHexAddr := types.NewFromBase58(usdtAddrBase58).ToGoEthAddr()

	decimals, err := cli.GetTRC20Decimals(ctx, usdtHexAddr)
	assert.NoError(err, "should not error on GetTRC20Decimals")
	assert.NotEmpty(decimals, "decimals should not be empty")

	t.Log("decimals", decimals)
}

func Test_GetTRC20Name(t *testing.T) {
	assert := testassert.New(t)

	ctx := context.Background()
	cli := NewJSONRPCClient(rpcURL)

	usdtHexAddr := types.NewFromBase58(usdtAddrBase58).ToGoEthAddr()

	name, err := cli.GetTRC20Name(ctx, usdtHexAddr)
	assert.NoError(err, "should not error on GetTRC20Name")
	assert.NotEmpty(name, "name should not be empty")

	t.Log("name", name)
}

func Test_GetTRC20Symbol(t *testing.T) {
	assert := testassert.New(t)

	ctx := context.Background()
	cli := NewJSONRPCClient(rpcURL)

	usdtHexAddr := types.NewFromBase58(usdtAddrBase58).ToGoEthAddr()

	symbol, err := cli.GetTRC20Symbol(ctx, usdtHexAddr)
	assert.NoError(err, "should not error on GetTRC20Symbol")
	assert.NotEmpty(symbol, "symbol should not be empty")

	t.Log("symbol", symbol)
}

func Test_GetToken0(t *testing.T) {
	assert := testassert.New(t)

	ctx := context.Background()
	cli := NewJSONRPCClient(rpcURL)

	jmUsdtPairHexAddr := types.NewFromBase58(jmUsdtJustMoneyPairAddrBase58).ToGoEthAddr()

	token0, err := cli.GetToken0(ctx, jmUsdtPairHexAddr)
	assert.NoError(err, "should not error on GetToken0")
	assert.NotEmpty(token0, "token0 should not be empty")

	usdtHexAddr := types.NewFromBase58(usdtAddrBase58).ToGoEthHex()
	assert.Equal(usdtHexAddr, token0, "token0 should match expected address")

	t.Log("token0", token0)
}

func Test_GetToken1(t *testing.T) {
	assert := testassert.New(t)

	ctx := context.Background()
	cli := NewJSONRPCClient(rpcURL)

	jmUsdtPairHexAddr := types.NewFromBase58(jmUsdtJustMoneyPairAddrBase58).ToGoEthAddr()

	token1, err := cli.GetToken1(ctx, jmUsdtPairHexAddr)
	assert.NoError(err, "should not error on GetToken1")
	assert.NotEmpty(token1, "token1 should not be empty")

	jmHexAddr := types.NewFromBase58(jmAddrBase58).ToGoEthHex()
	assert.Equal(jmHexAddr, token1, "token1 should match expected address")

	t.Log("token1", token1)
}

func Test_GetPairToken(t *testing.T) {
	assert := testassert.New(t)

	ctx := context.Background()
	cli := NewJSONRPCClient(rpcURL)

	sunswapV1PairHexAddr := types.NewFromBase58(sunswapV1PairAddrBase58).ToGoEthAddr()

	pairToken, err := cli.GetPairToken(ctx, sunswapV1PairHexAddr)
	assert.NoError(err, "should not error on GetPairToken")
	assert.NotEmpty(pairToken, "pairToken should not be empty")

	expectedTokenHex := types.NewFromBase58(sunswapV1TestTokenAddrBase58).ToGoEthHex()
	assert.Equal(expectedTokenHex, pairToken, "pairToken should match expected address")

	t.Log("pairToken", pairToken)
}

func Test_GetPair(t *testing.T) {
	assert := testassert.New(t)

	ctx := context.Background()
	cli := NewJSONRPCClient(rpcURL)

	factoryHex := types.NewFromBase58(sunswapV2FactoryAddrBase58).ToGoEthAddr()
	tokenAHex := types.NewFromBase58(usdtAddrBase58).ToGoEthAddr()
	tokenBHex := types.NewFromBase58(sunAddrBase58).ToGoEthAddr()

	pairHex, err := cli.GetPair(ctx, factoryHex, tokenAHex, tokenBHex)
	assert.NoError(err, "should not error on GetPair")
	assert.NotEmpty(pairHex, "pair address should not be empty")

	// expectedPairHex: = types.NewFromBase58(expectedPairAddr).ToGoEthAddr()
	// assert.Equal(expectedPairHex, pairHex, "pair address should match expected")

	t.Log("pair", pairHex)
}

func Test_GetReserves(t *testing.T) {
	assert := testassert.New(t)

	ctx := context.Background()
	cli := NewJSONRPCClient(rpcURL)

	jmUsdtPairHexAddr := types.NewFromBase58(jmUsdtJustMoneyPairAddrBase58).ToGoEthAddr()

	reserve0, reserve1, blockTimestampLast, err := cli.GetReserves(ctx, jmUsdtPairHexAddr)
	assert.NoError(err, "should not error on GetReserves")
	assert.NotEmpty(reserve0, "reserve0 should not be empty")
	assert.NotEmpty(reserve1, "reserve1 should not be empty")
	assert.NotEmpty(blockTimestampLast, "blockTimestampLast should not be empty")

	t.Log("reserve0: ", reserve0, "reserve1: ", reserve1, "blockTimestampLast:", blockTimestampLast)
}
