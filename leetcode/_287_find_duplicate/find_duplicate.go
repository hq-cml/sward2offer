/*
 * 找出数组的重复数字
 *
 * 题一：找出数组中重复的数字
 * 题目：在一个长度为n的数组里的所有数字都在0到n-1的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，
 * 也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。例如，如果输入长度为7的数组{2, 3, 1, 0, 2, 5, 3}，
 * 那么对应的输出是重复的数字2或者3。
 *
 * 题二：不修改数组找出重复的数字
 * 题目：在一个长度为n+1的数组里的所有数字都在1到n的范围内，所以数组中至
 * 少有一个数字是重复的。请找出数组中任意一个重复的数字，但不能修改输入的
 * 组。例如，如果输入长度为8的数组{2, 3, 5, 4, 3, 2, 6, 7}，那么对应的
 * 输出是重复的数字2或者3。
 */
package _287_find_duplicate

import "errors"

/*
 * 难度：4*
 * 这两类题目，感觉面试中短时间很难想到。
 * 自己直接实现了一个空间换时间的思路。这个方法的空间复杂度和时间复杂度都是O(N)，
 * 如果需要改进空间复杂度，也可以考虑使用bitmap改进型，及2个bit，表示一位。
 *
 * 另一种朴素的思路是排序，不过这个显然复杂度太高了
 */

// 题一：思路1
// 空间换时间，利用hash来实现判重
// 由于题目的特点，可以利用数组替换成为 hash
// 返回，exist标识是否重复，d标识重复的数，如果存在
func FindDuplicate0(s []int) (d int, exist bool) {
	n := len(s)
	//基础校验
	if n == 0 {
		return -1, false
	}
	for i := 0; i < n; i++ {
		if s[i] < 0 || s[i] >= n {
			return -1, false
		}
	}

	//正式逻辑
	m := make([]bool, n)
	for i := 0; i < n; i++ {
		if !m[s[i]] {
			m[s[i]] = true
			continue
		}
		return s[i], true
	}

	return -1, false
}

// 题一：思路2利用题目的特点：
// 如果不存在重复，那么0到n-1应该均匀分布在[0: n-1]上
// 通过不断互换下标和对应值，从而达到节省了空间的目的
func FindDuplicate1(s []int) (d int, exist bool) {
	n := len(s)
	//基础校验
	if n == 0 {
		return -1, false
	}
	for i := 0; i < n; i++ {
		if s[i] < 0 || s[i] >= n {
			return -1, false
		}
	}

	//实施交换
	//如果不存在重复，那么应该s[i] == i
	//反之就是如果s[i] != i，就可能存在重复
	for i := 0; i < n; i++ {
		for s[i] != i { //紧着第i位，一直交换，直到s[i]==i为止（颇为巧妙）
			x := s[i]
			if x == s[x] {
				return x, true
			}

			//swap
			s[x], s[i] = s[i], s[x]
		}
	}

	return -1, false
}

// 题二：
// 采用了一种类似于二分查找的思路，只不过划分的抓手不是下标，而是范围区间
// 比如分为1-4, 4-7两个区间，然后看哪个区间的长度超过n/2，
// 超过，则意味着重复，则对该区域存在重复继续探测
func FindDuplicate2(s []int) (int, bool, error) {
	n := len(s)
	if n <= 1 { //数组为空或者仅有一个元素，则必然不重复
		return -1, false, nil
	}
	//检验数据合法性
	for i := 0; i < n; i++ {
		if s[i] < 1 || s[i] > (n-1) {
			return -1, false, errors.New("Ivalid number")
		}
	}

	//正式逻辑，beg和end并非下标，而是数字的范围
	beg := 1
	end := n - 1
	for beg <= end { //n >= 2，所以end最小是1
		if beg == end {
			return beg, true, nil
		}

		mid := (end-beg)>>1 + beg
		num := countRange(s, beg, mid)
		if num <= (mid - beg + 1) {
			//重复数字，不可能在此，则换另一半进行探测
			beg = mid + 1
		} else {
			//重复有可能在此，则继续探测
			end = mid
		}
	}
	return -1, false, nil
}

// 数组元素符合给定区间大小的个数
func countRange(s []int, min, max int) int {
	i := 0
	for _, v := range s {
		if v >= min && v <= max {
			i++
		}
	}
	return i
}
