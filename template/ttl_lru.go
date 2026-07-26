package template

import "time"

// TTLLRUCache 是 LRUCache 的「纯惰性过期」扩展版本。
//
// 与 lru_cache_ttl.go（最小堆 + 主动清理）相比，本实现只用「惰性删除」一种策略：
//   - 只有 Get 命中某个 key 时才检查它是否过期，过期则当场删除并返回 -1；
//   - Put 超容量时按 LRU 直接淘汰尾部，不主动扫描过期节点。
//
// 优点：实现极简，仅 map + 双向链表，没有额外的堆与版本号，是面试中最常见的写法。
// 代价：过期但未被再次访问的节点会一直占用容量，直到被 LRU 挤出尾部或偶然被 Get 命中。

type TTLNode struct {
	key, val    int
	prev, next  *TTLNode
	expiredTime time.Time // 绝对过期时刻
}

type TTLLRUCache struct {
	capacity   int
	ttl        time.Duration
	head, tail *TTLNode
	cache      map[int]*TTLNode
	now        func() time.Time // 可注入时钟，默认 time.Now，便于测试
}

// NewTTLLRUCache capacity 为容量，ttl 为每个条目的存活时长
func NewTTLLRUCache(capacity int, ttl time.Duration) *TTLLRUCache {
	ca := &TTLLRUCache{
		capacity: capacity,
		ttl:      ttl,
		head:     &TTLNode{},
		tail:     &TTLNode{},
		cache:    make(map[int]*TTLNode, capacity),
		now:      time.Now,
	}
	ca.head.next = ca.tail
	ca.tail.prev = ca.head
	return ca
}

// Get 命中且未过期：提到表头并返回 value；命中但已过期：惰性删除后返回 -1；不存在：返回 -1。
// 采用「访问续期」语义：命中会顺带刷新过期时刻，因此持续被访问的活跃数据不会过期。
func (ca *TTLLRUCache) Get(key int) int {
	node, ok := ca.cache[key]
	if !ok {
		return -1
	}

	now := ca.now()
	// 惰性删除：命中即检查，过期当场清除
	if now.After(node.expiredTime) {
		ca.removeNode(node)
		delete(ca.cache, key)
		return -1
	}

	ca.moveToHead(node)
	node.expiredTime = now.Add(ca.ttl) // 访问续期
	return node.val
}

// Put 存在则更新值并重置过期时刻；不存在则新建。超容量时按 LRU 淘汰尾部。
func (ca *TTLLRUCache) Put(key int, value int) {
	now := ca.now()
	if node, ok := ca.cache[key]; ok {
		// 命中即更新；即便该 key 已过期（惰性策略下仍留在 map 中），也按命中处理：
		// 更新值并重置过期时刻，相当于「复活」。
		node.val = value
		node.expiredTime = now.Add(ca.ttl)
		ca.moveToHead(node)
		return
	}

	node := &TTLNode{key: key, val: value, expiredTime: now.Add(ca.ttl)}
	ca.cache[key] = node
	ca.addToHead(node)

	// 超容量：纯惰性策略下直接按 LRU 删尾部（尾部最久未访问），不额外判断是否过期；
	// 过期但仍在链表中部的节点不会被这里清理，待其被 Get 命中或被 LRU 挤到尾部时才淘汰。
	if len(ca.cache) > ca.capacity {
		rmd := ca.removeTail()
		delete(ca.cache, rmd.key)
	}
}

func (ca *TTLLRUCache) addToHead(node *TTLNode) {
	node.next = ca.head.next
	node.prev = ca.head

	ca.head.next.prev = node
	ca.head.next = node
}

func (ca *TTLLRUCache) removeNode(node *TTLNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (ca *TTLLRUCache) moveToHead(node *TTLNode) {
	ca.removeNode(node)
	ca.addToHead(node)
}

func (ca *TTLLRUCache) removeTail() *TTLNode {
	node := ca.tail.prev
	ca.removeNode(node)
	return node
}
