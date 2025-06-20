package types

import (
	"testing"
)

func TestAddress_ToBase58(t *testing.T) {
	type fields struct {
		address string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Address{
				address: tt.fields.address,
			}
			if got := a.ToBase58(); got != tt.want {
				t.Errorf("ToBase58() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddress_ToHex(t *testing.T) {
	tests := []struct {
		name   string
		fields *Address
		want   string
	}{
		{
			name:   "TQn9Y2khEsLJW1ChVWFMSMeRDow5KcbLSE to calleable hex",
			fields: NewFromBase58("TQn9Y2khEsLJW1ChVWFMSMeRDow5KcbLSE"),
			want:   "41a2726afbecbd8e936000ed684cef5e2f5cf43008",
		},
		{
			name:   "TL1AbcRiZExBUUbzuQXQ1ru7548cUczeyo to calleable hex",
			fields: NewFromBase58("TL1AbcRiZExBUUbzuQXQ1ru7548cUczeyo"),
			want:   "416e1015312f6bd4507a5d8052859afcd191c8ff13",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := tt.fields
			if got := a.ToHex(); got != tt.want {
				t.Errorf("ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromBase58(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TQn9Y2khEsLJW1ChVWFMSMeRDow5KcbLSE",
			args: args{
				input: "TQn9Y2khEsLJW1ChVWFMSMeRDow5KcbLSE",
			},
			want: "TQn9Y2khEsLJW1ChVWFMSMeRDow5KcbLSE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFromBase58(tt.args.input).ToBase58()
			if got != tt.want {
				t.Errorf("NewFromBase58() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromHex(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "41a2726afbecbd8e936000ed684cef5e2f5cf43008",
			args: args{
				input: "41a2726afbecbd8e936000ed684cef5e2f5cf43008",
			},
			want: "TQn9Y2khEsLJW1ChVWFMSMeRDow5KcbLSE",
		},
		{
			name: "a2726afbecbd8e936000ed684cef5e2f5cf43008",
			args: args{
				input: "a2726afbecbd8e936000ed684cef5e2f5cf43008",
			},
			want: "TQn9Y2khEsLJW1ChVWFMSMeRDow5KcbLSE",
		},
		{
			name: "41fba3416f7aac8ea9e12b950914d592c15c884372",
			args: args{
				input: "41fba3416f7aac8ea9e12b950914d592c15c884372",
			},
			want: "TYukBQZ2XXCcRCReAUguyXncCWNY9CEiDQ",
		},
		{
			name: "Empty string",
			args: args{
				input: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFromHex(tt.args.input).ToBase58()
			if got != tt.want {
				t.Errorf("NewFromHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trimZeroes(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Basic case",
			args: args{
				address: "000000000000000000000000a614f803b6fd780986a42c78ec9c7f77e6ded13c",
			},
			want: "a614f803b6fd780986a42c78ec9c7f77e6ded13c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimZeroes(tt.args.address); got != tt.want {
				t.Errorf("trimZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
