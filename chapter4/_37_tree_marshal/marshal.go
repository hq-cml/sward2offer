/*
 * 面试题37：序列化二叉树
 * 题目：请实现两个函数，分别用来序列化和反序列化二叉树。
 */
package _37_tree_marshal

import (
	"fmt"
	"github.com/hq-cml/sward2offer/common"
	"strconv"
	"strings"
)

//思路：
//思路1：转换成前序和中序遍历，然后将两个序列返回，作为序列化。反序列化则是之前做过的，根据前序和中序遍历进行重建。
//思路2：引入叶子节点空指针的特殊字符，这样就可以只用一个前序序列，也能够重建。这仍然是一个递归的思路。

//字符串序列化
//难度：3*
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

//反序列化：
//本质上仍是递归的应用
//难度:5*
func TreeUnMarshal(str string) *common.TreeNode {
	//过滤掉最后逗号
	str = strings.Trim(str, ",")

	//idx作为str字符串的索引值
	idx := 0
	var root *common.TreeNode
	unmarshal(str, &root, &idx)
	return root
}

//关键是结束条件！
//本质上，这还是要结合二叉树先序遍历的特点
//root是二级指针，
//关键是idx的使用，这是一个指针，所以他是随着构建，在不断向后
func unmarshal(str string, root **common.TreeNode, idx *int) {
	//退出条件：*idx最大值就是len(str)-1
	if *idx > len(str)-1 {
		return
	}

	//求出根的字符
	tmpIdx := strings.Index(str[*idx:], ",")
	if tmpIdx == -1 {
		//出现这种情况，说明只有一个$，即整个树是nil
		return
	}

	if str[*idx:*idx+tmpIdx] == "$" { //$,
		//当前这一枝儿，是空的
		*root = nil
		*idx += 2
		return
	} else {
		//当前这一枝儿，非空，有数字，比如：123,
		val, _ := strconv.Atoi(str[*idx : *idx+tmpIdx])
		*root = &common.TreeNode{
			Val: val,
		}
		*idx = *idx + tmpIdx + 1 //idx指向了下一个数字（可能是数字，也可能是$）
	}

	//递归的构建左子树（idx是指针，所以*idx是可以不断增长的）
	unmarshal(str, &(*root).Left, idx)

	//递归得构建右子树
	unmarshal(str, &(*root).Right, idx)
}
