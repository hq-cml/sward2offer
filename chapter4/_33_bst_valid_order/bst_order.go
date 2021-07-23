/*
 * 面试题33：二叉搜索树的后序遍历序列
 * 题目：输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。
 * 如果是则返回true，否则返回false。假设输入的数组的任意两个数字都互不相同。
 */

//二叉搜索树的概念：
//一个递归的概念，若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值；
//若它的右子树不空，则右子树上所有结点的值均大于它的根结点的值；
//它的左、右子树也分别为二叉排序树。
package _33_bst_valid_order

//思路：
//二叉搜索树本身就是递归的概念，所以这里也递归的去处理
func CheckValidOrder(order []int) bool {
    //如果空树或者仅一个节点，认为合法
    if len(order) <= 1 {
        return true
    }

    //二叉搜索树的root，必然是最后一个节点
    rootIdx := len(order) - 1
    rootVal := order[rootIdx]

    //找到左右节点分界线
    idx := 0
    for ; idx < len(order)-1; idx ++ {
        if order[idx] < rootVal {
            continue
        }
        break
    }

    //检查分界线之后，是否存在违背定义的节点
    for i:=idx; i<len(order); i++ {
        if order[i] < rootVal {
            return false
        }
    }

    //递归校验左右，牛逼
    return CheckValidOrder(order[0:idx]) &&
        CheckValidOrder(order[idx: rootIdx])
}