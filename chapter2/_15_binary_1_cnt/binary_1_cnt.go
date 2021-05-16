package _15_binary_1_cnt

//利用1不断左移，按位与操作，统计个数
func Calc1(n int64) int {
	var x int64 = 1
	cnt := 0
	for x != 0 {
		if x & n != 0 {
			cnt ++
		}
		x = x << 1
	}
	return cnt
}

//利用减1后与操作，速度更快
func Calc2(n int64) int {
	cnt := 0
	for n != 0  {
		cnt ++
		n = n & (n-1)
	}
	return cnt
}