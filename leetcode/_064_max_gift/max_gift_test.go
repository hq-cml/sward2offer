package _064_max_gift

import "testing"

func TestMaxGiftValue(t *testing.T) {
	type args struct {
		values  []int
		rows    int
		cols    int
		recurse bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				values:  []int{1, 10, 3, 8, 12, 2, 9, 6, 5, 7, 4, 11, 3, 7, 16, 5},
				rows:    4,
				cols:    4,
				recurse: true,
			},
			want: 53,
		},
		{
			name: "case2",
			args: args{
				values:  []int{1, 10, 3, 8, 12, 2, 9, 6, 5, 7, 4, 11, 3, 7, 16, 5},
				rows:    4,
				cols:    4,
				recurse: false,
			},
			want: 53,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxGiftValue(tt.args.values, tt.args.rows, tt.args.cols, tt.args.recurse); got != tt.want {
				t.Errorf("MaxGiftValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
