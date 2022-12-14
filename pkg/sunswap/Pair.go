package sunswap

import (
	"errors"
	"github.com/kattana-io/tron-objects-api/pkg/api"
)

type Pair struct {
	api     *api.Api
	address api.Address
}

func New(api *api.Api, address api.Address) *Pair {
	return &Pair{
		api:     api,
		address: address,
	}
}

func (s *Pair) GetTokenAddress() (*api.Address, error) {
	res, err := s.api.GetPairToken(s.address.ToHex())
	if err != nil {
		return api.EmptyAddress(), err
	}
	if res == "" {
		return api.EmptyAddress(), errors.New("returned nil address")
	}
	return api.FromHex(res), nil
}
