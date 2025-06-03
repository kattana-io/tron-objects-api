package rest

import (
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

func (t *Token) TryToGetDecimals(try int64) (int32, bool) {
	if try > maxRetry {
		return 0, false
	}

	decimals, err := t.api.GetTokenDecimals(t.address.ToHex())

	if err != nil {
		try += 1
		return t.TryToGetDecimals(try)
	} else {
		return decimals, true
	}
}

func (t *Token) GetDecimals() (int32, error) {
	return t.api.GetTokenDecimals(t.address.ToHex())
}

func (t *Token) GetSymbol() (string, error) {
	return t.api.GetTokenSymbol(t.address.ToHex())
}

func (t *Token) GetName() (string, error) {
	return t.api.GetTokenSymbol(t.address.ToHex())
}
