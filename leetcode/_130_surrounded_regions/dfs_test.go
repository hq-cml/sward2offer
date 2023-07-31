package _130_surrounded_regions

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		arr [][]byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				arr: [][]byte{
					{
						'X','X','X','X',
					},
					{
						'X','O','O','X',
					},
					{
						'X','X','O','X',
					},
					{
						'X','O','X','X',
					},
				},
			},
		},
		{
			name: "case1",
			args: args{
				arr: [][]byte{
					{
						'X','X','X','X','X','X',
					},
					{
						'X','O','O','X','O','X',
					},
					{
						'X','X','O','X','O','X',
					},
					{
						'X','O','X','X','O','X',
					},
				},
			},
		},
	}
	for _, tt := range tests {
		Solve(tt.args.arr)
		for _, line := range tt.args.arr {
			fmt.Println(string(line))
		}
	}
}