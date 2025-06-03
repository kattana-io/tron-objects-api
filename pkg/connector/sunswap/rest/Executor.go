package rest

import (
	"github.com/kattana-io/tron-objects-api/pkg/api/rest"
	rest2 "github.com/kattana-io/tron-objects-api/pkg/client/rest"
	"github.com/kattana-io/tron-objects-api/pkg/types"
	"math/big"
)

func CanSell(tokenIn, tokenOut *types.Address, amountIn *big.Int) (bool, error) {
	impl := rest.NewAPI("", nil, rest2.NewTrongridURLProvider())
	factory := NewFactory(impl)
	pairAddress, err := factory.GetPair(tokenIn, tokenOut)
	if err != nil {
		return false, err
	}
	pair := NewPair(impl, *pairAddress)
	token0, err := pair.GetToken0Address()
	if err != nil {
		return false, err
	}
	flipped := token0.ToBase58() != tokenIn.ToBase58()
	resA, resB, err := pair.GetReserves()
	if err != nil {
		return false, err
	}
	if flipped {
		return resB.Cmp(amountIn) > 0, nil
	}
	return resA.Cmp(amountIn) > 0, nil
}
