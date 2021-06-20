package _63_stock_max_earnings

import "testing"

func TestMaxStock(t *testing.T) {
	type args struct {
		price []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				price: []int{9, 11, 8, 5, 7, 12, 16, 14},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxStock(tt.args.price); got != tt.want {
				t.Errorf("MaxStock() = %v, want %v", got, tt.want)
			}
		})
	}
}