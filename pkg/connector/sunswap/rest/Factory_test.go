package rest

import (
	"github.com/kattana-io/tron-objects-api/pkg/api/rest"
	rest2 "github.com/kattana-io/tron-objects-api/pkg/client/rest"
	"github.com/kattana-io/tron-objects-api/pkg/types"
	"testing"
)

func TestFactory_GetPair(t *testing.T) {
	type args struct {
		tokenIn  *types.Address
		tokenOut *types.Address
	}
	tests := []struct {
		name    string
		args    args
		want    *types.Address
		wantErr bool
	}{
		{
			name: "BTT-USDT",
			args: args{
				tokenIn:  types.NewFromBase58("TAFjULxiVgT4qWk6UZwjqwZXTSaGaqnVp4"),
				tokenOut: types.NewFromBase58("TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"),
			},
			want:    types.NewFromBase58("TLKyq7eJ4YKbs3TGEvoBJWkAXWYQKWo2Nn"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFactory(rest.NewAPI("", nil, rest2.NewTrongridURLProvider()))
			got, err := f.GetPair(tt.args.tokenIn, tt.args.tokenOut)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPair() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.ToBase58() != tt.want.ToBase58() {
				t.Errorf("GetPair() got = %v, want %v", got, tt.want)
			}
		})
	}
}
