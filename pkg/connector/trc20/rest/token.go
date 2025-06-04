package rest

import (
	"context"
	"github.com/kattana-io/tron-objects-api/pkg/api/rest"
	"github.com/kattana-io/tron-objects-api/pkg/types"
)

type Token struct {
	api     *rest.API
	address *types.Address
}

func New(imp *rest.API, address *types.Address) *Token {
	return &Token{
		api:     imp,
		address: address,
	}
}

const maxRetry = 5

func (t *Token) TryToGetDecimals(ctx context.Context, try int64) (int32, bool) {
	if try > maxRetry {
		return 0, false
	}

	decimals, err := t.api.GetTokenDecimals(t.address.ToHex())
	if err != nil {
		try += 1
		return t.TryToGetDecimals(ctx, try)
	} else {
		return decimals, true
	}
}

func (t *Token) GetDecimals(_ context.Context) (int32, error) {
	return t.api.GetTokenDecimals(t.address.ToHex())
}

func (t *Token) GetSymbol(_ context.Context) (string, error) {
	return t.api.GetTokenSymbol(t.address.ToHex())
}

func (t *Token) GetName(_ context.Context) (string, error) {
	return t.api.GetTokenName(t.address.ToHex())
}
