package _61_poker_straight

import (
	"fmt"
	"testing"
)

func TestCheck(t *testing.T) {
	type args struct {
		arr Arr
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				arr: Arr{
					arr: []int{2,3,1,6,4,5,},
				},
			},
			want: true,
		},
		{
			name: "case2",
			args: args{
				arr: Arr{
					arr: []int{0,0,2,3},
				},
			},
			want: true,
		},
		{
			name: "case3",
			args: args{
				arr: Arr{
					arr: []int{0,0,2,5},
				},
			},
			want: true,
		},
		{
			name: "case4",
			args: args{
				arr: Arr{
					arr: []int{0,0,2,6},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckIsStraght(tt.args.arr);
			fmt.Println(tt.args.arr)
			if  got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}