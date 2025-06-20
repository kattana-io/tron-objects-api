package rest

import (
	"fmt"
	"github.com/kattana-io/tron-objects-api/pkg/api/rest"
	rest2 "github.com/kattana-io/tron-objects-api/pkg/client/rest"
	"github.com/kattana-io/tron-objects-api/pkg/types"
	"testing"
)

func TestSunswapPair_GetTokenAddress(t *testing.T) {
	type fields struct {
		api     *rest.API
		address types.Address
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		error  bool
	}{
		{
			name: "USDT/TRX",
			fields: fields{
				api:     rest.NewAPI("", nil, rest2.NewTrongridURLProvider()),
				address: *types.NewFromBase58("TQn9Y2khEsLJW1ChVWFMSMeRDow5KcbLSE"),
			},
			want:  "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
			error: false,
		},
		{
			name: "KBC/TRX",
			fields: fields{
				api:     rest.NewAPI("", nil, rest2.NewTrongridURLProvider()),
				address: *types.NewFromBase58("TLP9cpp3B8WQNUXgbvjuYvvWUpmA4bzv4V"),
			},
			want:  "TUGrjLMegH5jvnaS2at6inkmNAWvtqTRFa",
			error: false,
		},
		{
			name: "BAGH/TRX",
			fields: fields{
				api:     rest.NewAPI("", nil, rest2.NewTrongridURLProvider()),
				address: *types.NewFromBase58("TLYE9Qz3Kue6EV8Na5aNki9Jkk8rZHp8Yo"),
			},
			want:  "TKZyr8jUu3aZZtUNQ6cRzML91oPUvUKxEJ",
			error: false,
		},
		{
			name: "BC/TRX",
			fields: fields{
				api:     rest.NewAPI("", nil, rest2.NewTrongridURLProvider()),
				address: *types.NewFromBase58("TLZbTD6Yg6iBWX6wKvYAsxE83vwweRVtuU"),
			},
			want:  "TJ7s4HjC1dYapZZnMq96VD6XcEEH3GKx1x",
			error: false,
		},
		{
			name: "META/TRX",
			fields: fields{
				api:     rest.NewAPI("", nil, rest2.NewTrongridURLProvider()),
				address: *types.NewFromBase58("TLaeHRNDoP2YaccQwqzYJdJjwfzJ6DcKR5"),
			},
			want:  "TPvyF8CD6eknF7hgfoZgZ9capryqQALQ52",
			error: false,
		},
		{
			name: "(Not sunswap pair) MEOX/JM",
			fields: fields{
				api:     rest.NewAPI("", nil, rest2.NewTrongridURLProvider()),
				address: *types.NewFromBase58("TMS2EaT8oKQcNmrbjArhi1umN1kFStRqrj"),
			},
			want:  "",
			error: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Pair{
				api:     tt.fields.api,
				address: tt.fields.address,
			}
			got, err := s.GetTokenAddress()
			hasErr := err != nil
			if got.ToBase58() != tt.want && hasErr != tt.error {
				t.Errorf("GetTokenAddress() = %v, want %v", got.ToBase58(), tt.want)
			}
		})
	}
}

func TestPair_GetReserves(t *testing.T) {
	type fields struct {
		api     *rest.API
		address types.Address
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "",
			fields: fields{
				api:     rest.NewAPI("", nil, rest2.NewTrongridURLProvider()),
				address: *types.NewFromBase58("TTdeCobmYxhfFBYUZbiQqbZ56zrFkSE5DG"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			s := &Pair{
				api:     tt.fields.api,
				address: tt.fields.address,
			}
			resA, resB, err := s.GetReserves()
			if err != nil {
				fmt.Println(resA.String(), resB.String())
			}
		})
	}
}
