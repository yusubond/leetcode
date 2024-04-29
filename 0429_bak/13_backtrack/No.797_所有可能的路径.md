## 797 所有可能的路径-中等

题目：

给你一个有 `n` 个节点的 **有向无环图（DAG）**，请你找出所有从节点 `0` 到节点 `n-1` 的路径并输出（**不要求按特定顺序**）

 `graph[i]` 是一个从节点 `i` 可以访问的所有节点的列表（即从节点 `i` 到节点 `graph[i][j]`存在一条有向边）。



> **示例 1：**
>
> ![img](https://assets.leetcode.com/uploads/2020/09/28/all_1.jpg)
>
> ```
> 输入：graph = [[1,2],[3],[3],[]]
> 输出：[[0,1,3],[0,2,3]]
> 解释：有两条路径 0 -> 1 -> 3 和 0 -> 2 -> 3
> ```
>
> **示例 2：**
>
> ![img](https://assets.leetcode.com/uploads/2020/09/28/all_2.jpg)
>
> ```
> 输入：graph = [[4,3,1],[3,2,4],[3],[4],[]]
> 输出：[[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]
> ```



分析：

这道题的解题思路是 DFS。每次从遍历可选的节点，有两种情况：

- 如果节点值等于`n-1`，那么表示找到一条路径，直接添加结果
- 如果节点值不等于 `n-1`，那么找出节点的候选列表中继续搜索

无论找到还未找到，在回溯后，都要取消候选。



```go
// date 2023/12/27
func allPathsSourceTarget(graph [][]int) [][]int {
    res := make([][]int, 0, 16)

    n := len(graph)
    var backtrack func(nums []int, path []int)
    backtrack = func(nums []int, path []int) {
        for _, v := range nums {
            if v == n-1 {
                path = append(path, v)
                one := make([]int, len(path))
                copy(one, path)
                res = append(res, one)
                path = path[:len(path)-1]
            } else {
                path = append(path, v)
                backtrack(graph[v], path)
                path = path[:len(path)-1]
            }
        }
    }

    backtrack(graph[0], []int{0})

    return res
}
```

