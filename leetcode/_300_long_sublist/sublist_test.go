/**
 *
 */
package _300_long_sublist

import "testing"

func TestSubList(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubList(tt.args.nums); got != tt.want {
				t.Errorf("SubList() = %v, want %v", got, tt.want)
			}
		})
	}
}
