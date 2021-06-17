package _56_get_number_cnt

import "math"

//题目1：数组中只出现一次的2个数字
//要求时间复杂度是O(n)，说明只能循环常数次
//空间复杂度O(1)，不能用map，否则用map统计次数是最快的
//利用异或，但是需要将整个数组分成两部分，各包含出现1次的数字
//先整体异或，求出一个值，那么这个值当中某一个为1的位，一定就是这两个唯一1次的数不一样的地方
//按这一位来区分数组，即可分割成两个数组，分别异或求得最后结果
//难度：3*
func FindNumerOnlyTwice(arr []int) (int, int){
    if len(arr) < 2 {
        return 0, 0
    }

    //首先异或一次
    t := 0
    for _, v := range arr {
        t = t^v
    }

    //找到二进制中第一个非0位
    n := 1
    for t & n ==0 {
        n = n << 1
    }

    //将数据分成两组，分别循环异或
    t1 := 0
    t2 := 0
    for _, v := range arr {
        if n & v == n {
            t1 = t1^v
        } else {
            t2 = t2^v
        }
    }

    return t1, t2
}

//题目2：数组中只出现一次的1个数字，其他都是3次
//最简的方案，还是利用map，来做一次遍历然后计数
//一个非常奇特的思路，将每一位分别相加，则个数不是3的倍数就是3的倍数余1（因为其他都是3次出现）
//难度：4*
func FindNumerOnlyOnce(arr []int) (int){
    //假设32位数字，统计每一位出现的次数
    bitMap := make([]int, 32)
    for _, v := range arr {
        n := 1
        for i:=0; i<32; i++ {
            if v & n == n { //第n位是1
                bitMap[i] ++
            }
            n = n<<1
        }
    }

    //针对每一位的统计，计算最终的值
    t := 0
    for k, v := range bitMap {
        if v % 3 == 1 {
            t += int(math.Pow(2, float64(k)))
        }
    }
    return t
}