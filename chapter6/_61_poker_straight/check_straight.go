package _61_poker_straight

import (
	"sort"
)

//实现sort.Interface接口，支持排序
type Arr struct {
	arr []int
}

func (a Arr) Len() int {
	return len(a.arr)
}

func (a Arr) Less(i, j int) bool {
	return a.arr[i] < a.arr[j]
}

func (a Arr) Swap(i, j int) {
	a.arr[i], a.arr[j] = a.arr[j], a.arr[i]
}

//计算是否是顺子，0表示大小鬼，可代替任意牌
//先排序，然后算数差的和，如果小于等于0的个数，就表示可以凑成顺子
//难度：3*
func CheckIsStraght(arr Arr) bool {
	if arr.Len() == 0 {
		return false
	}
	if arr.Len() == 1 {
		return true
	}
	//排序
	sort.Sort(arr)

	//计算0的个数以及匹克派数差
	zeroCnt := 0
	gapCnt := 0
	for i:=0; i<arr.Len()-1; i++ {
		if arr.arr[i] == 0 {
			zeroCnt ++
			continue
		}
		j := i+1

		gapCnt += (arr.arr[j] - arr.arr[i] - 1)
	}

	if zeroCnt >= gapCnt {
		return true
	}
	return false
}