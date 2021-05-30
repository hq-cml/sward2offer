package _31_stack_order

import "github.com/hq-cml/sward2offer/common"

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
        if ok && t == b[idx2]{
            stack.Pop()
            idx2 ++
            continue
        }

        if a[idx1] == b[idx2] {
            idx1 ++
            idx2 ++
        } else {
            stack.Push(a[idx1])
            idx1 ++
        }
    }
    if idx2 >= len(b) {
        return true
    }

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