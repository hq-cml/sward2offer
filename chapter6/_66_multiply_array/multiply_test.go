package _66_multiply_array

import (
	"github.com/hq-cml/sward2offer/temp"
	"reflect"
	"testing"
)

func TestMultiply(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				src: []int{1, 2, 3, 4},
			},
			want: []int{24, 12, 8, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := temp.Mul(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
