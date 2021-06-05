package _39_find_majority_number

import (
	"errors"
	"github.com/hq-cml/sward2offer/basic"
)

//采用一种类似于快排分区的方式，不断探索中值
//时间复杂度：O(logN)
//难度：3*
func FindMajority1(arr []int) (int, error){
	if len(arr) == 0 {
		return 0, errors.New("invalid")
	}

	beg := 0
	end := len(arr) - 1
	mid := len(arr)/2
	for beg <= end {
		idx := basic.Partition(arr, beg, end)
		if idx == mid {
			return arr[mid], nil
		}

		if idx < mid {
			beg = idx + 1
		} else {
			end = idx - 1
		}
	}

	return 0, errors.New("Something wrong")
}

//采用一次遍历，进行相同计数的方式，不同则减一
//因为超过一半，则最终剩下的必然是最多的数
//时间复杂度：N
//难度：2*
func FindMajority2(arr []int) (int, error){
	if len(arr) == 0 {
		return 0, errors.New("Invalid arr")
	}

	currNum := arr[0]
	currCnt := 1
	for i:=1; i<len(arr); i++ {
		if currNum == arr[i] {
			currCnt ++
		} else {
			currCnt --
		}

		if currCnt == 0 {
			currNum = arr[i]
		}
	}

	return currNum, nil
}