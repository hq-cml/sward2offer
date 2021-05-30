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