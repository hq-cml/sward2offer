/*
// 面试题61：扑克牌的顺子
// 题目：从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。
// 2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王可以看成任意数字。

这道题本质上利用一个数组可以解决。首先大小王可以抽象成0，然后其他每个元素是1-13：
先对数组排序，求出0的个数
求各个元素的gap，最终要看整体的gap的和，与0个数的比较（0可以匹配任意数字）
如果0的个数>gap总和，则是一个顺子，否则不是
 */
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