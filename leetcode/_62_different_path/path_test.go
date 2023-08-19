package _62_different_path

import (
	"github.com/hq-cml/sward2offer/temp"
	"testing"
)

func TestCalcPath(t *testing.T) {
	type args struct {
		m int
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
				m: 3,
				n: 2,
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				m: 7,
				n: 3,
			},
			want: 28,
		},
		{
			name: "case3",
			args: args{
				m: 3,
				n: 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := temp.CalcPath(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("CalcPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
