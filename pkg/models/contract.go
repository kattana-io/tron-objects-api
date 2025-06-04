package models

import "math/big"

type Reserves struct {
	Reserve0           *big.Int `json:"reserve0"`
	Reserve1           *big.Int `json:"reserve1"`
	BlockTimestampLast uint32   `json:"block_timestamp_last"`
}
