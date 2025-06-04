package jsonrpc

import (
	"context"
	"github.com/kattana-io/tron-objects-api/pkg/api/jsonrpc"
	"github.com/kattana-io/tron-objects-api/pkg/connector/justmoney"
	"github.com/kattana-io/tron-objects-api/pkg/types"
)

type Pair struct {
	api  *jsonrpc.API
	addr types.Address
}

func New(impl *jsonrpc.API, addr types.Address) *Pair {
	return &Pair{
		api:  impl,
		addr: addr,
	}
}

func (s *Pair) Token0(ctx context.Context) (*types.Address, error) {
	res, err := s.api.GetToken0(ctx, s.addr.ToGoEthAddr())
	if err != nil {
		return types.NewEmptyAddress(), err
	}
	if res == "" {
		return types.NewEmptyAddress(), justmoney.ErrNotJustMoneyPair
	}
	return types.NewFromHex(res), nil
}

func (s *Pair) Token1(ctx context.Context) (*types.Address, error) {
	res, err := s.api.GetToken1(ctx, s.addr.ToGoEthAddr())
	if err != nil {
		return types.NewEmptyAddress(), err
	}
	if res == "" {
		return types.NewEmptyAddress(), justmoney.ErrNotJustMoneyPair
	}
	return types.NewFromHex(res), nil
}
