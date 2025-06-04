package api

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/kattana-io/tron-objects-api/pkg/models"
)

type API interface {
	GetBlockNumber(ctx context.Context) (uint64, error)
	GetBlockByNum(ctx context.Context, num int64) (*models.Block, error)
	GetTransactionByHash(ctx context.Context, hash common.Hash) (*models.Transaction, error)
	GetTransactionByBlockNum(ctx context.Context, num int64) (*models.Transaction, error)
	GetContractInfo(ctx context.Context, contract common.Address) (*models.ContractInfo, error)

	GetTRC20Decimal(ctx context.Context, token common.Address) (uint8, error)
	GetTRC20Name(ctx context.Context, token common.Address) (string, error)
	GetTRC20Symbol(ctx context.Context, token common.Address) (string, error)

	GetToken0(ctx context.Context, pair common.Address) (string, error)
	GetToken1(ctx context.Context, pair common.Address) (string, error)
	GetPair(ctx context.Context, factory, tokenA, tokenB common.Address) (string, error)
	GetPairToken(ctx context.Context, contract common.Address) (string, error)
	GetReserves(ctx context.Context, pair common.Address) (*models.Reserves, error)
}
