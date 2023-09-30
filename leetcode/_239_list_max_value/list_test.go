package _239_list_max_value

import (
	"fmt"
	"testing"
)

func TestNewMaxList(t *testing.T) {
	tests := []struct {
		name string
		want *MaxList
	}{
		{
			name: "case1",
			want: NewMaxList(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMaxList()
			got.Insert(2)
			got.Insert(3)
			got.Insert(4)
			fmt.Println(got.Max())
			fmt.Println(got.Pop())
			fmt.Println(got.Pop())
			fmt.Println(got.Max())
		})
	}
}

func TestNewMaxList2(t *testing.T) {
	tests := []struct {
		name string
		want *MaxList
	}{
		{
			name: "case1",
			want: NewMaxList(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := MaxListArr{
				Base: []int{},
				Help: []int{},
			}
			l.PushBack(2)
			l.PushBack(3)
			l.PushBack(4)
			l.PushBack(3)
			l.PushBack(2)
			fmt.Println(l.Max())
			fmt.Println(l.PopFront())
			fmt.Println(l.Max())
			fmt.Println(l.PopFront())
			fmt.Println(l.Max())
			fmt.Println(l.PopFront())
			fmt.Println(l.Max())
			fmt.Println(l.PopFront())
			fmt.Println(l.Max())
			//fmt.Println(l.PopFront())
			//fmt.Println(got.Max())
		})
	}
}
