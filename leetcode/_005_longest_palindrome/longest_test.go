package _005_longest_palindrome

import "testing"

func TestLongest(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				s: "abc",
			},
			want: "c",
		},
		{
			name: "case2",
			args: args{
				s: "aba",
			},
			want: "aba",
		},
		{
			name: "case3",
			args: args{
				s: "aaa",
			},
			want: "aaa",
		},
		{
			name: "case4",
			args: args{
				s: "babad",
			},
			want: "aba",
		},
		{
			name: "case5",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Longest(tt.args.s); got != tt.want {
				t.Errorf("Longest() = %v, want %v", got, tt.want)
			}
		})
	}
}