package _155_min_stack

import (
	"fmt"
	"testing"
)

func TestNewMinStack(t *testing.T) {
	tests := []struct {
		name string
		want MinStack
	}{
		{
			name: "case1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewMinStack()
			fmt.Println(stack.Min())
			stack.Push(3)
			stack.Push(5)
			stack.Push(2)
			stack.Push(10)
			fmt.Println(stack.Min())
			stack.Pop()
			fmt.Println(stack.Min())
			stack.Pop()
			fmt.Println(stack.Min())
			stack.Pop()
			fmt.Println(stack.Min())
			stack.Push(4)
			fmt.Println(stack.Min())
			stack.Pop()
			fmt.Println(stack.Min())
			stack.Push(1)
			fmt.Println(stack.Min())
			stack.Pop()
			fmt.Println(stack.Min())
			stack.Pop()
			fmt.Println(stack.Min())
		})
	}
}
