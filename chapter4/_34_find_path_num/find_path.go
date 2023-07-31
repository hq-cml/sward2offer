/*
 * 面试题34：二叉树中和为某一值的路径
 * 题目：输入一棵二叉树和一个整数，打印出二叉树中结点值的和为输入整数的所有路径。
 * 从树的根结点开始往下一直到叶结点所经过的结点形成一条路径。
 */
package _34_find_path_num

import (
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
func FindSumPath(root *common.TreeNode, num int) []int{
	var path []int
	if root == nil {
		return nil
	}
	path, ok := findSumPath(root, num, 0, path)
	if !ok {
		return nil
	}
	return path
}

func findSumPath(root *common.TreeNode, num, curr int, path []int) ([]int, bool) {
	curr += root.Val
	path = append(path, root.Val)
	// 当前累计值，等于需要值，则找到路径成功
	if curr == num {
		return path, true
	}
	// 未找到路径，则继续探测，先做后右
	if root.Left != nil {
		pathLeft, ok := findSumPath(root.Left, num, curr, path)
		if ok {
			return pathLeft, ok
		}
	}
	if root.Right != nil {
		pathRight, ok := findSumPath(root.Right, num, curr, path)
		if ok {
			return pathRight, ok
		}
	}
	// 如果整个探测失败，需要回退
	path = path[0: len(path)-1]
	return path, false
}

