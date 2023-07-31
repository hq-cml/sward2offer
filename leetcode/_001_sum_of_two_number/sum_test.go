package _001_sum_of_two_number

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		arr []int
		sum int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name:  "case1",
			args:  args{
				arr: []int{2,7,11,15},
				sum: 9,
			},
			want:  0,
			want1: 1,
		},
		{
			name:  "case2",
			args:  args{
				arr: []int{3,2,4},
				sum: 6,
			},
			want:  1,
			want1: 2,
		},
		{
			name:  "case3",
			args:  args{
				arr: []int{3,3},
				sum: 6,
			},
			want:  0,
			want1: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Sum(tt.args.arr, tt.args.sum)
			if got != tt.want {
				t.Errorf("Sum() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Sum() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}