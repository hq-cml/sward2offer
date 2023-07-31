/*
 * 面试题57（一）：和为s的两个数字
 * 题目：输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们
 * 的和正好是s。如果有多对数字的和等于s，输出任意一对即可。
 *
 * 面试题57（二）：为s的连续正数序列
 * 题目：输入一个正数s，打印出所有和为s的连续正数序列（至少含有两个数）。
 * 例如输入15，由于1+2+3+4+5=4+5+6=7+8=15，所以结果打印出3个连续序列1～5、
 * 4～6和7～8。
 */
package _57_sum_is_s

//题目1:
//给一个递增数组，然后求出和为给定数字的任意两个成员
//思路：因为是递增序列，考虑到特点，设计两个指针，从两头不断逼近，调整
//     首尾加起来之和与s比较，大了则前移尾指针，小了则后移头指针。
//难度：3*
func FindTwoNum(arr []int, sum int) (int, int, bool) {
	//异常输入
	if sum < 0 {
		return 0, 0, false
	}

	length := len(arr)

	//滑动窗口
	i := 0
	j := length - 1
	for i < j {
		curr := arr[i] + arr[j]

		// 找到
		if curr == sum {
			return arr[i], arr[j], true
		}

		//左边界右移
		if curr < sum {
			i++
		}

		//右边界左移
		if curr > sum {
			j--
		}
	}
	return 0, 0, false
}

//题目2:
//和为s的连续递增序列，列举出全部
//思路和题目1类似，但两个指针均是从头开始，不断试探，但是区别在于要列举出全部组合
//所以这里要明确退出条件：如果small>sum/2，则序列之后必然已经超过sum，可以退出
//难度：3*
// 书上退出条件比较难想到，自己实现的退出条件更容易理解一些
func FindSequence(sum int) [][]int {
	if sum < 1 || sum == 2{
		return nil
	}
	if sum == 1 {
		return [][]int{{1}}
	}
	i:=1
	j:=2
	curr := i+j
	var ret [][]int

	// i追上j，作为退出条件，效率略低但是更容易想到
	for i!=j {
		if curr < sum {
			j++
			curr += j
		} else if curr == sum{
			var s []int
			for x:=i; x<=j; x++ {
				s = append(s, x)
			}
			ret = append(ret, s)
			j ++
			curr += j
		} else {
			curr -= i
			i++
		}
	}
	return ret
}
