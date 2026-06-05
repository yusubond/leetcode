## 590 N叉树的后序遍历-简单

给定一个 n 叉树的根节点 root ，返回 其节点值的 后序遍历 。

n 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。



分析：

算法1：递归

```go
// date 2022/10/12
func postorder(root *Node) []int {
    if root == nil {
        return []int{}
    }
    res := make([]int, 0, 16)
    for i := 0; i < len(root.Children); i++ {
        res = append(res, postorder(root.Children[i])...)
    }
    res = append(res, root.Val)
    return res
}
```



算法2：迭代【推荐该算法】

```go
// date 2023/10/31
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func postorder(root *Node) []int {
    if root == nil {
        return []int{}
    }
    res := make([]int, 0, 16)
    stack := make([]*Node, 0, 16)
    stack = append(stack, root)
    visited := make(map[*Node]bool, 16)
    for len(stack) != 0 {
        cur := stack[len(stack)-1]
        // 如果当前节点为叶子节点，或者所有子结点已经遍历过，直接追加结果并更新栈
        if len(cur.Children) == 0 || visited[cur] {
            res = append(res, cur.Val)
            stack = stack[:len(stack)-1]
            continue
        }
        // 入栈当前节点的所有子结点
        for i := len(cur.Children)-1; i >= 0; i-- {
            if cur.Children[i] != nil {
                stack = append(stack, cur.Children[i])
            }
        }
        // 标记当前结点的所有子结点已经入栈
        visited[cur] = true
    }
    return res
}
```

