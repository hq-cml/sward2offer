package common
import (
	"fmt"
	"testing"
)

func TestNewStackInt(t *testing.T) {
	tests := []struct {
		name string
		want Stack
	}{
		{
			name: "case1",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStackInt()
			s.Push(1)
			s.Push(2)
			s.Push(3)

			fmt.Println(s.Top())
			fmt.Println("Len:", s.Len())

			fmt.Println(s.Pop())
			fmt.Println(s.Top())
			fmt.Println("Len:", s.Len())

			fmt.Println(s.Pop())
			fmt.Println(s.Pop())
			fmt.Println("Len:", s.Len())

			//nil
			fmt.Println(s.Pop())
		})
	}
}