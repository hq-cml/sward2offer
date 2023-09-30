/*
 * 面试题32（一）：不分行从上往下打印二叉树
 * 题目：从上往下打印出二叉树的每个结点，同一层的结点按照从左到右的顺序打印。
 *
 * 面试题32（二）：分行从上到下打印二叉树
 * 题目：从上到下按层打印二叉树，同一层的结点按从左到右的顺序打印，每一层
 * 打印到一行。
 *
 * 面试题32（三）：之字形打印二叉树
 * 题目：请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺
 * 序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，
 * 其他行以此类推。
 */
package _102_print_tree

import (
	"container/list"
	"fmt"
	"github.com/hq-cml/sward2offer/common"
)

// 思路：
// 引入队列，从root开始入队列
// 开始出队列操作，每次出队列，就将出队列元素的左右节点，分别放入队列
// 如此往复，直到队列为空
// 难度：3*
func PrintTreeIn1Line(root *common.TreeNode) error {
	if root == nil {
		return nil
	}

	mylist := list.New()
	mylist.PushBack(root)
	for mylist.Len() != 0 {
		e := mylist.Front()
		mylist.Remove(e)
		node := e.Value.(*common.TreeNode)
		if node == nil {
			continue
		}
		fmt.Print(node.Val, " ")
		mylist.PushBack(node.Left)
		mylist.PushBack(node.Right)
	}
	return nil
}

// 思路：分层打印二叉树，每行从左到右
// 再上一个程序的基础上，增加两个计数器，分别为当前行节点个数和下一行节点个数
// 每次打印完毕之后，当前行数减一，如果减到了0，则输出换行符
// 同时每次队列加入左右节点的时候，下一行个数自增
// 难度：4*
func PrintTreeInMultiLine(root *common.TreeNode) error {
	if root == nil {
		return nil
	}

	mylist := list.New()
	mylist.PushBack(root)
	currentLevelCnt := 1
	nextLevelCnt := 0
	for mylist.Len() != 0 {
		e := mylist.Front()
		mylist.Remove(e)
		node := e.Value.(*common.TreeNode)
		if node == nil {
			continue
		}

		if node.Left != nil {
			mylist.PushBack(node.Left)
			nextLevelCnt++
		}

		if node.Right != nil {
			mylist.PushBack(node.Right)
			nextLevelCnt++
		}

		fmt.Print(node.Val, " ")
		currentLevelCnt--
		// 妙。。。
		if currentLevelCnt == 0 {
			fmt.Println() //换行
			currentLevelCnt = nextLevelCnt
			nextLevelCnt = 0
		}
	}

	return nil
}

// 思路：蛇形线，按行打印二叉树
// 根据特点，将队列换成栈，且是同时使用两个栈
// 并且，每一行左右节点入栈的顺序是不同的，是先左后右还是先右后左，利用行号和2取模来判断
// 难度：5*
func PrintTreeSnakeLine(root *common.TreeNode) error {
	if root == nil {
		return nil
	}

	mystack1 := common.NewStack(true)
	mystack2 := common.NewStack(true)
	mystack1.Push(root)
	currentLevel := 1 //层号与左右的关系
	currentLevelCnt := 1
	nextLevelCnt := 0
	for mystack1.Len() != 0 {
		e, _ := mystack1.Pop()
		node := e.(*common.TreeNode)
		if node == nil {
			continue
		}

		var node1 *common.TreeNode
		var node2 *common.TreeNode
		// 顺序很精妙！是先左后右还是先右后左
		if currentLevel%2 == 1 {
			node1 = node.Right
			node2 = node.Left
		} else {
			node1 = node.Left
			node2 = node.Right
		}

		if node1 != nil {
			mystack2.Push(node1)
			nextLevelCnt++
		}
		if node2 != nil {
			mystack2.Push(node2)
			nextLevelCnt++
		}

		fmt.Print(node.Val, " ")
		currentLevelCnt--
		if currentLevelCnt == 0 {
			fmt.Println() //换行
			currentLevelCnt = nextLevelCnt
			nextLevelCnt = 0
			currentLevel++                          //行数自增
			mystack1, mystack2 = mystack2, mystack1 //对换两个栈，永远将stack1作为判断依据
		}
	}

	return nil
}
