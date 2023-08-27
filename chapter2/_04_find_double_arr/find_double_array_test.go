package _04_find_double_arr

import (
	"testing"
)

func TestFindDoubleArray(t *testing.T) {
	type args struct {
		arr      []int
		row      int
		col      int
		needFind int
	}

	a := []int{
		1, 2, 8, 9,
		2, 4, 9, 12,
		4, 7, 10, 13,
		6, 8, 11, 15}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				arr:      a,
				row:      4,
				col:      4,
				needFind: 7,
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				arr:      a,
				row:      4,
				col:      4,
				needFind: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDoubleArray(tt.args.arr, tt.args.row, tt.args.col, tt.args.needFind); got != tt.want {
				t.Errorf("FindDoubleArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
