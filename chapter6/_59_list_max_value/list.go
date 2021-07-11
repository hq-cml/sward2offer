/*
// 面试题59（二）：队列的最大值
// 题目：给定一个数组和滑动窗口的大小，请找出所有滑动窗口里的最大值。例如，
// 如果输入数组{2, 3, 4, 2, 6, 2, 5, 1}及滑动窗口的大小3，那么一共存在6个
// 滑动窗口，它们的最大值分别为{4, 4, 6, 6, 6, 5}，


这道题目本质上，是实现一个O(1)复杂度求最大值的队列。作者给出两个思路：
第一个是参考之前的题目，结合O(1)复杂对求栈最大值 & 两个栈模拟队列
第二个是一个新的思路：
两个队列，一个主队列，一个辅助队列，关键在辅助队列
如果新增的元素，大于辅助队列的尾巴则不断清空尾巴，直到不满足之后，入辅助队列尾巴
如果删除的元素，正好是辅助队列的头，则辅助队列和主队列都清空头。
我先实现了第一种。
 */

package _59_list_max_value

import "fmt"

//O(1)复杂度实现一个List，支持Max操作

//方法1：利用Maxstack结构来实现
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

func (ml *MaxList) Insert(value int)  {
    fmt.Println("Insert", value)
    ml.stack1.Push(value)
}

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

    if !ok1 && !ok2 {
        return 0, false
    }

    if ok1 && ok2 {
        if mx1 > mx2 {
            return mx1, true
        } else {
            return mx2, true
        }
    }

    if ok1 {
        return mx1, true
    } else {
        return mx2, true
    }
}
