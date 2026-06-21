## 226 翻转二叉树-中等

题目：

给你一棵二叉树的根节点 `root` ，翻转这棵二叉树，并返回其根节点。



**解题思路**

- DFS，详见解法1。具体是直接深度优先搜索，修改每个节点的左右指针。
- BFS，详见解法2。将节点加入队列，依次翻转队列中每个节点的左右指针。

```go
// date 2023/12/28
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 解法1 dfs
func invertTree(root *TreeNode) *TreeNode {

    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        or := root.Right
        root.Right = root.Left
        root.Left = or

        dfs(root.Right)
        dfs(root.Left)
    }

    dfs(root)

    return root
}

// 解法2
// BFS
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    if root.Left == nil && root.Right == nil {
        return root
    }

    queue := make([]*TreeNode, 0, 16)
    queue = append(queue, root)

    for len(queue) != 0 {
        n := len(queue)
        for i := 0; i < n; i++ {
            cur := queue[i]
            or := cur.Right
            cur.Right = cur.Left
            cur.Left = or

            if cur.Left != nil {
                queue = append(queue, cur.Left)
            }
            if cur.Right != nil {
                queue = append(queue, cur.Right)
            }
        }

        queue = queue[n:]
    }

    return root
}
```
