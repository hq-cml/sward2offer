/*
 * 面试题30：包含min函数的栈
 * 题目：定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的min
 * 函数。在该栈中，调用min、push及pop的时间复杂度都是O(1)。

引入辅助栈
 */
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