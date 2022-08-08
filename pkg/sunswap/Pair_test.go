package sunswap

import (
	"github.com/kattana-io/tron-objects-api/pkg/api"
	"github.com/kattana-io/tron-objects-api/pkg/url"
	"testing"
)

func TestSunswapPair_GetTokenAddress(t *testing.T) {
	type fields struct {
		api     *api.Api
		address api.Address
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "USDT/TRX",
			fields: fields{
				api:     api.NewApi("", nil, url.NewTrongridUrlProvider()),
				address: *api.FromBase58("TQn9Y2khEsLJW1ChVWFMSMeRDow5KcbLSE"),
			},
			want: "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Pair{
				api:     tt.fields.api,
				address: tt.fields.address,
			}
			if got, _ := s.GetTokenAddress(); got.ToBase58() != tt.want {
				t.Errorf("GetTokenAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
