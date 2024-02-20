## 589 N叉树的前序遍历-简单

题目：

给定一个 n 叉树的根节点  root ，返回 其节点值的 前序遍历 。

n 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。



**解题思路**

算法1：递归

```go
// date 2022/10/21
func preorder(root *Node) []int {
    res := make([]int, 0, 16)
    if root == nil {
        return res
    }
    res = append(res, root.Val)
    for i := 0; i < len(root.Children); i++ {
        res = append(res, preorder(root.Children[i])...)
    }
    return res
}
```



算法2：迭代

```go
// date 2022/10/19
func preorder(root *Node) []int {
    res := make([]int, 0, 16)
    if root == nil {
        return res
    }
    stack := make([]*Node, 0, 16)
    stack = append(stack, root)

    for len(stack) != 0 {
        cur := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        res = append(res, cur.Val)
        for i := len(cur.Children)-1; i >= 0; i-- {  // 直接逆序存放，出栈时就是最左边的那个
            if cur.Children[i] != nil {
                stack = append(stack, cur.Children[i])
            }
        }
    }
    return res
}
```

