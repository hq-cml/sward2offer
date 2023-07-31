/*
 * tx项目遇到的问题:
 *  1. 图形的深度遍历
 *  2. 同时判断是否存在环
 */
package _traverse_graph

import "fmt"

// 结构定义
type TopologyNode struct {
	Val int
	DownStreas []*TopologyNode
}

// 利用递归的方式，实现深度遍历
// passedBy用于记录是否访问过，如果访问过不应该继续
func Traverse(node *TopologyNode, passedBy map[int]struct{}, help StackIntIntf) (bool, string){
	var exist bool
	var circlePath string

	if node == nil {
		return exist, circlePath
	}

	// 判断辅助栈中是否存在当前节点，如果存在则说明存在环
	if help.Exist(node.Val) {
		circlePath = fmt.Sprintf("%v->%v", help.String(), node.Val)
		exist = true
		return exist, circlePath
	}

	// 节点已经访问过，则不应该继续
	// 这个边界这是有可能存在的（比如存在环的时候，或者存在公共的下游节点时候）
	if _, ok := passedBy[node.Val]; ok {
		return exist, circlePath
	}
	fmt.Println(node.Val)
	help.Push(node.Val)
	passedBy[node.Val] = struct{}{}
	for _, nd := range node.DownStreas {
		e, path := Traverse(nd, passedBy, help)
		if circlePath == "" {
			circlePath = path
		}
		exist = exist || e
	}
	// 程序结束前，需要将当前节点出栈
	help.Pop()
	return exist, circlePath
}

// ********自己实现一个简单的辅助栈
// StackIntIntf 整形栈接口类型
type StackIntIntf interface {
	Len() int         // 栈长度
	Push(int)         // 压栈
	Top() (int, bool) // 栈顶元素
	Pop() (int, bool) // 出栈
	String() string   // 打印全部栈元素
	Exist(int) bool   // 判断某个元素是否存在于栈中
}

// NewStackInt 新建
func NewStackInt() StackIntIntf {
	return &StackInt{
		m: map[int]int{},
	}
}

// StackInt 栈实现类型
type StackInt struct {
	s []int
	m map[int]int // key: 元素 val：个数
}

// Len 栈长度
func (stk *StackInt) Len() int {
	return len(stk.s)
}

// Push 压栈
func (stk *StackInt) Push(value int) {
	stk.s = append(stk.s, value)
	stk.m[value]++
}

// Top 栈顶元素，栈不会变化
func (stk *StackInt) Top() (int, bool) {
	if len(stk.s) == 0 {
		return 0, false
	}
	return stk.s[len(stk.s)-1], true
}

// Pop 出栈
func (stk *StackInt) Pop() (int, bool) {
	if len(stk.s) == 0 {
		return 0, false
	}
	value := stk.s[len(stk.s)-1]
	stk.s = stk.s[:len(stk.s)-1]
	if _, ok := stk.m[value]; ok {
		stk.m[value]--
	}
	return value, true
}

// String 打印全部栈元素
func (stk *StackInt) String() string {
	return fmt.Sprintf("%v", stk.s)
}

// Exist Exist
func (stk *StackInt) Exist(i int) bool {
	if _, ok := stk.m[i]; ok {
		return stk.m[i] > 0
	} else {
		return false
	}
}
