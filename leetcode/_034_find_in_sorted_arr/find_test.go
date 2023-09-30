package _034_find_in_sorted_arr

import "testing"

func Test_findFirstK(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				arr: []int{1, 2, 3, 3, 3, 3, 4, 5},
				k:   3,
			},
			want: 2,
		},
		{
			name: "case2",
			args: args{
				arr: []int{1, 2, 3, 3, 3, 4, 5},
				k:   3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFirstK(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("findFirstK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLastK(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				arr: []int{1, 2, 3, 3, 3, 3, 4, 5},
				k:   3,
			},
			want: 5,
		},
		{
			name: "case2",
			args: args{
				arr: []int{1, 2, 3, 3, 3, 4, 5},
				k:   3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLastK(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("findLastK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcCnt(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				arr: []int{1, 2, 3, 3, 3, 3, 4, 5},
				k:   3,
			},
			want: 4,
		},
		{
			name: "case2",
			args: args{
				arr: []int{1, 2, 3, 3, 3, 4, 5},
				k:   3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcCnt(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("CalcCnt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindNumberLost(t *testing.T) {
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
				arr: []int{0, 1, 2, 3, 5, 6, 7},
			},
			want: 4,
		},
		{
			name: "case2",
			args: args{
				arr: []int{0, 1, 2, 3, 5, 6, 7, 8},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindNumberLost(tt.args.arr); got != tt.want {
				t.Errorf("FindNumberLost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindNumberSameAsIndex(t *testing.T) {
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
				arr: []int{-3, -1, 1, 3, 5},
			},
			want: 3,
		},
		{
			name: "case2",
			args: args{
				arr: []int{-3, -1, 1, 2, 4, 5},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindNumberSameAsIndex(tt.args.arr); got != tt.want {
				t.Errorf("FindNumberSameAsIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
