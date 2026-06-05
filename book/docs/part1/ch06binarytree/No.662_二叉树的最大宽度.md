## 662 二叉树的最大宽度-中等

题目：

给你一棵二叉树的根节点 root ，返回树的 最大宽度 。

树的 最大宽度 是所有层中最大的 宽度 。

每一层的 宽度 被定义为该层最左和最右的非空节点（即，两个端点）之间的长度。将这个二叉树视作与满二叉树结构相同，两端点间会出现一些延伸到这一层的 null 节点，这些 null 节点也计入长度。

题目数据保证答案将会在  32 位 带符号整数范围内。



分析：

因为空节点也算数，所以可通过对节点进行编号进行计算，设父结点编号为 `i`，其左右子节点分别为`2i`和`2i+1`，那么每层的最大跨度就是最大的编号，减去最小的编号。

这里使用深度优先搜索，进行编号，先遍历左子树，再遍历右子树。这样最小的编号就会先被填充。

在遍历到每个节点时，求当前节点与同层最左侧的节点跨度，如果大于结果集，就更新结果。

这样的代码可读性更高。

```go
// date 2023/10/26
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func widthOfBinaryTree(root *TreeNode) int {
    ans := 0
    levelMinIdx := map[int]int{}
    var dfs func(root *TreeNode, depth, idx int)
    dfs = func(root *TreeNode, depth, idx int) {
        if root == nil {
            return
        }
        if _, ok := levelMinIdx[depth]; !ok {
            levelMinIdx[depth] = idx
        }
        curNodeWidth := idx - levelMinIdx[depth] + 1
        if curNodeWidth > ans {
            ans = curNodeWidth
        }
        dfs(root.Left, depth+1, idx*2)
        dfs(root.Right, depth+1, idx*2+1)
    }

    dfs(root, 1, 1)

    return ans
}
```

