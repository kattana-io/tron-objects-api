package trc20

import (
	"github.com/kattana-io/tron-objects-api/pkg/api"
	"github.com/kattana-io/tron-objects-api/pkg/url"
	"testing"
)

func TestToken_GetDecimals(t1 *testing.T) {
	type fields struct {
		api     *api.API
		address *api.Address
	}
	tests := []struct {
		name    string
		fields  fields
		want    int32
		wantErr bool
	}{
		{
			name: "USDT Decimals",
			fields: fields{
				api:     api.NewAPI("", nil, url.NewTrongridURLProvider()),
				address: api.FromBase58("TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"),
			},
			want:    6,
			wantErr: false,
		},
		{
			name: "YBT Decimals",
			fields: fields{
				api:     api.NewAPI("", nil, url.NewTrongridURLProvider()),
				address: api.FromBase58("TPGx2NszcNRvSDiDM3c1YXM5XSvcRwrHLP"),
			},
			want:    6,
			wantErr: false,
		},
		{
			name: "Tiger Decimals",
			fields: fields{
				api:     api.NewAPI("", nil, url.NewTrongridURLProvider()),
				address: api.FromBase58("TQW9R6Ps1bXDJbTND6nTT5C5z2jDSbTDm8"),
			},
			want:    6,
			wantErr: false,
		},
		{
			name: "Tiger Decimals",
			fields: fields{
				api:     api.NewAPI("", nil, url.NewTrongridURLProvider()),
				address: api.FromBase58("TQW9R6Ps1bXDJbTND6nTT5C5z2jDSbTDm8"),
			},
			want:    6,
			wantErr: false,
		},
		{
			name: "ZSTokenName (ZS) Decimals",
			fields: fields{
				api:     api.NewAPI("", nil, url.NewTrongridURLProvider()),
				address: api.FromBase58("TMTc7Z9zG9oRTUN6Yhmb7kGa6PCJmV2GhE"),
			},
			want:    6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Token{
				api:     tt.fields.api,
				address: tt.fields.address,
			}
			got, err := t.GetDecimals()
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetDecimals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t1.Errorf("GetDecimals() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//nolint:dupl
func TestToken_GetName(t1 *testing.T) {
	type fields struct {
		api     *api.API
		address *api.Address
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "USDT",
			fields: fields{
				api:     api.NewAPI("", nil, url.NewTrongridURLProvider()),
				address: api.FromBase58("TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"),
			},
			want:    "USDT",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Token{
				api:     tt.fields.api,
				address: tt.fields.address,
			}
			got, err := t.GetName()
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t1.Errorf("GetName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//nolint:dupl
func TestToken_GetSymbol(t1 *testing.T) {
	type fields struct {
		api     *api.API
		address *api.Address
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "USDT",
			fields: fields{
				api:     api.NewAPI("", nil, url.NewTrongridURLProvider()),
				address: api.FromBase58("TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"),
			},
			want:    "USDT",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Token{
				api:     tt.fields.api,
				address: tt.fields.address,
			}
			got, err := t.GetSymbol()
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetSymbol() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t1.Errorf("GetSymbol() got = %v, want %v", got, tt.want)
			}
		})
	}
}
