package _58_reverse_string

//题目1：翻转每个单词
//思路比较常见，就是两次翻转
//难度：2*
func ReverseWord(str string) string {
    if len(str) == 0 {
        return str
    }
    //整体翻转
    b := []byte(str)
    reverse(b, 0, len(b)-1)

    //逐个单词翻转
    start := 0
    end := 0
    for end < len(b) {
        if b[end] == ' ' {
            reverse(b, start, end-1)
            start = end+1
            end = end+1
            continue
        } else {
            end++
        }
    }

    return string(b)
}

//题目2：坐旋转字符串
//思路1：我个人的思路，拼接字符串，然后直接后移。这种方式需要消耗O(n)空间
//思路2：仍然利用翻转思路，三次：先翻转两个子部分，然后整体翻转
//难度：3*
func SwapLeft(str string, n int) string {
    if len(str) < n {
        return str
    }
    //先翻转局部
    b := []byte(str)
    reverse(b, 0, n-1)
    reverse(b, n, len(b)-1)

    //整体翻转
    reverse(b, 0, len(b)-1)

    return string(b)
}

func reverse(str []byte, beg, end int) {
    for beg < end {
        str[beg], str[end] = str[end], str[beg]
        beg ++
        end --
    }
}