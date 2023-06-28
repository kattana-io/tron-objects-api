package sunswap

import (
	"github.com/kattana-io/tron-objects-api/pkg/api"
	"github.com/kattana-io/tron-objects-api/pkg/url"
	"testing"
)

func TestFactory_GetPair(t *testing.T) {
	type args struct {
		tokenIn  *api.Address
		tokenOut *api.Address
	}
	tests := []struct {
		name    string
		args    args
		want    *api.Address
		wantErr bool
	}{
		{
			name: "BTT-USDT",
			args: args{
				tokenIn:  api.FromBase58("TAFjULxiVgT4qWk6UZwjqwZXTSaGaqnVp4"),
				tokenOut: api.FromBase58("TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"),
			},
			want:    api.FromBase58("TLKyq7eJ4YKbs3TGEvoBJWkAXWYQKWo2Nn"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFactory(api.NewAPI("", nil, url.NewTrongridURLProvider()))
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
