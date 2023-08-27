/**
 *
 */
package _322_coin_changes

import "testing"

func TestCoinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				coins:  []int{2, 5, 10, 1},
				amount: 27,
			},
			want: 4,
		},
		{
			name: "case1",
			args: args{
				coins:  []int{10, 5, 7, 2},
				amount: 14,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CoinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("CoinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
