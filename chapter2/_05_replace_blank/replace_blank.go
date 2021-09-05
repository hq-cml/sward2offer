/*
 * 面试题5：替换空格
 * 题目：请实现一个函数，把字符串中的每个空格替换成"%20"。例如输入“We are happy.”，
 * 则输出“We%20are%20happy.”。
 */
package _05_replace_blank

//这是一个简单的urlencode，有两种考虑:
//如果不考虑空间问题，那么直接再申请一个字符串的空间，逐个拷贝遇到空格则替换是最简单的。
//如果要求原地址替换，需要不断遇到空格后移出空间，再替换，效率低，时间复杂度是O(N^2)，所以考虑倒着从尾巴替换。
//先计算出空格的数量，则新字符串是的长度是原长度+空格数*2，然后设计两个指针，从尾巴上进行替换，时间复杂度是O(N)。
func ReplaceBlank(s []byte) ([]byte, error) {
    if s == nil {
        return s, nil
    }

    //先统计空格的数量
    cnt := 0
    for _, v := range s {
        if string(v) == " " {
            cnt ++
        }
    }
    if cnt == 0 {
        return s, nil
    }

    //将需要延长的空间，在原来数组后面续上
    //每个空格替换后延长2个字符
    add := make([]byte, cnt * 2)
    p1 := len(s) - 1
    s = append(s, add ...)  //原理上来说，这里有可能也不是原址替换了，这里只是说明这个意思
    p2 := len(s) - 1

    //从尾部开始替换
    for p1 < p2 {
        if s[p1] != ' ' {
            s[p2] = s[p1]
            p2 --
        } else {
            s[p2-2], s[p2-1], s[p2] = '%', '2', '0'
            p2 -= 3
        }
        p1 --
    }
    return s, nil
}