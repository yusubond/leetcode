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



**分析**

这道题就是一个标准的 AOV 网的拓扑排序问题，关于 AOV 网详见 README。

拓扑排序问题的解决思路主要是循环执行以下两步，直到不存在入度为零的顶点为止。

- 选择一个入度为零的顶点并将其加入列表
- 从网中删除该顶点和其所有边（即将其出度指向的节点的入度减一）

循环结束，如果拓扑排序中的节点数小于网中的节点数，那么该 AOV 网中有回路；否则没有。



```go
// date 2023/12/28
func canFinish(numCourses int, prerequisites [][]int) bool {
    n := len(prerequisites)
    if n == 0 {
        return true
    }
    // in dress
    in := make([]int, numCourses)
    // out node list
    out := make([][]int, numCourses)
    // next is the sort list of aov
    next := make([]int, 0, numCourses)

    // 计算每个节点的入度和出度列表
    for _, v := range prerequisites {
        // [0, 1] 1->0
        in[v[0]]++
        out[v[1]] = append(out[v[1]], v[0])
    }

    // 1. 先把所有入度为零的完成
    for i := 0; i < numCourses; i++ {
        if in[i] == 0 {
            next = append(next, i)
        }
    }

    // 2. 遍历 next 列表，更新其他节点
    for i := 0; i != len(next); i++ {
        cur := next[i]
        // the free list of cur
        v := out[cur]
        if len(v) != 0 {
            for _, vv := range v {
                // cur 课程已经完成, 那么对它依赖的课程入度减一
                in[vv]--
                if in[vv] == 0 {
                    next = append(next, vv)
                }
            }
        }
    }
    // 如果不存在回环，那么其拓扑排序中的节点数等于总数
    return len(next) == numCourses
}
```

