## 460 LFU缓存-困难

题目：

请你为 [最不经常使用（LFU）](https://baike.baidu.com/item/缓存算法)缓存算法设计并实现数据结构。

实现 `LFUCache` 类：

- `LFUCache(int capacity)` - 用数据结构的容量 `capacity` 初始化对象
- `int get(int key)` - 如果键 `key` 存在于缓存中，则获取键的值，否则返回 `-1` 。
- `void put(int key, int value)` - 如果键 `key` 已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量 `capacity` 时，则应该在插入新项之前，移除最不经常使用的项。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 **最久未使用** 的键。

为了确定最不常使用的键，可以为缓存中的每个键维护一个 **使用计数器** 。使用计数最小的键是最久未使用的键。

当一个键首次插入到缓存中时，它的使用计数器被设置为 `1` (由于 put 操作)。对缓存中的键执行 `get` 或 `put` 操作，使用计数器的值将会递增。

函数 `get` 和 `put` 必须以 `O(1)` 的平均时间复杂度运行。



**解题思路**

```go
// date 2024/01/09
type LFUNode struct {
    key, val int
    frequency int
    pre, next *LFUNode
}

type LFUList struct {
    size int
    head, tail *LFUNode
}

type LFUCache struct {
    cache  map[int]*LFUNode
    list   map[int]*LFUList
    min    int
    cap    int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:   capacity,
		min:   0,
		cache: make(map[int]*LFUNode),
		list:  make(map[int]*LFUList),
	}
}


func (this *LFUCache) Get(key int) int {
    node, ok := this.cache[key]
    if !ok {
        return -1
    }    
    this.list[node.frequency].removeNode(node)
    node.frequency++
    if _, ok = this.list[node.frequency]; !ok {
        this.list[node.frequency] = NewLFUList()
    }
    oldList := this.list[node.frequency]
    oldList.addToFront(node)
    
    if node.frequency-1 == this.min {
        if minList, ok := this.list[this.min]; !ok || minList.isEmpty() {
            this.min = node.frequency
        }
    }
    return node.val
}


func (this *LFUCache) Put(key int, value int)  {
    if this.cap == 0 {
        return
    }
    if node, ok := this.cache[key]; ok {
        node.val = value
        this.Get(key)
        return
    }
    if this.cap == len(this.cache) {
        minList := this.list[this.min]
        rmd := minList.removeFromRear()
        delete(this.cache, rmd.key)
    }
    node := &LFUNode{key: key, val: value, frequency: 1}
    this.min = 1
    if _, ok := this.list[this.min]; !ok {
        this.list[this.min] = NewLFUList()
    }
    minList := this.list[this.min]
    minList.addToFront(node)
    this.cache[key] = node
}

func NewLFUList() *LFUList {
	res := &LFUList{
		size: 0,
		head: &LFUNode{0, 0, 0, nil, nil},
		tail: &LFUNode{0, 0, 0, nil, nil},
	}
    res.head.next = res.tail
    res.tail.pre = res.head
    return res
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



/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
```

