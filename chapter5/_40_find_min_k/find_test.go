package _40_find_min_k

import (
	"fmt"
	"testing"
)

func TestFindMinK1(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name:    "case1",
			args:    args{
				arr: []int{4,1,3,2,6,5,7,8},
				k:   4,
			},
			wantErr: false,
		},
		{
			name:    "case2",
			args:    args{
				arr: []int{4,1,3,2,6,5,7,8},
				k:   5,
			},
			wantErr: false,
		},
		{
			name:    "case3",
			args:    args{
				arr: []int{4,1,3,2,6,5,7,8},
				k:   1,
			},
			wantErr: false,
		},
		{
			name:    "case4",
			args:    args{
				arr: []int{4,1,4,4,4,3,2,6,5,7,8},
				k:   4,
			},
			wantErr: false,
		},
		{
			name:    "case5",
			args:    args{
				arr: []int{4,1,4,4,4,3,2,6,5,7,8},
				k:   5,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindMinK1(tt.args.arr, tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindMinK1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestFindMinK2(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name:    "case1",
			args:    args{
				arr: []int{4,1,3,2,6,5,7,8},
				k:   4,
			},
			wantErr: false,
		},
		{
			name:    "case2",
			args:    args{
				arr: []int{4,1,3,2,6,5,7,8},
				k:   5,
			},
			wantErr: false,
		},
		{
			name:    "case3",
			args:    args{
				arr: []int{4,1,3,2,6,5,7,8},
				k:   1,
			},
			wantErr: false,
		},
		{
			name:    "case4",
			args:    args{
				arr: []int{4,1,4,4,4,3,2,6,5,7,8},
				k:   4,
			},
			wantErr: false,
		},
		{
			name:    "case5",
			args:    args{
				arr: []int{4,1,4,4,4,3,2,6,5,7,8},
				k:   5,
			},
			wantErr: false,
		},
		{
			name:    "case6",
			args:    args{
				arr: []int{4,1,4,4,4,3,2,6,5,7,8},
				k:   6,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindMinK2(tt.args.arr, tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindMinK2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}