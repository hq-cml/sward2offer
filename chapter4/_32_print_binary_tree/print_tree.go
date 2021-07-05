/*
 * 面试题32（一）：不分行从上往下打印二叉树
// 题目：从上往下打印出二叉树的每个结点，同一层的结点按照从左到右的顺序打印。

// 面试题32（二）：分行从上到下打印二叉树
// 题目：从上到下按层打印二叉树，同一层的结点按从左到右的顺序打印，每一层
// 打印到一行。

// 面试题32（三）：之字形打印二叉树
// 题目：请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺
// 序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，
// 其他行以此类推。
 */

package _32_print_binary_tree

import (
    "container/list"
    "fmt"
    "github.com/hq-cml/sward2offer/common"
    "reflect"
)

//把二叉树打印成一个序列，每行从左到右
//引入队列，活用循环
func PrintTreeIn1Line(root *common.TreeNode) error {
    if root == nil {
        return nil
    }

    mylist := list.New()
    mylist.PushBack(root)
    for mylist.Len() != 0 {
        e := mylist.Front()
        mylist.Remove(e)
        node, ok := e.Value.(*common.TreeNode)
        if !ok {
            return fmt.Errorf("Wrong type:%v", reflect.TypeOf(e.Value))
        }
        if node == nil {
            continue
        }
        fmt.Print(node.Val, " ")
        mylist.PushBack(node.Left)
        mylist.PushBack(node.Right)
    }
    return nil
}

//分层打印二叉树，每行从左到右
//再上一程序的基础上，增加每行的个数
func PrintTreeInMultiLine(root *common.TreeNode) error {
    if root == nil {
        return nil
    }

    mylist := list.New()
    mylist.PushBack(root)
    currentLevelCnt := 1
    nextLevelCnt := 0
    for mylist.Len() != 0 {
        e := mylist.Front()
        mylist.Remove(e)
        node, ok := e.Value.(*common.TreeNode)
        if !ok {
            return fmt.Errorf("Wrong type:%v", reflect.TypeOf(e.Value))
        }
        if node == nil {
            continue
        }

        if node.Left != nil {
            mylist.PushBack(node.Left)
            nextLevelCnt ++
        }

        if node.Right != nil {
            mylist.PushBack(node.Right)
            nextLevelCnt ++
        }

        fmt.Print(node.Val, " ")
        currentLevelCnt --
        if currentLevelCnt == 0 {
            fmt.Println() //换行
            currentLevelCnt = nextLevelCnt
            nextLevelCnt = 0
        }
    }

    return nil
}

//蛇形线，按行打印二叉树
//根据特点，将队列换成栈，且考虑左右关系
func PrintTreeSnakeLine(root *common.TreeNode) error {
    if root == nil {
        return nil
    }

    mystack1 := common.NewStack(true)
    mystack2 := common.NewStack(true)
    mystack1.Push(root)
    currentLevel := 1
    currentLevelCnt := 1
    nextLevelCnt := 0
    for mystack1.Len() != 0 {
        e, _ := mystack1.Pop()
        node, ok := e.(*common.TreeNode)
        if !ok {
            return fmt.Errorf("Wrong type:%v", reflect.TypeOf(node))
        }
        if node == nil {
            continue
        }

        var node1 *common.TreeNode
        var node2 *common.TreeNode
        if currentLevel % 2 == 1 {
            node1 = node.Right
            node2 = node.Left
        } else {
            node1 = node.Left
            node2 = node.Right
        }

        if node1 != nil {
            mystack2.Push(node1)
            nextLevelCnt ++
        }

        if node2 != nil {
            mystack2.Push(node2)
            nextLevelCnt ++
        }

        fmt.Print(node.Val, " ")
        currentLevelCnt --
        if currentLevelCnt == 0 {
            fmt.Println() //换行
            currentLevelCnt = nextLevelCnt
            nextLevelCnt = 0
            currentLevel ++
            mystack1, mystack2 = mystack2, mystack1 //对换两个栈
        }
    }

    return nil
}