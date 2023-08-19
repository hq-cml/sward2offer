package _055_jump_game

import (
	"testing"
)

func TestJump(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				arr: []int{2, 3, 1, 1, 4},
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				arr: []int{3, 2, 1, 0, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Jump(tt.args.arr); got != tt.want {
				t.Errorf("Jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
