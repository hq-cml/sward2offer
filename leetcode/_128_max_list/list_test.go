package _128_max_list

import "testing"

func TestLongestConsecutive(t *testing.T) {
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
				nums: []int{100, 4, 200, 1, 3, 2},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestConsecutive(tt.args.nums); got != tt.want {
				t.Errorf("LongestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
