## 108 将有序数组转换为二叉搜索树-简单

题目：

给你一个整数数组 nums，其中元素已经按升序排列，请你将其转换为一个 高度平衡 的二叉搜索树。

高度平衡二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1」的二叉树。



分析：

跟【从中序和后序序列构造二叉树】类似，为了保持高度平衡，从有序数组中去中间值作为根，依次递归生成左右子树。

```go
// date 2023/10/23
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    mid := len(nums)/2

    root := &TreeNode{Val: nums[mid]}
    root.Left = sortedArrayToBST(nums[0:mid])
    root.Right = sortedArrayToBST(nums[mid+1:])

    return root
}
```

