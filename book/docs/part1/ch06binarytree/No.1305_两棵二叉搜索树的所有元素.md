## 1305 两棵二叉搜索树的所有元素-中等

题目：

给你 root1 和 root2 这两棵二叉搜索树。请你返回一个列表，其中包含 两棵树 中的所有整数并按 升序 排序。.



分析：

前序遍历，变成两个有序数组，然后归并。

```go
// date 2023/10/30
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    s1 := bst2array(root1)
    s2 := bst2array(root2)
    ans := make([]int, 0, 16)
    i, j := 0, 0
    k1, k2 := len(s1), len(s2)
    for i < k1 || j < k2 {
        if i >= k1 {
            break
        }
        if j >= k2 {
            break
        }
        if s1[i] < s2[j] {
            ans = append(ans, s1[i])
            i++
        } else {
            ans = append(ans, s2[j])
            j++
        }
    }

    if i < k1 {
        ans = append(ans, s1[i:]...)
    }
    if j < k2 {
        ans = append(ans, s2[j:]...)
    }

    return ans
}

func bst2array(r1 *TreeNode) []int {
    ans := make([]int, 0, 16)

    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)
        ans = append(ans, root.Val)
        dfs(root.Right)
    }

    dfs(r1)

    return ans
}
```
