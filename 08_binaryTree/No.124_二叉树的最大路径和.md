## 124 二叉树的最大路径和-中等

题目：

二叉树中的 **路径** 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 **至多出现一次** 。该路径 **至少包含一个** 节点，且不一定经过根节点。

**路径和** 是路径中各节点值的总和。

给你一个二叉树的根节点 `root` ，返回其 **最大路径和** 。



> **示例 1：**
>
> ![img](https://assets.leetcode.com/uploads/2020/10/13/exx1.jpg)
>
> ```
> 输入：root = [1,2,3]
> 输出：6
> 解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
> ```
>
> **示例 2：**
>
> ![img](https://assets.leetcode.com/uploads/2020/10/13/exx2.jpg)
>
> ```
> 输入：root = [-10,9,20,null,null,15,7]
> 输出：42
> 解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
> ```



**解题思路**

这道题可用 BFS 解。

所谓路径，就是树中节点可以连接起来，且有起点和终点。所以对于某个节点而言，最大路径和就是其节点值加上左右子树的最大路径和。

需要注意的是，BFS 返回的时候，选取左右子树其中最大的即可，因为要有起点或者终点。

```go
// date 2024/01/16
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    ans := math.MinInt32

    var bfs func(root *TreeNode) int
    bfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
      	// 如果左右子树的总和为负数，可以不加入路径
        left := max(bfs(root.Left), 0)
        right := max(bfs(root.Right), 0)
        
        sum := root.Val + left + right
        ans = max(ans, sum)

        return root.Val + max(left, right)
    }

    bfs(root)

    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
```

