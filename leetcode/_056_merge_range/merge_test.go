package _056_merge_range

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		arr [][2]int
	}
	tests := []struct {
		name string
		args args
		want [][2]int
	}{
		{
			name: "case1",
			args: args{
				arr: [][2]int{{15, 18}, {1, 3}, {8, 10}, {2, 6}},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Merge(tt.args.arr)
			fmt.Println(got)
		})
	}
}
