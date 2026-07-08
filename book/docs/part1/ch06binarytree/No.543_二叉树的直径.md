## 543 二叉树的直径-简单

题目：

给你一棵二叉树的根节点，返回该树的 **直径** 。

二叉树的 **直径** 是指树中任意两个节点之间最长路径的 **长度** 。这条路径可能经过也可能不经过根节点 `root` 。

两节点之间路径的 **长度** 由它们之间边数表示。

> **示例 1：**
>
> ```
> 输入：root = [1,2,3,4,5]
> 输出：3
> 解释：3，取路径 [4,2,1,3] 或 [5,2,1,3] 的长度。
> ```
>
> **示例 2：**
>
> ```
> 输入：root = [1,2]
> 输出：1
> ```

**关键观察：直径 = 左右子树深度之和的全局最大值**

- 对于任意节点，**经过该节点的"局部直径" = 左子树深度 + 右子树深度**（深度以边数计，叶子节点深度为 0）。
- 全局直径 = 所有节点局部直径的最大值。路径**可能不经过根节点**，因此必须在每个节点都计算一次，而不能只在根节点算——这是本题最容易出错的地方。
- **后序遍历（DFS 自底向上）**天然适合此题：先递归得到左右子树的深度，用它们的和更新全局答案，再向上返回"当前节点为根的子树深度"（`1 + max(左深, 右深)`）。
- 该题与「二叉树的最大深度」（No.104）共享同一递归骨架，区别仅在于多了一个"在返回深度之前用 `左深+右深` 更新直径"的步骤。

**解法 1：DFS 后序遍历 + 全局变量**

每次递归返回子树深度的同时，用 `左深 + 右深` 更新全局最大值。

```go
// date 2023/10/25
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    var ans int

    var depth func(*TreeNode) int
    depth = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        l := depth(node.Left)
        r := depth(node.Right)
        // 经过当前节点的局部直径 = 左深 + 右深
        if l+r > ans {
            ans = l + r
        }
        // 向上返回子树深度（边数）
        if l > r {
            return l + 1
        }
        return r + 1
    }

    depth(root)
    return ans
}
```

**解法 1 变体：更简洁的写法**

逻辑完全一致，利用 `max` 函数减少分支。

```go
// date 2026/07/08
func diameterOfBinaryTree(root *TreeNode) int {
    ans := 0

    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        l := dfs(node.Left)
        r := dfs(node.Right)
        ans = max(ans, l+r) // 更新全局直径
        return max(l, r) + 1 // 返回子树深度
    }

    dfs(root)
    return ans
}
```

**核心要点：为什么空节点返回 0**

叶子节点的左右孩子均为 `nil`，`depth(nil) = 0`，于是叶子处 `l + r = 0 + 0 = 0`，叶子向上返回 `1`。这使得深度以**边数**计量，直径自然也是边数——与题意一致。

**复杂度**

| 解法 | 时间 | 空间 | 备注 |
|---|---|---|---|
| DFS 后序遍历 | O(n) | O(h) | 每个节点访问一次；递归栈深度最坏 O(n)，平衡树 O(log n) |

> `h` 为树高。最坏情况（链表状树）递归栈深度为 O(n)；平均平衡树为 O(log n)。

**相邻题**

- [No.104 二叉树的最大深度](./No.104_二叉树的最大深度.md)：深度计算的三种写法（递归 / DFS / BFS），直径题的核心递归骨架即来源于此。
- [No.1522 N叉树的直径](./No.1522_N叉树的直径.md)：将"左右子树"推广为"最深的两棵子树"，排序取前二，核心思路同源。
- [No.124 二叉树中的最大路径和](./No.124_二叉树中的最大路径和.md)：同为"DFS 返回单边贡献 + 全局更新"模式，只是把边数换成了节点值之和，可作为同框架的进阶对照。
