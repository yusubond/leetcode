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

**两个关键点：**

1. 为什么必须用「双向」链表？淘汰时要删尾节点（`removeTail` 需要 O(1)），`moveToHead` 也需要先把节点摘下来再插头，而摘下操作 `node.prev.next = node.next` 必须知道前驱。单链表无法在 O(1) 内摘除任意节点。

2. 链表节点为什么也要存 `key`？这是最容易漏的一点。淘汰尾部节点时，除了从链表中删除，还要从 `map` 中删除对应条目，而 `map` 用 key 索引。因此节点必须记住自己的 key，否则删尾时拿不到 key 就无法更新 map。

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



**运行示例**（capacity = 2）：

| 操作 | 链表（头→尾） | map | 说明 |
|------|--------------|-----|------|
| put(1,1) | 1 | {1} | 插入 |
| put(2,2) | 2 1 | {1,2} | 插入 |
| get(1)=1 | 1 2 | {1,2} | 1 被访问，提到头部 |
| put(3,3) | 3 1 | {1,3} | 超容量，淘汰尾部 2 |
| get(2)=-1 | 3 1 | {1,3} | 2 已被逐出 |
| put(4,4) | 4 3 | {3,4} | 超容量，淘汰尾部 1 |
| get(1)=-1 | 4 3 | {3,4} | 1 已被逐出 |
| get(3)=3 | 3 4 | {3,4} | 3 被访问，提到头部 |
| get(4)=4 | 4 3 | {3,4} | 4 被访问，提到头部 |

**复杂度分析：**

- 时间复杂度：`get` / `put` 均为 **O(1)**。两者只涉及一次 map 查找 + 常数次链表指针操作；淘汰也只是删尾节点 + 删 map 条目，均为常数时间。
- 空间复杂度：**O(capacity)**。map 与双向链表各存一份节点（指针），元素总数始终不超过 capacity。

**相邻题目：**

- **460. LFU 缓存**（最不经常使用）：LRU 的进阶版，在「时间」维度之外再叠加「访问频次」维度，`template/lfu_cache.go` 中有实现。
- **432. 全 O(1) 的数据结构**：同样是「哈希表 + 双向链表」的组合思想，只不过按频次而非访问时间维护节点顺序。

