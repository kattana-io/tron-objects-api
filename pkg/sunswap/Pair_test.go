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
		{
			name: "KBC/TRX",
			fields: fields{
				api:     api.NewApi("", nil, url.NewTrongridUrlProvider()),
				address: *api.FromBase58("TLP9cpp3B8WQNUXgbvjuYvvWUpmA4bzv4V"),
			},
			want: "TUGrjLMegH5jvnaS2at6inkmNAWvtqTRFa",
		},
		{
			name: "BAGH/TRX",
			fields: fields{
				api:     api.NewApi("", nil, url.NewTrongridUrlProvider()),
				address: *api.FromBase58("TLYE9Qz3Kue6EV8Na5aNki9Jkk8rZHp8Yo"),
			},
			want: "TKZyr8jUu3aZZtUNQ6cRzML91oPUvUKxEJ",
		},
		{
			name: "BC/TRX",
			fields: fields{
				api:     api.NewApi("", nil, url.NewTrongridUrlProvider()),
				address: *api.FromBase58("TLZbTD6Yg6iBWX6wKvYAsxE83vwweRVtuU"),
			},
			want: "TJ7s4HjC1dYapZZnMq96VD6XcEEH3GKx1x",
		},
		{
			name: "META/TRX",
			fields: fields{
				api:     api.NewApi("", nil, url.NewTrongridUrlProvider()),
				address: *api.FromBase58("TLaeHRNDoP2YaccQwqzYJdJjwfzJ6DcKR5"),
			},
			want: "TPvyF8CD6eknF7hgfoZgZ9capryqQALQ52",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Pair{
				api:     tt.fields.api,
				address: tt.fields.address,
			}
			if got, _ := s.GetTokenAddress(); got.ToBase58() != tt.want {
				t.Errorf("GetTokenAddress() = %v, want %v", got.ToBase58(), tt.want)
			}
		})
	}
}
