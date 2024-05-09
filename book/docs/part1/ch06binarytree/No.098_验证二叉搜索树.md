## 98 验证二叉搜索树-中等

题目：

给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。



**解题思路**

- 迭代版本的中序遍历，详见解法1。二叉搜索树的中序遍历为递增序列。
- 递归得到中序遍历数组，然后判断数组，详见解法2。
- 带区间值的DFS，详见解法3。

```go
// date 2023/10/22
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 解法1
// 迭代 中序遍历
func isValidBST(root *TreeNode) bool {
    stack := make([]*TreeNode, 0, 16)
    preVal := math.Inf(-1)
    for root != nil || len(stack) != 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        cur := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if float64(cur.Val) <= preVal {
            return false
        }
        preVal = float64(cur.Val)
        root = cur.Right
    }
    return true
}

// date 2020/03/21
// 解法2
// 遍历其中序遍历序列，并判断其序列是否为递增
func isValidBST(root *TreeNode) bool {
    nums := inOrder(root)
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] >= nums[i+1] { return false }
    }
    return true
}

func inOrder(root *TreeNode) []int {
    if root == nil { return []int{} }
    res := make([]int, 0)
    if root.Left != nil { res = append(res, inOrder(root.Left)...) }
    res = append(res, root.Val)
    if root.Right != nil { res = append(res, inOrder(root.Right)...) }
    return res
}

// 解法3：递归
// 递归判断当前结点的值是否位于上下边界之中
// 时间复杂度O(N)
func isValidBST(root *TreeNode) bool {
  if root == nil { return true }
  var isValid func(root *TreeNode, min, max float64) bool
  isValid = func(root *TreeNode, min, max float64) bool {
    if root == nil { return true }
    if float64(root.Val) <= min || float64(root.Val) >= max { return false }
    return isValid(root.Left, min, float64(root.Val)) && isValid(root.Right, float64(root.Val), max)
  }
  return isValid(root, math.Inf(-1), math.Inf(0))
}
```

