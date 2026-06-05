## 1669 合并两个链表-中等

题目：

给你两个链表 `list1` 和 `list2` ，它们包含的元素分别为 `n` 个和 `m` 个。

请你将 `list1` 中下标从 `a` 到 `b` 的全部节点都删除，并将`list2` 接在被删除节点的位置。

下图中蓝色边和节点展示了操作后的结果：

![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/11/28/fig1.png)

请你返回结果链表的头指针。



**提示：**

- `3 <= list1.length <= 104`
- `1 <= a <= b < list1.length - 1`
- `1 <= list2.length <= 104`



**解题思路**：

因为题目中已经说明 a、b的范围，a 比大于1，所以新链表的表头一定是 list1。

只需要考虑链表的区间删除即可。思路是遍历 list1 找到 a 的前驱，然后遍历 list2 找到最后一个节点，然后继续遍历 list1 找到 b，最后把 a 的前驱下一跳指向 list2，list2 的表尾节点指向 b 的后驱节点即可。

```go
// date 2024/01/12
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    apre, cur := list1, list1
    for a > 0 {
        apre = cur
        cur = cur.Next
        a--
        b--
    }
    apre.Next = list2
    l2Last := list2
    for l2Last.Next != nil {
        l2Last = l2Last.Next
    }
    for b > 0 {
        cur = cur.Next
        b--
    }
    l2Last.Next = cur.Next
    return list1
}
```

