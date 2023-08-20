/*
*

	*Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。

请你实现 Trie 类：

Trie() 初始化前缀树对象。
void insert(String word) 向前缀树中插入字符串 word 。
boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。
*/
package _208_trie_tree

// 这个树真是鼎鼎大名，如雷贯耳
// 本质上就是一个26叉树，
type Trie struct {
	Root *Node
}

type Node struct {
	C      byte
	Child  [26]*Node
	IsWord bool // 以当前node结尾，是否作为一个合法单词
}

func Constructor() Trie {
	return Trie{
		Root: &Node{},
	}
}

// 本质上，是一种边查找，边构建的过程
func (this *Trie) Insert(word string) {
	root := this.Root
	for _, c := range []byte(word) {
		if root.Child[c-'a'] == nil {
			root.Child[c-'a'] = &Node{
				C: c,
			}
		}
		root = root.Child[c-'a']
	}
	root.IsWord = true // 当前最终节点，是一个合法单词
}

// 本质上，单纯的查找，需要是合法单词，所以需要关注IsWord
func (this *Trie) Search(word string) bool {
	root := this.Root
	for _, c := range []byte(word) {
		if root.Child[c-'a'] == nil {
			return false
		}
		root = root.Child[c-'a']
	}
	return root.IsWord
}

// 本质上，也是查找，和Search不同的是不需要关注IsWord
func (this *Trie) StartsWith(prefix string) bool {
	root := this.Root
	for _, c := range []byte(prefix) {
		if root.Child[c-'a'] == nil {
			return false
		}
		root = root.Child[c-'a']
	}
	return true
}
