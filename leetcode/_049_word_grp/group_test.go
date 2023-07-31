package _049_word_grp

import (
	"fmt"
	"testing"
)

func Test_sortStr(t *testing.T) {
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
				s: "bca",
			},
			want: "abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortStr(tt.args.s); got != tt.want {
				t.Errorf("sortStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "case1",
			args: args{
				arr: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Group(tt.args.arr)
			fmt.Println(got)
		})
	}
}
