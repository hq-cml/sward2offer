package _64_sum_1_to_n

import (
	"github.com/hq-cml/sward2offer/temp"
	"testing"
)

func TestSum1ton_ArithmeticProgression(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				n: 5,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum1ton_ArithmeticProgression(tt.args.n); got != tt.want {
				t.Errorf("Sum1ton_ArithmeticProgression() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum1ton_Recurse(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				n: 5,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := temp.Sum(tt.args.n); got != tt.want {
				t.Errorf("Sum1ton_Recurse() = %v, want %v", got, tt.want)
			}
		})
	}
}
