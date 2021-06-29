/*
 // 面试题5：替换空格
// 题目：请实现一个函数，把字符串中的每个空格替换成"%20"。例如输入“We are happy.”，
// 则输出“We%20are%20happy.”。

这是一个简单的urlencode，看到这种题目，需要考虑的问题是C语言的字符串特性，末尾是'\0'，这个题目有两种考虑，如果不考虑空间问题，那么直接再申请一个字符串的空间，逐个拷贝遇到空格则替换是最简单的。但是往往这个题肯定要求原地址替换，那么就存在技巧性了，如果重头开始，则需要考虑替换过程中的覆盖问题，而且效率低，时间复杂度是O(N平方)，所以考虑倒着从尾巴替换。
先计算出空格的数量，则新字符串是的长度是原长度+空格数*2，然后设计两个指针，从尾巴上进行替换，时间复杂度是O(N)。

 */
package _05_replace_blank

// blank => %20
func ReplaceBlank(s []byte) ([]byte, error) {
    if s == nil {
        return s, nil
    }

    cnt := 0
    for _, v := range s {
        if string(v) == " " {
            cnt ++
        }
    }
    if cnt == 0 {
        return s, nil
    }

    add := make([]byte, cnt * 2)
    p1 := len(s) - 1
    s = append(s, add ...) // 续上空白的长度
    p2 := len(s) - 1

    for p1 < p2 {
        if string(s[p1]) != " " {
            s[p2] = s[p1]
            p2 --
        } else {
            s[p2-2], s[p2-1], s[p2] = byte('%'), byte('2'), byte('0')
            p2 -= 3
        }
        p1 --
    }
    return s, nil
}