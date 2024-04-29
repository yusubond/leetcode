## 545 二叉树的边界-中等

题目：

给定一个二叉树，按逆时针返回其边界。



分析：

分别求其左边界，右边界，和所有的叶子节点，用visited辅助去重，算法如下：

```go
// date 2020/02/23
func boundaryOfBinaryTree(root *TreeNode) []int {
  if root == nil { return []int{} }
  res := make([]int, 0)
  left, right := make([]*TreeNode, 0), make([]*TreeNode, 0)
  visited := make(map[*TreeNode]int, 0)
  res = append(res, root.Val)
  visited[root] = 1
  // find left node
  pre := root.Left
  for pre != nil {
    left = append(left, pre)
    visited[pre] = 1
    if pre.Left != nil {
      pre = pre.Left
    } else {
      pre = pre.Right
    }
  }
  // find right node
  pre = root.Right
  for pre != nil {
    right = append(right, pre)
    visited[pre] = 1
    if pre.Right != nil {
      pre = pre.Right
    } else {
      pre = pre.Left
    }
  }
  leafs := findLeafs(root)
  // make res
  for i := 0; i < len(left); i++ {
    res = append(res, left[i].Val)
  }
  for i := 0; i < len(leafs); i++ {
    if _, ok := visited[leafs[i]]; ok { continue }
    res = append(res, ,leafs[i].Val)
  }
  for i := len(right) - 1; i >= 0; i-- {
    res = append(res, right[i].Val)
  }
  return res
}

func findLeafs(root *TreeNode) []*TreeNode {
  res := make([]*TreeNode, 0)
  if root == nil { return res }
  if root.Left == nil && root.Right == nil {
    res = append(res, root)
    return res
  }
  res = append(res, findLeafs(root.Left)...)
  res = append(res, findLeafs(root.Right)...)
  return res
}
```

