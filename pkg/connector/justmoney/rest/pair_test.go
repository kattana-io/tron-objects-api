package rest

import (
	"github.com/kattana-io/tron-objects-api/pkg/api/rest"
	rest2 "github.com/kattana-io/tron-objects-api/pkg/client/rest"
	"github.com/kattana-io/tron-objects-api/pkg/types"
	"testing"
)

func TestPair_Token0(t *testing.T) {
	type fields struct {
		api     *rest.API
		address types.Address
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
				api:     rest.NewAPI("", nil, rest2.NewTrongridURLProvider()),
				address: *types.NewFromBase58("TYA7DfE44XFsZEpBm7M2HAmEgU5kCtDDXg"),
			},
			want:    "TNUC9Qb1rRpS5CbWLmNMxXBjyFoydXjWFR",
			wantErr: false,
		},
		{
			name: "TRX/USDT",
			fields: fields{
				api:     rest.NewAPI("", nil, rest2.NewTrongridURLProvider()),
				address: *types.NewFromBase58("TYA7DfE44XFsZEpBm7M2HAmEgU5kCtDDXg"),
			},
			want:    "TNUC9Qb1rRpS5CbWLmNMxXBjyFoydXjWFR",
			wantErr: false,
		},
		{
			name: "PROS/TRX",
			fields: fields{
				api:     rest.NewAPI("", nil, rest2.NewTrongridURLProvider()),
				address: *types.NewFromBase58("TA7hPWMWPWoadfFKWpTAdPYVZd3SNdtBDE"),
			},
			want:    "TFf1aBoNFqxN32V2NQdvNrXVyYCy9qY8p1",
			wantErr: false,
		},
		{
			name: "(Not a justmoney pair) AOT/TRX",
			fields: fields{
				api:     rest.NewAPI("", nil, rest2.NewTrongridURLProvider()),
				address: *types.NewFromBase58("TUer6gnscMcEX8Pid2FtwBCCS4coKRCrtL"),
			},
			want:    "1111",
			wantErr: true,
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
		api     *rest.API
		address types.Address
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
				api:     rest.NewAPI("", nil, rest2.NewTrongridURLProvider()),
				address: *types.NewFromBase58("TYA7DfE44XFsZEpBm7M2HAmEgU5kCtDDXg"),
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
