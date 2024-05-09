## 979 在二叉树中分配硬币-中等

题目：

给你一个有 `n` 个结点的二叉树的根结点 `root` ，其中树中每个结点 `node` 都对应有 `node.val` 枚硬币。整棵树上一共有 `n` 枚硬币。

在一次移动中，我们可以选择两个相邻的结点，然后将一枚硬币从其中一个结点移动到另一个结点。移动可以是从父结点到子结点，或者从子结点移动到父结点。

返回使每个结点上 **只有** 一枚硬币所需的 **最少** 移动次数。



分析：

从流通的角度计算。

```go
// date 2023/10/28
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distributeCoins(root *TreeNode) int {
    var ans int

    // 以每个节点目标为1开始计算,
    // 计算该节点所在的子树每个节点都满足1，所需要或者贡献出去的金币数
    var dfs func(root *TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        // 每个节点最终都要有 1 个
        // 1 表示需要 1 个, 0 表示不需要, 负数表示需要贡献出去的
        need := 1 - root.Val
        l := dfs(root.Left)  // left  贡献或者需要
        r := dfs(root.Right) // right 贡献或者需要的

        // 在当前节点的视角看来，左右之间的不平衡就是需要 流桶 的次数
        ans += abs(l) + abs(r)
        
        // 跳出当前节点的视角，从外部看该节点
        // 左右相加，加上节点本身，就是该节点整体对外需要或者贡献的
        return l + r + need
    }

    dfs(root)

    return ans
}

func abs(x int) int {
    if x > 0 {
        return x
    }
    return -x
}
```

