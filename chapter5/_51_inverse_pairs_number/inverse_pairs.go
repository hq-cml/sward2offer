/*
 * 面试题51：数组中的逆序对
 * 题目：在数组中的两个数字如果前面一个数字大于后面的数字，则这两个数字组
 * 成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
 */
package _51_inverse_pairs_number

//思路1：
//  暴力的循环，复杂度O(N^2)
//思路2：
//  先拆分成最细粒度子数组，然后利用归并排序逐个算逆序数。。。
//  本质上就是将逆序数的统计，融合在了归并排序中去实现
//  正常人不可能会想到，复杂度O(N*logN)
//  难度：6*
func InverseParis(data []int) int {
	if len(data) == 0 {
		return 0
	}
	dst := make([]int, len(data))
	copy(dst, data) //深拷贝

	cnt := inversePairs(data, dst, 0, len(data)-1)
	return cnt
}

//递归
//  比较烧脑的地方是src和dst的参数位置，是来回调换的
//  因为这里是递归，每次都是以dst为目标，src作为源数据
func inversePairs(src, dst []int, start, end int) int {
	//递归的结束条件
	if start == end {
		dst[start] = src[start]
		return 0
	}

	mid := (end - start) / 2
	//递归左右
	//比较烧脑的：dst和src的参数位置，是调换了的
	//因为这里是分别将前后src的前后两部分，分别排好序，才好最终填入dst
	left := inversePairs(dst, src, start, start+mid)
	right := inversePairs(dst, src, start+mid+1, end)

	//到了此处，递归已经结束，此时src的前半部分和后半部分，已经排序完毕
	//归并排序是针对dst进行处理的，这也是为什么需要调换的原因
	i := start + mid //从后向前归并，i和j分别指向两节的末尾
	j := end

	indexCopy := end
	cnt := 0
	//归并处理的同时，计算逆序数
	for i >= start && j >= start+mid+1 {
		if src[i] > src[j] {
			dst[indexCopy] = src[i]
			cnt += (j - start - mid) //累积逆序数，真是很抽象，只可举例
			indexCopy--
			i--
		} else {
			dst[indexCopy] = src[j]
			indexCopy--
			j--
		}
	}

	//剩余的处理，逆序数不会再有变化
	for i >= start {
		dst[indexCopy] = src[i]
		indexCopy--
		i--
	}
	for j >= start+mid+1 {
		dst[indexCopy] = src[j]
		indexCopy--
		j--
	}

	return cnt + left + right
}