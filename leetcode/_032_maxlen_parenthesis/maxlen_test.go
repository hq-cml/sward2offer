package _032_maxlen_parenthesis

import (
	"testing"
)

func TestCalcMax(t *testing.T) {
	type args struct {
		arr []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				arr: []byte(")()())"),
			},
			want: 4,
		}, {
			name: "case2",
			args: args{
				arr: []byte("(()"),
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcMax(tt.args.arr); got != tt.want {
				t.Errorf("CalcMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
