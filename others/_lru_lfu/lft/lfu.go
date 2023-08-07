package lft

/*
 * LFU实现
 * 2 1 2 1 2 3 4 -> 删除3
 */

// LFUCache 具体的缓存
type LFUCache struct {
	Cap         int           // 容量
	Cache       map[int]*Node // 实际的缓存key=>Node
	KeyToFreq   map[int]int   // 映射：key=>频率
	FreqToNodes map[int]*List // 逆映射：频率 => Node list（相同频率的Node组成队列，按照FIFO的顺序操作）
	MinFreq     int           // 记录当前最小的频率（即下一个将要被替换的频率）
}

// 拥有头尾节点（拥有头尾节点的好处是，对于首元素和尾元素的处理，都不需要特殊对待）
type List struct {
	cnt  int
	head *Node
	tail *Node
}

// 缓存的每个数据节点，包括了key和val，以及前后指针
type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

func NewLFU(cap int) *LFUCache {
	return &LFUCache{
		Cap:         cap,
		Cache:       make(map[int]*Node),
		KeyToFreq:   make(map[int]int),
		FreqToNodes: make(map[int]*List),
	}
}

func NewList() *List {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &List{
		head: head,
		tail: tail,
	}
}

// 队尾追加一个node
func (l *List) appendNode(node *Node) {
	pre := l.tail.prev
	pre.next = node
	node.next = l.tail
	l.tail.prev = node
	node.prev = pre
	l.cnt++
}

// 删除节点
func (l *List) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	l.cnt--
}

func (l *List) isEmpty() bool {
	return l.cnt == 0
}

// 从valueMap 中查找相应的键值对如果存在，也需要更新下key对应的频率
func (lf *LFUCache) Get(key int) int {
	//1.从value map 中获取
	v, ok := lf.Cache[key]
	if !ok {
		return -1
	}

	//2.需要跟新FreqMap
	lf.incrFreq(key)
	return v.value
}

// 增加key对应的频率
func (lf *LFUCache) incrFreq(key int) {
	//根据key 获取 freq，并自增
	freq, _ := lf.KeyToFreq[key] // 忽略ok值，如果不存在，freq为0
	newFreq := freq + 1
	lf.KeyToFreq[key] = newFreq

	//移除 freqtokeys 建议用链表删除(下面代码使用栈实现)
	node := lf.Cache[key]
	lf.FreqToNodes[freq].removeNode(node)

	//需要判断freqToKeys的列表，如果空了，需要移除这个freq
	if lf.FreqToNodes[freq].isEmpty() {
		delete(lf.FreqToNodes, freq)
		//如果这个freq是 minFreq 需要更新一下minFreq
		if freq == lf.MinFreq {
			lf.MinFreq = newFreq
		}
	}

	//将 key 加入到freq+1 对应的列表中
	if _, ok := lf.FreqToNodes[newFreq]; !ok {
		lf.FreqToNodes[newFreq] = NewList()
	}
	lf.FreqToNodes[newFreq].appendNode(node)
}

func (lf *LFUCache) Put(key, val int) int {
	//key 存在，需要更新freqMap
	if _, exist := lf.Cache[key]; exist {
		//更新最新的值
		lf.Cache[key].value = val
		//更新频率
		lf.incrFreq(key)
	} else {
		//如果不存在, 需要判断容量是不是满了
		//如果容量满了，需要根据最小的minFreq，先删除一个key
		if len(lf.Cache) >= lf.Cap {
			lf.removeMinFreq(lf.MinFreq)
		}

		//如果容量未满 valueMap里面新加一个key,并且加入到freqMap key=1 对应的列表中
		node := &Node{
			key:   key,
			value: val,
		}
		//更新 KV
		lf.Cache[key] = node

		//有新增的key进来，更新 KF
		lf.KeyToFreq[key] = 1

		//有新增的key进来，MinFreq必然 = 1
		lf.MinFreq = 1

		//将新的节点append到FreqToNodes对应位置
		if _, ok := lf.FreqToNodes[lf.MinFreq]; !ok {
			lf.FreqToNodes[lf.MinFreq] = NewList()
		}
		lf.FreqToNodes[lf.MinFreq].appendNode(node)
	}
	return 0
}

// 从最小频率对应的数组中，删除最初加入的key
func (lf *LFUCache) removeMinFreq(minFreq int) {
	//删除FreqToNodes对应项（第一个节点）
	node := lf.FreqToNodes[minFreq].head.next
	delKey := node.key
	lf.FreqToNodes[minFreq].removeNode(node)
	if lf.FreqToNodes[minFreq].isEmpty() {
		delete(lf.FreqToNodes, minFreq)
	}

	//删除 KV 表
	delete(lf.Cache, delKey)

	//更新 KF表
	delete(lf.KeyToFreq, delKey)
}
