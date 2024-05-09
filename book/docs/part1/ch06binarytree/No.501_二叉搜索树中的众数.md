## 501 二叉搜索树中的众数-中等

题目：

给你一个含重复值的二叉搜索树（BST）的根节点 `root` ，找出并返回 BST 中的所有 [众数](https://baike.baidu.com/item/众数/44796)（即，出现频率最高的元素）。

如果树中有不止一个众数，可以按 **任意顺序** 返回。

假定 BST 满足如下定义：

- 结点左子树中所含节点的值 **小于等于** 当前节点的值
- 结点右子树中所含节点的值 **大于等于** 当前节点的值
- 左子树和右子树都是二叉搜索树



分析：

充分利用二叉搜索数的特点，其中序遍历序列为非递减序列。

所以，递归中序遍历，统计每个元素出现的次数，一旦形成结果，就更新结果。

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
func findMode(root *TreeNode) []int {
    ans := make([]int, 0, 16)

    var base, count, maxCount int

    var updateNodeVal func(x int)
    updateNodeVal = func(x int) {
        if x == base {
            count++
        } else {
            base, count = x, 1
        }
        // check count and maxcount
        if count == maxCount {
            // find a answer
            ans = append(ans, x)
        } else if count > maxCount {
            // find new answer
            maxCount = count
            ans = []int{x}   // 重新定义结果集
        }
    }

    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)
        updateNodeVal(root.Val)
        dfs(root.Right)
    }

    dfs(root)

    return ans
}
```

