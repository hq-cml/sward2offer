package _54_k_th_Node
/*
 * 二叉搜索树的第K大节点
 */
import "github.com/hq-cml/sward2offer/common"

//二叉搜索树
//特性：中序遍历就是顺序遍历
//难度：2*
func KthNode(root *common.TreeNode, k int) *common.TreeNode {
    if root == nil || k < 0 {
        return nil
    }

    return kthNode(root, &k)
}

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