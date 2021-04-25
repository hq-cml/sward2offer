package _06_revert_print_list

import (
    "fmt"
)

type Node struct {
    I int
    next *Node
}

func (l *Node)insert(i int) *Node {
    if l == nil {
        l = &Node{
            I:    i,
            next: nil,
        }
    } else {
        l = &Node{
            I:    i,
            next: l,
        }
    }
    return l
}

func (l *Node)print() {
    p := l
    for p != nil {
        fmt.Print(p.I, ", ")
        p = p.next
    }
    fmt.Println()
}

func RevertPrintList (l *Node) {
    if l == nil {
        return
    }

    RevertPrintList(l.next)
    fmt.Print(l.I, ", ")
}