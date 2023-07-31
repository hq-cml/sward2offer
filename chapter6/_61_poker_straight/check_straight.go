/*
 * 面试题61：扑克牌的顺子
 * 题目：从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。
 * 2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王可以看成任意数字。
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

//思路：
//这道题本质上是排序。大小王可以抽象成0，然后其他每个元素是1-13：
//先对数组排序，求出0的个数
//求各个元素的gap，最终要看整体的gap的和，与0个数的比较（0可以匹配任意数字）
//如果0的个数>gap总和，则是一个顺子，否则不是
//例子：12450，0可以代替3，所以可以是顺子
//     12560,0不够，5-2-1=2，至少需要两个0才能凑齐
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
		if arr.arr[i] == 0 {  //0的个数
			zeroCnt ++
			continue
		}

		gapCnt += (arr.arr[i+1] - arr.arr[i] - 1) //gap数累计（-1是因为天然就应该自增，这个不能算）
	}

	if zeroCnt >= gapCnt {
		return true
	}
	return false
}