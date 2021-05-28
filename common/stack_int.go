package common

import (
	"container/list"
)

type StackInt interface {
	Len() int
	Push(int)
	Top() (int, bool)
	Pop() (int, bool)
}

type StackIntImpl struct {
	l list.List    //开箱即用
}

func NewStackInt() StackInt {
	return &StackIntImpl{
	}
}

func (s *StackIntImpl) Len() int {
	return s.l.Len()
}

func (s *StackIntImpl) Push(v int) {
	s.l.PushFront(v)
}

func (s *StackIntImpl) Top() (int, bool) {
	if s.l.Len() == 0 {
		return 0, false
	}
	i, ok := s.l.Front().Value.(int)
	if !ok {
		return 0, false
	}
	return i, true
}

func (s *StackIntImpl) Pop() (int, bool) {
	if s.l.Len() == 0 {
		return 0, false
	}
	front := s.l.Front()
	s.l.Remove(front)

	i, ok := front.Value.(int)
	if !ok {
		return 0, false
	}
	return i, true
}