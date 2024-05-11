## 146 LRU 缓存-中等

题目：

请你设计并实现一个满足 LRU（最少最近使用）缓存约束的数据结构。

实现 LRUCache 类：

- `LRUCache(int capacity)` 以正整数作为容量 capacity 初始化 LRU 缓存
- `int get(int key)` 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1
- `void put(int key, int value)` 如果关键字 key 已经存在，则变更其数据值 value；如果不存在，则向缓存中插入该组`key-value`。如果插入操作导致关键字数量超过 capacity，则应该逐出最久未使用的关键字。

函数 get 和 put 必须在 O(1) 的平均时间复杂度运行。



分析：

这里重点理解 LRU 的机制。LRU（Least Recently Used）即，最近最少使用，其核心是：

1. 最近。最近访问的要保留在缓存中。
2. 最少。并不是值访问次数最少，而是很久没有访问，且超过缓存容量的时候，可以删除。比如容量为 2 的缓存。A 访问次数为100，最近一次访问是 B 数据，从访问次数为2，那么 新插入的数据会取代 A，而不是 B。因为最近访问的 B 数据位于队头，而 A 位于队尾。

所谓，热数据也是这个道理，按访问时间，越是最近访问的，越是位于链表的头部。



要求 get 和 put 函数在 O(1) 时间内完成，那么考虑双向链表来实现。同时，用 head, tail 做头和尾的伪节点。这样在头部插入、去除尾部会很方便。

辅助 map 结构实现 O(1)查找。

`get`函数：

如果不存在，直接返回 -1；如果存在，把缓存 node 移到 头部。

`put`函数：

如果不存在，直接在头部插入，插入后超过容量的，直接删除尾部。

如果存在，直接更新数据，并移到头部。

```go
// 辅助函数
addToHead(node)
removeNode(node)
moveToHead(node)
removeTail()
```



```go
// date 2023/10/16
type MyNode struct {
    key, val int
    prev, next *MyNode
}

type LRUCache struct {
    size int
    cache map[int]*MyNode
    head, tail *MyNode
}


func Constructor(capacity int) LRUCache {
    lr := LRUCache{
        size: capacity,
        cache: make(map[int]*MyNode, capacity),
        head: &MyNode{0,0,nil,nil},
        tail: &MyNode{0,0,nil,nil},
    }
    lr.head.next = lr.tail
    lr.tail.prev = lr.head
    return lr
}


func (this *LRUCache) Get(key int) int {
    node, ok := this.cache[key]
    if !ok {
        return -1
    }
    this.moveToHead(node)
    return node.val
}


func (this *LRUCache) Put(key int, value int)  {
    node, ok := this.cache[key]
    if !ok {
        // add
        node := &MyNode{key, value, nil, nil}
        this.cache[key] = node
        this.addToHead(node)
        if len(this.cache) > this.size {
            // remove tail
            rmd := this.removeTail()
            delete(this.cache, rmd.key)
        }
    } else {
        node.val = value
        this.moveToHead(node)
    }
}

func (this *LRUCache) addToHead(node *MyNode) {
    node.prev = this.head
    node.next = this.head.next
    
    this.head.next.prev = node
    this.head.next = node
}

func (this *LRUCache) removeNode(node *MyNode) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

func(this *LRUCache) moveToHead(node *MyNode) {
    this.removeNode(node)
    this.addToHead(node)
}

func (this *LRUCache) removeTail() *MyNode {
    node := this.tail.prev
    this.removeNode(node)
    return node
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
```

