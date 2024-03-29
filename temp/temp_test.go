package temp

import (
	"math/rand"
	"testing"
)

func TestA(t *testing.T) {
	a := 2 + rand.Intn(2)
	t.Log(a)
}

func TestReplace(t *testing.T) {
	type args struct {
		arr [][]int
		src int
		dst int
		i   int
		j   int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				arr: [][]int{
					{0, 2, 0, 0, 0},
					{0, 2, 2, 2, 0},
					{0, 1, 2, 2, 0},
					{0, 2, 0, 0, 0},
				},
				src: 2,
				dst: 3,
				i:   2,
				j:   2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
