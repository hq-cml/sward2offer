package _54_k_th_Node
/*
 * 面试题54：二叉搜索树的第k个结点
 * 题目：给定一棵二叉搜索树，请找出其中的第k大的结点。
 */
import "github.com/hq-cml/sward2offer/common"

//思路：
// 二叉搜索树特性：中序遍历就是顺序遍历
//难度：2*
func KthNode(root *common.TreeNode, k int) *common.TreeNode {
    if root == nil || k < 0 {
        return nil
    }

    return kthNode(root, &k)
}

//中序遍历，灵魂在于这个计数器pk
func kthNode(root *common.TreeNode, pk *int) *common.TreeNode {
    if root == nil {
        return nil
    }

    if left := kthNode(root.Left, pk); left != nil {
        return left
    }

    *pk --
    if *pk == 0 {
        return root
    }

    return kthNode(root.Right, pk)
}