package rest

import (
	"errors"
	"github.com/kattana-io/tron-objects-api/pkg/api/rest"
	"github.com/kattana-io/tron-objects-api/pkg/types"
)

type Pair struct {
	api     *rest.API
	address types.Address
}

func New(impl *rest.API, address types.Address) *Pair {
	return &Pair{
		api:     impl,
		address: address,
	}
}

func (s *Pair) Token0() (*types.Address, error) {
	res, err := s.api.GetToken0(s.address.ToHex())
	if err != nil {
		return types.NewEmptyAddress(), err
	}
	if res == "" {
		return types.NewEmptyAddress(), errors.New(" not a justmoney pair")
	}
	return types.NewFromHex(res), nil
}

func (s *Pair) Token1() (*types.Address, error) {
	res, err := s.api.GetToken1(s.address.ToHex())
	if err != nil {
		return types.NewEmptyAddress(), err
	}
	if res == "" {
		return types.NewEmptyAddress(), errors.New(" not a justmoney pair")
	}
	return types.NewFromHex(res), nil
}
