/*
 * 给定一个长度为 n 的整数数组height。有n条垂线，第 i 条线的两个端点是(i, 0)和(i, height[i])。
 * 找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。返回容器可以储存的最大水量。
 *
 * 例子：
 * 输入：[1,8,6,2,5,4,8,3,7]
 * 输出：49
 *
 * 输入：[1,1]
 * 输出: 1
 */
package _011_max_water

// 思路1：暴力法，复杂度是O(n^2)
// 思路2：感觉上双指针应该可行
//
//	通常双指针有两种方式，第一种是都在开始位置，第二种是一头一尾
//	本题选用第二种方式可以解决问题
func MaxWater(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	i, j := 0, len(arr)-1
	max := 0
	for i < j {
		water := calcWater(arr, i, j)
		if water > max {
			max = water
		}
		// 木桶原理，取决于最短板，此时如果移动长的板只可能变得更少！
		if arr[i] < arr[j] {
			// i更矮，则应该右移i，因为如果左移j，无论怎么移动，都不可能得到更多的水
			i++
		} else {
			// j更矮，则应该左移j，因为如果右移i，无论怎么移动，都不可能得到更多的水
			j--
		}
	}
	return max
}

func calcWater(arr []int, i, j int) int {
	if i >= j {
		return 0
	}
	height := arr[i]
	if height > arr[j] {
		height = arr[j]
	}
	return (j - i) * height
}
