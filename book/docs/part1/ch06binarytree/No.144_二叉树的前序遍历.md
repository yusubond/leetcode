## 144 二叉树的前序遍历-简单

题目：

给你二叉树的根节点 `root` ，返回它节点值的 **前序** 遍历。



分析：

算法1：递归

```go
// root -> left -> right
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    res := make([]int, 0, 16)
    res = append(res, root.Val)
    if root.Left != nil {
        res = append(res, preorderTraversal(root.Left)...)
    }
    if root.Right != nil {
        res = append(res, preorderTraversal(root.Right)...)
    }
    return res
}
```



算法2：迭代

迭代的思想就是利用节点仅有左右指针，不断的搜索，同时借助 stack 数据结构，保存已经遍历过的节点。

其主要思路为：

1. 先取根节点值加入结果集，并将根节点入栈；持续检查其左子树，重复1，直到左子树为空【此时栈中保留的是一系列的左子树节点】
2. 出栈一次，检查最后一个左子树的右子树，重复1,2

```go
// date 2023/10/18
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    stack := make([]*TreeNode, 0, 16)
    res := make([]int, 0, 16)
    for root != nil || len(stack) != 0 {
        // 不断地查找左子树
        for root != nil {
            // 先将根节点加入结果集
            res = append(res, root.Val)
            // 根节点已经遍历过，入栈，后续出栈检查其右子树
            stack = append(stack, root)
            root = root.Left
        }
        // 此时栈中最后一个元素就是没有左子树节点
        // 取出并检查其右子树
        if len(stack) != 0 {
            root = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            root = root.Right
        }
    }
    return res
}
```

