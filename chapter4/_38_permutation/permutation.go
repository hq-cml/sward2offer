/*
 * 面试题38：字符串的排列
 * 题目：输入一个字符串，打印出该字符串中字符的所有排列。例如输入字符串abc，
 * 则打印出由字符a、b、c所能排列出来的所有字符串abc、acb、bac、bca、cab和cba。
 */
package _38_permutation

import "fmt"

//思路：
//总体来说是一个递归思想，分为第一个字符和剩余字符，然后不断和后面交换。。。
//难度：5*
func Permutation(str string) {
	if len(str) == 0 {
		return
	}

	permutaion([]byte(str), 0)
}

//本质上是递归，很抽象
//首字母和后面的字母逐个替换
//结果分析：
//  abc：首字母a相当于完全没有换，仅仅递归bc
//  acb
//  bac：首字母和第二字母b交换，然后递归ac
//  bca
//  cba：首字母和第三字母c交换，然后递归ab
//  cab
func permutaion(arr []byte, beginIdx int) {
	//递归的结束条件：beginIdx最大值就是len(arr)-1
	//递归的结束动作：打印一次组合（此时一次组合已经完毕）
	if beginIdx >= len(arr)-1 {
		fmt.Println(string(arr))
		return
	}

	//beginIdx最初==0
	for i := beginIdx; i < len(arr); i++ {
		//swap，首字母和后面的字母逐个替换（beginIdx从0开始，所以第一次相当于没有换，不会遗漏）
		arr[beginIdx], arr[i] = arr[i], arr[beginIdx]

		//递归
		permutaion(arr, beginIdx+1)

		//swap，换回来，如果不换回来，就乱了，就无法达到逐个替换的目的，因为，比如：
		//一开始是1和2换，然后应该是1和3换，以此类推...
		//如果这里不换回来，就没法达到下一步1和3换的目的（变成了2和3换）
		arr[i], arr[beginIdx] = arr[beginIdx], arr[i]
	}
}
