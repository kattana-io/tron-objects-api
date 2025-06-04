package jsonrpc

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kattana-io/tron-objects-api/pkg/api/jsonrpc"
	jsonrpc2 "github.com/kattana-io/tron-objects-api/pkg/client/jsonrpc"
	"github.com/kattana-io/tron-objects-api/pkg/types"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

const (
	rpcURL = "https://api.trongrid.io/jsonrpc"
)

type fields struct {
	api     *jsonrpc.API
	address types.Address
}

func TestPair_Token0(t *testing.T) {
	ctx := context.Background()
	logger := zap.NewNop()
	ethcli, err := ethclient.DialContext(ctx, rpcURL)
	assert.NoError(t, err)
	rpccli := jsonrpc2.NewJSONRPCClient(rpcURL)

	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "TRX/USDT",
			fields: fields{
				api:     jsonrpc.NewAPI(rpcURL, logger, rpccli, ethcli),
				address: *types.NewFromBase58("TYA7DfE44XFsZEpBm7M2HAmEgU5kCtDDXg"),
			},
			want:    "TNUC9Qb1rRpS5CbWLmNMxXBjyFoydXjWFR",
			wantErr: false,
		},
		{
			name: "TRX/USDT",
			fields: fields{
				api:     jsonrpc.NewAPI(rpcURL, logger, rpccli, ethcli),
				address: *types.NewFromBase58("TYA7DfE44XFsZEpBm7M2HAmEgU5kCtDDXg"),
			},
			want:    "TNUC9Qb1rRpS5CbWLmNMxXBjyFoydXjWFR",
			wantErr: false,
		},
		{
			name: "PROS/TRX",
			fields: fields{
				api:     jsonrpc.NewAPI(rpcURL, logger, rpccli, ethcli),
				address: *types.NewFromBase58("TA7hPWMWPWoadfFKWpTAdPYVZd3SNdtBDE"),
			},
			want:    "TFf1aBoNFqxN32V2NQdvNrXVyYCy9qY8p1",
			wantErr: false,
		},
		{
			name: "(Not a justmoney pair) AOT/TRX",
			fields: fields{
				api:     jsonrpc.NewAPI(rpcURL, logger, rpccli, ethcli),
				address: *types.NewFromBase58("TUer6gnscMcEX8Pid2FtwBCCS4coKRCrtL"),
			},
			want:    "1111",
			wantErr: true,
		},
	}
	for _, tt := range tests { //nolint:dupl
		t.Run(tt.name, func(t *testing.T) {
			s := &Pair{
				api:  tt.fields.api,
				addr: tt.fields.address,
			}
			got, err := s.Token0(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("Token0() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.ToBase58() != tt.want {
				t.Errorf("Token0() got = %v, want %v", got.ToBase58(), tt.want)
			}

			t.Logf("Token0(): %+v", got)
		})
	}
}

func TestPair_Token1(t *testing.T) {
	ctx := context.Background()
	logger := zap.NewNop()
	ethcli, err := ethclient.DialContext(ctx, rpcURL)
	assert.NoError(t, err)
	rpccli := jsonrpc2.NewJSONRPCClient(rpcURL)

	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "TRX/USDT",
			fields: fields{
				api:     jsonrpc.NewAPI(rpcURL, logger, rpccli, ethcli),
				address: *types.NewFromBase58("TYA7DfE44XFsZEpBm7M2HAmEgU5kCtDDXg"),
			},
			want:    "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
			wantErr: false,
		},
	}
	for _, tt := range tests { //nolint:dupl
		t.Run(tt.name, func(t *testing.T) {
			s := &Pair{
				api:  tt.fields.api,
				addr: tt.fields.address,
			}
			got, err := s.Token1(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("Token1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.ToBase58() != tt.want {
				t.Errorf("Token1() got = %v, want %v", got.ToBase58(), tt.want)
			}

			t.Logf("Token1(): %+v", got)
		})
	}
}
