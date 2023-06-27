package sunswap

import (
	"github.com/kattana-io/tron-objects-api/pkg/api"
	"github.com/kattana-io/tron-objects-api/pkg/url"
	"math/big"
	"reflect"
	"testing"
	"time"
)

const ZeroAddress = "T9yD14Nj9j7xAB4dbGeiX9h8unkKHxuWwb"

func TestRouter_swapETHToToken(t *testing.T) {
	type args struct {
		token        api.Address
		amountOutMin *big.Int
		to           api.Address
		deadline     *big.Int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "TRX->USDT",
			args: args{
				token:        *api.FromBase58(WTRXBase58),
				amountOutMin: big.NewInt(10 ^ 18),
				to:           *api.FromBase58(ZeroAddress),
				deadline:     big.NewInt(time.Now().Add(time.Minute).Unix()),
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := api.NewAPI("", nil, url.NewTrongridURLProvider())
			r := NewRouter(impl)
			got, err := r.SwapETHToToken(tt.args.token, tt.args.amountOutMin, tt.args.to, tt.args.deadline)
			if (err != nil) != tt.wantErr {
				t.Errorf("swapETHToToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapETHToToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
