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
