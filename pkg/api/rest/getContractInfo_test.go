package rest

import (
	"github.com/kattana-io/tron-objects-api/pkg/client/rest"
	"github.com/kattana-io/tron-objects-api/pkg/types"
	"go.uber.org/zap"
	"testing"
)

func TestApi_GetContractInfo(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	type fields struct {
		endpoint string
		log      *zap.Logger
		provider rest.APIURLProvider
	}
	type args struct {
		token *types.Address
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Sunswap factory origin",
			fields: fields{
				endpoint: "https://api.trongrid.io",
				log:      logger,
				provider: rest.NewTrongridURLProvider(),
			},
			args: args{
				token: types.NewFromHex("41a2726afbecbd8e936000ed684cef5e2f5cf43008"),
			},
			want:    "41eed9e56a5cddaa15ef0c42984884a8afcf1bdebb",
			wantErr: false,
		},
		{
			name: "Justmoney factory origin",
			fields: fields{
				endpoint: "https://api.trongrid.io",
				log:      logger,
				provider: rest.NewTrongridURLProvider(),
			},
			args: args{
				token: types.NewFromHex("412e12cbc3bacb0a80aef95b320773a7c64bc4372d"),
			},
			want:    "411294fce9b5932fc8c289fa1059a8b0b0a9f6daf5",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &API{
				endpoint: tt.fields.endpoint,
				log:      tt.fields.log,
				provider: tt.fields.provider,
			}
			got, err := a.GetContractInfo(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetContractInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.SmartContract.OriginAddress != tt.want {
				t.Errorf("Mismatch response: %s got %s", got.SmartContract.OriginAddress, tt.want)
			}
		})
	}
}
