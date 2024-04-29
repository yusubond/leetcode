## 988 从叶子节点开始的最小字符串

题目：

给定一颗根结点为 `root` 的二叉树，树中的每一个结点都有一个 `[0, 25]` 范围内的值，分别代表字母 `'a'` 到 `'z'`。

返回 ***按字典序最小** 的字符串，该字符串从这棵树的一个叶结点开始，到根结点结束*。

> 注**：**字符串中任何较短的前缀在 **字典序上** 都是 **较小** 的：
>
> - 例如，在字典序上 `"ab"` 比 `"aba"` 要小。叶结点是指没有子结点的结点。 

节点的叶节点是没有子节点的节点。



**解题思路**

直接携带已有的字符进行深度优先遍历，将子节点的字符加载已有字符的头部。

到达叶子节点时，进行比较更新。

```go
// date 2024/02/28
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) string {
	var ans string
	var dfs func(root *TreeNode, has string)
	dfs = func(root *TreeNode, has string) {
		if root == nil {
			return
		}
		has = string('a'+root.Val) + has
		if root.Left == nil && root.Right == nil {
			if ans == "" {
				ans = has
			} else if ans > has {
				ans = has
			}
			return
		}
		if root.Left != nil {
			dfs(root.Left, has)
		}
		if root.Right != nil {
			dfs(root.Right, has)
		}
	}

	dfs(root, "")

	return ans
}
```

