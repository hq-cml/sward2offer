/*
 * 面试题7：重建二叉树
 * 题目：输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输
 * 入的前序遍历和中序遍历的结果中都不含重复的数字。例如输入前序遍历序列{1, 2, 4, 7, 3, 5, 6, 8}和
 * 中序遍历序列{4, 7, 2, 1, 5, 3, 8, 6}，则重建出如下的二叉树并输出它的头结点。
 *       1
 *    2      3
 *  4      5    6
 *    7       8
 */
package _07_build_tree

import (
	"errors"
	"github.com/hq-cml/sward2offer/common"
)

//pre: 1, 2, 4, 7, 3, 5, 6, 8
//mid: 4, 7, 2, 1, 5, 3, 8, 6
func BuildTree(pre, mid []int) (*common.TreeNode, error) {
	if len(pre) != len(mid) {
		return nil, errors.New("valid pre & mid")
	}

	length := len(pre)
	if length == 0 {
		return nil, nil
	}

	head := pre[0]
	idx := 0
	for i, v := range mid {
		if v == head {
			idx = i
			break
		}
	}

	node := common.NewNode(head)
	var err1 error
	var err2 error
	node.Left, err1 = BuildTree(pre[1 : idx+1], mid[0 : idx])
	node.Right, err2 = BuildTree(pre[idx+1:], mid[idx+1:])
	if err1 != nil || err2 != nil {
		return nil, errors.New(err1.Error() + err2.Error())
	}
	return node, nil
}
