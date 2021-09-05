package common

import (
	"container/list"
)

type Stack interface {
	Len() int
	Push(value interface{})
	Top() (interface{}, bool)
	Pop() (interface{}, bool)
}

//self标识是否自行实现栈
func NewStack(self bool) Stack {
	if !self {
		return &StackListImpl{}
	} else {
		return &StackSelfImpl{
			s: []interface{}{},
		}
	}
}

// 利用container/list实现
type StackListImpl struct {
	l list.List //开箱即用
}

func (s *StackListImpl) Len() int {
	return s.l.Len()
}

func (s *StackListImpl) Push(v interface{}) {
	s.l.PushFront(v)
}

func (s *StackListImpl) Top() (interface{}, bool) {
	if s.l.Len() == 0 {
		return nil, false
	}
	return s.l.Front().Value, true
}

func (s *StackListImpl) Pop() (interface{}, bool) {
	if s.l.Len() == 0 {
		return nil, false
	}
	front := s.l.Front()
	s.l.Remove(front)
	return front.Value, true
}

//自己实现
type StackSelfImpl struct {
	s []interface{}
}

func (stack *StackSelfImpl) Len() int {
	return len(stack.s)
}

func (stack *StackSelfImpl) Push(value interface{}) {
	stack.s = append(stack.s, value)
}

//栈顶元素，但是栈不会变化
func (stack *StackSelfImpl) Top() (interface{}, bool) {
	if len(stack.s) == 0 {
		return nil, false
	}
	return stack.s[len(stack.s)-1], true
}

func (stack *StackSelfImpl) Pop() (interface{}, bool) {
	if len(stack.s) == 0 {
		return nil, false
	}
	value := stack.s[len(stack.s)-1]
	stack.s = stack.s[:len(stack.s)-1]
	return value, true
}
