package template

import (
	"container/heap"
	"time"
)

// TTLCache 是在 LRUCache 基础上增加过期时间（TTL，Time To Live）的扩展实现。
// 与 LRU 的差别：每个 key-value 自带存活时长 ttl，超过后视为过期应被淘汰。
//
// 过期数据何时删除？采用与 Redis 一致的「惰性 + 主动」组合策略：
//   - 惰性删除：Get/Put 命中某 key 时检查其 expireAt，过期则当场删除；
//   - 主动清理：用最小堆按 expireAt 排序，Put 导致超容量时批量弹出过期项。
//
// 堆不支持高效删除任意节点，采用「懒删除」：节点带 version 版本号，堆里存
// (expireAt, key, version)；弹出堆顶时若 key 已不在 map 或 version 不一致，
// 说明是已被 LRU 淘汰/已更新的脏条目，直接丢弃。

type (
	// ttlNode 双向链表节点，相比 CacheNode 多了 expireAt 与 version
	ttlNode struct {
		key, val  int
		version   int64     // 版本号：堆的懒删除校验
		expireAt  time.Time // 绝对过期时刻
		prev, next *ttlNode
	}

	// ttlHeapEntry 最小堆条目
	ttlHeapEntry struct {
		key      int
		version  int64
		expireAt time.Time
	}

	ttlHeap []ttlHeapEntry

	// TTLCache define
	TTLCache struct {
		capacity   int
		ttl        time.Duration
		cache      map[int]*ttlNode
		head, tail *ttlNode
		expHeap    *ttlHeap
		now        func() time.Time // 可注入时钟，便于测试/演示
	}
)

func (h ttlHeap) Len() int            { return len(h) }
func (h ttlHeap) Less(i, j int) bool  { return h[i].expireAt.Before(h[j].expireAt) }
func (h ttlHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *ttlHeap) Push(x any)         { *h = append(*h, x.(ttlHeapEntry)) }
func (h *ttlHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// NewTTLCache define
func NewTTLCache(capacity int, ttl time.Duration) *TTLCache {
	c := &TTLCache{
		capacity: capacity,
		ttl:      ttl,
		cache:    make(map[int]*ttlNode, capacity),
		expHeap:  &ttlHeap{},
		head:     &ttlNode{},
		tail:     &ttlNode{},
		now:      time.Now,
	}
	c.head.next = c.tail
	c.tail.prev = c.head
	return c
}

// isExpired 节点是否已过期：expireAt <= now 即过期
func (n *ttlNode) isExpired(now time.Time) bool { return !n.expireAt.After(now) }

// Get define
// 命中且未过期：提到表头并返回 value；命中但已过期：惰性删除后返回 -1；不存在：返回 -1
func (c *TTLCache) Get(key int) int {
	node, ok := c.cache[key]
	if !ok {
		return -1
	}
	if node.isExpired(c.now()) {
		// 命中但已过期：惰性删除，返回未命中
		c.removeNode(node)
		delete(c.cache, key)
		// 堆中的旧条目留作脏数据，待主动清理时丢弃
		return -1
	}
	c.moveToHead(node)
	return node.val
}

// Put define
// 存在且未过期：更新值并重置过期时间（version++ 后入堆）；否则当作新插入。
// 超容量时先主动清理所有过期项，仍超容量再按 LRU 淘汰表尾。
func (c *TTLCache) Put(key, value int) {
	now := c.now()
	if node, ok := c.cache[key]; ok && !node.isExpired(now) {
		// 命中且未过期：更新值，重置 TTL，version++
		node.val = value
		node.expireAt = now.Add(c.ttl)
		node.version++
		heap.Push(c.expHeap, ttlHeapEntry{key, node.version, node.expireAt})
		c.moveToHead(node)
		return
	}

	// 不存在 或 已过期：先清掉旧节点（已过期的情况），再当新节点插入
	if node, ok := c.cache[key]; ok {
		c.removeNode(node)
		delete(c.cache, key)
	}
	node := &ttlNode{key: key, val: value, expireAt: now.Add(c.ttl)}
	c.cache[key] = node
	c.addToHead(node)
	heap.Push(c.expHeap, ttlHeapEntry{key, node.version, node.expireAt})

	// 超容量：先主动清理过期项，仍超则按 LRU 淘汰表尾
	if len(c.cache) > c.capacity {
		c.evictExpired(now)
	}
	if len(c.cache) > c.capacity {
		tail := c.removeTail()
		delete(c.cache, tail.key)
	}
}

// evictExpired 弹出堆中所有已过期（且未脏）的节点并从缓存删除
func (c *TTLCache) evictExpired(now time.Time) {
	for c.expHeap.Len() > 0 {
		top := (*c.expHeap)[0]
		if top.expireAt.After(now) {
			break // 堆顶未过期，后面的 expireAt 更大，不可能过期
		}
		heap.Pop(c.expHeap)
		node, ok := c.cache[top.key]
		if !ok || node.version != top.version {
			continue // 脏条目：已被 LRU 淘汰或已被更新，丢弃
		}
		c.removeNode(node)
		delete(c.cache, top.key)
	}
}

// Len 返回当前缓存条目数（含可能已过期但尚未惰性清理的项）
func (c *TTLCache) Len() int { return len(c.cache) }

// --- 双向链表辅助函数（与 LRU 一致） ---

func (c *TTLCache) addToHead(node *ttlNode) {
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}

func (c *TTLCache) removeNode(node *ttlNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *TTLCache) moveToHead(node *ttlNode) {
	c.removeNode(node)
	c.addToHead(node)
}

func (c *TTLCache) removeTail() *ttlNode {
	node := c.tail.prev
	c.removeNode(node)
	return node
}
