## 210 课程表 II-中等

题目：

现在你总共有 `numCourses` 门课需要选，记为 `0` 到 `numCourses - 1`。给你一个数组 `prerequisites` ，其中 `prerequisites[i] = [ai, bi]` ，表示在选修课程 `ai` 前 **必须** 先选修 `bi` 。

- 例如，想要学习课程 `0` ，你需要先完成课程 `1` ，我们用一个匹配来表示：`[0,1]` 。

返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 **任意一种** 就可以了。如果不可能完成所有课程，返回 **一个空数组** 。



> **示例 1：**
>
> ```
> 输入：numCourses = 2, prerequisites = [[1,0]]
> 输出：[0,1]
> 解释：总共有 2 门课程。要学习课程 1，你需要先完成课程 0。因此，正确的课程顺序为 [0,1] 。
> ```
>
> **示例 2：**
>
> ```
> 输入：numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
> 输出：[0,2,1,3]
> 解释：总共有 4 门课程。要学习课程 3，你应该先完成课程 1 和课程 2。并且课程 1 和课程 2 都应该排在课程 0 之后。
> 因此，一个正确的课程顺序是 [0,1,2,3] 。另一个正确的排序是 [0,2,1,3] 。
> ```
>
> **示例 3：**
>
> ```
> 输入：numCourses = 1, prerequisites = []
> 输出：[0]
> ```



**分析：**

把每门课看成顶点，先修关系 `[a, b]` 看成一条有向边 `b -> a`（先学 b 才能学 a）。问题转化为：对这个有向图做**拓扑排序**。

关键结论：能完成所有课程 ⟺ 该有向图是 **DAG（有向无环图）** ⟺ 存在拓扑序。

下面给出两种经典解法。

### 方法一：BFS（Kahn 算法）✅ 推荐

核心思想是「不断摘掉入度为 0 的叶子」：

1. 统计每个顶点的**入度** `inDegree`（有多少条边指向它，即还有几门前置课没学）；
2. 把入度为 0 的顶点入队（没有前置，可以直接学）；
3. 出队一个顶点加入结果，并「删掉」它的所有出边——后继的入度减 1；若后继入度变为 0 则入队；
4. 重复直到队空。若结果包含全部顶点，即为一个拓扑序；否则说明存在环，返回空数组。

```go
// date 2026/07/12
func findOrder(numCourses int, prerequisites [][]int) []int {
    // 建图：[a, b] 表示先修 b 才能学 a，即 b -> a
    inDegree := make([]int, numCourses)    // 入度：还有多少前置没学
    adjacency := make([][]int, numCourses) // 邻接表：学完 i 后可解锁的课
    for _, p := range prerequisites {
        a, b := p[0], p[1]
        inDegree[a]++
        adjacency[b] = append(adjacency[b], a)
    }

    // 入度为 0 的课程可以直接学，先入队
    queue := make([]int, 0, numCourses)
    for i := 0; i < numCourses; i++ {
        if inDegree[i] == 0 {
            queue = append(queue, i)
        }
    }

    order := make([]int, 0, numCourses)
    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]
        order = append(order, cur)
        for _, next := range adjacency[cur] {
            inDegree[next]--
            if inDegree[next] == 0 {
                queue = append(queue, next)
            }
        }
    }

    if len(order) == numCourses {
        return order
    }
    return []int{} // 存在环，无法完成
}
```

**运行示例 2：** `numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]`

图：`0->1, 0->2, 1->3, 2->3`，初始入度 `inDegree = [0,1,1,2]`。

| 出队节点 | 后继变化 | 队列 | order |
|---------|---------|------|-------|
| 初始 | — | `[0]` | `[]` |
| 0 | 1 的入度 1→0，2 的入度 1→0 | `[1,2]` | `[0]` |
| 1 | 3 的入度 2→1 | `[2]` | `[0,1]` |
| 2 | 3 的入度 1→0 | `[3]` | `[0,1,2]` |
| 3 | 无后继 | `[]` | `[0,1,2,3]` |

`len(order)=4 == numCourses`，返回 `[0,1,2,3]`（题目给的 `[0,2,1,3]` 也合法，差别仅在 1、2 谁先出队）。

### 方法二：DFS（三色标记 + 逆后续）

原理：DFS 中当一个顶点的**所有后继都访问完毕**时（即顶点「完成」），把它压入栈。因为后继一定先于当前顶点完成，所以**从栈顶到栈底天然是逆拓扑序，翻转即得拓扑序**。

判环用三色标记：`0`=未访问，`1`=访问中（在递归栈里），`2`=已完成。若 DFS 过程中遇到处于「访问中」的节点，说明出现回边，即存在环。

```go
func findOrderDFS(numCourses int, prerequisites [][]int) []int {
    adjacency := make([][]int, numCourses)
    for _, p := range prerequisites {
        adjacency[p[1]] = append(adjacency[p[1]], p[0]) // b -> a
    }

    color := make([]int, numCourses) // 0 未访问，1 访问中，2 已完成
    var stack []int                  // 节点完成时入栈

    var dfs func(int) bool
    dfs = func(node int) bool {
        if color[node] == 1 {
            return false // 回边，存在环
        }
        if color[node] == 2 {
            return true // 已处理，跳过
        }
        color[node] = 1 // 标记为访问中
        for _, next := range adjacency[node] {
            if !dfs(next) {
                return false
            }
        }
        color[node] = 2 // 后继全部处理完，标记已完成
        stack = append(stack, node)
        return true
    }

    for i := 0; i < numCourses; i++ {
        if !dfs(i) {
            return []int{} // 存在环
        }
    }

    // 栈是逆拓扑序，翻转得到拓扑序
    for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
        stack[i], stack[j] = stack[j], stack[i]
    }
    return stack
}
```

### 复杂度分析

| 方法 | 时间 | 空间 |
|------|------|------|
| BFS（Kahn） | O(V+E) | O(V+E) |
| DFS | O(V+E) | O(V+E) |

两种方法都遍历了每个顶点和每条边各一次。BFS 的入度数组、DFS 的颜色数组和递归栈均为 O(V)，邻接表为 O(V+E)。

**如何选择：** BFS 直观且天然按层输出，求「最小拓扑序」（字典序最小）时配合优先队列即可；DFS 思路更接近图的遍历本质，判环也很自然。面试中 BFS（Kahn）更常考。

### 相邻题

- [207 课程表](../../part2/ch14bfs/No.207_课程表.md)：模型完全相同，只需判断能否完成（返回布尔值）。BFS 中判断 `len(order)==numCourses`；DFS 中判断是否出现环即可。
- [1462 课程表 IV](./No.1462_课程表4.md)：多次查询某门课是否是另一门课的先修，可用拓扑排序过程中维护可达关系（传递闭包），或 Floyd 预处理。
- [1136 平行课程](https://leetcode.cn/problems/parallel-courses/)：拓扑排序的同时记录层数（最大层数即最少学期数），是 210 的直接扩展。

> 本题在 [第 10 章 广度优先搜索](../../part2/ch14bfs/No.210_课程表2.md) 中也有一份 BFS 版题解；拓扑排序本质是图论算法，故同时收录于本图论章节。
