## 725 分隔链表-中等

题目：

给你一个头节点为 head 的单链表和一个整数 k，请你设计一个算法将链表分隔为 k 个连续的部分。

每部分的长度尽可能的相等：任意两部分的长度差距不能超过 1。这可能导致有些部分为 null。【也就是不够分的时候，用空链表填充。】

这 k 个部分应该按照在链表中出现的顺序排列，并且排列在前面的部分的长度应该大于或等于排在后面的长度。

请返回一个由上述 k 部分组成的数组。



分析：

如果把链表的长度记为 x，那么该问题可以转化为 如何尽可能把 x 等分 k 份。这个问题容易解决，直接对 x 进行 k 求余和整除

```go
base := x / k
rmd := x % k  // remainder
```

然后构造有 k 个元素的数组，其中前 rmd 个 元素大小为 base + 1，其他元素为 base。

那么该数组表示的就是最终结果集各个链表的元素个数。根据该数组就可以的链表分隔了。



```go
// date 2023/10/12
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
    res := make([]*ListNode, k, k)
    if head == nil {
        return res
    }
    pre := head
    total := 0
    for pre != nil {
        total++
        pre = pre.Next
    }

    sglen := getSingleLenRes(total, k)
    pre = head
    for i, x := range sglen {
        dy := &ListNode{}
        dpre := dy
        for x > 0 {
            dpre.Next = pre
            pre = pre.Next
            dpre = dpre.Next
            x--
        }
        dpre.Next = nil
        res[i] = dy.Next
    }
    return res
}

func getSingleLenRes(total, k int) []int {
    res := make([]int, k, k)
    idx := k

    base := total / k
    rmd := total % k
    for k > 0 {
        if rmd > 0 {
            res[idx-k] = base + 1
            rmd--
        } else {
            res[idx-k] = base
        }
        k--
    }
    return res
}
```

