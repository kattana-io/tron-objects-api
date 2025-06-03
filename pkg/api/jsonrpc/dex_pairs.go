package jsonrpc

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/kattana-io/tron-objects-api/pkg/models"
	"go.uber.org/zap"
)

func (a *API) GetToken0(ctx context.Context, pair common.Address) (string, error) {
	token0, err := a.rpcclient.GetToken0(ctx, pair)
	if err != nil {
		a.log.Error("error getting token0", zap.String("pair", pair.String()), zap.String("error", err.Error()))
		return "", err
	}
	return token0, nil
}

func (a *API) GetToken1(ctx context.Context, pair common.Address) (string, error) {
	token1, err := a.rpcclient.GetToken1(ctx, pair)
	if err != nil {
		a.log.Error("error getting token1", zap.String("pair", pair.String()), zap.String("error", err.Error()))
		return "", err
	}
	return token1, nil
}

func (a *API) GetPair(ctx context.Context, factory, tokenA, tokenB common.Address) (string, error) {
	pair, err := a.rpcclient.GetPair(ctx, factory, tokenA, tokenB)
	if err != nil {
		a.log.Error("error getting pair",
			zap.String("factory", factory.String()),
			zap.String("tokenA", tokenA.String()),
			zap.String("tokenB", tokenB.String()),
			zap.String("error", err.Error()))
		return "", err
	}
	return pair, nil
}

func (a *API) GetPairToken(ctx context.Context, contract common.Address) (string, error) {
	pairToken, err := a.rpcclient.GetPairToken(ctx, contract)
	if err != nil {
		a.log.Error("error getting pair token address", zap.String("contract", contract.String()),
			zap.String("error", err.Error()))
		return "", err
	}
	return pairToken, nil
}

func (a *API) GetReserves(ctx context.Context, pair common.Address) (*models.Reserves, error) {
	reserve0, reserve1, blockTimestampLast, err := a.rpcclient.GetReserves(ctx, pair)
	if err != nil {
		a.log.Error("error getting reserves", zap.String("pair", pair.String()), zap.String("error", err.Error()))
		return nil, err
	}
	return &models.Reserves{
		Reserve0:           reserve0,
		Reserve1:           reserve1,
		BlockTimestampLast: blockTimestampLast,
	}, nil
}
