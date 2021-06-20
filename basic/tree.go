package basic

import "github.com/hq-cml/sward2offer/common"

//求一课二叉树的指定节点的路径
//假设所有的节点值不重复
//这是先序遍历递归算法的一个应用
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