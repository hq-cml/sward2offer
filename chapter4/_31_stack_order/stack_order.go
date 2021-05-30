package _31_stack_order

import "github.com/hq-cml/sward2offer/common"

/*
 *
 */
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