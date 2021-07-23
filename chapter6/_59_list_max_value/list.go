/*
 * 实现Max为O(1)复杂度的List
 */
package _59_list_max_value

import "fmt"

//O(1)复杂度实现一个List，支持Max操作
//方法1：
//利用Maxstack结构来实现，即两个栈实现一个队列
type MaxList struct {
    stack1 *MaxStack
    stack2 *MaxStack
}
func NewMaxList() *MaxList {
    return &MaxList{
        stack1: NewMaxStack(),
        stack2: NewMaxStack(),
    }
}

//对尾插入
func (ml *MaxList) Insert(value int)  {
    fmt.Println("Insert", value)
    ml.stack1.Push(value)
}

//对头弹出
func (ml *MaxList) Pop() (int, bool) {
    if ml.stack2.Len() > 0 {
        return ml.stack2.Pop()
    }
    if  ml.stack1.Len() > 0 {
        for ml.stack1.Len() > 0 {
            v, _ := ml.stack1.Pop()
            fmt.Println("Swap", v)
            ml.stack2.Push(v)
        }
        return ml.stack2.Pop()
    }
    return 0, false
}

func (ml *MaxList) Max() (int, bool){
    mx1, ok1 := ml.stack1.Max()
    mx2, ok2 := ml.stack2.Max()

    //两个栈均为空
    if !ok1 && !ok2 {
        return 0, false
    }

    //两个栈均非空
    if ok1 && ok2 {
        if mx1 > mx2 {
            return mx1, true
        } else {
            return mx2, true
        }
    }

    //有一个为空
    if ok1 {
        return mx1, true
    } else {
        return mx2, true
    }
}
