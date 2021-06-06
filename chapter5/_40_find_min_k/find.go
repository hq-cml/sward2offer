package _40_find_min_k

import (
	"errors"
	"github.com/hq-cml/sward2offer/basic"
	"github.com/hq-cml/sward2offer/common"
)

func FindMinK1(arr []int, k int) ([]int, error) {
	if len(arr) == 0 {
		return nil, errors.New("invalid")
	}

	if k <=0 || k > len(arr) {
		return nil, errors.New("invalid")
	}

	if k == len(arr) {
		return arr, nil
	}

	beg := 0
	end := len(arr) - 1
	for beg <= end {
		idx := basic.Partition(arr, beg, end)
		if idx == k-1 {
			return arr[:k], nil
		}

		if idx > k-1 {
			end = idx - 1
		} else {
			beg = idx + 1
		}
	}

	return nil, errors.New("Something wrong")
}

//利用大顶堆
func FindMinK2(arr []int, k int) ([]int, error) {
	if len(arr) == 0 {
		return nil, errors.New("invalid")
	}

	if k <=0 || k > len(arr) {
		return nil, errors.New("invalid")
	}

	if k == len(arr) {
		return arr, nil
	}

	//大顶堆
	heap := common.NewHeapInt(nil, false)
	for _, v := range arr {
		if heap.Len() < k {
			heap.Push(v)
		} else {
			top := heap.Get(0)
			if top > v {
				heap.Remove(0)
				heap.Push(v)
			}
		}
	}

	return heap.All(), nil
}