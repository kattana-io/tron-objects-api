package url

import "testing"

func TestTrongridUrlProvider_TriggerConstantContract(t *testing.T) {
	type fields struct {
		ApiKey string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "triggerconstantcontract",
			fields: fields{},
			want:   "https://api.trongrid.io/wallet/triggerconstantcontract",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &TrongridUrlProvider{
				ApiKey: tt.fields.ApiKey,
			}
			if got := n.TriggerConstantContract(); got != tt.want {
				t.Errorf("TriggerConstantContract() = %v, want %v", got, tt.want)
			}
		})
	}
}
