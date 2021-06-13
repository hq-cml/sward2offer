package _50_first_once_char

//第一个只出现一次的字符
//最核心的思路：空间换时间！
func FindFirst(str string) byte {
    m := make([]int, 256)
    for i, _ := range m {
        m[i] = -1
    }

    for pos, c := range str {
        if m[c] == -1 { //第一次出现，存储其位置
            m[c] = pos
        } else if m[c] >= 0 {
            m[c] = -2   //已经出现过一次了，则c不再有效
        } else {
            //已经出现多次，则什么也不敢
        }
    }

    //扫描出只出现1次的，且最先出现的
    minIdx := -1
    var firstChar byte
    for c, idx := range m {
        if idx < 0 {   //idx==-1或者-2，说明未出现或者出现过多次
            continue
        }
        if minIdx == -1 || idx < minIdx {
            minIdx = idx
            firstChar = byte(c)
        }
    }
    return firstChar
}

//书上还有题目二，计算流，思路和第一题是一致的，只是换成面向对象写法