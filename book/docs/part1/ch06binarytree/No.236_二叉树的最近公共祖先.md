## 236 二叉树的最近公共祖先-中等

题目：

给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为："对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。"

> **示例 1：**
>
> ```
> 输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
> 输出：3
> 解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
> ```
>
> **示例 2：**
>
> ```
> 输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
> 输出：5
> 解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义，一个节点可以是它自己的祖先。
> ```
>
> **示例 3：**
>
> ```
> 输入：root = [1,2], p = 1, q = 2
> 输出：1
> ```

**关键观察：递归返回值的三种含义**

递归函数 `dfs(node)` 在 node 的子树中查找 p 和 q，返回值有三种可能：

- 返回 `nil`：当前子树中**没有** p 也没有 q。
- 返回 `p`（或 `q`）：当前子树中**找到了 p（或 q）**，但尚未找到两者的公共祖先。
- 返回某个非 p、非 q 的节点：该节点就是**已经找到的最近公共祖先**，直接向上透传。

后序遍历中，左右子树各自返回结果。对当前节点 `root`：

- 若 `left != nil && right != nil`：说明 p 和 q **分别位于左右子树**（或一个就是 root 本身），root 就是 LCA。
- 若只有一边非 nil：说明 p 和 q **都在同一侧**（或者只找到了一个），把该侧结果向上传递。
- 若两边都是 nil：返回 nil，表示当前子树不包含 p 或 q。

这个算法之所以正确，是因为**后序遍历保证了第一个收到"左右均非 nil"的节点一定是深度最大的公共祖先**——更浅的祖先最多只有一侧子树能同时包含 p 和 q。

**解法 1：递归（后序遍历）**

```go
// date 2023/10/24
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || root == p || root == q {
        return root
    }
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)
    // 左右各有一个 → root 就是 LCA
    if left != nil && right != nil {
        return root
    }
    // 只有一侧有 → 向上透传
    if left == nil {
        return right
    }
    return left
}
```

**解法 2：哈希表存储父节点（迭代）**

另一种思路：用哈希表记录每个节点的父节点，然后从 p 向上走到根并标记访问过的节点，再从 q 向上走，第一个遇到已标记的节点即为 LCA。这种写法不依赖递归，适合深度很大时避免栈溢出。

```go
// date 2026/07/08
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    parent := map[*TreeNode]*TreeNode{}
    // BFS/DFS 构建父节点映射
    var dfs func(*TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        if node.Left != nil {
            parent[node.Left] = node
            dfs(node.Left)
        }
        if node.Right != nil {
            parent[node.Right] = node
            dfs(node.Right)
        }
    }
    dfs(root)

    // 从 p 向上标记所有祖先
    visited := map[*TreeNode]bool{}
    for p != nil {
        visited[p] = true
        p = parent[p]
    }
    // 从 q 向上找第一个已标记的祖先
    for q != nil {
        if visited[q] {
            return q
        }
        q = parent[q]
    }
    return nil
}
```

**解法对比**

| 解法 | 时间 | 空间 | 适用场景 |
|---|---|---|---|
| 递归（后序遍历） | O(n) | O(h) | 常规情况，代码最简洁 |
| 哈希表 + 向上标记 | O(n) | O(n) | 树极深时避免递归栈溢出；或需要多次查询同一棵树（父节点表可复用） |

> 两者均只遍历一次整棵树。递归版空间取决于树高 h（最坏 O(n)）；哈希表版额外存储全部 n 个节点的父指针，但递归栈深度为 O(h)。

**与 BST 版本的区别**

注意本题是**普通二叉树**，无法利用节点值的大小关系。如果树是二叉搜索树（BST），则有更简洁的解法：利用 BST 的性质，通过比较 `root.Val` 与 `p.Val`、`q.Val` 的大小来判断 p 和 q 位于哪一侧，无需遍历整棵树。详见相邻题 No.235。

**相邻题**

- [No.235 二叉搜索树的最近公共祖先](./No.235_二叉搜索树的最近公共祖先.md)：BST 版本，利用 BST 有序性一次遍历即可，不必递归整棵树。
- [No.1644 二叉树的最近公共祖先 II](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree-ii/)：p 或 q 可能不在树中——需要在递归中额外传递"是否找到"的信息，不能仅靠 `root == p || root == q` 判断。
- [No.1650 二叉树的最近公共祖先 III](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree-iii/)：节点带有 parent 指针——转化为两个链表的相交问题，双指针 O(1) 空间即可。
- [No.1676 二叉树的最近公共祖先 IV](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree-iv/)：查找多个节点（而非两个）的 LCA——递归逻辑与本题相同，只是将 `root == p || root == q` 改为集合查找。
