package _022_generate_parenthesis

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case1",
			args: args{
				n: 3,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Generate(tt.args.n)
			fmt.Println(got)
		})
	}
}
