package basic

import (
	"github.com/hq-cml/sward2offer/temp"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr  []int
		need int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				arr:  []int{1, 2, 3, 4},
				need: 5,
			},
			want: -1,
		},
		{
			name: "case2",
			args: args{
				arr:  []int{1, 2, 3, 4},
				need: 0,
			},
			want: -1,
		},
		{
			name: "case3",
			args: args{
				arr:  []int{1, 2, 3},
				need: 5,
			},
			want: -1,
		},
		{
			name: "case4",
			args: args{
				arr:  []int{1, 2, 3},
				need: 0,
			},
			want: -1,
		},
		{
			name: "case5",
			args: args{
				arr:  []int{1, 2, 3, 4, 5},
				need: 4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := temp.BinarySearch1(tt.args.arr, tt.args.need); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
