## 1462 课程表4-中等

题目：

你总共需要上 `numCourses` 门课，课程编号依次为 `0` 到 `numCourses-1` 。你会得到一个数组 `prerequisite` ，其中 `prerequisites[i] = [ai, bi]` 表示如果你想选 `bi` 课程，你 **必须** 先选 `ai` 课程。

- 有的课会有直接的先修课程，比如如果想上课程 `1` ，你必须先上课程 `0` ，那么会以 `[0,1]` 数对的形式给出先修课程数对。

先决条件也可以是 **间接** 的。如果课程 `a` 是课程 `b` 的先决条件，课程 `b` 是课程 `c` 的先决条件，那么课程 `a` 就是课程 `c` 的先决条件。

你也得到一个数组 `queries` ，其中 `queries[j] = [uj, vj]`。对于第 `j` 个查询，您应该回答课程 `uj` 是否是课程 `vj` 的先决条件。

返回一个布尔数组 `answer` ，其中 `answer[j]` 是第 `j` 个查询的答案。



> **示例 1：**
>
> ![img](https://assets.leetcode.com/uploads/2021/05/01/courses4-1-graph.jpg)
>
> ```
> 输入：numCourses = 2, prerequisites = [[1,0]], queries = [[0,1],[1,0]]
> 输出：[false,true]
> 解释：课程 0 不是课程 1 的先修课程，但课程 1 是课程 0 的先修课程。
> ```
>
> **示例 2：**
>
> ```
> 输入：numCourses = 2, prerequisites = [], queries = [[1,0],[0,1]]
> 输出：[false,false]
> 解释：没有先修课程对，所以每门课程之间是独立的。
> ```
>
> **示例 3：**
>
> ![img](https://assets.leetcode.com/uploads/2021/05/01/courses4-3-graph.jpg)
>
> ```
> 输入：numCourses = 3, prerequisites = [[1,2],[1,0],[2,0]], queries = [[1,0],[1,2]]
> 输出：[true,true]
> ```



**分析：**

多次查询「u 是否为 v 的先修（含间接）」，本质是在 DAG 上求**传递闭包**。先一次性预处理出所有可达关系，之后每个查询 O(1) 回答。

> **注意边方向**：本题 `[a, b]` 表示 **a 是 b 的先修**，即边 `a -> b`（先修在前）。这与 [207](./No.207_课程表.md) / [210](./No.210_课程表2.md) 中 `[a, b]` 表示「学 a 需先学 b」方向**相反**，建图时要留意，不要照搬。

几种朴素思路为何不可行：

1. **拓扑序号比较**：拓扑排序里入度为 0 的节点是「批量」入队的，它们彼此无依赖却先后排了序，所以不能用拓扑序的相对位置判断先修关系。
2. **每个查询做一次 BFS/DFS**：查询次数多时会超时。
3. **为每个节点存「全部前置依赖」的列表**：稠密图下空间爆炸，超内存。

正解是 **拓扑排序过程中增量维护传递闭包 `isPre`**，或直接用 **Floyd 求传递闭包**。

### 方法一：拓扑排序 + 传递闭包传播 ✅ 推荐

`isPre[i][j] = true` 表示 i 是 j 的先修（含间接）。在拓扑序中处理每条边 `cur -> next` 时：

- `cur` 显然是 `next` 的直接先修：`isPre[cur][next] = true`；
- **关键传递**：凡是在拓扑序中排在 `cur` 之前的（即 `cur` 的所有先修），也都是 `next` 的先修。由于按拓扑序处理，`cur` 的先修关系此时已全部算完，所以直接把 `cur` 这一列拷给 `next`：`isPre[j][next] |= isPre[j][cur]`。

```go
// date 2026/07/12
func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
    // [a, b] 表示 a 是 b 的先修，即 a -> b
    inDegree := make([]int, numCourses)
    adjacency := make([][]int, numCourses)
    for _, p := range prerequisites {
        a, b := p[0], p[1]
        inDegree[b]++
        adjacency[a] = append(adjacency[a], b)
    }

    // isPre[i][j] = true 表示 i 是 j 的先修（含间接）
    isPre := make([][]bool, numCourses)
    for i := range isPre {
        isPre[i] = make([]bool, numCourses)
    }

    queue := make([]int, 0, numCourses)
    for i := 0; i < numCourses; i++ {
        if inDegree[i] == 0 {
            queue = append(queue, i)
        }
    }

    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]
        for _, next := range adjacency[cur] {
            isPre[cur][next] = true
            // cur 的所有先修，也是 next 的先修（传递闭包传播）
            for j := 0; j < numCourses; j++ {
                if isPre[j][cur] {
                    isPre[j][next] = true
                }
            }
            inDegree[next]--
            if inDegree[next] == 0 {
                queue = append(queue, next)
            }
        }
    }

    ans := make([]bool, len(queries))
    for i, q := range queries {
        ans[i] = isPre[q[0]][q[1]]
    }
    return ans
}
```

**运行示例 3：** `numCourses = 3, prerequisites = [[1,2],[1,0],[2,0]], queries = [[1,0],[1,2]]`

边：`1->2, 1->0, 2->0`，入度 `inDegree = [2,0,1]`。

| 出队 | 处理边 | isPre 更新 |
|------|--------|-----------|
| 1 | 1→2 | `isPre[1][2]=true`；inDegree[2] 1→0，入队 |
| 1 | 1→0 | `isPre[1][0]=true`；inDegree[0] 2→1 |
| 2 | 2→0 | `isPre[2][0]=true`；传播：`isPre[1][2]` 为真 → `isPre[1][0]=true`；inDegree[0] 1→0，入队 |
| 0 | 无出边 | — |

查询 `[1,0]` → `isPre[1][0]=true`；`[1,2]` → `isPre[1][2]=true`，返回 `[true,true]`。

### 方法二：Floyd 传递闭包

把问题看成图的可达性：`reach[i][j]=true` 表示存在路径 `i -> ... -> j`。先用直接边初始化，再跑 Floyd（中间点 k）补全间接可达：

```go
func checkIfPrerequisiteFloyd(numCourses int, prerequisites [][]int, queries [][]int) []bool {
    reach := make([][]bool, numCourses) // reach[i][j] = i 可达 j（i 是 j 的先修）
    for i := range reach {
        reach[i] = make([]bool, numCourses)
    }
    for _, p := range prerequisites {
        reach[p[0]][p[1]] = true
    }

    // Floyd 求传递闭包
    for k := 0; k < numCourses; k++ {
        for i := 0; i < numCourses; i++ {
            for j := 0; j < numCourses; j++ {
                if reach[i][k] && reach[k][j] {
                    reach[i][j] = true
                }
            }
        }
    }

    ans := make([]bool, len(queries))
    for i, q := range queries {
        ans[i] = reach[q[0]][q[1]]
    }
    return ans
}
```

### 复杂度分析

| 方法 | 预处理时间 | 单次查询 | 空间 |
|------|-----------|---------|------|
| 拓扑排序 + 传播 | O(V·E)，稠密图 O(V³) | O(1) | O(V²) |
| Floyd 传递闭包 | O(V³) | O(1) | O(V²) |

两种方法空间都是 O(V²) 的闭包矩阵。Floyd 代码更短、恒定 O(V³)；拓扑传播在稀疏图（E 远小于 V²）时更快，且与 207/210 一脉相承。numCourses ≤ 100（典型 LeetCode 数据规模）时两者都轻松通过。

### 相邻题

- [207 课程表](./No.207_课程表.md)：判环，本题在拓扑排序基础上多维护一层传递闭包。
- [210 课程表 II](./No.210_课程表2.md)：输出拓扑序。
- [2192 所有祖先节点](https://leetcode.cn/problems/all-ancestors-of-a-node-in-a-directed-acyclic-graph/)：同样的「DAG 传递闭包」模型，要求输出每个节点的全部祖先。
