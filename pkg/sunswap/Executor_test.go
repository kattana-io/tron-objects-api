package sunswap

import (
	"github.com/kattana-io/tron-objects-api/pkg/api"
	"math/big"
	"testing"
)

func TestCanSell(t *testing.T) {
	type args struct {
		tokenIn  *api.Address
		tokenOut *api.Address
		amountIn *big.Int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "USDT-SUN",
			args: args{
				tokenIn:  api.FromBase58("TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"),
				tokenOut: api.FromBase58("TSSMHYeV2uE9qYH95DqyoCuNCzEL1NvU3S"),
				amountIn: big.NewInt(1000000000000000000),
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CanSell(tt.args.tokenIn, tt.args.tokenOut, tt.args.amountIn)
			if (err != nil) != tt.wantErr {
				t.Errorf("CanSell() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CanSell() got = %v, want %v", got, tt.want)
			}
		})
	}
}
