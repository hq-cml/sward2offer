/*
 * 面试题45：把数组排成最小的数
 * 题目：输入一个正整数数组，把数组里所有数字拼接起来排成一个数，打印能拼
 * 接出的所有数字中最小的一个。例如输入数组{3, 32, 321}，则打印出这3个数
 * 字能排成的最小数字321323。
 */
package _45_arrange_array_for_min

import (
	"fmt"
	"sort"
	"strconv"
)

type MyArr struct {
	arr []int
}

//思路：
//这个题目关键是找到一种拼接后比大小的方法，进而搞出一种排序效果。
//本质上这个可以看成是一种另类的排序算法，两个数字m,n
//如果mn<nm，则定义为m<n
//如果mn>nm，则定义为m>n
//如果mn==nm，则定义为m==n
//而且，这里面还隐含这交换律
//同时，由于数字的拼接，可能溢出，所以要作为字符串来处理避免溢出的问题！
//难度：4*
func ArrangeMin(arr []int) string {
	//校验
	for _, i := range arr {
		if i < 0 {
			return ""
		}
	}

	//装入MyArr容器，进行排序
	a := MyArr{
		arr: arr,
	}
	sort.Sort(&a)

	//拼装最后结果
	str := ""
	for _, i := range a.arr {
		str += fmt.Sprintf("%d", i)
	}
	return str
}

//MyArr实现sort.Interface
//Len is the number of elements in the collection.
func (arr *MyArr) Len() int {
	return len(arr.arr)
}

// Swap swaps the elements with indexes i and j.
func (arr *MyArr) Swap(i, j int) {
	arr.arr[i], arr.arr[j] = arr.arr[j], arr.arr[i]
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (arr *MyArr) Less(i, j int) bool {
	s1 := strconv.Itoa(arr.arr[i])
	s2 := strconv.Itoa(arr.arr[j])

	return (s1 + s2) < (s2 + s1) //利用字符串的比较，来防止溢出
}
