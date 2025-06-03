package jsonrpc

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

// GetToken0 retrieves the token0() address from a pair contract.
func (c *JSONRPCClient) GetToken0(ctx context.Context, pair common.Address) (string, error) {
	methodSig := c.encodeFunctionSelector(token0Method + "()")
	params := []any{
		CallParams{
			To:   pair.String(),
			Data: methodSig,
		},
		"latest",
	}

	resultHex, err := call[any, string](ctx, c.url, ethCallMethod, params)
	if err != nil {
		return "", fmt.Errorf("error calling 'eth_call' for %s(): %w", token0Method, err)
	}
	if resultHex == nil {
		return "", fmt.Errorf("empty response from 'eth_call' for %s()", token0Method)
	}
	return c.unpackAddress(*resultHex, token0Method)
}

// GetToken1 retrieves the token1() address from a pair contract.
func (c *JSONRPCClient) GetToken1(ctx context.Context, pair common.Address) (string, error) {
	methodSig := c.encodeFunctionSelector(token1Method + "()")
	params := []any{
		CallParams{
			To:   pair.String(),
			Data: methodSig,
		},
		"latest",
	}

	resultHex, err := call[any, string](ctx, c.url, ethCallMethod, params)
	if err != nil {
		return "", fmt.Errorf("error calling 'eth_call' for %s(): %w", token1Method, err)
	}
	if resultHex == nil {
		return "", fmt.Errorf("empty response from 'eth_call' for %s()", token1Method)
	}
	return c.unpackAddress(*resultHex, token1Method)
}

// GetPairToken retrieves the tokenAddress() address from a pair contract.
func (c *JSONRPCClient) GetPairToken(ctx context.Context, contract common.Address) (string, error) {
	methodSig := c.encodeFunctionSelector(tokenAddressMethod + "()") // "0x9d76ea58" // function selector for tokenAddress()
	fmt.Println("methodSig", methodSig)
	params := []any{
		CallParams{
			To:   contract.String(),
			Data: methodSig,
		},
		"latest",
	}

	resultHex, err := call[any, string](ctx, c.url, ethCallMethod, params)
	if err != nil {
		return "", fmt.Errorf("error calling 'eth_call' for %s(): %w", tokenAddressMethod, err)
	}
	if resultHex == nil {
		return "", fmt.Errorf("empty response from 'eth_call' for %s()", tokenAddressMethod)
	}
	return c.unpackAddress(*resultHex, tokenAddressMethod)
}

// GetReserves retrieves the getReserves() reserves from a pair contract.
//
//nolint:gocritic
func (c *JSONRPCClient) GetReserves(
	ctx context.Context,
	pair common.Address,
) (*big.Int, *big.Int, uint32, error) {
	methodData, err := getReservesABI.Pack(getReservesMethod)
	if err != nil {
		return nil, nil, 0, err
	}
	params := []any{
		CallParams{
			To:   pair.String(),
			Data: hexutil.Encode(methodData),
		},
		"latest",
	}

	resultHex, err := call[any, string](ctx, c.url, ethCallMethod, params)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("error calling 'eth_call' for %s(): %w", getReservesMethod, err)
	}
	if resultHex == nil {
		return nil, nil, 0, fmt.Errorf("empty response from 'eth_call' for %s()", getReservesMethod)
	}
	if len(*resultHex) < 2 {
		return nil, nil, 0, fmt.Errorf("empty or invalid response from 'eth_call' for %s()", getReservesMethod)
	}

	// Decode the hex result into bytes
	resultData, err := hexutil.Decode(*resultHex)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("failed to decode result hex: %w", err)
	}

	// Unpack using ABI
	out, err := getReservesABI.Unpack(getReservesMethod, resultData)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("failed to unpack result: %w", err)
	}
	if len(out) != 3 {
		return nil, nil, 0, fmt.Errorf("unexpected unpacked output length")
	}

	return out[0].(*big.Int), out[1].(*big.Int), out[2].(uint32), nil
}

// GetPair from factory contract.
func (c *JSONRPCClient) GetPair(ctx context.Context, factory, tokenA, tokenB common.Address) (string, error) {
	methodData, err := sunswapV2FactoryABI.Pack(getPairMethod, tokenA, tokenB)
	if err != nil {
		return "", err
	}
	params := []any{
		CallParams{
			To:   factory.String(),
			Data: hexutil.Encode(methodData),
		},
		"latest",
	}

	resultHex, err := call[any, string](ctx, c.url, ethCallMethod, params)
	if err != nil {
		return "", fmt.Errorf("error calling 'eth_call' for %s(): %w", getPairMethod, err)
	}
	if resultHex == nil {
		return "", fmt.Errorf("empty response from 'eth_call' for %s()", getPairMethod)
	}
	return c.unpackAddress(*resultHex, getPairMethod)
}
