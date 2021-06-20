package _59_list_max_value

import "github.com/hq-cml/sward2offer/common"

//参照第30题，实现一个时间复杂度O(1)的MaxStack
type MaxStack struct {
    s1 common.StackInt  //主栈
    s2 common.StackInt  //从栈，辅助栈
}

func NewMaxStack() *MaxStack {
    return &MaxStack{
        s1 :common.NewStackInt(),
        s2 :common.NewStackInt(),
    }
}

func (s *MaxStack) Len() int{
    return s.s1.Len()
}

func (s *MaxStack) Top() (int, bool) {
    return s.s1.Top()
}

func (s *MaxStack) Push(v int) {
    s.s1.Push(v)
    if s.s2.Len() == 0 {
        s.s2.Push(v)
    } else {
        top, _ := s.s2.Top()
        if top < v {
            s.s2.Push(v)
        } else {
            s.s2.Push(top)
        }
    }
}

func (s *MaxStack) Pop() (int, bool) {
    s.s2.Pop()
    return s.s1.Pop()
}

func (s *MaxStack) Max() (int, bool) {
    return s.s2.Top()
}