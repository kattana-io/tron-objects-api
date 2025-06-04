package justmoney

import "github.com/kattana-io/tron-objects-api/pkg/types"

type Pair interface {
	Token0() (*types.Address, error)
	Token1() (*types.Address, error)
}
