package jsonrpc

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// GetTRC20Decimals retrieves the decimals() value from a TRC20 contract.
// Uses JSON-RPC eth_call method.
// Docs: https://developers.tron.network/reference/eth_call
func (c *JSONRPCClient) GetTRC20Decimals(ctx context.Context, token common.Address) (uint8, error) {
	methodData, err := trc20ABI.Pack(decimalsMethod)
	if err != nil {
		return 0, err
	}
	params := []any{
		CallParams{
			To:   token.String(),
			Data: hexutil.Encode(methodData),
		},
		"latest",
	}

	// Call eth_call expecting a hex-encoded result (e.g. "0x0000000000000000000000000000000000000000000000000000000000000006")
	resultHex, err := call[any, string](ctx, c.url, ethCallMethod, params)
	if err != nil {
		return 0, fmt.Errorf("error calling 'eth_call' for %s(): %w", decimalsMethod, err)
	}
	if resultHex == nil || len(*resultHex) < 3 {
		return 0, fmt.Errorf("invalid or empty response from 'eth_call' for %s()", decimalsMethod)
	}

	decimals, err := c.unpackInteger(*resultHex, decimalsMethod)
	if err != nil {
		return 0, err
	}
	//nolint:gosec
	return uint8(decimals.Uint64()), nil
}

// GetTRC20Name retrieves the name() value from a TRC20 contract.
func (c *JSONRPCClient) GetTRC20Name(ctx context.Context, token common.Address) (string, error) { //nolint:dupl
	methodData, err := trc20ABI.Pack(nameMethod)
	if err != nil {
		return "", err
	}
	params := []any{
		CallParams{
			To:   token.String(),
			Data: hexutil.Encode(methodData),
		},
		"latest",
	}

	resultHex, err := call[any, string](ctx, c.url, ethCallMethod, params)
	if err != nil {
		return "", fmt.Errorf("error calling 'eth_call' for %s(): %w", nameMethod, err)
	}
	if resultHex == nil {
		return "", fmt.Errorf("invalid response from 'eth_call' for %s(): %+v", nameMethod, resultHex)
	}

	return c.unpackString(*resultHex, nameMethod)
}

// GetTRC20Symbol retrieves the symbol() value from a TRC20 contract.
func (c *JSONRPCClient) GetTRC20Symbol(ctx context.Context, token common.Address) (string, error) { //nolint:dupl
	methodData, err := trc20ABI.Pack(symbolMethod)
	if err != nil {
		return "", err
	}
	params := []any{
		CallParams{
			To:   token.String(),
			Data: hexutil.Encode(methodData),
		},
		"latest",
	}

	resultHex, err := call[any, string](ctx, c.url, ethCallMethod, params)
	if err != nil {
		return "", fmt.Errorf("error calling 'eth_call' for %s(): %w", symbolMethod, err)
	}
	if resultHex == nil {
		return "", fmt.Errorf("invalid response from 'eth_call' for %s(): %+v", symbolMethod, resultHex)
	}

	return c.unpackString(*resultHex, symbolMethod)
}
