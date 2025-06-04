package jsonrpc

import (
	"context"
	"github.com/kattana-io/tron-objects-api/pkg/api/jsonrpc"
	"github.com/kattana-io/tron-objects-api/pkg/types"
)

// Token is implementation of trc20 token.
type Token struct {
	api  *jsonrpc.API
	addr *types.Address
}

func New(impl *jsonrpc.API, token *types.Address) *Token {
	return &Token{
		api:  impl,
		addr: token,
	}
}

const maxRetry = 5

// TryToGetDecimals is retry to get decimals of trc20 token.
func (t *Token) TryToGetDecimals(ctx context.Context, try int64) (int32, bool) {
	if try > maxRetry {
		return 0, false
	}

	decimals, err := t.api.GetTRC20Decimal(ctx, t.addr.ToGoEthAddr())
	if err != nil {
		try += 1
		return t.TryToGetDecimals(ctx, try)
	} else {
		return int32(decimals), true
	}
}

// GetDecimals returns decimals of trc20 token.
func (t *Token) GetDecimals(ctx context.Context) (int32, error) {
	decimals, err := t.api.GetTRC20Decimal(ctx, t.addr.ToGoEthAddr())
	if err != nil {
		return 0, err
	}
	return int32(decimals), nil
}

// GetSymbol returns symbol of trc20 token.
func (t *Token) GetSymbol(ctx context.Context) (string, error) {
	return t.api.GetTRC20Symbol(ctx, t.addr.ToGoEthAddr())
}

// GetName returns name of trc20 token.
func (t *Token) GetName(ctx context.Context) (string, error) {
	return t.api.GetTRC20Name(ctx, t.addr.ToGoEthAddr())
}
