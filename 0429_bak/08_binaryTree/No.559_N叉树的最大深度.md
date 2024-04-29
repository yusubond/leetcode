## 559 N叉树的最大深度-简单

题目：

给定一个 N 叉树，找到其最大深度。

最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。

N 叉树输入按层序遍历序列化表示，每组子节点由空值分隔（请参见示例）。



**解题思路**

```go
// date 2024/02/20
// 解法1，递归
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
    ans := 0
    var dfs func(root *Node) int
    dfs = func(root *Node) int {
        if root == nil {
            return 0
        }
        if len(root.Children) == 0 {
            if ans < 1 {
                ans = 1
            }
            return 1
        }
        temp := 0
        for i := 0; i < len(root.Children); i++ {
            res := dfs(root.Children[i])
            if res > temp {
                temp = res
            }
        }
        temp++
        if temp > ans {
            ans = temp
        }
        return temp
    }

    dfs(root)

    return ans
}

// 解法2，bfs，层序遍历
func maxDepth(root *Node) int {
    if root == nil {
        return 0
    }
    if len(root.Children) == 0 {
        return 1
    }
    ans := 0
    queue := make([]*Node, 0, 16)
    queue = append(queue, root)

    for len(queue) != 0 {
        n := len(queue)
        
        for i := 0; i < n; i++ {
            cur := queue[i]
            for j := 0; j < len(cur.Children); j++ {
                if cur.Children[j] != nil {
                    queue = append(queue, cur.Children[j])
                }
            }
        }
        queue = queue[n:]
        
        ans++
    }

    return ans
}
```

