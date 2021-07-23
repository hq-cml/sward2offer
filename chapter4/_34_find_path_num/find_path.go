/*
 * 面试题34：二叉树中和为某一值的路径
 * 题目：输入一棵二叉树和一个整数，打印出二叉树中结点值的和为输入整数的所
 * 有路径。从树的根结点开始往下一直到叶结点所经过的结点形成一条路径。
 */
package _34_find_path_num

import (
    "fmt"
    "github.com/hq-cml/sward2offer/common"
)

//思路：
//这道题，还是在变相得考察二叉树的遍历！
//注意：这里的实现中，path没有回退机制，其实这是利用了一个golang的slice特性
//func tt(a []int) {
//    a = append(a, 1)
//    fmt.Println(a)
//}
//a  := []int{1,2,3}
//fmt.Println(a) //1,2,3
//tt(a)          //1,2,3,1
//fmt.Println(a) //外层感知不到 1,2,3
func FindPath(root *common.TreeNode, num int) {
    path := []int{} //维护一个数组作为路径记录
    findPath(root, num, path)
}

func findPath(node *common.TreeNode, num int, path []int) {
    if node == nil {
        return
    }

    path = append(path, node.Val)
    if sum(path) == num &&
        node.Left == nil && node.Right == nil {
        fmt.Println("The Path:", path)
    }

    findPath(node.Left, num, path)
    findPath(node.Right, num, path)
}

func sum(arr []int) int {
    num := 0
    for _, v := range arr {
        num += v
    }
    return num
}
