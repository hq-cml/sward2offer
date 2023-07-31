package basic

import "github.com/hq-cml/sward2offer/common"

//给定一颗二叉树root，再给定一个值，要求在二叉树中查找，如果能够找到节点值和给定值相同，给出路径
//假设所有的节点值不重复
//本质上：这是先序遍历递归算法的一个应用
//注意：这里的实现中，path没有回退机制，其实这是利用了一个golang的slice特性
//interview
func FindPath(root *common.TreeNode, need int, path []int) ([]int, bool) {
    if root == nil {
        return nil, false
    }

    data := root.Val
    path = append(path, data)
    //找到
    if data == need {
        return path, true
    }

    ret, ok := FindPath(root.Left, need, path)
    if ok {
        return ret, true
    }
    ret, ok = FindPath(root.Right, need, path)
    if ok {
        return ret, true
    }

    return nil, false
}