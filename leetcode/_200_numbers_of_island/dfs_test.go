package _200_numbers_of_island

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		arr [][]byte
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				arr: [][]byte{
					{
						1, 1, 1, 1, 0,
					},
					{
						1, 1, 0, 1, 0,
					},
					{
						1, 1, 0, 0, 0,
					},
					{
						0, 0, 0, 0, 0,
					},
				},
			},
			want: 1,
		},
		{
			name: "case2",
			args: args{
				arr: [][]byte{
					{
						1, 1, 0, 0, 0,
					},
					{
						1, 1, 0, 0, 0,
					},
					{
						0, 0, 1, 0, 0,
					},
					{
						0, 0, 0, 1, 1,
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IslandNum(tt.args.arr)
			if got != tt.want {
				t.Errorf("IslandNum() = %v, want %v", got, tt.want)
			}
			fmt.Println(got)
		})
	}
}
