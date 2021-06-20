package _60_n_dice_sum

import (
	"fmt"
	"math"
)

const (
	DiceMaxNumber = 6
)

//极其烧脑，利用两个数组，不断替换，来算出每一轮的各个可能总和的个数
//在当前本轮，总和为数字i的个数，是上一轮中数字和为i-1,i-2...i-6的个数之后
//并且，i-6需要保证>=0，否则也不能参与计算
//难度：5*
func PrintProbability(n int) {
	if n < 1 {
		return
	}
	probablities := [2][]int {}
	probablities[0] = make([]int, DiceMaxNumber * n + 1)
	probablities[1] = make([]int, DiceMaxNumber * n + 1)
	idxFlag := 0

	//只有一个骰子，出现次数为1
	for i := 1; i <= DiceMaxNumber; i ++ {
		probablities[idxFlag][i] = 1
	}

	//2个及更多的骰子
	for k := 2; k <= n; k ++ {
		idxNew := 1-idxFlag
		for i:=0; i<k; i++ {
			probablities[idxNew][i] = 0 //这些都不可能出现
		}

		//从数字k开始，一直到k*DiceMaxNumber，是这一轮所有可能出现的次数
		for i:=k; i <= DiceMaxNumber * k; i++ {
			probablities[idxNew][i] = 0 //初始化
			//在本轮，总和为数字i的个数，是上一轮中数字和为i-1,i-2...i-6的个数之后
			//并且，i-6需要保证>=0，否则也不能参与计算
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
		fmt.Printf("%v ratio is: %v\n", i, ratio)
	}
}