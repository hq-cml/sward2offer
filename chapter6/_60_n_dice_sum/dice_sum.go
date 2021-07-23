/*
 * 面试题60：n个骰子的点数
 * 题目：把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。
 */
package _60_n_dice_sum

import (
	"fmt"
	"math"
)

const (
	DiceMaxNumber = 6
)

//思路：
//本质上是一道考察建模的题目，基本思想是利用递归来实现所有出现的数字的次数然后除以总次数。
//n个骰子，可能出现的数字分部：n~6n，总数量为6^n
//极其烧脑，利用两个数组，不断替换，来算出每一轮的各个可能总和的个数
//在当前轮k，即第k个骰子，总和为数字i的个数，是上一轮k-1中数字和为i-1,i-2...i-6的个数之和
//并且，i-6需要保证>=0，否则也不能参与计算
//非常抽象，举个例子：
//  假设有2个骰子，则数字可能是2-12，且2(1次),3(2次),4(3次)....
//  此时引入第3个骰子，则出现5的次数，应该是第2个骰子出现2,3,4的次数总和，即1+2+3
//说白了就是一个一个骰子的算个数，每一轮都是利用上一轮的计算结果
//难度：5*
func PrintProbability(n int) {
	if n < 1 {
		return
	}
	probablities := [2][]int{}
	probablities[0] = make([]int, DiceMaxNumber * n + 1) //idx==0，弃用
	probablities[1] = make([]int, DiceMaxNumber * n + 1)
	idxFlag := 0

	//如果只有一个骰子，那么出现次数为1
	for i := 1; i <= DiceMaxNumber; i ++ {
		probablities[idxFlag][i] = 1
	}

	//假设有第2个及更多的骰子
	for k := 2; k <= n; k ++ {
		idxNew := 1-idxFlag
		for i:=0; i<k; i++ {
			probablities[idxNew][i] = 0 //这些都不可能出现
		}

		//从数字k开始，一直到k*DiceMaxNumber，是这一轮所有可能出现的次数
		for i:=k; i <= DiceMaxNumber * k; i++ {
			probablities[idxNew][i] = 0 //初始化
			//在本轮，总和为数字i的个数，是上一轮中数字和为i-1,i-2...i-6的个数之后
			//并且，i-j需要保证>=0，否则也不能参与计算（其实更严谨的应该是i-j > k-1)
			for j:=1; j<=i && j <= DiceMaxNumber; j++{
				probablities[idxNew][i] += probablities[idxFlag][i-j]
			}
		}

		idxFlag = idxNew //数组索引来回交替
	}

	//打印所有的概率
	total := math.Pow(DiceMaxNumber, float64(n))
	for i := n; i<=n*DiceMaxNumber; i ++ {
		ratio := float64(probablities[idxFlag][i])*100/total
		fmt.Printf("%v 出现次数：%v, 概率：%2.2f %%\n", i, probablities[idxFlag][i], ratio)
	}
}