## 437 路径总和3-中等

题目：

给定一个二叉树的根节点 `root` ，和一个整数 `targetSum` ，求该二叉树里节点值之和等于 `targetSum` 的 **路径** 的数目。

**路径** 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。



分析：

算法1：前缀和【推荐该算法】

因为题目中的路径并不需要从根节点开始，也不需要从叶子节点结束。只要路径上任意连续的值求和等于目标就是其中一个解。

所以可以在深度优先搜索的时候，将从根节点到当前节点（包含当前节点）的路径总和（定义为 cur ）保存起来（保存到map preSum中）。保存起来有两个用处：

1）如果 cur - targetSum == 0，那么这就是一个解。

2）如果 cur - targetSum 作为 key，在 preSum 中存在，那么表示在这条路径存在一个解。

需要注意的是，当结束当前节点的遍历时，需要将前缀和从保存路径和的 map 中去掉。因为当前节点的所有可能已经查看过了，如果还在 map 中存在，会影响其他节点的判断。

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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	ans := 0
	pathPreSum := map[int]int{0: 1}

	var dfs func(root *TreeNode, sum int)
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		sum += root.Val
		ans += pathPreSum[sum-targetSum]

		pathPreSum[sum]++
		dfs(root.Left, sum)
		dfs(root.Right, sum)
		pathPreSum[sum]--
	}

	dfs(root, 0)

	return ans
}
```



算法2：

深度优先搜索。

定义`rootSum()`

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
func pathSum(root *TreeNode, targetSum int) int {
    res := 0
    if root == nil {
        return res
    }

    res = rootSum(root, targetSum)
    res += pathSum(root.Left, targetSum)
    res += pathSum(root.Right, targetSum)

    return res
}

// 以 root 为起点满足条件的路径总和
func rootSum(root *TreeNode, tg int) int {
    res := 0
    if root == nil {
        return res
    }
    if root.Val == tg {
        res++
    }
    res += rootSum(root.Left, tg - root.Val)
    res += rootSum(root.Right, tg - root.Val)
    return res
}
```

