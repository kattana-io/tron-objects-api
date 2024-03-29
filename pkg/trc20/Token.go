package trc20

import (
	"github.com/kattana-io/tron-objects-api/pkg/api"
)

type Token struct {
	api     *api.API
	address *api.Address
}

func New(imp *api.API, address *api.Address) *Token {
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
