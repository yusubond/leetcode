## TTL LRU 缓存（带超时时间）- 扩展题

> 本篇是 [146. LRU 缓存](./No.146_LRU缓存.md) 的姊妹篇，在其基础上引入「过期时间」维度。146 的淘汰只看「访问时间」，本篇额外引入「存活时长 TTL」——超时即视为失效。这是工程中（Redis、Caffeine、Guava Cache 等）的真实需求，也是 LRU 类题目的常见进阶问法。
>
> 题号说明：这是 LeetCode 之外的设计扩展题（非标准编号）。本篇给出**两种方案对比**：方案一「纯惰性删除」（面试最常用，对应 `template/ttl_lru.go`）与方案二「惰性 + 最小堆主动清理」（工程完整，对应 `template/lru_cache_ttl.go`）。

题目：

请你设计并实现一个**支持过期时间（TTL，Time To Live）** 的 LRU 缓存。在 [146. LRU 缓存](./No.146_LRU缓存.md) 的两个约束（容量淘汰 + 高效读写）之外，每个 key-value 还带有存活时长 `ttl`：一旦存活超过 `ttl`，该条目视为过期，应当被淘汰。

实现一个 TTL 缓存类：

- `Constructor(capacity, ttl)` 以容量 `capacity`（条数）和存活时长 `ttl` 初始化缓存。
- `int get(key)`：若 key **不存在**或**已过期**，返回 `-1`（视为未命中）；命中则返回 value。
- `void put(key, value)`：插入或更新，并把该 key 的过期时刻（重）置为「当前时间 + ttl」。插入导致数量超过 `capacity` 时需淘汰数据。

有两个语义点需要特别说明——它们正是本篇两种方案的分野：

- **get 是否刷新过期时间**：刷新即为「滑动过期」（持续被访问的活跃数据不过期），不刷新即为「固定生命周期」。
- **超容量时是否优先淘汰过期数据**：优先淘汰能更及时回收内存，但需要额外结构。

`get` / `put` 的时间复杂度应尽量接近 O(1)（容许对数级别）。



分析：

LRU 部分完全复用 146 的「map + 双向链表」：map 做 O(1) 查找，双向链表按访问时间维护顺序，表头更新、表尾淘汰。本篇的核心新问题是——**过期的数据，什么时候删？**

这有两条路线：

- **惰性删除（lazy / passive）**：只有当 `get`/`put` 命中某个 key 时，才检查它是否过期、过期才删。优点是实现极简、`get` 严格 O(1)；缺点是过期却长期不被访问的 key 会一直占着内存，直到被 LRU 挤出尾部。
- **主动删除（active / proactive）**：额外维护一个按过期时间排序的结构，定期或在写入前批量清理已过期数据。优点是内存回收及时；缺点是需要额外结构与清理成本。

成熟的实现（如 Redis）两者都用：惰性兜底 + 定期采样清理。由这两条路线衍生出本篇的两种典型方案，下文逐一给出并对比：

- **方案一（纯惰性）**：只用惰性删除，且 `get` 续期。极简，是面试中最常见的写法。
- **方案二（惰性 + 最小堆主动清理）**：在惰性基础上加最小堆批量清过期，`get` 不续期。工程上更完整。



## 方案一：纯惰性删除（面试最常用）

只采用「惰性删除」：`get` 命中时检查过期，过期当场删；`put` 超容量时按 LRU 直接删尾部，**不主动扫描过期节点**。并采用「访问续期」语义——`get` 命中顺带刷新过期时刻，因此持续被访问的活跃数据不会过期。

优点是极简：相比 146，节点只多了一个 `expiredTime` 字段，没有额外的堆与版本号。代价是：过期但未被再次访问的节点会一直占用容量，直到被 LRU 挤出尾部或偶然被 `get` 命中。

```go
// date 2026/07/26
// 方案一：纯惰性删除（对应 template/ttl_lru.go）
import "time"

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

// get：命中且未过期→提到表头并返回 value；命中但已过期→惰性删除后返回 -1；不存在→返回 -1。
// 采用「访问续期」语义：命中顺带刷新过期时刻，活跃数据不会过期。
func (ca *TTLLRUCache) Get(key int) int {
	node, ok := ca.cache[key]
	if !ok {
		return -1
	}
	now := ca.now()
	if now.After(node.expiredTime) { // 惰性删除：命中即检查
		ca.removeNode(node)
		delete(ca.cache, key)
		return -1
	}
	ca.moveToHead(node)
	node.expiredTime = now.Add(ca.ttl) // 访问续期
	return node.val
}

// put：存在则更新值并重置过期时刻（即便该 key 已过期，也按命中「复活」）；不存在则新建。
// 超容量时按 LRU 直接删尾部，不主动判断是否过期。
func (ca *TTLLRUCache) Put(key int, value int) {
	now := ca.now()
	if node, ok := ca.cache[key]; ok {
		node.val = value
		node.expiredTime = now.Add(ca.ttl)
		ca.moveToHead(node)
		return
	}
	node := &TTLNode{key: key, val: value, expiredTime: now.Add(ca.ttl)}
	ca.cache[key] = node
	ca.addToHead(node)
	if len(ca.cache) > ca.capacity {
		rmd := ca.removeTail()
		delete(ca.cache, rmd.key)
	}
}

// --- 双向链表辅助函数（与 146 一致） ---
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
```

**复杂度分析：**

- 时间复杂度：`get` / `put` 均为 **O(1)**，只有一次 map 查找 + 常数次链表指针操作 + O(1) 的过期判断。
- 空间复杂度：**O(capacity)**，map 与双向链表各存一份节点。



## 方案二：惰性 + 最小堆主动清理（工程完整）

在方案一的惰性删除之上，再加一个**按过期时间排序的最小堆**：`get` 命中时仍做惰性检查（不续期），`put` 触发超容量时则用最小堆**批量弹出所有过期项**，确保「优先淘汰过期数据」而非误伤未过期的热数据。

**两个关键点**（对应 146 的两个关键点，体会 TTL 带来的新增量）：

1. **过期数据何时删？** `get` 命中时做惰性检查（O(1)）保证读路径不退化；`put` 触发超容量时，用最小堆批量弹出所有过期项，确保「优先淘汰过期数据」而非误伤未过期的热数据。

2. **最小堆如何删除「已被 LRU 淘汰 / 已被更新」的节点？** 堆不支持 O(log n) 以下删除任意节点。采用**懒删除（lazy deletion）**：节点带 `version` 版本号，每次 `put` 更新时 `version++`；堆里存的是 `(expireAt, key, version)` 三元组。弹出堆顶时，若该 `key` 已不在 map（被 LRU 淘汰）或 `version` 不一致（被更新过），即为脏条目，直接丢弃，继续弹下一个。这样避免了「堆中真删除」的高昂代价。

数据结构相比 146 的增量：节点新增 `expireAt`（绝对过期时刻）与 `version`（版本号）；缓存新增一个最小堆 `expHeap`（按 `expireAt` 排序）和一个可注入的时钟 `now`（便于测试与演示过期）。

```go
// date 2026/07/26
// 方案二：惰性 + 最小堆主动清理（对应 template/lru_cache_ttl.go）
import (
	"container/heap"
	"time"
)

// ttlNode 相比 146 的节点，多了 expireAt 与 version
type ttlNode struct {
	key, val   int
	version    int64     // 版本号：堆的懒删除校验
	expireAt   time.Time // 绝对过期时刻
	prev, next *ttlNode
}

type ttlHeapEntry struct {
	key      int
	version  int64
	expireAt time.Time
}

type ttlHeap []ttlHeapEntry

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

type TTLCache struct {
	capacity   int
	ttl        time.Duration
	cache      map[int]*ttlNode
	head, tail *ttlNode
	expHeap    *ttlHeap
	now        func() time.Time // 可注入时钟，便于测试/演示；默认 time.Now
}

func TTLConstructor(capacity int, ttl time.Duration) *TTLCache {
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

// isExpired：expireAt <= now 即过期
func (n *ttlNode) isExpired(now time.Time) bool { return !n.expireAt.After(now) }

// get：命中且未过期→提到表头并返回 value；命中但已过期→惰性删除后返回 -1；不存在→返回 -1
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

// put：存在且未过期→更新值并重置过期时间（version++ 后入堆）；否则当作新插入。
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

// --- 双向链表辅助函数（与 146 一致） ---
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
```



**运行示例**（capacity = 2，ttl = 3s，注入可控时钟；表中「过期时刻」= 写入时刻 + 3）：

| 时刻 | 操作 | 链表（头→尾） | map | 返回 | 说明 |
|------|------|--------------|-----|------|------|
| t=0 | put(1,1) | 1 | {1} | — | 插入，1 过期时刻=3 |
| t=1 | put(2,2) | 2 1 | {1,2} | — | 插入，2 过期时刻=4 |
| t=2 | get(1) | 1 2 | {1,2} | 1 | 未过期，提到头部 |
| t=3 | put(3,3) | 3 2 | {2,3} | — | 超容量→**先 evictExpired 清掉过期的 1**，再插 3；2 未过期被保留 |
| t=3 | get(1) | 3 2 | {2,3} | -1 | 1 已被主动清理 |
| t=3 | get(2) | 2 3 | {2,3} | 2 | 2 未过期（过期时刻=4） |
| t=7 | get(3) | — | {} | -1 | 3 过期时刻=6 已过期，**惰性删除** |

两个值得对比 146 的瞬间：

- **t=3 put(3,3)**：同样的操作序列，146 会淘汰「最久未使用」的尾部 2；而本方案先批量清掉过期的 1，于是 2 被保留。这正是「优先淘汰过期数据」带来的差别。
- **t=7 get(3)**：即便容量没满，过期数据也会在读时被惰性删除——这是 TTL 独有、146 完全没有的语义。



**复杂度分析：**

- 时间复杂度：`get` 为 **O(1)**，只涉及一次 map 查找 + 常数次链表指针操作 + O(1) 的过期判断；`put` 为**均摊 O(log n)**，主要开销是 `heap.Push` 的 O(log n)，而 `evictExpired` 中每个过期项的弹出均可摊到它当初入堆的那次 `put`，整体均摊仍是 O(log n)。
- 空间复杂度：map 与双向链表 **O(capacity)**；最小堆 **O(P)**，其中 P 为累计 `put` 次数（含待清理的脏条目）。脏条目在轮到堆顶时被丢弃，长期规模与有效条目同阶。若需要把堆空间严格压到 O(capacity)，可改用「索引堆」（节点记录自身在堆中的下标）实现 O(log n) 的真删除。



## 两种方案对比

| 维度 | 方案一 纯惰性 | 方案二 惰性 + 最小堆 |
|------|--------------|---------------------|
| 数据结构 | map + 双向链表 | 再加 最小堆 + version |
| get 复杂度 | O(1) | O(1) |
| put 复杂度 | O(1) | 均摊 O(log n) |
| get 是否续期 | 是（滑动过期） | 否（固定生命周期） |
| 过期回收时机 | 仅 get 命中时 | get 命中 + put 超容量时批量 |
| 过期未访问节点 | 占内存，直到被 LRU 挤出 | put 时主动清理 |
| 实现复杂度 | 极简 | 中等（堆 + 懒删除） |
| 适用场景 | 面试 / 小规模缓存 | 工程 / 需及时回收内存 |
| 对应文件 | `template/ttl_lru.go` | `template/lru_cache_ttl.go` |

**关键差异**（capacity = 2，ttl = 3：t=0 先后 put(1,1)、put(2,2)，两者过期时刻都 = 3；t=4 时 1、2 均已过期，此时 put(3,3) 触发超容量）：

- **方案一**：put(3,3) 直接按 LRU 删尾部 1，但 2 同样过期却仍留在缓存里占着一个位置——要等下一次 put 才可能被挤出。
- **方案二**：put(3,3) 先 evictExpired 把过期的 1、2 一次性清掉，再插入 3，最终缓存里只剩 3。

换言之，方案一**被动**地随访问 / 淘汰清理过期数据，方案二**主动**地在写入时批量回收。面试中若只要求「能用、好写」，方案一足够；若追问「过期数据长期占内存怎么办」，再引出方案二。



**相邻题目：**

- **[146. LRU 缓存](./No.146_LRU缓存.md)**（本篇基础）：去掉 TTL，只按访问时间淘汰，是理解本篇的前置。`template/lru_cache.go` 中有实现。
- **460. LFU 缓存**：在「时间」与「存活时长」之外，再叠加「访问频次」维度，是 LRU 的另一条进阶路线。`template/lfu_cache.go` 中有实现。
- **1797. 设计一个验证系统**（Design Authentication Manager）：纯 TTL 机制（token 到期），不含 LRU，可作为「仅过期、不按访问淘汰」的简化对照，体会两者各自解决的一半问题。
