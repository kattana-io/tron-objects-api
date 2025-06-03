package rest

import "testing"

func TestNodeUrlProvider_TriggerConstantContract(t *testing.T) {
	type fields struct {
		host string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "triggerconstantcontract",
			fields: fields{
				host: "https://rpc.ankr.com/http/tron",
			},
			want: "https://rpc.ankr.com/http/tron/walletsolidity/triggerconstantcontract",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NodeURLProvider{
				host: tt.fields.host,
			}
			if got := n.TriggerConstantContract(); got != tt.want {
				t.Errorf("TriggerConstantContract() = %v, want %v", got, tt.want)
			}
		})
	}
}
