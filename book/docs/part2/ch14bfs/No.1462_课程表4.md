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



分析：

这道题也是 AOV 网拓扑排序的思路，需要注意两点：

1）通过拓扑排序，从已排序的数组中查找相对位置关系不可行。因为 aov 网的拓扑排序中对入度为零的节点是依次添加的的，入度为零的节点没有依赖关系，但是依次添加形成了顺序，所以不可行。

2）通过构造 aov 网，对全网进行 BFS 搜索会超时。

3）通过构造 aov 网，进而构造前置依赖 map 列表，会超时内存限制。



综合来看，这道题的解法是拓扑排序和 BFS。具体这些步骤来完成：

1）根据依赖关系，构建每个节点的入度和出度节点关系

2）构建二维数组 `isPre[i][j]`，表示课程`i`是课程`j`的前置依赖

3）开始拓扑排序，排序过程中根据出度节点更新`isPre二维数组`。需要注意的是，如果课程`cur`是课程`v`的前置依赖，那么所有的课程中，`cur`的前置依赖也是`v`的前置依赖。

```go
// date 2024/01/02
func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
    in := make([]int, numCourses)
    out := make(map[int][]int)
    for _, v := range prerequisites {
        in[v[1]]++
        _, ok := out[v[0]]
        if !ok {
            out[v[0]] = []int{}
        }
        out[v[0]] = append(out[v[0]], v[1])
    }
    queue := make([]int, 0, 16)
    for i, v := range in {
        if v == 0 {
            queue = append(queue, i)
        }
    }
    isPre := make([][]bool, numCourses)
    for i := 0; i < numCourses; i++ {
        isPre[i] = make([]bool, numCourses)
    }

    for len(queue) > 0 {
        n := len(queue)

        for i := 0; i < n; i++ {
            cur := queue[i]
            if ov, ok := out[cur]; ok && len(ov) != 0 {
                for _, v := range ov {
                    isPre[cur][v] = true

                    // cur 是 v 的前置依赖
                    // 那么，所有的课程中 cur 的前置依赖也是 v 的前置依赖
                    for j := 0; j < numCourses; j++ {
                        isPre[j][v] = isPre[j][v] || isPre[j][cur]
                    }
                    in[v]--
                    if in[v] == 0 {
                        queue = append(queue, v)
                    }
                }
            }
        }

        queue = queue[n:]
    }

    ans := make([]bool, 0, 16)
    for _, v := range queries {
        ans = append(ans, isPre[v[0]][v[1]])
    }

    return ans
}
```

