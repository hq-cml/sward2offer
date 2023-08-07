package _65_add_without_add

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				i: 101,
				j: 202,
			},
			want: 303,
		},
		{
			name: "case2",
			args: args{
				i: -5,
				j: 11,
			}, want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
