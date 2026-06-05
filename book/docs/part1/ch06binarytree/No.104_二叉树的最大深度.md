## 104二叉树的最大深度-简单

题目：

给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

**说明:** 叶子节点是指没有子节点的节点。



分析：

算法1：递归

```go
// 算法一: 递归，采用自底向上的递归思想
// 时间复杂度O(N)，空间复杂度O(NlogN)
// 自底向上的递归
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    l, r := maxDepth(root.Left), maxDepth(root.Right)
    if l > r {
        return l+1
    }
    return r+1
}
```



算法2：dfs深度优先搜索

这里跟递归算法其实是一样的，只不过中间过程记录一个深度值 depth, 初始为零。

- 当遇到空节点，直接返回 depth；否则递增
- 如果当前节点为叶子节点，返回depth
- 否则分别搜索左子树和右子树，返回其中的最大值

```go

// 算法2 dfs深度优先搜索
// 时间复杂度O(N)，空间复杂度O(NlogN)
// 自顶向下的递归
func maxDepth(root *TreeNode) int {
    var dfs func(root *TreeNode, depth int) int
    dfs = func(root *TreeNode, depth int) int {
        if root == nil {
            return depth
        }
        depth++
        l, r := dfs(root.Left, depth), dfs(root.Right, depth)
        if l > r {
            return l
        }
        return r
    }

    return dfs(root, 0)
}
```



算法3：bfs广度优先搜索

所谓广度优先搜索，其实就是层序遍历的思路，当遍历到最后一层时，即为最大深度。

```go
// 算法3
// bfs广度优先搜索
// 时间复杂度O(N)，空间复杂度O(N)
func maxDepth(root *TreeNode) int {
  if root == nil {return 0}
  queue := make([]*TreeNode, 0)
  queue = append(queue, root)
  depth, n := 0, 0
  // 判断是否还有新的一层
  for len(queue) != 0 {
    n = len(queue)
    // 从上一层中取节点，并查看其左右子树
    for i := 0; i < n; i++ {
      if queue[i].Left != nil {
        queue = append(queue, queue[i].Left)
      }
      if queue[i].Right != nil {
        queue = append(queue, queue[i].Right)
      }
    }
    queue = queue[n:]
    depth++
  }
  return depth
}
```

