/*
 * 面试题9：用两个栈实现队列
 * 题目：用两个栈实现一个队列。队列的声明如下，请实现它的两个函数appendTail
 * 和deleteHead，分别完成在队列尾部插入结点和在队列头部删除结点的功能。
 */
package _09_two_stack_list

import (
	"github.com/hq-cml/sward2offer/common"
)

type List struct {
	s1 common.Stack
	s2 common.Stack
}

//常规思路
func NewList() *List {
	return &List{
		s1: common.NewStack(false),
		s2: common.NewStack(false),
	}
}

func (l *List)AppendTail(v interface{}) {
	l.s1.Push(v)
}

func (l *List)DeleteHead() (v interface{}, exist bool) {
	if l.s2.Len() > 0 {
		return l.s2.Pop()
	}

	if l.s1.Len() > 0 {
		for l.s1.Len() > 0 {
			v, _ := l.s1.Pop()
			l.s2.Push(v)
		}

		return l.s2.Pop()
	}
	return nil, false
}