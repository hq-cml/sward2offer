package _13_active_range

import (
	"testing"
)

func Test_calc(t *testing.T) {
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
				n: 101,
			},
			want: 2,
		},
		{
			name: "case2",
			args: args{
				n: 111,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want != calc(tt.args.n) {
				t.Errorf("%v != %v", tt.want, calc(tt.args.n))
			}
		})
	}
}

func TestCalcRange(t *testing.T) {
	type args struct {
		rows int
		cols int
		max  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				rows: 5,
				cols: 5,
				max:  100,
			},
			want: 25,
		},
		{
			name: "case2",
			args: args{
				rows: 100,
				cols: 100,
				max:  5,
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcRange(tt.args.rows, tt.args.cols, tt.args.max); got != tt.want {
				t.Errorf("CalcRange() = %v, want %v", got, tt.want)
			}
		})
	}
}