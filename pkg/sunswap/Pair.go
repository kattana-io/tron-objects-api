package sunswap

import (
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/kattana-io/tron-objects-api/pkg/api"
	decoder "github.com/mingjingc/abi-decoder"
	"math/big"
)

type Pair struct {
	api     *api.API
	address api.Address
}

const getReservesAbi = `[{"constant":true,"inputs":[],"name":"getReserves","outputs":[{"internalType":"uint112","name":"_reserve0","type":"uint112"},{"internalType":"uint112","name":"_reserve1","type":"uint112"},{"internalType":"uint32","name":"_blockTimestampLast","type":"uint32"}],"payable":false,"stateMutability":"view","type":"function"}]`

func New(impl *api.API, address api.Address) *Pair {
	return &Pair{
		api:     impl,
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

func (s *Pair) GetReserves() (*big.Int, *big.Int, error) {
	res, err := s.api.GetPairReserves(s.address.ToHex())
	if err != nil {
		return nil, nil, err
	}
	dec := decoder.NewABIDecoder()
	dec.SetABI(getReservesAbi)
	method := dec.ABI().Methods["getReserves"]
	bts, err := hexutil.Decode("0x" + res)
	if err != nil {
		return nil, nil, err
	}
	data, err := method.Outputs.Unpack(bts)
	if err != nil {
		return nil, nil, err
	}
	if len(data) > 0 {
		return data[0].(*big.Int), data[1].(*big.Int), nil
	}
	return nil, nil, errors.New("no data")
}
