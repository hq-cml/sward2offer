/**
 * 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
 * 实现 LRUCache 类：
 * LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
 * int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
 * void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
 * 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
 */
package _146_lru_cache

type LRUCache struct {
	cache map[int]*Node // key=>value
	list  *Dlist        // 维护一个双向队列，每次最新访问的元素放在队尾，删除从头部开始
	cap   int
}

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

// 拥有头尾节点的链表（拥有头尾节点的好处是，对于首元素和尾元素的处理，都不需要特殊对待）
type Dlist struct {
	head *Node
	tail *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache: map[int]*Node{},
		list:  NewDlist(),
		cap:   capacity,
	}
}

func NewDlist() *Dlist {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &Dlist{
		head: head,
		tail: tail,
	}
}

// 在队列尾部插入
func (l *Dlist) appendTail(node *Node) {
	curr := l.tail.prev
	curr.next.prev = node
	node.next = curr.next
	curr.next = node
	node.prev = curr
}

// 删除双向链表中的一个节点
func (n *Node) Remove() {
	n.next.prev = n.prev
	n.prev.next = n.next
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}

	// 更新缓存优先级队列
	node.Remove()
	this.list.appendTail(node)

	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	// 曾经有缓存，刷新调整
	node, ok := this.cache[key]
	if ok {
		node.val = value
		node.Remove()
		this.list.appendTail(node)
		return
	}

	// 判断是否超长
	if len(this.cache) == this.cap {
		toDel := this.list.head.next
		toDel.Remove()
		delete(this.cache, toDel.key)
	}

	newNode := &Node{
		key: key,
		val: value,
	}
	this.cache[key] = newNode
	this.list.appendTail(newNode)
}
