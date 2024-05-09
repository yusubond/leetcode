## 257 二叉树的所有路径-简单

题目：

给你一个二叉树的根节点 `root` ，按 **任意顺序** ，返回所有从根节点到叶子节点的路径。

**叶子节点** 是指没有子节点的节点。



分析：

深搜，记录了访问每个节点的路径，如果到达叶子结点，直接追加结果，并且返回时去掉当前结果；

如果不是叶子结点，继续深搜。

```go
// date 2023/10/24
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    res := make([]string, 0, 16)
    path := make([]int, 0, 16)

    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        path = append(path, root.Val)
        defer func() {
            path = path[:len(path)-1]
        }()
        if root.Left == nil && root.Right == nil {
            res = append(res, path2Str(path))
            return
        }
        if root.Left != nil {
            dfs(root.Left)
        }
        if root.Right != nil {
            dfs(root.Right)
        }
    }

    dfs(root)

    return res
}

func path2Str(nums []int) string {
    var res string
    for i, v := range nums {
        if i == 0 {
            res = fmt.Sprintf("%d", v)
        } else {
            res += fmt.Sprintf("->%d", v)
        }
    }
    return res
}
```

