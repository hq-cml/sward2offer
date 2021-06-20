package _60_n_dice_sum

import "testing"

func TestPrintProbability(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				n: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintProbability(tt.args.n)
		})
	}
}