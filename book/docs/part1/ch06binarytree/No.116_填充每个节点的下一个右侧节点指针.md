## 116 填充每个节点的下一个右侧节点-中等

题目：

给定一个 **完美二叉树** ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

```
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
```

填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 `NULL`。

初始状态下，所有 next 指针都被设置为 `NULL`。



分析：


```go
// date 2023/10/23
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    if root == nil {
        return nil
    }
	queue := make([]*Node, 0, 16)
    queue = append(queue, root)

    for len(queue) != 0 {
        n := len(queue)
        for i := 0; i < n; i++ {
            cur := queue[i]
            if i + 1 < n {
                cur.Next = queue[i+1]
            }
            if cur.Left != nil {
                queue = append(queue, cur.Left)
            }
            if cur.Right != nil {
                queue = append(queue, cur.Right)
            }
        }
        queue = queue[n:]
    }
    return root
}
```
