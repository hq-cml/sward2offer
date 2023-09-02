package _033_rotate_arr

import "testing"

func TestFind(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				arr:    []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		},
		{
			name: "case2",
			args: args{
				arr:    []int{4, 5, 6, 7, 0, 1, 2},
				target: 4,
			},
			want: 0,
		},
		{
			name: "case3",
			args: args{
				arr:    []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			want: -1,
		},
		{
			name: "case4",
			args: args{
				arr:    []int{1},
				target: 0,
			},
			want: -1,
		},
		{
			name: "case5",
			args: args{
				arr:    []int{3, 1},
				target: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
