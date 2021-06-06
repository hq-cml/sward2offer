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

			h := NewHeapInt(arr, false) //大顶堆
			//将h初始化为堆序列
			fmt.Println("Init:", h.All())
			h.Push(3)
			fmt.Println("Add 3:", h.All())

			fmt.Printf("maximum: %d\n", h.Get(0))

			for h.Len() > 0 {
				fmt.Printf("Pop:%d . Remain:%v\n", h.Remove(0), h.All())
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
			h := NewHeapInt(nil, true)

			//将h初始化为堆序列
			fmt.Println("Init:", h.All())
			h.Push(3)
			fmt.Println("Add 3:", h.All())
			h.Push( 2)
			h.Push( 1)
			h.Push( 5)
			h.Push( 100)
			h.Push( 3)
			h.Push( 6)
			h.Push( 4)
			h.Push( 5)
			fmt.Println("Fix:", h.All())
			fmt.Printf("minimum: %d\n", h.Get(0))

			for h.Len() > 0 {
				fmt.Printf("Pop:%d . Remain:%v\n", h.Pop(), h.All())
			}
		})
	}
}