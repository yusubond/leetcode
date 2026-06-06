## 654 最大二叉树-中等

题目：

给定一个不重复的整数数组 nums 。 最大二叉树 可以用下面的算法从 nums 递归地构建:

创建一个根节点，其值为 nums 中的最大值。
递归地在最大值 左边 的 子数组前缀上 构建左子树。
递归地在最大值 右边 的 子数组后缀上 构建右子树。
返回 nums 构建的 最大二叉树 。



分析：

找到数组中的最大值，然后递归构造。

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
func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    if len(nums) == 1 {
        return &TreeNode{Val: nums[0]}
    }
    // find the max
    idx := 0
    for i, v := range nums {
        if v > nums[idx] {
            idx = i
        }
    }
    root := &TreeNode{Val: nums[idx]}
    root.Left = constructMaximumBinaryTree(nums[:idx])
    root.Right = constructMaximumBinaryTree(nums[idx+1:])
    return root
}


// date 2024/10/10
// 补充 dfs 函数版本
func constructMaximumBinaryTree(nums []int) *TreeNode {
    var dfs func(left, right int) *TreeNode

    // [left, right)
    dfs = func(left, right int) *TreeNode {
        if left >= right {
            return nil
        }

        idx := left
        for i := left+1; i < right; i++ {
            if nums[idx] < nums[i] {
                idx = i
            }
        }

        root := &TreeNode{Val: nums[idx]}
        root.Left = dfs(left, idx)
        root.Right = dfs(idx+1, right)

        return root
    }

    return dfs(0, len(nums))
}
```

