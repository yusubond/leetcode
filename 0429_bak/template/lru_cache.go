package template

type (
	// CacheNode define
	// key, val 存储缓存的键值对
	// pre, next 指向双向链表的前驱和后驱节点
	CacheNode struct {
		key, val  int
		pre, next *CacheNode
	}

	// LRUCache define
	LRUCache struct {
		size, capacity int
		head, tail     *CacheNode
		cache          map[int]*CacheNode
	}
)

// NewLRUCache define
func NewLRUCache(cap int) *LRUCache {
	ca := &LRUCache{
		size:     0,
		capacity: cap,
		head:     &CacheNode{0, 0, nil, nil},
		tail:     &CacheNode{0, 0, nil, nil},
		cache:    make(map[int]*CacheNode, cap),
	}
	ca.head.next = ca.tail
	ca.tail.pre = ca.head
	return ca
}

// Get define
func (l *LRUCache) Get(key int) int {
	node, ok := l.cache[key]
	if !ok {
		return -1
	}
	l.moveToHead(node)
	return node.val
}

// Put define
func (l *LRUCache) Put(key, value int) {
	if node, ok := l.cache[key]; ok {
		// update node and move to head
		node.val = value
		l.moveToHead(node)
		return
	}

	// add node
	node := &CacheNode{key: key, val: value}
	l.cache[key] = node
	l.addToHead(node)
	l.size++
	if l.size > l.capacity {
		// remote tail
		rmd := l.removeTail()
		delete(l.cache, rmd.key)
		l.size--
	}
}

// add node to head
func (l *LRUCache) addToHead(node *CacheNode) {
	node.pre = l.head
	node.next = l.head.next

	l.head.next.pre = node
	l.head.next = node
}

// remove the node from linklist
func (l *LRUCache) removeNode(node *CacheNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

// move the node to head
func (l *LRUCache) moveToHead(node *CacheNode) {
	l.removeNode(node)
	l.addToHead(node)
}

// remove and return the tail node
func (l *LRUCache) removeTail() *CacheNode {
	node := l.tail.pre
	l.removeNode(node)
	return node
}
