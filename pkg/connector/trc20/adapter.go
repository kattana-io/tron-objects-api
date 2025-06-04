package trc20

import "context"

type TRC20 interface {
	TryToGetDecimals(ctx context.Context, try int64) (int32, bool)
	GetDecimals(ctx context.Context) (int32, error)
	GetSymbol(ctx context.Context) (string, error)
	GetName(ctx context.Context) (string, error)
}
