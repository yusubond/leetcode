## 255 验证二叉搜索树的前序遍历序列

题目：

给定一个 **无重复元素** 的整数数组 `preorder` ， *如果它是以二叉搜索树的**先序遍历**排列* ，返回 `true` 。



> **示例 1：**
>
> ![img](https://assets.leetcode.com/uploads/2021/03/12/preorder-tree.jpg)
>
> ```
> 输入: preorder = [5,2,1,3,6]
> 输出: true
> ```
>
> **示例 2：**
>
> ```
> 输入: preorder = [5,2,6,1,3]
> 输出: false
> ```



**解题思路**

前序遍历为：根-左-右

二叉搜索树的特点是，根节点大于左子树，小于右子树。

那么，对于给定的一个序列，如果当前节点值小于上一个节点值，那么该节点是上一个节点的左子树；如果大于上一个节点值，那么该节点是上一个节点的右子树。

基于这个发现，我们可以维护一个单调栈，从栈底到栈顶元素单调递减。



具体为：对给定的前序遍历序列，依次遍历每个值。如果栈为空，或者当前值小于栈顶值，那么肯定是根节点或者左子树，则入栈。如果当前值大于栈顶值，那么肯定是某个节点的右子树，则将栈内小于当前值的所有元素全部弹出，将当前值入栈。

最后弹出的肯定是当前值所在节点的父节点，保存它。

根据二叉树搜索树的特点，父节点右子树中的任何一个值都要大于父节点，即刚刚保存的值。

如果不满足，则该遍历序列不是二叉树搜索树正确的先序遍历序列。



```go
// date 2024/02/19
// 单调栈
func verifyPreorder(preorder []int) bool {
    stack := make([]int, 0, 16)
    min := math.MinInt32
    for _, v := range preorder {
        if v <= min {
            return false
        }

        for len(stack) > 0 && v > stack[len(stack)-1] {
            min = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        }

        stack = append(stack, v)
    }

    return true
}
```

