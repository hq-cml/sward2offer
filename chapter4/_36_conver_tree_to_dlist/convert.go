/*
 * 面试题36：二叉搜索树与双向链表
 * 题目：输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。要求
 * 不能创建任何新的结点，只能调整树中结点指针的指向。
 */

package _36_conver_tree_to_dlist

import "github.com/hq-cml/sward2offer/common"

//思路：
//根据二叉搜索的树的定义，由于这个特性，那么二叉树的中序遍历，其实就是链表的序列
//关于树的题目，首先先搞就是想到递归的思路
//返回双向链表的第一个节点地址
//沿用了treeNode结构，Left表示pre，Right表示post
//难度等级: 5*
func Convert(root *common.TreeNode) (*common.TreeNode) {
    var nowTail *common.TreeNode
    convert(root, &nowTail)

    //此时nowTail只想了队列尾部，需要不断向前移动到头部
    head := nowTail
    for head != nil && head.Left != nil {
        head = head.Left
    }
    return head
}

//第二个参数nowTail，即处理好的链表
//本质上这是一个中序遍历递归
//但是，由于指针比较多，所以非常晦涩
func convert(root *common.TreeNode, nowTail **common.TreeNode) {
    if root == nil {
        return
    }

    //递归左子树
    if root.Left != nil {
        convert(root.Left, nowTail)
    }

    //处理左节点善后
    root.Left = *nowTail
    if (*nowTail) != nil {
        (*nowTail).Right = root
    }

    //nowTail后移成为当前节点
    *nowTail = root

    //递归处理右子树
    if root.Right != nil {
        convert(root.Right, nowTail)
    }
}