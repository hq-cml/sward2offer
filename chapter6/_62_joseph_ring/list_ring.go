/*
// 面试题62：圆圈中最后剩下的数字
// 题目：0, 1, …, n-1这n个数字排成一个圆圈，从数字0开始每次从这个圆圈里
// 删除第m个数字。求出这个圆圈里剩下的最后一个数字。
 */
package _62_joseph_ring

import (
    "errors"
    "fmt"
    "github.com/hq-cml/sward2offer/common"
)

//约瑟夫环
//难度：2*
func CalclRing(n, m int) (int, error) {
    if n <= 0 || m <= 0 {
        return -1, errors.New("Invalid err")
    }
    l := common.NewList()
    l = l.PushNode(0)
    head := l
    var node *common.ListNode
    for i:=1; i<n; i++ { //构造环
        node = &common.ListNode{
            Val:  i,
            Next: nil,
        }
        l = l.PushNodePtr(node)
    }
    if node != nil {
        node.Next = head
    } else {
        head.Next = head  //单节点成环
    }

    var pre *common.ListNode
    p := head
    for p.Next != p { //剩下不止一个节点
        for i:=0; i<m-1; i++ {
            pre = p
            p = p.Next
        }
        pre.Next = p.Next
        fmt.Println("Del:", p.Val)
        p = p.Next
    }
    return p.Val, nil
}

//方法2：
//作者分析出一个公式。。。
//这就只能望洋兴叹了，欣赏欣赏了。。。