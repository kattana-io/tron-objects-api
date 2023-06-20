package sunswap

import (
	"github.com/kattana-io/tron-objects-api/pkg/api"
	"github.com/kattana-io/tron-objects-api/pkg/url"
	"math/big"
)

func CanSell(tokenIn, tokenOut *api.Address, amountIn *big.Int) (bool, error) {
	impl := api.NewAPI("", nil, url.NewTrongridURLProvider())
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
		return resB.Cmp(amountIn) == +1, nil
	}
	return resA.Cmp(amountIn) == +1, nil
}
