package justmoney

import "github.com/kattana-io/tron-objects-api/pkg/api"

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

func (s *Pair) Token0() (*api.Address, error) {
	res, err := s.api.GetToken0(s.address.ToHex())
	if err != nil {
		return api.FromHex("0x0"), err
	}
	return api.FromHex(res), nil
}

func (s *Pair) Token1() (*api.Address, error) {
	res, err := s.api.GetToken1(s.address.ToHex())
	if err != nil {
		return api.FromHex("0x0"), err
	}
	return api.FromHex(res), nil
}
