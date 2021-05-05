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

type StackImpl struct {
	l list.List    //开箱即用
}

func NewStack() Stack {
	return &StackImpl{
	}
}

func (s *StackImpl) Len() int {
	return s.l.Len()
}

func (s *StackImpl) Push(v interface{}) {
	s.l.PushFront(v)
}

func (s *StackImpl) Top() (interface{}, bool) {
	if s.l.Len() == 0 {
		return nil, false
	}
	return s.l.Front().Value, true
}

func (s *StackImpl) Pop() (interface{}, bool) {
	if s.l.Len() == 0 {
		return nil, false
	}
	front := s.l.Front()
	s.l.Remove(front)
	return front.Value, true
}