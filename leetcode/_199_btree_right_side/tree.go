/*
 * 一棵二叉树从右侧视角看过去，从上到下组成的数字
 * 例子1：
 *     1
 *   2    3     => 从右侧看应该是134
 * 6  5  3  4
 *
 * 例子2：
 *      1
 *   2     3     => 从右侧看应该是135
 * 6  5  nil nil
 */
package _199_btree_right_side

import (
	"container/list"
	"github.com/hq-cml/sward2offer/common"
)

// 思路：本质上，这是二叉树按层遍历的变种，利用一个队列
func RightSide(root *common.TreeNode) int {
	if root == nil {
		return -1
	}
	ret := 0
	currCnt := 1
	nextCnt := 0
	l := list.New()
	l.PushBack(root)
	for l.Len() != 0 {
		tmp := l.Front()
		node := tmp.Value.(*common.TreeNode)
		l.Remove(tmp)

		currCnt --
		if node.Left != nil {
			l.PushBack(node.Left)
			nextCnt ++
		}
		if node.Right != nil {
			l.PushBack(node.Right)
			nextCnt ++
		}

		// 需要换行的时候，即是本层的最右侧节点，所以进位相加
		if currCnt == 0 {
			currCnt = nextCnt
			nextCnt = 0
			ret = ret*10 + node.Val
		}
	}
	return ret
}