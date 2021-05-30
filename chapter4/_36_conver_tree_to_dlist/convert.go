package _36_conver_tree_to_dlist

import "github.com/hq-cml/sward2offer/common"

//难度等级: 5*
//返回双向链表的第一个节点地址
//沿用了treeNode结构，Left表示pre，Right表示post
func Convert(root *common.TreeNode) (*common.TreeNode) {
    var nowTail *common.TreeNode
    convert(root, &nowTail)

    head := nowTail
    for head != nil && head.Left != nil {
        head = head.Left
    }
    return head
}

//第二个参数nowTail，即处理好的链表
//本质上这是一个中序遍历递归
//但是，由于指针比较多，所以非常晦涩
func convert(node *common.TreeNode, nowTail **common.TreeNode) {
    if node == nil {
        return
    }

    //递归左子树
    if node.Left != nil {
        convert(node.Left, nowTail)
    }

    //处理左节点善后
    node.Left = *nowTail
    if (*nowTail) != nil {
        (*nowTail).Right = node
    }

    //nowTail后移成为当前节点
    *nowTail = node

    //递归处理右子树
    if node.Right != nil {
        convert(node.Right, nowTail)
    }
}