/*
// 面试题39：数组中出现次数超过一半的数字
// 题目：数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。例
// 如输入一个长度为9的数组{1, 2, 3, 2, 2, 2, 5, 4, 2}。由于数字2在数组中
// 出现了5次，超过数组长度的一半，因此输出2。

我最直观的思路，就是排序，然后取出中间那个数，这是一种思路，快速排序时间复杂度是O(NlogN)
作者给出了第一个思路，和快速排序类似，利用Partion函数，如果最终得到的a[mid]==正好是n/2，这说明找到。这种神仙思路面试很难找到的也很难实现。
第三种最巧妙，利用一个数字和她的计数器，来实现，遍历一次数组，相等加1，不等减1，为0则换下一个数字。
这个思路其实就是每次都在数组中去除不同的两个成员，直到最终剩下的那个肯定是超过半数的数字。
 */
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