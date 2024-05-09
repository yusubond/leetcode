## 235 二叉搜索数的最近公共祖先-中等

题目：

给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

[百度百科](https://baike.baidu.com/item/最近公共祖先/8918834?fr=aladdin)中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（**一个节点也可以是它自己的祖先**）。”



分析：

与 236 不同的是，二叉搜索数的节点值是有规律的，可以借助这个规律进行递归。

```go
// date 2023/10/24
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
        return root
    }
    if p == q {
        return p
    }
    if root.Val > q.Val && root.Val > p.Val {
        return lowestCommonAncestor(root.Left, p, q)
    }
    if root.Val < q.Val && root.Val < p.Val {
        return lowestCommonAncestor(root.Right, p, q)
    }
    return root
}
```

