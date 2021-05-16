package _15_binary_1_cnt

import (
	"testing"
)

func TestCalc1(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				15,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calc1(tt.args.n); got != tt.want {
				t.Errorf("Calc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalc2(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				15,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calc2(tt.args.n); got != tt.want {
				t.Errorf("Calc2() = %v, want %v", got, tt.want)
			}
		})
	}
}