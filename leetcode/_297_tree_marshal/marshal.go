/*
 * 面试题37：序列化二叉树
 * 题目：请实现两个函数，分别用来序列化和反序列化二叉树。
 */
package _297_tree_marshal

import (
	"fmt"
	"github.com/hq-cml/sward2offer/common"
	"github.com/spf13/cast"
	"strings"
)

//思路：
//思路1：转换成前序和中序遍历，然后将两个序列返回，作为序列化。反序列化则是之前做过的，根据前序和中序遍历进行重建。
//思路2：引入叶子节点空指针的特殊字符，这样就可以只用一个前序序列，也能够重建。这仍然是一个递归的思路。

// 字符串序列化
// 难度：3*
func TreeMarshal(root *common.TreeNode) string {
	if root == nil {
		return "$,"
	}

	//数字值转换成为字符串
	str := fmt.Sprintf("%v,", root.Val)

	return str +
		TreeMarshal(root.Left) +
		TreeMarshal(root.Right)
}

// 反序列化：
// 本质上仍是递归的应用
// 难度:5*
func TreeUnMarshal(str string) *common.TreeNode {
	//过滤掉最后逗号
	str = strings.Trim(str, ",")

	root, _ := unMarshal(str)
	return root
}

// 递归，不断消耗掉str，直到为空
// 第二返回值表示的是剩余的str
func unMarshal(str string) (*common.TreeNode, string) {
	if len(str) == 0 {
		return nil, str
	}
	tmp := strings.Split(str, ",")
	if tmp[0] == "$" {
		return nil, strings.Join(tmp[1:], ",")
	}
	root := &common.TreeNode{
		Val: cast.ToInt(tmp[0]),
	}

	// 递归左右
	root.Left, str = unMarshal(strings.Join(tmp[1:], ","))
	root.Right, str = unMarshal(str)
	return root, str
}
