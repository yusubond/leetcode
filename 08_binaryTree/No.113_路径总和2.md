## 113 路径总和2-中等

题目：

给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点。



分析：

算法1:

这里采用深度优先搜索，当遍历到叶子节点且节点的值等于剩下的值，那么就直接追加结果，并返回。如果不是，递归遍历左右子树。

但是这里的写法相对麻烦一些，因为是自底向上的递推；其实可以采用自顶向下的递推，利用path记录已经遍历过的节点。详见算法2。

```go
// date 2022/10/24
func pathSum(root *TreeNode, targetSum int) [][]int {
    var dfs func(root *TreeNode, v int) ([][]int, bool)
    dfs = func(root *TreeNode, v int) ([][]int, bool) {
        if root == nil {
            return nil, false
        }
        res := make([][]int, 0, 16)
        if root.Left == nil && root.Right == nil && root.Val == v {
            t := []int{root.Val}
            res = append(res, t)
            return res, true
        }
        v -= root.Val
        l, ok1 := dfs(root.Left, v)
        r, ok2 := dfs(root.Right, v)
        if ok1 {
            for i := 0; i < len(l); i++ {
                t := make([]int, 0, 16)
                t = append(t, root.Val)
                t = append(t, l[i]...)
                res = append(res, t)
            }
        }
        if ok2 {
            for i := 0; i < len(r); i++ {
                t := make([]int, 0, 16)
                t = append(t, root.Val)
                t = append(t, r[i]...)
                res = append(res, t)
            }
        }
        if ok1 || ok2 {
            return res, true
        }
        return nil, false
    }

    asn, _ := dfs(root, targetSum)
    return asn
}
```



算法2：【推荐该算法】

同样采用深度优先搜索，`path` 记录已经遍历过的节点值。

1. 当遍历到节点时，直接追加值到path中
2. 如果遇到叶子节点且满足条件，直接将path追加到最终结果，
3. 因为每次遍历，整个路径会变化，因为需要返回前将1中增加的值去掉。

```go
// date 2022/10/24
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
    ans := make([][]int, 0, 16)
    path := make([]int, 0, 16)

    var dfs func(root *TreeNode, sum int)
    dfs = func(root *TreeNode, sum int) {
        if root == nil {
            return
        }
        path = append(path, root.Val)
        sum -= root.Val

        defer func() {
            path = path[:len(path)-1]
        }()

        if root.Left == nil && root.Right == nil && sum == 0 {
            one := make([]int, len(path))
            copy(one, path)
            ans = append(ans, one)
            return
        }

        dfs(root.Left, sum)
        dfs(root.Right, sum)
    }

    dfs(root, targetSum)

    return ans
}
```

