package template

import "fmt"

type (
	// LFUNode define
	LFUNode struct {
		key, val  int
		frequency int
		pre, next *LFUNode
	}
	// LFUList define
	LFUList struct {
		size       int // the total nodes in lfu list
		head, tail *LFUNode
	}
	// LFUCache define
	LFUCache struct {
		cache    map[int]*LFUNode // key is node key
		list     map[int]*LFUList // key is frequency
		capacity int
		min      int
	}
)

// NewLFUList new a double linked list
func NewLFUList() *LFUList {
	list := &LFUList{
		size: 0,
		head: &LFUNode{0, 0, 0, nil, nil},
		tail: &LFUNode{0, 0, 0, nil, nil},
	}
	list.head.next = list.tail
	list.tail.pre = list.head
	return list
}

// addToHead define
// node is update or new, so insert to front
// 新节点肯定是被访问或者新插入的
func (l *LFUList) addToFront(node *LFUNode) {
	node.pre = l.head
	node.next = l.head.next

	l.head.next.pre = node
	l.head.next = node
	l.size++
}

func (l *LFUList) removeFromRear() *LFUNode {
	node := l.tail.pre
	l.removeNode(node)
	return node
}

func (l *LFUList) removeNode(node *LFUNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
	l.size--
}

func (l *LFUList) isEmpty() bool {
	return l.size == 0
}

func NewLFUCache(cap int) *LFUCache {
	return &LFUCache{
		capacity: cap,
		min:      0,
		cache:    make(map[int]*LFUNode),
		list:     make(map[int]*LFUList),
	}
}

func (l *LFUCache) Put(key, value int) {
	if l.capacity == 0 {
		return
	}
	// update node's frequency if exist
	if node, ok := l.cache[key]; ok {
		node.val = value
		// 处理过程与 get 一样，直接复用
		l.Get(key)
		return
	}

	// not exist
	// 如果不存在且缓冲满了，需要删除
	if l.capacity == len(l.cache) {
		minList := l.list[l.min]
		rmd := minList.removeFromRear()
		delete(l.cache, rmd.key)
	}

	// new node, insert to map
	node := &LFUNode{key: key, val: value, frequency: 1}
	// the min change to 1, once create a new node
	l.min = 1
	if _, ok := l.list[l.min]; !ok {
		l.list[l.min] = NewLFUList()
	}
	oldList := l.list[l.min]
	oldList.addToFront(node)
	// insert node to all cache
	l.cache[key] = node
}

func (l *LFUCache) Get(key int) int {
	node, ok := l.cache[key]
	if !ok {
		return -1
	}
	// first remote from the old frequency, then move to the new double list
	l.list[node.frequency].removeNode(node)
	node.frequency++
	if _, ok = l.list[node.frequency]; !ok {
		l.list[node.frequency] = NewLFUList()
	}
	oldList := l.list[node.frequency]
	oldList.addToFront(node)

	// if l.min is empty update
	if node.frequency-1 == l.min {
		if minList, ok := l.list[l.min]; !ok || minList.isEmpty() {
			l.min = node.frequency
		}
	}
	return node.val
}

func (l *LFUCache) ShowLFUNode() {
	for k, v := range l.cache {
		fmt.Printf("node key %d, value = %+v\n", k, v)
	}
}
