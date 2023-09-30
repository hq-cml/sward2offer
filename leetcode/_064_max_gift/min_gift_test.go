/**
 *
 */
package _064_max_gift

import "testing"

func TestMinGift(t *testing.T) {
	type args struct {
		values [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				values: [][]int{
					{1, 3, 1},
					{1, 5, 1},
					{4, 2, 1},
				},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinGift(tt.args.values); got != tt.want {
				t.Errorf("MinGift() = %v, want %v", got, tt.want)
			}
		})
	}
}
