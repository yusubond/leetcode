## 426 将二叉搜索树转化为排序后的双向链表-中等

题目：

将一个 **二叉搜索树** 就地转化为一个 **已排序的双向循环链表** 。

对于双向循环列表，你可以将左右孩子指针作为双向循环链表的前驱和后继指针，第一个节点的前驱是最后一个节点，最后一个节点的后继是第一个节点。

特别地，我们希望可以 **就地** 完成转换操作。当转化完成以后，树中节点的左指针需要指向前驱，树中节点的右指针需要指向后继。还需要返回链表中最小元素的指针。



> **示例 1：**
>
> ```
> 输入：root = [4,2,5,1,3] 
> 
> 
> 输出：[1,2,3,4,5]
> 
> 解释：下图显示了转化后的二叉搜索树，实线表示后继关系，虚线表示前驱关系。
> ```
>
> **示例 2：**
>
> ```
> 输入：root = [2,1,3]
> 输出：[1,2,3]
> ```
>
> **示例 3：**
>
> ```
> 输入：root = []
> 输出：[]
> 解释：输入是空树，所以输出也是空链表。
> ```
>
> **示例 4：**
>
> ```
> 输入：root = [1]
> 输出：[1]
> ```



**解题思路**

根据访问根节点的顺序不同，树有三种遍历方式，前序、中序和后序。这道题给出的是二叉搜索树，而且要求给出排序的双向链表，那么递归版中序更合适。

因为递归中序遍历的时候，可以最先达到值最小的那个节点，把它保存为 first。

同时，第一个点也是 last 节点。然后递归遍历，将节点依次追加到 last 即可。

最后，再把 last 的后继指向 first，first 的前驱指向 last。

```go
// date 2024/01/11
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 * }
 */

func treeToDoublyList(root *Node) *Node {
    if root == nil {
        return root
    }
    var first, last *Node

    var inorder func(root *Node)
    inorder =  func(root *Node) {
        if root == nil {
            return
        }
        inorder(root.Left)
        if last != nil {
            last.Right = root
            root.Left = last
        } else {
            // last == nil
            // this is first
            first = root
        }
        // save root
        last = root
        inorder(root.Right)
    }
    inorder(root)
    last.Right = first
    first.Left = last
    return first
}
```

