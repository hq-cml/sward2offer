package _03_find_duplicate

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
