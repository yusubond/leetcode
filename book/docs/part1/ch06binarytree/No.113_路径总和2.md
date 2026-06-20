## 113 路径总和2-中等

题目：

给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点。



分析：

深度优先搜索+回溯算法，`path` 记录已经遍历过的节点值。

1. 当遍历节点时，直接追加值到 path 中
2. 如果遇到叶子节点且满足条件，直接将 path 追加到最终结果，注意深拷贝。
3. 因为每次遍历，整个路径会变化，撤销当前选择即可。

```go
// date 2022/10/24
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
    ans := make([][]int, 0, 16)
    path := make([]int, 0, 16)

    var dfs func(root *TreeNode, sum int)
    dfs = func(root *TreeNode, sum int) {
        if root == nil {
            return
        }
        path = append(path, root.Val)
        sum -= root.Val

        defer func() {
            path = path[:len(path)-1]
        }()

        if root.Left == nil && root.Right == nil && sum == 0 {
            one := make([]int, len(path))
            copy(one, path)
            ans = append(ans, one)
            return
        }

        dfs(root.Left, sum)
        dfs(root.Right, sum)
    }

    dfs(root, targetSum)

    return ans
}
```

