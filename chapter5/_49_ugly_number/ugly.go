package _49_ugly_number

//获得第idx个丑数
//第一个丑数是1，其他的丑数，都是有且仅有2,3,5作为因子的数字
//1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15 ...
//动态规划的思路：
//  首先递归的分析问题，然后反过来自底向上解决问题
//  通常会需要辅助空间，存储中间结果
//难度：5*
func GetUglyNumber(idx int) int {
    if idx <= 0 {
        return 0
    }

    //空间换时间，保留数据的中间状态，加速
    uglyNumbers := make([]int, idx)
    uglyNumbers[0] = 1
    nextIdx := 1

    p2 := 0  //以2为因子，到目前的最大指向
    p3 := 0  //以3为因子，到目前的最大指向
    p5 := 0  //以5为因子，到目前的最大指向

    for nextIdx < idx {
        min := min(uglyNumbers[p2] * 2,
            uglyNumbers[p3] * 3,
            uglyNumbers[p5] * 5)
        uglyNumbers[nextIdx] = min

        //加速探索过程，因为丑数是有序的
        //所以如果*因子仍然小于min，则可以快速递进
        for uglyNumbers[p2] * 2 <= min {
            p2 ++
        }
        for uglyNumbers[p3] * 3 <= min {
            p3 ++
        }
        for uglyNumbers[p5] * 5 <= min {
            p5 ++
        }
        nextIdx ++
    }
    return uglyNumbers[idx-1]
}

func min(a, b, c int) int {
    m := a
    if a > b {
        m = b
    }
    if m < c {
        return m
    }
    return c
}