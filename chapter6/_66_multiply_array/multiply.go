package _66_multiply_array

//给定数组A[0,1...,n-1]，求数组B[0,1...,n-1]，不能用除法
//其中B的元素B[i] = A[0]*..A[i-1]*A[i+1]*...A[n-1]
//思路：
//题目要求不能用除法，否则就很容易了
//常规思路就是两次循环，则复杂度O(n^2)
//一个改进思路，是购置一个逻辑二维矩阵
//然后分别按照从上到下和从下到上的顺序计算，复杂度降到O(n)
//难度：4*，这种思路实际中很难想到
func Multiply(src []int) []int {
    if len(src) == 0 {
        return src
    }
    if len(src) == 1 {
        return []int{1}
    }

    dst := make([]int, len(src))
    dst[0] = 1

    //先自上而下
    for i:=1; i<len(src); i++ {
        dst[i] = dst[i-1] * src[i-1]
    }

    //再自下而上
    tmp := 1
    for i:=len(src)-2; i >= 0; i-- {
        tmp *= src[i+1]
        dst[i] *= tmp
    }
    return dst
}