## 95 不同的二叉搜索树-中等

题目：

给你一个整数 `n` ，请你生成并返回所有由 `n` 个节点组成且节点值从 `1` 到 `n` 互不相同的不同 **二叉搜索树** 。可以按 **任意顺序** 返回答案。



分析：

题目看着很难，实际利用递归思想没有那么难。

从序列`1..n`取出数字`1` 作为当前树的根节点，那么就有i-1个元素用来构造左子树，二另外的i+1..n个元素用来构造右子树。最后我们将得到G(i-1)棵不同的左子树，G(n-i)棵不同的右子树，其中G为卡特兰数。

这样就可以以i为根节点，和两个可能的左右子树序列，遍历一遍左右子树序列，将左右子树和根节点链接起来。

```go
// date 2022/10/01
func generateTrees(n int) []*TreeNode {
    var genWithNode func(start, end int) []*TreeNode
    genWithNode = func(start, end int) []*TreeNode {
        res := make([]*TreeNode, 0, 16)
        if start > end {
            // 空树
            res = append(res, nil)
            return res
        }
        for i := start; i <= end; i++ {
            left := genWithNode(start, i-1)
            right := genWithNode(i+1, end)
            for _, l := range left {
                for _, r := range right {
                    res = append(res, &TreeNode{Val: i, Left: l, Right: r})
                }
            }
        }
        return res
    }
    return genWithNode(1, n)
}
```

