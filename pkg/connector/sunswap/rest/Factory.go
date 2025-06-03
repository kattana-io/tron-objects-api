package rest

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	tronAbi "github.com/kattana-io/go-tron/abi"
	"github.com/kattana-io/tron-objects-api/pkg/api/rest"
	"github.com/kattana-io/tron-objects-api/pkg/types"
)

type Factory struct {
	api     *rest.API
	address *types.Address
	abi     tronAbi.ABI
}

//nolint:lll
const sunswapV2FactoryABI = `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"token0","type":"address"},{"indexed":true,"internalType":"address","name":"token1","type":"address"},{"indexed":false,"internalType":"address","name":"pair","type":"address"},{"indexed":false,"internalType":"uint256","name":"","type":"uint256"}],"name":"PairCreated","type":"event"},{"inputs":[],"name":"INIT_CODE_PAIR_HASH","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"allPairs","outputs":[{"internalType":"address","name":"pair","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"allPairsLength","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"tokenA","type":"address"},{"internalType":"address","name":"tokenB","type":"address"}],"name":"createPair","outputs":[{"internalType":"address","name":"pair","type":"address"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"feeTo","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"feeToSetter","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"tokenA","type":"address"},{"internalType":"address","name":"tokenB","type":"address"}],"name":"getPair","outputs":[{"internalType":"address","name":"pair","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"setFeeTo","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"setFeeToSetter","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
const sunswapV2FactoryAddress = "TKWJdrQkqHisa1X8HUdHEfREvTzw4pMAaY"
const factoryOwner = "4128fb7be6c95a27217e0e0bff42ca50cd9461cc9f"

func NewFactory(impl *rest.API) *Factory {
	contract := tronAbi.ABI{}
	err := contract.UnmarshalJSON([]byte(sunswapV2FactoryABI))
	if err != nil {
		return nil
	}

	return &Factory{
		api:     impl,
		address: types.NewFromBase58(sunswapV2FactoryAddress),
		abi:     contract,
	}
}

func (f *Factory) GetPair(tokenIn, tokenOut *types.Address) (*types.Address, error) {
	addrA, err := tokenIn.ToGoTronAddr()
	if err != nil {
		return nil, err
	}
	addrB, err := tokenOut.ToGoTronAddr()
	if err != nil {
		return nil, err
	}
	bts := f.abi.Functions["getPair"].Encode(addrA, addrB)
	pair, err := f.api.GetFactoryPair(f.address, types.NewFromHex(factoryOwner), hexutil.Encode(bts)[2:])
	if err != nil {
		return nil, err
	}
	return pair, nil
}
