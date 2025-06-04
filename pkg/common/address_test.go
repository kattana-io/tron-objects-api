package common

import (
	"testing"
)

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
