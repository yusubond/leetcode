## 207 课程表-中等

题目：

你这个学期必须选修 `numCourses` 门课程，记为 `0` 到 `numCourses - 1` 。

在选修某些课程之前需要一些先修课程。 先修课程按数组 `prerequisites` 给出，其中 `prerequisites[i] = [ai, bi]` ，表示如果要学习课程 `ai` 则 **必须** 先学习课程 `bi` 。

- 例如，先修课程对 `[0, 1]` 表示：想要学习课程 `0` ，你需要先完成课程 `1` 。

请你判断是否可能完成所有课程的学习？如果可以，返回 `true` ；否则，返回 `false` 。



> **示例 1：**
>
> ```
> 输入：numCourses = 2, prerequisites = [[1,0]]
> 输出：true
> 解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
> ```
>
> **示例 2：**
>
> ```
> 输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
> 输出：false
> 解释：总共有 2 门课程。学习课程 1 之前，你需要先完成课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
> ```



**分析：**

把每门课看成顶点，先修关系 `[a, b]` 看成有向边 `b -> a`。能否完成所有课程 ⟺ 这个有向图是 **DAG（无环）**。

所以本题等价于**有向图判环**：只要能做一次完整的拓扑排序（访问到全部顶点），就无环、返回 `true`；否则有环、返回 `false`。

> 本题是 [210 课程表 II](./No.210_课程表2.md) 的判定版本——210 要输出顺序，207 只需回答能否完成。判环是关键，下面给出两种解法。

### 方法一：BFS（Kahn 算法）✅ 推荐

与 210 完全相同的框架，只是不再需要保存顺序，**计数**能出队的顶点个数即可。

```go
// date 2026/07/12
func canFinish(numCourses int, prerequisites [][]int) bool {
    inDegree := make([]int, numCourses)    // 入度
    adjacency := make([][]int, numCourses) // 邻接表
    for _, p := range prerequisites {
        a, b := p[0], p[1]
        inDegree[a]++
        adjacency[b] = append(adjacency[b], a)
    }

    // 入度为 0 的顶点先入队
    queue := make([]int, 0, numCourses)
    for i := 0; i < numCourses; i++ {
        if inDegree[i] == 0 {
            queue = append(queue, i)
        }
    }

    visited := 0
    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]
        visited++
        for _, next := range adjacency[cur] {
            inDegree[next]--
            if inDegree[next] == 0 {
                queue = append(queue, next)
            }
        }
    }
    // 全部访问到 → 无环
    return visited == numCourses
}
```

**运行示例 2（有环）：** `numCourses = 2, prerequisites = [[1,0],[0,1]]`

图：`0->1` 且 `1->0`，入度 `inDegree = [1, 1]`。没有入度为 0 的顶点，队列为空，循环一次都不进，`visited = 0`。`0 != 2` → 返回 `false`。

对比示例 1（无环）：`[[1,0]]`，入度 `inDegree = [0, 1]`。课程 0 入度为 0 先出队，把 1 的入度减到 0，1 也出队，`visited = 2 == numCourses` → 返回 `true`。

> **注意自环**：LeetCode 测试用例含 `[5,5]` 这类自环边（课程 5 是自己的前置）。它会让 `inDegree[5]` 始终 ≥ 1，永远进不了队，从而被正确判为有环。

### 方法二：DFS（三色标记判环）

三色标记：`0`=未访问，`1`=访问中（在递归栈里），`2`=已完成。DFS 过程中若遇到「访问中」的节点，说明出现回边，即存在环。

```go
func canFinishDFS(numCourses int, prerequisites [][]int) bool {
    adjacency := make([][]int, numCourses)
    for _, p := range prerequisites {
        adjacency[p[1]] = append(adjacency[p[1]], p[0]) // b -> a
    }

    color := make([]int, numCourses) // 0 未访问，1 访问中，2 已完成
    var dfs func(int) bool           // 返回 false 表示发现环
    dfs = func(node int) bool {
        if color[node] == 1 {
            return false // 回边，存在环
        }
        if color[node] == 2 {
            return true // 已处理，跳过
        }
        color[node] = 1
        for _, next := range adjacency[node] {
            if !dfs(next) {
                return false
            }
        }
        color[node] = 2
        return true
    }

    for i := 0; i < numCourses; i++ {
        if !dfs(i) {
            return false
        }
    }
    return true
}
```

> 与 210 的 DFS 相比，这里**不需要收集顺序**，所以省去了栈和翻转——发现环立即返回 `false`，全部跑完返回 `true`。

### 复杂度分析

| 方法 | 时间 | 空间 |
|------|------|------|
| BFS（Kahn） | O(V+E) | O(V+E) |
| DFS | O(V+E) | O(V+E) |

每个顶点和每条边都只访问有限次。两种方法等价，面试中 BFS 更直观，是首选。

### 相邻题

- [210 课程表 II](./No.210_课程表2.md)：同一模型，要求输出一种学习顺序（返回数组）。
- [1462 课程表 IV](./No.1462_课程表4.md)：多次查询先修关系，需在拓扑排序中维护可达性（传递闭包）。
- [1136 平行课程](https://leetcode.cn/problems/parallel-courses/)：拓扑排序 + 分层，求最少学期数。
