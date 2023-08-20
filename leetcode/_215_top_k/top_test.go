/**
 *
 */
package _215_top_k

import "testing"

func TestFindKthHeap(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				nums: []int{3, 2, 1, 5, 6, 4},
				k:    2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindKthHeap(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("FindKthHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}
