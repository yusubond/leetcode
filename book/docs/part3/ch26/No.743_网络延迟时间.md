## 743 网络延迟时间-中等

题目：

有 `n` 个网络节点，标记为 `1` 到 `n`。给你一个列表 `times`，表示信号经过**有向边**的传递时间。`times[i] = (ui, vi, wi)`，其中 `ui` 是源节点，`vi` 是目标节点，`wi` 是一个信号从源节点传递到目标节点的时间。

现在，从某个节点 `k` 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 `-1`。



> **示例 1：**
>
> ```
> 输入：times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
> 输出：2
> 解释：第 1 秒，信号到达节点 1 和节点 3；第 2 秒，信号由节点 3 到达节点 4。
> ```
>
> **示例 2：**
>
> ```
> 输入：times = [[1,2,1]], n = 2, k = 1
> 输出：1
> ```
>
> **示例 3：**
>
> ```
> 输入：times = [[1,2,1]], n = 2, k = 2
> 输出：-1
> 解释：信号只能从节点 2 出发，但没有任何边从节点 2 出去，节点 1 收不到信号。
> ```



**分析：**

信号总是沿着**最短路径**最快到达每个节点。因此「所有节点都收到信号」的时刻，等于起点 `k` 到达**最远节点**的最短距离：

$$ans = \max_{i \in [1,n]} dist[k][i]$$

若存在某个节点 `dist[k][i] == ∞`（不可达），则返回 `-1`。

这正是一个带权有向图、**边权非负**的**单源最短路径（SSSP）**问题，标准解法是 **Dijkstra**。下面给出三种实现。

### 方法一：Dijkstra + 小顶堆 ✅ 推荐

核心思想：**贪心 + 优先队列**。维护起点到各点的最短距离 `dist`，每次从堆中取出**当前距离最小**的未确定节点 `u`，用它去松弛（relax）邻居 `v`：若 `dist[u] + w(u,v) < dist[v]`，则更新 `dist[v]` 并入堆。用堆优化后，每条边最多被松弛并入堆一次。

注意出堆时的「陈旧记录」过滤：堆中可能存在同一节点的多个旧距离，当 `cur.d > dist[cur.u]` 时直接跳过，避免重复处理。

```go
// date 2026/07/13
import (
	"container/heap"
	"math"
)

func networkDelayTime(times [][]int, n int, k int) int {
	type edge struct{ to, w int }
	graph := make([][]edge, n+1) // 节点编号 1..n
	for _, t := range times {
		graph[t[0]] = append(graph[t[0]], edge{t[1], t[2]})
	}

	const inf = math.MaxInt32 / 2
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = inf
	}
	dist[k] = 0

	h := &minHeap{}
	heap.Init(h)
	heap.Push(h, state{0, k})

	for h.Len() > 0 {
		cur := heap.Pop(h).(state)
		if cur.d > dist[cur.u] {
			continue // 陈旧记录，已有更短路径，跳过
		}
		for _, e := range graph[cur.u] {
			if nd := cur.d + e.w; nd < dist[e.to] {
				dist[e.to] = nd
				heap.Push(h, state{nd, e.to})
			}
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		if dist[i] == inf {
			return -1 // 有节点不可达
		}
		if dist[i] > ans {
			ans = dist[i]
		}
	}
	return ans
}

type state struct{ d, u int }
type minHeap []state

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].d < h[j].d }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(state)) }
func (h *minHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}
```

### 方法二：Dijkstra + 数组（O(V²)）

**算法本质：贪心 + 「已确定集」。** Dijkstra 维护一个关键不变量：每一步从「未确定」的节点中挑出 `dist` 最小的那个 `u`，此时 `dist[u]` 已经是该节点**最终**的最短距离，可以「盖棺定论」移入已确定集，再用它去更新（松弛）邻居。重复 n 轮即可。

> **为什么挑出最小 `dist` 的节点就能盖棺定论？** 这依赖**边权非负**的前提——因为所有边权 ≥ 0，任何经由别的未确定节点绕路到达 `u` 的路径，长度都不可能小于 `dist[u]`（绕路只会更长或相等），所以 `dist[u]` 不可能再被更新变小。**这也正是 Dijkstra 不能处理负权边的根本原因**（若有负权边，绕路反而更短，贪心选择性质被破坏）。

**与堆版本唯一的区别，在于「如何挑出最小 dist 的节点」：**

- 堆版本：用小顶堆维护，挑最小是 O(log V)，总体 O((V+E)log V)，适合**稀疏图**（E 较小）；
- 数组版本：每轮线性扫描一遍 `dist` 找最小，挑最小是 O(V)，总体 O(V²)，适合**稠密图**或 `n` 较小时。

**为什么数组版更适合稠密图？** 稠密图 E≈V²，堆版本会退化到 O(V²log V)，反而比数组的 O(V²) 更慢；而且数组版省去了手写堆的代码量。反过来，稀疏图 E≈V 时堆版本 O(Vlog V) 占优。本题 n≤100，O(V²) 也就 10⁴ 量级，常数极小，写数组版既快又简洁。

**完整流程：**

1. **建图**：邻接矩阵 `g[u][v]` 存 `u→v` 的边权，无边处填 `inf`（便于直接相加做比较，无需特判）。
2. **初始化**：`dist` 全填 `inf`（暂时认为都不可达），起点 `dist[k]=0`；`visited` 全 false。
3. **主循环 n 轮**，每轮做三件事：
   - **找最小**：线性扫描「未访问且 `dist` 最小」的节点 `u`；若一个都找不到（`u==-1`，说明剩余节点都不可达），提前 `break`。
   - **盖棺定论**：`visited[u]=true`，`u` 的最短距离就此确定。
   - **松弛**：用 `u` 作中转，遍历所有未访问邻居 `v`，若 `dist[u]+g[u][v] < dist[v]` 则更新。
4. **汇总答案**：遍历 `dist`，遇 `inf` 返回 `-1`，否则取最大值。

> **复杂度来源：** 外层 n 轮，每轮「找最小」O(V) +「松弛」O(V)，合计 O(V²)；邻接矩阵占 O(V²) 空间。

```go
func networkDelayTimeArray(times [][]int, n int, k int) int {
	const inf = math.MaxInt32 / 2 // 除以 2，防止后续 dist[u]+g[u][v] 相加时溢出

	// —— 1) 建图：邻接矩阵 g[u][v] = u->v 的边权；无边处填 inf ——
	g := make([][]int, n+1)
	for i := range g {
		g[i] = make([]int, n+1)
		for j := range g[i] {
			g[i][j] = inf
		}
	}
	for _, t := range times {
		u, v, w := t[0], t[1], t[2]
		g[u][v] = w
	}

	// —— 2) 初始化 dist 与 visited ——
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = inf // 暂时认为起点到各点都不可达
	}
	dist[k] = 0                   // 起点到自身距离为 0
	visited := make([]bool, n+1)  // 标记节点是否已「盖棺定论」（最短距离已确定）

	// —— 3) 主循环：共 n 轮，每轮确定一个节点的最短距离 ——
	for i := 1; i <= n; i++ {
		// 3.1 线性扫描「未确定」节点，挑出 dist 最小的 u
		u := -1
		for j := 1; j <= n; j++ {
			if !visited[j] && (u == -1 || dist[j] < dist[u]) {
				u = j
			}
		}
		if u == -1 {
			break // 剩余未确定节点的 dist 都是 inf（不可达），提前结束
		}

		// 3.2 盖棺定论：依赖「边权非负」，此时 dist[u] 已是其最终最短距离
		visited[u] = true

		// 3.3 松弛：尝试用 u 作为中转，更新其未确定邻居 v 的最短距离
		for v := 1; v <= n; v++ {
			if !visited[v] && dist[u]+g[u][v] < dist[v] {
				dist[v] = dist[u] + g[u][v]
			}
		}
	}

	// —— 4) 汇总：所有节点都收到信号的时刻 = 最远节点的最短距离 ——
	ans := 0
	for i := 1; i <= n; i++ {
		if dist[i] == inf {
			return -1 // 存在不可达节点，信号无法传遍全网
		}
		if dist[i] > ans {
			ans = dist[i]
		}
	}
	return ans
}
```

### 方法三：Bellman-Ford（n−1 轮边松弛）

直接对**边集**重复松弛 n−1 轮（最短路径最多 n−1 条边）。无需建图、代码最短，且可处理负权边（本题边权非负，Dijkstra 更优）。加一个 `updated` 标记，若某轮无任何更新则提前收敛。

```go
func networkDelayTimeBF(times [][]int, n int, k int) int {
	const inf = math.MaxInt32 / 2
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = inf
	}
	dist[k] = 0

	for i := 1; i < n; i++ { // 松弛 n-1 轮
		updated := false
		for _, t := range times {
			u, v, w := t[0], t[1], t[2]
			if dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				updated = true
			}
		}
		if !updated {
			break // 提前收敛
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		if dist[i] == inf {
			return -1
		}
		if dist[i] > ans {
			ans = dist[i]
		}
	}
	return ans
}
```

**运行示例 1：** `times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2`

图：`2→1(w1)`、`2→3(w1)`、`3→4(w1)`，从节点 2 出发，方法一的堆过程如下：

| 出堆节点 u (距离 d) | 松弛邻居结果 |
|---|---|
| 2 (0) | dist[1]=1，dist[3]=1 |
| 1 (1) | 无改进 |
| 3 (1) | dist[4]=2 |
| 4 (2) | 无改进 |

最终 `dist = [_, 1, 0, 1, 2]`，最大值 = **2**，与示例一致。

### 复杂度分析

| 方法 | 时间 | 空间 | 适用场景 |
|------|------|------|----------|
| Dijkstra + 堆 | O((V+E)log V) | O(V+E) | 稀疏图首选 |
| Dijkstra + 数组 | O(V²) | O(V²) | 稠密图 / n 很小 |
| Bellman-Ford | O(VE) | O(V) | 可处理负权、代码最短 |

三种方法都遍历了所有顶点和边有限次。本题 n≤100、E≤600，三种都能轻松通过。

**如何选择：** 面试中一般写**方法一（堆）**；若题目明确 n 很小且不想手写堆，用方法二；Bellman-Ford 在「存在负权边」「只能对边松弛」等场景才有不可替代性。

### 相邻题

- [1631 最小体力消耗路径](https://leetcode.cn/problems/path-with-minimum-effort/)：Dijkstra 变体，把松弛条件改成「路径上最大边权最小」。
- [1514 概率最大路径](https://leetcode.cn/problems/path-with-maximum-probability/)：Dijkstra 变体，松弛方向取乘积最大。
- [1976 到达目的地的方案数](https://leetcode.cn/problems/number-of-ways-to-arrive-at-destination/)：Dijkstra 求最短路的同时统计方案数。
- [505 迷宫 II](https://leetcode.cn/problems/the-maze-ii/)（会员）：在网格上跑 Dijkstra 求最短滚动距离。
