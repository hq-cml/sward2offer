package temp

import (
	"github.com/hq-cml/sward2offer/common"
	"math"
)

func Good1(root *common.TreeNode) int {
	var ret int
	var good func(*common.TreeNode, int)
	good = func(root *common.TreeNode, currMax int) {
		if root == nil {
			return
		}
		if root.Val >= currMax {
			ret++
			currMax = root.Val
		}
		good(root.Left, currMax)
		good(root.Right, currMax)
	}
	good(root, math.MinInt)
	return ret
}
