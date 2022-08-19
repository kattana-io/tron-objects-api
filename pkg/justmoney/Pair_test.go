package justmoney

import (
	"github.com/kattana-io/tron-objects-api/pkg/api"
	"github.com/kattana-io/tron-objects-api/pkg/url"
	"testing"
)

func TestPair_Token0(t *testing.T) {
	type fields struct {
		api     *api.Api
		address api.Address
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "TRX/USDT",
			fields: fields{
				api:     api.NewApi("", nil, url.NewTrongridUrlProvider()),
				address: *api.FromBase58("TYA7DfE44XFsZEpBm7M2HAmEgU5kCtDDXg"),
			},
			want:    "TNUC9Qb1rRpS5CbWLmNMxXBjyFoydXjWFR",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Pair{
				api:     tt.fields.api,
				address: tt.fields.address,
			}
			got, err := s.Token0()
			if (err != nil) != tt.wantErr {
				t.Errorf("Token0() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.ToBase58() != tt.want {
				t.Errorf("Token0() got = %v, want %v", got.ToBase58(), tt.want)
			}
		})
	}
}

func TestPair_Token1(t *testing.T) {
	type fields struct {
		api     *api.Api
		address api.Address
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "TRX/USDT",
			fields: fields{
				api:     api.NewApi("", nil, url.NewTrongridUrlProvider()),
				address: *api.FromBase58("TYA7DfE44XFsZEpBm7M2HAmEgU5kCtDDXg"),
			},
			want:    "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Pair{
				api:     tt.fields.api,
				address: tt.fields.address,
			}
			got, err := s.Token1()
			if (err != nil) != tt.wantErr {
				t.Errorf("Token1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.ToBase58() != tt.want {
				t.Errorf("Token1() got = %v, want %v", got.ToBase58(), tt.want)
			}
		})
	}
}
