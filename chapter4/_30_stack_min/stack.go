/*
 * 面试题30：包含min函数的栈
 * 题目：定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的min
 * 函数。在该栈中，调用min、push及pop的时间复杂度都是O(1)。
 */
package _30_stack_min

import "github.com/hq-cml/sward2offer/common"

type MinStack struct {
	s1 common.StackInt  //主栈
	s2 common.StackInt  //辅助栈
}

//思路：
//引入辅助栈，永远保存当前最小元素
//方案1：入栈的时候进行检测，如果入栈的元素比富辅助栈top还小，则辅助栈入栈该元素，否则入栈当前top元素；出栈就是主栈和辅助栈无脑出栈即可。简单但浪费空间
//方案2：入栈的时候进行检测，如果入栈的元素<=辅助栈top，则辅助栈入栈该元素；出栈就是主栈出栈，如果出栈元素==辅助栈top，则辅助栈出栈。省空间
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