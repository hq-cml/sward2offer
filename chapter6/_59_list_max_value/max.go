package _59_list_max_value

//滑动窗口最大值
import (
    "container/list"
)

//方法1：利用MaxList来实现
//将前面多次的成果，组装起来
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
        if k <= (windowLen-1) {
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

    max, ok := help.Max()
    if ok {
        ret = append(ret, max)
    }
    return ret, true
}

//方法2：
// 利用滑动窗口的特点，结合一个两头开口的队列，来实现
//func FindWindowMaxSlide(arr []int, windowLen int) ([]int, bool)  {
//    if arr == nil || windowLen <= 0 {
//        return nil, false
//    }
//    if len(arr) < windowLen {
//        return nil, false
//    }
//
//    l := list.New()
//    for i:=0; i<windowLen; i++ {
//        for l.Len() > 0 && !compareTail(l, arr[i]){
//            tail := l.Back()
//            l.Remove(tail)
//        }
//        l.PushBack(arr[i])
//    }
//
//    for i:=windowLen; i<len(arr); i++ {
//
//    }
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