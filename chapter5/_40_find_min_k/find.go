/*
 * 面试题40：最小的k个数
 * 题目：输入n个整数，找出其中最小的k个数。例如输入4、5、1、6、2、7、3、8
 * 这8个数字，则最小的4个数字是1、2、3、4。
 */
package _40_find_min_k

import (
	"errors"
	"github.com/hq-cml/sward2offer/basic"
	"github.com/hq-cml/sward2offer/common"
)

//思路：
//TopK的问题，首先想到用大顶堆或者小顶堆
//其次作者还给了一个利用快排的partition函数的思路，这个基本上很难想到

//思路1：
//采用快排的partition函数的思路，这个基本上很难想到
//难度：5*
func FindMinK1(arr []int, k int) ([]int, error) {
	if len(arr) == 0 {
		return nil, errors.New("invalid")
	}

	if k <= 0 || k > len(arr) {
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

//思路2：
//利用大顶堆，复杂度：K*logN
func FindMinK2(arr []int, k int) ([]int, error) {
	//异常情况
	if len(arr) == 0 {
		return nil, errors.New("invalid")
	}

	//异常情况
	if k <= 0 || k > len(arr) {
		return nil, errors.New("invalid")
	}

	//特殊情况，k和数组长度相同
	if k == len(arr) {
		return arr, nil
	}

	//大顶堆
	heap := common.NewHeapInt(nil, false)
	for _, v := range arr {
		if heap.Len() < k {
			//数量还没到k，直接push
			heap.Push(v)
		} else {
			//数量达到k，是否push
			top := heap.Top()
			if top > v { //有更小的出现
				heap.Remove(0) //堆底层实现是一个数组，0元素，就是top元素
				heap.Push(v)
			}
		}
	}

	return heap.All(), nil
}
