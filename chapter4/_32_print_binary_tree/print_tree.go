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
        fmt.Print(node.I, " ")
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

        fmt.Print(node.I, " ")
        currentLevelCnt --
        if currentLevelCnt == 0 {
            fmt.Println() //换行
            currentLevelCnt = nextLevelCnt
            nextLevelCnt = 0
        }
    }

    return nil
}