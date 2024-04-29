## 1522 N叉树的直径-中等

题目：

给定一棵 `N 叉树` 的根节点 `root` ，计算这棵树的直径长度。

N 叉树的直径指的是树中任意两个节点间路径中 **最长** 路径的长度。这条路径可能经过根节点，也可能不经过根节点。

*（N 叉树的输入序列以层序遍历的形式给出，每组子节点用 null 分隔）*



> **示例 1：**
>
> ![img](https://assets.leetcode.com/uploads/2020/07/19/sample_2_1897.png)
>
> ```
> 输入：root = [1,null,3,2,4,null,5,6]
> 输出：3
> 解释：直径如图中红线所示。
> ```
>
> **示例 2：**
>
> **![img](https://assets.leetcode.com/uploads/2020/07/19/sample_1_1897.png)**
>
> ```
> 输入：root = [1,null,2,null,3,4,null,5,null,6]
> 输出：4
> ```
>
> **示例 3：**
>
> ![img](https://assets.leetcode.com/uploads/2020/07/19/sample_3_1897.png)
>
> ```
> 输入: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
> 输出: 7
> ```



**解题思路**

所谓求直径，就是遍历每个节点，求其左右子树深度的最大和。

可用深度优先搜索解决。具体为：dfs 返回左右子树的最大值，并且返回前更新全局解 ans。

```go
// date 2024/02/20
// dfs
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func diameter(root *Node) int {
    // 就是求每个节点的子树深度最大和
    ans := 0

    var dfs func(root *Node) int
    dfs = func(root *Node) int {
        if root == nil {
            return 0
        }
        if len(root.Children) == 0 {
            return 1
        }
        chil := make([]int, 0, 16)
        for i := 0; i < len(root.Children); i++ {
            res := dfs(root.Children[i])
            chil = append(chil, res)
        }
        sort.Slice(chil, func(i, j int) bool {
            return chil[i] > chil[j]
        })
        n := len(chil)
        temp := 0
        if n == 1 {
            if chil[0] > ans {
                ans = chil[0]
            }
            temp = chil[0]+1
        } else if n > 1 {
            if chil[0] + chil[1] > ans {
                ans = chil[0] + chil[1]
            }
            temp = chil[0]+1
        }
        return temp
    }

    dfs(root)

    return ans
}
```

