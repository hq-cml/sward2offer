package common

import (
	//"container/heap"
	"fmt"
	//"sort"
	"testing"
)

func TestSortHeap1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "case1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := []int{2, 1, 5, 100, 3, 6, 4, 5}
			fmt.Println("Org:", arr)

			h := NewMaxHeapInt(arr)
			//将h初始化为堆序列
			fmt.Println("Init:", h.Arr())
			h.Put(3)
			fmt.Println("Add 3:", h.Arr())

			fmt.Printf("minimum: %d\n", h.Get(0))

			for h.Len() > 0 {
				fmt.Printf("Pop:%d . Remain:%v\n", h.Remove(0), h.Arr())
			}
		})
	}
}

func TestSortHeap2(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "case1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewMinHeapInt(nil)

			//将h初始化为堆序列
			fmt.Println("Init:", h.Arr())
			h.Put(3)
			fmt.Println("Add 3:", h.Arr())
			h.Put( 2)
			h.Put( 1)
			h.Put( 5)
			h.Put( 100)
			h.Put( 3)
			h.Put( 6)
			h.Put( 4)
			h.Put( 5)
			fmt.Println("Fix:", h.Arr())
			fmt.Printf("minimum: %d\n", h.Get(0))

			for h.Len() > 0 {
				fmt.Printf("Pop:%d . Remain:%v\n", h.Remove(0), h.Arr())
			}
		})
	}
}