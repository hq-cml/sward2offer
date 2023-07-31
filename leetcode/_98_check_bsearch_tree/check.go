/*
 * 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
 * 有效 二叉搜索树定义如下：
 * 节点的左子树只包含 小于 当前节点的数。
 * 节点的右子树只包含 大于 当前节点的数。
 * 所有左子树和右子树自身必须也是二叉搜索树。
 *
 * 输入：root = [2,1,3] => 输出：true
 * 输入：root = [5,1,4,null,null,3,6] => 输出：false
 * 解释：根节点的值是 5 ，但是右子节点的值是 4 。
 */
package _98_check_bsearch_tree

// 0表示null
func Check(node []int) bool {
	length := len(node)
	if length <= 1 {
		return true
	}

	for i := 0; i < length; i++ {
		v := node[i]
		left := 2*i + 1
		right := 2*i + 2
		if v == 0 {
			continue // 空
		}
		if left <= length-1 && node[left] != 0 && node[left] > v {
			return false
		}
		if right <= length-1 && node[right] != 0 && node[right] < v {
			return false
		}
	}
	return true
}
