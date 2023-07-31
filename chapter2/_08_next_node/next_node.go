/*
 * 面试题8：二叉树中序遍历的下一个结点
 * 题目：给定一棵二叉树和其中的一个结点，如何找出中序遍历顺序的下一个结点？
 * 树中的结点除了有两个分别指向左右子结点的指针以外，还有一个指向父结点的指针。
 */
package _08_next_node

type Node struct {
	I      int
	Left   *Node
	Right  *Node
	Parant *Node
}

//思路，通过画图，简化问题，有如下规律：
//1. 如果节点拥有右子树，那么返回右子树的最左边的子节点。
//2. 如果节点无右子树，且它是其父节点的左子树，那么直接返回父节点
//3. 如果节点是其父节点的右子树，那么就一直向上找，直到某个节点是其父节点的左子树，然后返回这个父节点。
// 都不符合，返回空。
// 注意这3个规律，是有顺序的，不能随意的颠倒。逻辑上他们覆盖了所有的情况，并给出了处理办法。
func NextNode(find *Node) *Node {
	if find == nil {
		return nil
	}

	p := find
	//1. 如果节点拥有右子树，那么返回右子树的最左边的子节点。
	if p.Right != nil {
		p = p.Right
		for p.Left != nil { //一直向左走
			p = p.Left
		}
		return p
	}
	// 节点没有右子树
	if p.Parant == nil {
		return nil // 节点是root
	}

	//2. 如果节点是他父节点的左子树，那么直接返回父节点
	if p.Parant.Left == p {
		return p.Parant
	}

	//3. 如果节点是他父节点的右子树，那么就一直向上找，直到某个节点是其父节点的左子树，然后返回这个父节点
	for p.Parant != nil && p.Parant.Right == p {
		p = p.Parant
	}
	return p.Parant
}
