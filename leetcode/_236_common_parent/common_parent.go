/*
 * 面试题68：树中两个结点的最低公共祖先
 * 题目：输入根节点和两个树结点，求它们的最低公共祖先。
 */
package _236_common_parent

import (
	"github.com/hq-cml/sward2offer/basic"
	"github.com/hq-cml/sward2offer/common"
)

// 做了leetCode之后得到的牛逼思路，基于递归！
// 这题非常的抽象！所谓公共祖先，而且还要求最低，则应该保证：
// 节点左子树包含p或者q && 右子树包含p或者q，这样的节点才是最低且公共
// 如果某个节点左子树或者右子树，同时包括p和q，那么这个节点就不是最低公共
// 所以，本题的不要把LowestCommonAncestor理解为求p和q的公共最低，而应该理解为：
// 在root这颗树及其子树总，找到一个p或者q就可以
func LowestCommonAncestor(root, p, q *common.TreeNode) *common.TreeNode {
	if root == nil {
		return nil
	}

	// 正好某个根节点了
	if root == p || root == q {
		return root
	}

	// 递归左右，得到结果
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)

	// 如果左右结果都非空
	// 说明左子树存在一个p或者q，同时右子树存在一个p或者q
	// 说明真正的公共节点正是自身root
	if left != nil && right != nil {
		return root
	}

	// 左结果非空，说明公共节点在左侧，返回
	if left != nil {
		return left
	}

	// 右结果非空，说明公共节点在右侧，返回
	if right != nil {
		return right
	}

	// left == nil && right == nil，说明不存在公共节点
	return nil
}

// 以前的思路：
//  1. 如果是二叉搜索树，这是一个递归的思路。如果两个节点都比根小，则递归左子树；
//     如果都比根大，递归右子树；如果一大一小，答案就是根。（利用了二叉搜索树的性质）
//  2. 如果是包含指向父节点的指针，那么这个题就平移成了求两个链表的首个汇合节点。这个前面有
//  3. 如果上述特殊条件都没有，那么就分别求出从根到两个节点各自的路径，然后从头开始找到第一个共同元素即可。
func FindCommonParent(root *common.TreeNode, num1, num2 int) (int, bool) {
	path1 := []int{}
	path2 := []int{}

	path1, ok1 := basic.FindPath(root, num1, path1)
	path2, ok2 := basic.FindPath(root, num2, path2)

	if !ok1 || !ok2 {
		return -1, false
	}

	for k, v := range path1 {
		if k > len(path2)-1 {
			//paht1长
			return path1[k-1], true
		}
		//第一次出现不同，则上一个节点是公共节点
		if v != path2[k] {
			return path1[k-1], true
		}
	}

	//到这里说明path1比path2短，则公共节点应该是path1的最后一个元素
	return path1[len(path1)-1], true
}
