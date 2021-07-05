/*
// 面试题37：序列化二叉树
// 题目：请实现两个函数，分别用来序列化和反序列化二叉树。

两个思路：
第一个思路是转换成前序和中序遍历，然后将两个序列返回，作为序列化。反序列化则是之前做过的，根据前序和中序遍历进行重建。
第二个思路是引入叶子节点空指针的特殊字符，这样就可以只用一个前序序列，也能够重建。这仍然是一个递归的思路。
 */
package _37_tree_marshal

import (
    "fmt"
    "github.com/hq-cml/sward2offer/common"
    "strconv"
    "strings"
)

//字符串序列化
func TreeMarshal(root *common.TreeNode) string {
    if root == nil {
        return "$,"
    }

    str := fmt.Sprintf("%v,", root.Val)

    return str +
        TreeMarshal(root.Left) +
        TreeMarshal(root.Right)
}

//反序列化：
//本质上仍是递归的应用
//难度:4
func TreeUnMarshal(str string) *common.TreeNode {
    str = strings.Trim(str, ",")
    var root *common.TreeNode
    idx := 0
    unmarshal(str, &root, &idx)
    return root
}

//关键是结束条件！
func unmarshal(str string, node **common.TreeNode, idx *int) {
    if *idx > len(str)-1 {
        return
    }

    tmpIdx := strings.Index(str[*idx:], ",")
    if tmpIdx == -1 {
        return
    }
    if str[*idx: *idx + tmpIdx] == "$" {
        *node = nil
        *idx += 2
        return    //结束，给右子树留一点。。。
    } else {
        val, _ := strconv.Atoi(str[*idx: *idx + tmpIdx])
        *node = &common.TreeNode{
            Val:   val,
        }
        *idx = *idx + tmpIdx + 1
    }

    unmarshal(str, &(*node).Left, idx)
    unmarshal(str, &(*node).Right, idx)
}