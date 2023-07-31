/*
 * 给你一个整数数组nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
 * 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
 *
 * 输入：nums = [1,2,3]
 * 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
 *
 * 输入：nums = [0]
 * 输出：[[],[0]]
 */
package _078_sub_set

// 思路：直觉上，能够想到递归思路，类似于全排列
// 如果长度为len，所有的子集则是依次选取0,1,2...len个字符，作为全部子集
func PickSet(arr []int) [][]int {
	var ret [][]int
	length := len(arr)
	// 依次选取0,1,2...len个字符
	for n := 0; n <= length; n++ {
		ret = append(ret, pickSet(arr, n)...)
	}
	return ret
}

// 数组arr中，选择n个字符，所有的组合
func pickSet(arr []int, n int) [][]int {
	var ret [][]int
	if n == 0 {
		ret = append(ret, []int{})
		return ret
	}
	if n == len(arr) {
		ret = append(ret, arr)
		return ret
	}
	// 逐个取出一个字符，然后剩下的递归
	for i := 0; i < len(arr); i++ {
		v := arr[i]
		var left []int
		//这里很巧妙，直觉上应该把arr[:i]也算上
		//但实际上，如果算上会出现重复，不算则刚好
		//left = append(left, arr[:i]...)
		left = append(left, arr[i+1:]...)
		tmp := pickSet(left, n-1)
		for _, s := range tmp {
			s = append(s, v)
			ret = append(ret, s)
		}
	}
	return ret
}
