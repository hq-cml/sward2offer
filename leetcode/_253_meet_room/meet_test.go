/**
 *
 */
package _253_meet_room

import "testing"

func TestMeet(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				intervals: [][]int{
					{1, 4},
					{2, 8},
					{3, 4},
					{5, 7},
					{5, 9},
				},
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				intervals: [][]int{
					{0, 30},
					{5, 10},
					{15, 20},
				},
			},
			want: 2,
		},
		{
			name: "case3",
			args: args{
				intervals: [][]int{
					{7, 10},
					{2, 4},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Meet(tt.args.intervals); got != tt.want {
				t.Errorf("Meet() = %v, want %v", got, tt.want)
			}
		})
	}
}
