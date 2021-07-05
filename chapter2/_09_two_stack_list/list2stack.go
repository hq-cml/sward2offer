/*
 * 两个队列，实现一个栈
 */
package _09_two_stack_list

import (
    "container/list"
)

type Stack struct {
    l1 list.List
    l2 list.List
}

//思路：
//始终保持一个队列是空的，
//插入操作：直插入非空队列
//删除操作：将非空的队列前面的值，逐个删除插入另一个空队列，只最后留一个删除，作为实际删除
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

    //吐出n-1个元素，插入空队列，留下最后一个删除
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