/*
 * 面试题59（一）：滑动窗口的最大值
 * 题目：给定一个数组和滑动窗口的大小，请找出所有滑动窗口里的最大值。例如，
 * 如果输入数组{2, 3, 4, 2, 6, 2, 5, 1}及滑动窗口的大小3，那么一共存在6个
 * 滑动窗口，它们的最大值分别为{4, 4, 6, 6, 6, 5}，
 *
 * 面试题59（二）：队列的最大值
 * 题目：给定一个数组和滑动窗口的大小，请找出所有滑动窗口里的最大值。例如，
 * 如果输入数组{2, 3, 4, 2, 6, 2, 5, 1}及滑动窗口的大小3，那么一共存在6个
 * 滑动窗口，它们的最大值分别为{4, 4, 6, 6, 6, 5}，
 * PS：这题也是编程之美3.7
 */

package _59_list_max_value

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
        if k < (windowLen-1) { //窗口不满，直接进入
            help.Insert(v)
            continue
        }
        help.Insert(v)
        max, ok := help.Max()
        if ok {
            ret = append(ret, max)
        }
        help.Pop()
    }

    return ret, true
}



//方法2：
//利用自己实现的O(1)复杂度的队列来实现
//作者给的思路太晦涩，我自己实现了一套
func FindWindowMaxSlide(arr []int, windowLen int) ([]int, bool)  {
    if arr == nil || windowLen <= 0 {
        return nil, false
    }
    if len(arr) < windowLen {
        return nil, false
    }
    ret := []int{}
    l := NewMaxListArr()
    for k, v := range arr {
        if k < (windowLen-1) { //窗口不满，直接进入
            l.PushBack(v)
            continue
        }
        l.PushBack(v)
        max, ok := l.Max()
        if ok {
            ret = append(ret, max)
        }
        l.PopFront()
    }
    return ret, true
}


//方法3：
// 本题还可以用维护一个大顶堆的方式来实现！
// 同样可以O(1)复杂度取最大值，但是队列变更的复杂度O(logN)