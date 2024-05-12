## 打家劫舍3-中等

题目：

小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 `root` 。

除了 `root` 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果 **两个直接相连的房子在同一天晚上被打劫** ，房屋将自动报警。

给定二叉树的 `root` 。返回 ***在不触动警报的情况下** ，小偷能够盗取的最高金额* 。



**解题思路**

第一，不能简单地隔层计算，然后求最大值。因为如果大值出现在首尾，则行不通。



那么，考虑动态规划。

对于每个节点，都有两种选择：1）选中，那么其值就是当前节点值和左右子节点的不被选择的值，之和。

2）不选中，那么其值就是左节点中或选中或不选中的最大值，加上右子树选中或者不选中的最大值，之和。



```go
// date 2024/03/01
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
    var dfs func(root *TreeNode) [2]int
    dfs = func(root *TreeNode) [2]int {
        if root == nil {
            return [2]int{0, 0}
        }
        l, r := dfs(root.Left), dfs(root.Right)
        // select root
        v1 := root.Val + l[1] + r[1]
        // no root
        v2 := max(l[0], l[1]) + max(r[0], r[1])
        return [2]int{v1, v2}
    }

    res := dfs(root)

    return max(res[0], res[1])
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
```

