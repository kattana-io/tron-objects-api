package trc20

import (
	"github.com/kattana-io/tron-objects-api/pkg/api"
)

type Token struct {
	api     *api.Api
	address *api.Address
}

func New(api *api.Api, address *api.Address) *Token {
	return &Token{
		api:     api,
		address: address,
	}
}

func (t *Token) TryToGetDecimals(try int64) (int32, bool) {
	if try > 5 {
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
