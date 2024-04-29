## 138 随机链表的复制-中等

题目：

给你一个长度为 `n` 的链表，每个节点包含一个额外增加的随机指针 `random` ，该指针可以指向链表中的任何节点或空节点。

构造这个链表的 **[深拷贝](https://baike.baidu.com/item/深拷贝/22785317?fr=aladdin)**。 深拷贝应该正好由 `n` 个 **全新** 节点组成，其中每个新节点的值都设为其对应的原节点的值。新节点的 `next` 指针和 `random` 指针也都应指向复制链表中的新节点，并使原链表和复制链表中的这些指针能够表示相同的链表状态。**复制链表中的指针都不应指向原链表中的节点** 。

例如，如果原链表中有 `X` 和 `Y` 两个节点，其中 `X.random --> Y` 。那么在复制链表中对应的两个节点 `x` 和 `y` ，同样有 `x.random --> y` 。

返回复制链表的头节点。

用一个由 `n` 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 `[val, random_index]` 表示：

- `val`：一个表示 `Node.val` 的整数。
- `random_index`：随机指针指向的节点索引（范围从 `0` 到 `n-1`）；如果不指向任何节点，则为 `null` 。

你的代码 **只** 接受原链表的头节点 `head` 作为传入参数



分析：

算法1：【推荐该算法】

哈希表+回溯

哈希表的 key 和 value 分别是原链表中的节点 和 结果集中的节点，用于表示那些原来的节点已经被创建，

```go
// date 2023/10/18
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil {
        return head
    }
    cache := make(map[*Node]*Node, 1024)
    var deepcopy func(h *Node) *Node
    deepcopy = func(h *Node) *Node {
        if h == nil {
            return h
        }
        old, ok := cache[h]
        if ok {
            return old
        }
        node := &Node{Val: h.Val}
        cache[h] = node
        node.Next = deepcopy(h.Next)
        node.Random = deepcopy(h.Random)
        return node
    }

    return deepcopy(head)
}
```





算法2：

原链表遍历两次：

第一次：根据 next 节点，创建结果集，并将原链表中的节点和索引存入 map。新创建的节点，存到数组。

第二次，遍历原链表，根据 random 修改结果集中的 random 指针。

```go
// date 2023/10/18
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil {
        return head
    }

    dumy := &Node{}
    dpre := dumy
    resA := make([]*Node, 0, 1024)
    pre := head
    cache := make(map[*Node]int, 1024)
    index := int(0)
    for pre != nil {
        cache[pre] = index
        nd := &Node{Val: pre.Val}
        dpre.Next = nd
        dpre = dpre.Next
        pre = pre.Next

        resA = append(resA, nd)
        index++
    }

    pre = head
    index = 0
    for pre != nil {
        if pre.Random != nil {
            if idx, ok := cache[pre.Random]; ok {
                resA[index].Random = resA[idx]
            }
        }
        pre = pre.Next
        index++
    }

    return dumy.Next
}
```



