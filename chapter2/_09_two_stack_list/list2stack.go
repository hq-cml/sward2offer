package _09_two_stack_list

import (
    "container/list"
)

type Stack struct {
    l1 list.List
    l2 list.List
}

// 永远向非空的栈压入数据，如果均为空，则任意
func (s *Stack)Push(v interface{}) {
    if s.l1.Len() > 0 {
        s.l1.PushBack(v)
    } else {
        s.l2.PushBack(v)
    }
}

func (s *Stack)Pop()(v interface{}, exist bool) {
    if s.l1.Len() == 0 && s.l2.Len() == 0 {
        return nil, false
    }

    var src *list.List
    var dst *list.List
    if s.l1.Len() > 0 {
        src, dst = &s.l1, &s.l2
    } else {
        src, dst = &s.l2, &s.l1
    }

    n := src.Len()
    for n > 1 {
        head := src.Front()
        dst.PushBack(head.Value)
        src.Remove(head)
        n --
    }
    head := src.Front()
    return src.Remove(head), true
}