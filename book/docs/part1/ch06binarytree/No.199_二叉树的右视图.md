## 199 二叉树的右视图-中等

题目：

给定一个二叉树的根节点 `root`，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

> **示例 1：**
>
> ```
> 输入：root = [1,2,3,null,5,null,4]
> 输出：[1,3,4]
> 解释：
>       1     ← 看到 1
>      / \
>     2   3   ← 看到 3
>      \   \
>       5   4 ← 看到 4
> ```
>
> **示例 2：**
>
> ```
> 输入：root = [1,2,3,4,null,null,null,5]
> 输出：[1,3,4,5]
> 解释：右视图按层依次为 [1,3,4,5]
> ```
>
> **示例 3：**
>
> ```
> 输入：root = [1,null,3]
> 输出：[1,3]
> ```
>
> **示例 4：**
>
> ```
> 输入：root = []
> 输出：[]
> ```

**关键观察：每层的最右节点**

站在右侧看二叉树，每一层只能看到**该层最右边的节点**。因此问题等价于：**取每一层的最后一个节点（或按层序遍历反序取第一个）**。

两种经典思路：

- **BFS（层序遍历）**：逐层访问，取每层最右侧节点。可以左→右遍历取最后一个，也可以右→左遍历取第一个。
- **DFS（深度优先）**：按"根 → 右 → 左"的顺序遍历，**每层第一个访问到的节点**就是从右侧看到的节点。用一个 `depth` 参数跟踪当前深度，当 `len(res) == depth` 时说明这是该层首次访问，加入结果。

**解法 1：BFS 层序遍历（左→右，取最后一个）**

最直观的写法：标准层序遍历，每层遍历结束后取最后一个节点。

```go
// date 2026/07/08
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    res := make([]int, 0)
    if root == nil {
        return res
    }
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        n := len(queue)
        // 每层最后一个节点即右视图节点
        res = append(res, queue[n-1].Val)
        for i := 0; i < n; i++ {
            node := queue[i]
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        queue = queue[n:]
    }
    return res
}
```

**解法 1 变体：BFS（右→左，取第一个）**

先入队右孩子再入队左孩子，则每层第一个节点就是右视图节点，无需等到该层遍历结束。

```go
// date 2023/10/23
func rightSideView(root *TreeNode) []int {
    res := make([]int, 0)
    if root == nil {
        return res
    }
    stack := make([]*TreeNode, 0)
    stack = append(stack, root)
    for len(stack) != 0 {
        n := len(stack)
        // 右→左入队，每层第一个即为最右节点
        res = append(res, stack[0].Val)
        for i := 0; i < n; i++ {
            if stack[i].Right != nil {
                stack = append(stack, stack[i].Right)
            }
            if stack[i].Left != nil {
                stack = append(stack, stack[i].Left)
            }
        }
        stack = stack[n:]
    }
    return res
}
```

**解法 2：DFS 深度优先（根→右→左，记录每层首个）**

按"根 → 右子树 → 左子树"的顺序 DFS。对于每一层深度 `depth`，第一个到达的节点一定是最右侧的节点。当 `len(res) == depth` 时说明该层尚未记录，当前节点即为该层的右视图节点。

```go
// date 2026/07/08
func rightSideView(root *TreeNode) []int {
    res := make([]int, 0)

    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, depth int) {
        if node == nil {
            return
        }
        // 当前层首次访问 → 该节点即右侧看到的节点
        if len(res) == depth {
            res = append(res, node.Val)
        }
        // 先右后左，保证右侧节点先被访问
        dfs(node.Right, depth+1)
        dfs(node.Left, depth+1)
    }

    dfs(root, 0)
    return res
}
```

**复杂度**

| 解法 | 时间 | 空间 | 备注 |
|---|---|---|---|
| BFS 层序遍历 | O(n) | O(w) | w 为树的最大宽度，最坏 O(n) |
| DFS 深度优先 | O(n) | O(h) | h 为树高，最坏 O(n)，平衡树 O(log n) |

> 两种解法渐进复杂度相同。BFS 适合宽树，DFS 适合深树（递归栈更省空间）。DFS 写法更简洁，无需显式维护队列。

**相邻题**

- [No.102 二叉树的层序遍历](./No.102_二叉树的层序遍历.md)：BFS 层序遍历的母题，左右视图的 BFS 解法均基于此。
- [No.103 二叉树的锯齿形层序遍历](./No.103_二叉树的锯齿形层序遍历.md)：同为层序遍历变体，交替方向取节点。
- [No.637 二叉树的层平均值](./No.637_二叉树的层平均值.md)：同一层遍历框架，只是将"取最右节点"替换为"取层平均值"。
- [No.513 找树左下角的值](https://leetcode.cn/problems/find-bottom-left-tree-value/)：左视图的对应题——找最后一层的最左节点，DFS 改为"左→右"遍历即可。
