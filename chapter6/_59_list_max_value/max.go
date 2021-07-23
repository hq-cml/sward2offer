/*
 * 面试题59（一）：滑动窗口的最大值
 * 题目：给定一个数组和滑动窗口的大小，请找出所有滑动窗口里的最大值。例如，
 * 如果输入数组{2, 3, 4, 2, 6, 2, 5, 1}及滑动窗口的大小3，那么一共存在6个
 * 滑动窗口，它们的最大值分别为{4, 4, 6, 6, 6, 5}，
 *
 * 面试题59（二）：队列的最大值 //TODO 题目一模一样
 * 题目：给定一个数组和滑动窗口的大小，请找出所有滑动窗口里的最大值。例如，
 * 如果输入数组{2, 3, 4, 2, 6, 2, 5, 1}及滑动窗口的大小3，那么一共存在6个
 * 滑动窗口，它们的最大值分别为{4, 4, 6, 6, 6, 5}，
 */

package _59_list_max_value

//滑动窗口最大值
import (
    "container/list"
)

//方法1：
//思路：主要是需要一种能够用O(1)复杂度获取最大值的容器
//利用MaxList来实现：结合之前的O(1)复杂度求栈最大值 & 两个栈模拟队列
//将前面多次的成果，组装起来，在stack.go和list.go中
//难度：3*
func FindWindowMax(arr []int, windowLen int) ([]int, bool)  {
    if arr == nil {
        return nil, false
    }
    if len(arr) < windowLen {
        return nil, false
    }

    help := NewMaxList()

    ret := []int{}
    for k, v := range arr {
        if k <= (windowLen-1) { //窗口不满，直接进入
            help.Insert(v)
            continue
        }
        max, ok := help.Max()
        if ok {
            ret = append(ret, max)
        }
        help.Pop()
        help.Insert(v)
    }

    //最后一个元素入队列之后，窗口最大值还未识别，这里补上
    max, ok := help.Max()
    if ok {
        ret = append(ret, max)
    }
    return ret, true
}

//方法2：//TODO
//思路：参考O(1)复杂度获取Max的Stack的实现
//设置两个队列，一个主队列，一个辅助队列，关键在辅助队列
//如果新增的元素，大于辅助队列的尾巴则不断清空尾巴，直到不满足之后，入辅助队列尾巴
//如果删除的元素，正好是辅助队列的头，则辅助队列和主队列都清空头。
//我先实现了第一种。
// 利用滑动窗口的特点，结合一个两头开口的队列，来实现
//func FindWindowMaxSlide(arr []int, windowLen int) ([]int, bool)  {
//   if arr == nil || windowLen <= 0 {
//       return nil, false
//   }
//   if len(arr) < windowLen {
//       return nil, false
//   }
//
//   l := list.New()
//   for i:=0; i<windowLen; i++ {
//       for l.Len() > 0 && !compareTail(l, arr[i]){
//           tail := l.Back()
//           l.Remove(tail)
//       }
//       l.PushBack(arr[i])
//   }
//
//   for i:=windowLen; i<len(arr); i++ {
//
//   }
//
//}

func compareTail(l *list.List, v int) bool {
    tail := l.Back()
    last := tail.Value.(int)
    if last > v {
        return true
    } else {
        return false
    }
}