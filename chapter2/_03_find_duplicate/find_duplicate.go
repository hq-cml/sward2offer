package _03_find_duplicate

import "errors"

//空间换时间，利用数组作为hash
//返回，exist标识是否重复，d标识重复的数，如果存在
func FindDuplicate0(s []int) (d int, exist bool){
	n := len(s)
	if n == 0 { return -1, false}
	for i:=0; i<n; i++ {
		if s[i] < 0 || s[i] >= n {
			return -1, false
		}
	}

	m := make([]bool, n)
	for i:=0; i<n; i++ {
		if ! m[s[i]] {
			m[s[i]] = true
			continue
		}
		return s[i], true
	}

	return -1, false
}

//利用题目的特点：
//如果不存在重复，那么0到n-1应该均匀分布在[0: n-1]上
//通过不断互换下标和对应值，从而达到节省了空间的目的
func FindDuplicate1(s []int) (d int, exist bool){
	n := len(s)
	if n == 0 { return -1, false}
	for i:=0; i<n; i++ {
		if s[i] < 0 || s[i] >= n {
			return -1, false
		}
	}

	for i:=0; i<n; i++ {
		for s[i] != i {    //紧着第i位，一直交换（颇为巧妙）
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

//2 3 5 4 3 2 6 7
//6 1 2 3 4 5 6 7
func countRange(s []int, min, max int) int {
	i := 0
	for _, v := range s {
		if v >= min && v <= max {
			i++
		}
	}
	return i
}

func FindDuplicate2 (s []int) (int, bool, error) {
	n := len(s)
	if n <= 1 { //数组为空或者仅有一个元素，则必然不重复
		return -1, false, nil
	}

	//检验数据合法性
	for i:=0; i<n; i++ {
		if s[i] < 1 || s[i] > (n-1) {
			return -1, false, errors.New("Ivalid number")
		}
	}

	beg := 1
	end := n-1

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