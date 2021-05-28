package _30_stack_min

import "github.com/hq-cml/sward2offer/common"

type MinStack struct {
	s1 common.StackInt  //主栈
	s2 common.StackInt  //从栈，辅助栈
}

func NewMinStack() MinStack {
	return MinStack{
		s1 :common.NewStackInt(),
		s2 :common.NewStackInt(),
	}
}

func (s *MinStack) Len() int{
	return s.s1.Len()
}

func (s *MinStack) Top() (int, bool) {
	return s.s1.Top()
}

func (s *MinStack) Push(v int) {
	s.s1.Push(v)
	if s.s2.Len() == 0 {
		s.s2.Push(v)
	} else {
		top, _ := s.s2.Top()
		if top > v {
			s.s2.Push(v)
		} else {
			s.s2.Push(top)
		}
	}
}

func (s *MinStack) Pop() (int, bool) {
	s.s2.Pop()
	return s.s1.Pop()
}

func (s *MinStack) Min() (int, bool) {
	return s.s2.Top()
}