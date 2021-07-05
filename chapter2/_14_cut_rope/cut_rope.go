/*
// 面试题14：剪绳子
// 题目：给你一根长度为n绳子，请把绳子剪成m段（m、n都是整数，n>1并且m≥1）。
// 每段的绳子的长度记为k[0]、k[1]、……、k[m]。k[0]*k[1]*…*k[m]可能的最大乘
// 积是多少？例如当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此
// 时得到最大的乘积18。
动态规划问题的四个特点：
问题通常是求一个最优解
大问题可以分解成小问题，大问题的最优解依赖小问题的最优解
一步步向下分解的过程中，会发现很多重叠（重复）的解
由于存在重叠，所以往往求解的过程应该从下到上（而分析问题的时候是从上到下），将下层结果先存起来

贪心算法和动态规划不同，当应用贪心的时候，每一步都是一个贪心的选择，确定能够得到最优解。
动态规划问题通常也可以用贪心算法，但是这需要一个证明的过程。所以还是动态规划稍微好一些。
 */

package _14_cut_rope

//动态规划，自底向上
func MaxCut(length int) int {
    //一些特殊情况
    if length <= 1 {
        return 0
    }
    if length == 2 {
        return 1
    }
    if length == 3 {
        return 2
    }

    arr := make([]int, length+1)
    //当长度小于3的时候，切不如不切
    arr[1] = 1
    arr[2] = 2
    arr[3] = 3

    for i := 4; i <= length; i++ {
        max := 0
        for j := 1; j < i; j ++ {
            if arr[j] * arr[i-j] > max {
                max = arr[j] * arr[i-j]
            }
        }
        arr[i] = max
    }

    return arr[length]
}