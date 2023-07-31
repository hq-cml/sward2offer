package _011_max_water

import "testing"

func TestMaxWater(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				arr: []int{1,1},
			},
			want: 1,
		},
		{
			name: "case2",
			args: args{
				arr: []int{1,8,6,2,5,4,8,3,7},
			},
			want: 49,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxWater(tt.args.arr); got != tt.want {
				t.Errorf("MaxWater() = %v, want %v", got, tt.want)
			}
		})
	}
}