package common

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func NewList() *ListNode {
    return nil
}

func (l *ListNode)Empty() bool {
    return l == nil
}

func (l *ListNode)Len() int {
    length := 0
    if l.Empty() {
        return length
    }
    for l != nil {
        l = l.Next
        length ++
    }
    return length
}

//入队列，队尾
func (l *ListNode)PushNode(v int) *ListNode {
    node := &ListNode{
        Val:  v,
        Next: nil,
    }
    if l == nil {
        return node
    }
    p := l
    for p.Next != nil {
        p = p.Next
    }
    p.Next = node

    return l
}

//出队列，队头
func (l *ListNode)PopNode() (int, *ListNode, bool) {
    if l == nil {
        return 0, nil, false
    }

    return l.Val, l.Next, true
}

//遍历
func (l *ListNode)Foreach(f func(*ListNode)) {
    p := l
    for p != nil {
        f(p)
        p = p.Next
    }
}

func NodePrint(node *ListNode) {
    if node == nil {
        return
    }
    if node.Next == nil {
        fmt.Println(node.Val)
    } else {
        fmt.Print(node.Val, " ")
    }
}


func (l *ListNode)PushNodePtr(node *ListNode) *ListNode {
    if l == nil {
        return node
    }
    p := l
    for p.Next != nil {
        p = p.Next
    }
    p.Next = node

    return l
}