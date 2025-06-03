package jsonrpc

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

//	const (
//		defaultDecimals = 18
//	)

func (a *API) GetTRC20Decimal(ctx context.Context, token common.Address) (uint8, error) {
	decimals, err := a.rpcclient.GetTRC20Decimals(ctx, token)
	if err != nil {
		a.log.Error("error getting decimals", zap.String("token", token.String()), zap.String("error", err.Error()))
		return 0, err
	}
	return decimals, nil
}

func (a *API) GetTRC20Name(ctx context.Context, token common.Address) (string, error) {
	name, err := a.rpcclient.GetTRC20Name(ctx, token)
	if err != nil {
		a.log.Error("error getting name", zap.String("token", token.String()), zap.String("error", err.Error()))
		return "", err
	}
	return name, nil
}

func (a *API) GetTRC20Symbol(ctx context.Context, token common.Address) (string, error) {
	symbol, err := a.rpcclient.GetTRC20Symbol(ctx, token)
	if err != nil {
		a.log.Error("error getting symbol", zap.String("token", token.String()), zap.String("error", err.Error()))
		return "", err
	}
	return symbol, nil
}
