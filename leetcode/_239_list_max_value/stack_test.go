package _239_list_max_value

import (
	"fmt"
	"testing"
)

func TestNewMaxStack(t *testing.T) {
	tests := []struct {
		name string
		want *MaxStack
	}{
		{
			name: "case1",
			want: NewMaxStack(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMaxStack()
			got.Push(4)
			got.Push(2)
			got.Push(5)
			got.Push(6)
			got.Push(1)
			fmt.Println(got.Max())
			fmt.Println(got.Pop())
			fmt.Println(got.Max())
			fmt.Println(got.Pop())
			fmt.Println(got.Max())
			fmt.Println(got.Pop())
			fmt.Println(got.Max())
		})
	}
}
