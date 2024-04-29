## 653 两数之和4-输入二叉搜索树-简单

题目：

给定一个二叉搜索树 `root` 和一个目标结果 `k`，如果二叉搜索树中存在两个元素且它们的和等于给定的目标结果，则返回 `true`。



分析：【推荐该算法】

深度搜索，记录已经遍历过的节点。

```go
// date 2023/10/26
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    var res bool
    has := make(map[int]struct{}, 16)
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        if _, ok := has[k-root.Val]; ok {
            res = true
            return
        }
        has[root.Val] = struct{}{}
        dfs(root.Left)
        dfs(root.Right)
    }

    dfs(root)

    return res
}
```

算法2：

因为二叉搜索树的前序遍历序列为递增序列，所以可以先将二叉搜索树转换为递增数组，然后从排序数组中求两数之和。

```go
// date 2023/10/26
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    nums := preOrder(root)
    i, j := 0, len(nums)-1
    for i < j {
        tg := nums[i] + nums[j]
        if tg == k {
            return true
        }
        if tg < k {
            i++
        }
        if tg > k {
            j--
        }
    }
    return false
}

func preOrder(root *TreeNode) []int {
    nums := make([]int, 0, 16)
    stack := make([]*TreeNode, 0, 16)
    for root != nil || len(stack) != 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }

        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        nums = append(nums, root.Val)
        root = root.Right
    }
    return nums
}
```

