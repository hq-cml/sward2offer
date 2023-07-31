package lru

/*
 * LRU实现
 * 利用一个队列来实现FIFO，
 * 2 1 2 1 2 3 4 -> 删除1
 */
// LRUCache 具体的缓存
type LRUCache struct {
	cache    map[int]*Node
	list     *List // 维护一个双向队列，每次最新访问的元素放在队尾，删除从头部开始
	capacity int   // 容量
}

// 拥有头尾节点的链表（拥有头尾节点的好处是，对于首元素和尾元素的处理，都不需要特殊对待）
type List struct {
	head *Node
	tail *Node
}

// 双向链表节点
type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

func NewList() *List {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &List{
		head: head,
		tail: head.next,
	}
}

// NewLRU 构建缓存容器
func NewLRU(capacity int) LRUCache {
	// 结构一个双向空链表(拥有头结点和尾结点，都是空节点)
	return LRUCache{
		cache:    make(map[int]*Node),
		list:     NewList(),
		capacity: capacity,
	}
}

// 在队列尾部插入
func (l *List) appendNode(node *Node) {
	curr := l.tail.prev
	curr.next.prev = node
	node.next = curr.next
	curr.next = node
	node.prev = curr
}

// 删除链表中的节点
func removeNode(node *Node) {
	node.next.prev = node.prev
	node.prev.next = node.next
}

func (lc *LRUCache) Get(key int) int {
	// 依据key 获取缓存
	v, ok := lc.cache[key]
	// 如果没有缓存, 返回-1
	if !ok {
		return -1
	}
	// 如果有缓存
	removeNode(v)         // 移除该缓存
	lc.list.appendNode(v) // 把缓存增加双向链表尾部
	return v.value
}

// Put 新建缓存
func (lc *LRUCache) Put(key int, value int) {
	// 曾经有缓存
	if v, ok := lc.cache[key]; ok { // v 是双链表中的节点
		v.value = value       // 更新链表节点中的值
		lc.cache[key] = v     // 更新缓存中映射关系
		removeNode(v)         // 移除该缓存
		lc.list.appendNode(v) // 把缓存增加双向链表尾部
		return
	}
	// 缓存超长 淘汰缓存
	if len(lc.cache) == lc.capacity {
		node := lc.list.head.next
		removeNode(node) // 删除该节点
		delete(lc.cache, node.key)
	}
	newNode := &Node{
		key:   key,
		value: value,
	}
	lc.cache[key] = newNode
	lc.list.appendNode(newNode)
}
