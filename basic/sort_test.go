package basic

import (
    "fmt"
    "testing"
)

func TestBubbleSort(t *testing.T) {
    type args struct {
        arr []int
    }
    tests := []struct {
        name string
        args args
    }{
        {
            name: "case1",
            args: args{
                arr: []int{5,1,3,4,2,7,8,6},
            },
        },
        {
            name: "case2",
            args: args{
                arr: []int{5,1,3,4,2,7,6},
            },
        },
        {
            name: "case3",
            args: args{
                arr: []int{1,2,3,4,5,6,7,8},
            },
        },
        {
            name: "case4",
            args: args{
                arr: []int{8,7,6,5,4,3,2,1},
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            BubbleSort(tt.args.arr)
            fmt.Println(tt.args.arr)
        })
    }
}


func TestQuickSort(t *testing.T) {
    type args struct {
        arr []int
    }
    tests := []struct {
        name string
        args args
    }{
        {
            name: "case1",
            args: args{
                arr: []int{5,1,3,4,2,7,8,6},
            },
        },
        {
            name: "case2",
            args: args{
                arr: []int{5,1,3,4,2,7,6},
            },
        },
        {
            name: "case3",
            args: args{
                arr: []int{1,2,3,4,5,6,7,8},
            },
        },
        {
            name: "case4",
            args: args{
                arr: []int{8,7,6,5,4,3,2,1},
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            QuickSort(tt.args.arr)
            fmt.Println(tt.args.arr)
        })
    }
}

func TestHeapSort(t *testing.T) {
    type args struct {
        arr []int
    }
    tests := []struct {
        name string
        args args
        want []int
    }{
        {
            name: "case1",
            args: args{
                arr: []int{5,1,3,4,2,7,8,6},
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := HeapSort(tt.args.arr)
            fmt.Println(got)
        })
    }
}

func TestSelectSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
        {
            name: "case1",
            args: args{
                arr: []int{5,1,3,4,2,7,8,6},
            },
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		    SelectSort(tt.args.arr)
            fmt.Println(tt.args.arr)
		})
	}
}

func TestInsertSort(t *testing.T) {
    type args struct {
        arr []int
    }
    tests := []struct {
        name string
        args args
    }{
        {
            name: "case1",
            args: args{
                arr: []int{5,1,3,4,2,7,8,6},
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            InsertSort(tt.args.arr)
            fmt.Println(tt.args.arr)
        })
    }
}

func TestMergeSort(t *testing.T) {
    type args struct {
        arr []int
    }
    tests := []struct {
        name string
        args args
        want []int
    }{
        {
            name: "case1",
            args: args{
                arr: []int{5,1,3,4,2,7,8,6},
            },
        },
        {
            name: "case2",
            args: args{
                arr: []int{5,1,3,4,2,7,6},
            },
        },
        {
            name: "case3",
            args: args{
                arr: []int{1,2,3,4,5,6,7,8},
            },
        },
        {
            name: "case4",
            args: args{
                arr: []int{8,7,6,5,4,3,2,1},
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            a := MergeSort(tt.args.arr)
            fmt.Println(a)
        })
    }
}

func Test_heapify(t *testing.T) {
    type args struct {
        arr         []int
        currIdx     int
        boundaryIdx int
    }
    tests := []struct {
        name string
        args args
    }{
        {
            name: "case1",
            args: args{
                arr: []int{1, 6, 7, 2, 3, 4, 5}, // 除了currIdx，均已满足大顶堆
                currIdx: 0,
                boundaryIdx: 6,
            },
        },
    }
    for _, tt := range tests {
        heapify(tt.args.arr, tt.args.currIdx, tt.args.boundaryIdx)
        fmt.Println(tt.args.arr)
    }
}

func Test_maxHeapify(t *testing.T) {
    type args struct {
        arr         []int
        boundaryIdx int
    }
    tests := []struct {
        name string
        args args
    }{
        {
            name: "case1",
            args: args{
                arr: []int{91, 60, 96, 13, 35, 65, 46, 65, 10, 30, 20, 31, 77, 81, 22},
                boundaryIdx: 14,
            },
        },
    }
    for _, tt := range tests {
        MaxHeapify(tt.args.arr, tt.args.boundaryIdx)
        fmt.Println(tt.args.arr)
    }
}

func TestSelfHeapSort(t *testing.T) {
    type args struct {
        arr []int
    }
    tests := []struct {
        name string
        args args
    }{
        {
            name: "case1",
            args: args{
                arr: []int{91, 60, 96, 13, 35, 65, 46, 65, 10, 30, 20, 31, 77, 81, 22},
            },
        },
    }
    for _, tt := range tests {
        SelfHeapSort(tt.args.arr)
        fmt.Println(tt.args.arr)
    }
}