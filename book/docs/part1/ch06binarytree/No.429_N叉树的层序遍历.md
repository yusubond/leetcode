## 429 N叉树的层序遍历-中等

题目：

给定一个 N 叉树，返回其节点值的*层序遍历*。（即从左到右，逐层遍历）。

树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。





分析：bfs

换汤不换药，还是老规矩，一层层的搜索。

```go
// date 2022/10/21
func levelOrder(root *Node) [][]int {
    res := make([][]int, 0, 16)
    if root == nil {
        return res
    }
    queue := make([]*Node, 0, 16)
    queue = append(queue, root)
    for len(queue) != 0 {
        n := len(queue)
        lRes := make([]int, 0, 16)
        for i := 0; i < n; i++ {
            cur := queue[i]
            lRes = append(lRes, cur.Val)
            for j := 0; j < len(cur.Children); j++ {
                if cur.Children[j] != nil {
                    queue = append(queue, cur.Children[j])
                }
            }
        }
        queue = queue[n:]
        res = append(res, lRes)
    }

    return res
}
```

