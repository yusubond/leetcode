## 111 二叉树的最小深度-简单

给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

**说明：**叶子节点是指没有子节点的节点。



算法1：递归

这里还是递归的思想，只不过返回左右子树其中的最小值。需要注意的是，如果不存在左子树，则返回右子树的最小深度并加1，同理右子树也是如此。

```go
// date 2020/03/21
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil {
        return 1+minDepth(root.Right)
    }
    if root.Right == nil {
        return 1+minDepth(root.Left)
    }
    l, r := minDepth(root.Left), minDepth(root.Right)
    if l > r {
        return r+1
    }
    return l+1
}
```



算法2：bfs【推荐该算法】

该算法与层序遍历类似，依次入队，判断当前节点是否同时没有左右子树，从而找到最近的一层，即为最小深度。

```go
// date 2022/10/19
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    queue := make([]*TreeNode, 0, 16)
    queue = append(queue, root)
    n := len(queue)
    var depth int
    for len(queue) != 0 {
      	depth++
        n = len(queue)
        for i := 0; i < n; i++ {
            cur := queue[i]
            // 已经找到了最近的一层
            if cur.Left == nil && cur.Right == nil {
                return depth
            }
            if cur.Left != nil {
                queue = append(queue, cur.Left)
            }
            if cur.Right != nil {
                queue = append(queue, cur.Right)
            }
        }
        queue = queue[n:]
    }
    return depth
}
```



算法3：dfs深度优先搜索

这里跟递归类似，只是写法不一样。

```go
func minDepth(root *TreeNode) int {
    var dfs func(root *TreeNode, depth int) int
    dfs = func(root *TreeNode, depth int) int {
        if root == nil {
            return depth
        }
        depth++
        if root.Left == nil && root.Right == nil {
            return depth
        }
        if root.Left == nil {
            return dfs(root.Right, depth)
        }
        if root.Right == nil {
            return dfs(root.Left, depth)
        }
        l, r := dfs(root.Left, depth), dfs(root.Right, depth)
        if l < r {
            return l
        }
        return r
    }

    return dfs(root, 0)
}
```

