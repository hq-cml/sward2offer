package _17_print_1_to_n

import "testing"


func TestPrintN_1(t *testing.T) {
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
				n: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintN_1(tt.args.n)
		})
	}
}

func TestPrintN_2(t *testing.T) {
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
				n: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintN_2(tt.args.n)
		})
	}
}