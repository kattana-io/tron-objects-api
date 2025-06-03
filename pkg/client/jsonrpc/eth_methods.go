package jsonrpc

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/kattana-io/tron-objects-api/pkg/models"
)

// GetBlockByNum See docs https://developers.tron.network/reference/eth_getblockbynumber
func (c *JSONRPCClient) GetBlockByNum(ctx context.Context, num int64, fullTxs bool) (*models.Block, error) {
	blockNumHex := fmt.Sprintf("0x%x", num)
	params := []any{blockNumHex, fullTxs}

	resp, err := call[any, models.Block](ctx, c.url, ethGetBlockByNumberMethod, params)
	if err != nil {
		return nil, fmt.Errorf("error calling %s: %w", ethGetBlockByNumberMethod, err)
	}
	return resp, nil
}

// GetTransactionByHash See docs https://developers.tron.network/reference/eth_gettransactionbyhash
func (c *JSONRPCClient) GetTransactionByHash(ctx context.Context, hash common.Hash) (*models.Transaction, error) {
	params := []any{hash.String()}

	resp, err := call[any, models.Transaction](ctx, c.url, ethGetTransactionByHashMethod, params)
	if err != nil {
		return nil, fmt.Errorf("error calling %s: %w", ethGetTransactionByHashMethod, err)
	}
	return resp, nil
}

// GetTransactionByBlockNumAndIndex See docs https://developers.tron.network/reference/eth_gettransactionbyblocknumberandindex
func (c *JSONRPCClient) GetTransactionByBlockNumAndIndex(
	ctx context.Context,
	num int64,
	index int64,
) (*models.Transaction, error) {
	blockNumHex := fmt.Sprintf("0x%x", num)
	indexHex := fmt.Sprintf("0x%x", index)
	params := []any{blockNumHex, indexHex}

	resp, err := call[any, models.Transaction](ctx, c.url, ethGetTransactionByBlockNumberAndIndexMethod, params)
	if err != nil {
		return nil, fmt.Errorf("error calling %s: %w", ethGetTransactionByBlockNumberAndIndexMethod, err)
	}
	return resp, nil
}

// GetCode See docs https://developers.tron.network/reference/eth_getcode
func (c *JSONRPCClient) GetCode(ctx context.Context, contract common.Address, tag string) (string, error) {
	params := []any{contract.String(), tag}

	runtimeCode, err := call[any, string](ctx, c.url, ethGetCodeMethod, params)
	if err != nil {
		return "", fmt.Errorf("error calling %s: %w", ethGetCodeMethod, err)
	}
	if runtimeCode == nil {
		return "", fmt.Errorf("empty response from %s", ethGetCodeMethod)
	}
	return *runtimeCode, nil
}
