## 199 二叉树的右视图-中等

题目：

给定一个二叉树的根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。



**解题思路**

层序遍历的最后一个节点值。

```go
// date 2023/10/23
func rightSideView(root *TreeNode) []int {
  res := make([]int, 0)
  if root == nil { return res }
  stack := make([]*TreeNode, 0)
  stack = append(stack, root)
  for len(stack) != 0 {
    n := len(stack)
    res = append(res, stack[0].Val)
    for i := 0; i < n; i++ {
      if stack[i].Right != nil { stack = append(stack, stack[i].Right) }
      if stack[i].Left != nil { stack = append(stack, stack[i].Left) }
    }
    stack = stack[n:]
  }
  return res
}
```
