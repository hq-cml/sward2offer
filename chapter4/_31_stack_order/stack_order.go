/*
 * 面试题31：栈的压入、弹出序列
 * 题目：输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是
 * 否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如序列1、2、3、4、5
 * 是某栈的压栈序列，序列4、5、3、2、1是该压栈序列对应的一个弹出序列，但
 * 4、3、5、1、2就不可能是该压栈序列的弹出序列。
 */
package _31_stack_order

import "github.com/hq-cml/sward2offer/common"

//思路：
//引入辅助栈，模拟操作
//遍历两个序列：不一致则入栈，序列1后移，一致则出栈，出现序列完毕栈仍然非空，则表示非法
func CheckValidSequence(a []int, b []int) bool {
    if len(a) != len(b) {
        return false
    }

    if len(a) == 0 {
        return true
    }

    stack := common.NewStackInt()

    idx1 := 0
    idx2 := 0

    for idx1 < len(a) && idx2 < len(b){
        t, ok := stack.Top()
        //如果栈非空，且栈顶元素和出栈序列首元素相当，则出栈且序列后移
        if ok && t == b[idx2]{
            stack.Pop()
            idx2 ++
            continue
        }

        //如果相等，相当于入栈再出栈，直接后移
        if a[idx1] == b[idx2] {
            idx1 ++
            idx2 ++
        } else {
            stack.Push(a[idx1])
            idx1 ++
        }
    }

    //通常都会是idx1结束，如果idx2也结束了，通常是两个序列完全一模一样
    if idx2 >= len(b) {
        return true
    }

    //对于剩余的栈元素和出栈序列进行最终判断
    for idx2 < len(b) {
        t, ok := stack.Top()
        if ok && t == b[idx2]{
            stack.Pop()
            idx2 ++
        } else {
            return false
        }
    }

    return true
}