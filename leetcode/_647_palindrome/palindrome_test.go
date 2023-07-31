package _647_palindrome

import "testing"

func TestCalcCnt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				s: "abc",
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				s: "aba",
			},
			want: 4,
		},
		{
			name: "case3",
			args: args{
				s: "aaa",
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcCnt(tt.args.s); got != tt.want {
				t.Errorf("CalcCnt() = %v, want %v", got, tt.want)
			}
		})
	}
}