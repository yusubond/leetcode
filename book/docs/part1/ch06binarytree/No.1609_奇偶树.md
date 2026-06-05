## 1609 奇偶树-中等

题目：

如果一棵二叉树满足下述几个条件，则可以称为 **奇偶树** ：

- 二叉树根节点所在层下标为 `0` ，根的子节点所在层下标为 `1` ，根的孙节点所在层下标为 `2` ，依此类推。
- **偶数下标** 层上的所有节点的值都是 **奇** 整数，从左到右按顺序 **严格递增**
- **奇数下标** 层上的所有节点的值都是 **偶** 整数，从左到右按顺序 **严格递减**

给你二叉树的根节点，如果二叉树为 **奇偶树** ，则返回 `true` ，否则返回 `false` 。



分析：

没啥好说的，就是基本的层序遍历。

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
func isEvenOddTree(root *TreeNode) bool {
    if root == nil {
        return false
    }

    queue := make([]*TreeNode, 0, 16)
    queue = append(queue, root)

    var preVal int
    isEvenIdx := false

    for len(queue) != 0 {
        n := len(queue)

        for i := 0; i < n; i++ {
            cur := queue[i]

            if isEvenIdx {
                if cur.Val % 2 == 1 {
                    return false
                }
            } else {
                if cur.Val % 2 == 0 {
                    return false
                }
            }

            if i == 0 {
                preVal = cur.Val
            } else {
                if isEvenIdx {
                    // 严格递减
                    if cur.Val >= preVal {
                        return false
                    }
                } else {
                    if cur.Val <= preVal {
                        return false
                    }
                }
                preVal = cur.Val
            }

            if cur.Left != nil {
                queue = append(queue, cur.Left)
            }
            if cur.Right != nil {
                queue = append(queue, cur.Right)
            }
        }

        queue = queue[n:]
        isEvenIdx = !isEvenIdx
    }

    return true
}
```

