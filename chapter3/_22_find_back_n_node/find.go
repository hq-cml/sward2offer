package _22_find_back_n_node

import (
    "errors"
    "github.com/hq-cml/sward2offer/common"
)

func Find(root *common.ListNode, k int)(v int, err error) {
    if root == nil {
        return 0, errors.New("Nil root")
    }
    if k <= 0 { //最后一个，是倒数第一个，不是倒数第0个
        return 0, errors.New("Invalid k")
    }

    p1 := root
    p2 := root

    var i int
    for i = 0; i < k-1; i++ {
        if p1 == nil {
            return 0, errors.New("k > len(root)")
        }
        p1 = p1.Next
    }

    if p1 == nil {
        return 0, errors.New("k > len(root)")
    }

    for p1.Next != nil {
        p2 = p2.Next
        p1 = p1.Next
    }

    return p2.Val, nil
}