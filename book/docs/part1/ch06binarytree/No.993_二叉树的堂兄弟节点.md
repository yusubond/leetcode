## 993 二叉树的堂兄弟节点-简单
题目：

在二叉树中，根节点位于深度 0 处，每个深度为 k 的节点的子节点位于深度 k+1 处。
如果二叉树的两个节点深度相同，但 父节点不同 ，则它们是一对堂兄弟节点。
我们给出了具有唯一值的二叉树的根节点 root ，以及树中两个不同节点的值 x 和 y 。
只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true 。否则，返回 false。



分析：

算法1：

通过查找 x,y 的父结点以及深度来进行判断。这里的算法需要遍历树两次。


```go
func isCousins(root *TreeNode, x int, y int) bool {
    xf, xd := findFatherAndDepth(root, x)
    yf, yd := findFatherAndDepth(root, y)
    return xd == yd && xf != yf
}

func findFatherAndDepth(root *TreeNode, v int) (*TreeNode, int) {
    if root == nil || root.Val == v {
        return nil, 0
    }
    if root.Left != nil && root.Left.Val == v {
        return root, 1
    }
    if root.Right != nil && root.Right.Val == v {
        return root, 1
    }
    l, lv := findFatherAndDepth(root.Left, v)
    r, rv := findFatherAndDepth(root.Right, v)
    if l != nil {
        return l, lv+1
    }
    return r, rv+1
}
```



算法2：【推荐该算法】

这里用深度优先搜索只遍历一次，搜索过程中记录x,y的深度和父结点。

```go
// date 2022/10/02
func isCousins(root *TreeNode, x int, y int) bool {
    if root == nil {return false}
    xd, yd := 0, 0  // 存储x,y的深度
    f := make(map[int]*TreeNode, 2) // 存储x,y父结点
    path := make([]*TreeNode, 0, 16)
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        path = append(path, root)
        defer func() {
            path = path[:len(path)-1]
        }()
        if root.Val == x {
            xd = len(path)
            if len(path) > 1 {
                f[x] = path[len(path)-2]
            }
        } else if root.Val == y {
            yd = len(path)
            if len(path) > 1 {
                f[y] = path[len(path)-2]
            }
        }
        dfs(root.Left)
        dfs(root.Right)
    }
    dfs(root)
    xf, ok1 := f[x]
    yf, ok2 := f[y]
    return ok1 && ok2 && xf != yf && xd == yd
}
```

另一个写法

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
func isCousins(root *TreeNode, x int, y int) bool {
    // 深度一样，且父结点不同

    xd, yd := 0, 0
    var xf, yf *TreeNode

    var dfs func(root *TreeNode, depth int)
    dfs = func(root *TreeNode, depth int) {
        if root == nil {
            return
        }
        if root.Left != nil {
            if root.Left.Val == x {
                xd = depth+1
                xf = root
            }
            if root.Left.Val == y {
                yd = depth+1
                yf = root
            }
        }
        if root.Right != nil {
            if root.Right.Val == x {
                xd = depth+1
                xf = root
            }
            if root.Right.Val == y {
                yd = depth+1
                yf = root
            }
        }
        dfs(root.Left, depth+1)
        dfs(root.Right, depth+1)
    }

    dfs(root, 0)

    return xd == yd && xf != nil && yf != nil && xf != yf
}
```

