package _52_list_first_common_node

import "github.com/hq-cml/sward2offer/common"

//链表的第一个公共节点
//思路一：蛮力朴素，时间复杂度O(m*n)，显然不符合要求
//思路二：因为公共节点，在最后肯定是汇合一处，希望能够到推
//       利用两个辅助栈，可以实现倒推的效果，两个队列遍历各自压入两个辅助栈，
//       然后开始分别从栈顶弹出，最后一个相同的节点，就是汇聚节点
//思路三：各自遍历，求出长度差，长的的先遍历差数个节点，然后一起向后遍历，直到相同为止
func FirstCommonNode(l1, l2 *common.ListNode) *common.ListNode {
    if l1 == nil || l2 == nil {
        return nil
    }
    len1 := l1.Len()
    len2 := l2.Len()
    var long *common.ListNode
    var short *common.ListNode
    var diff int
    if len1 > len2 {
        long, short = l1, l2
        diff = len1 - len2
    } else {
        long, short = l2, l1
        diff = len2 - len1
    }

    for diff > 0 {
        long = long.Next
        diff --
    }

    for long != short {
        long = long.Next
        short = short.Next
    }

    return long
}