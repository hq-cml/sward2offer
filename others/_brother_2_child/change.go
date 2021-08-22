/*
 * by面试题：二叉树的兄弟孩子表示法
 * 方式：二叉树的左孩子保留，右孩子作为左孩子的右孩子
 *
 * 处理前：
 *        1               前序： 1 2 4 5 3 6 7
 *     2     3            中：   4 2 5 1 6 3 7
 *   4   5  6  7          后：   4 5 2 6 7 3 1
 *
 * 处理后：
 *           1            前序： 1 2 4 5 3 6 7
 *         2              中：   4 5 2 6 7 3 1
 *       4   3            后：   5 4 7 6 3 2 1
 *        5   6
 *           7
 */
package _brother_2_child

import "github.com/hq-cml/sward2offer/common"

//递归方式
//难度：3*
func Change2ChildRecurse(root *common.TreeNode) {
    if root == nil {
        return
    }

    if root.Left != nil {
        //左子树非空，先递归处理左子树，否则断了
        Change2ChildRecurse(root.Left)
        //左子树处理完毕，将右子树赋值给左子树的右孩子
        root.Left.Right = root.Right
    } else {
        //左子空，直接将将右子树赋值给左子树的右孩子
        root.Left = root.Right
    }

    //右子树需要同样的递归处理
    Change2ChildRecurse(root.Right)

    //都处理完毕之后，根的右子树已经没用了，置空
    root.Right = nil
}

//非递归方式
//借助一个栈，模拟递归的过程
//这个比较烧脑，牛逼
//难度：4*
func Change2ChildNoRecursion(root *common.TreeNode) {
    myStack := common.NewStack(false)
    p := root //p是精髓

    for p != nil || myStack.Len()!=0 {
        //一路捅到最左边，先入栈
        for p != nil {
            myStack.Push(p)
            p = p.Left
        }

        //从最基层开始处理
        tmpNode, _ := myStack.Top() //只是获取栈顶元素，不出栈
        tmp, _ := tmpNode.(*common.TreeNode)
        if tmp.Right == nil {
            myStack.Pop() //扔出元素
        } else {
            p = tmp.Right
            if tmp.Left == nil {
                tmp.Left = p
            } else {
                tmp.Left.Right = p
            }
            tmp.Right = nil   //Don't forget to set right to NULL
        }
    }
}