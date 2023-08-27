/**
 *
 */
package _338_bit_cnt

import (
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				n: 2,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountBits(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
