/**
 *
 */
package _347_freq_topk

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		//{
		//	name: "case1",
		//	args: args{
		//		nums: []int{1, 1, 1, 2, 2, 3},
		//		k:    2,
		//	},
		//	want: []int{1, 2},
		//},
		{
			name: "case2",
			args: args{
				nums: []int{1, 2},
				k:    2,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TopKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
