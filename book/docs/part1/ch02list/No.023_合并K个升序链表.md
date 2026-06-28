## 023 合并K个升序链表-困难

题目：

给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。



**解题思路**

复习递归算法，掌握自底向上和自上向下的递归方法。

- 最朴素的算法，顺次两两合并，详见解法1。每次合并时前一条已经累计更长，后面的链表会被反复比较，时间复杂度 O(N·k)（N 为节点总数，k 为链表数），空间复杂度 O(1)。
- 分治算法，类似归并排序，详见解法2。时间复杂度 O(N log k)，空间复杂度 O(log k) 递归栈。
- 优先队列 + 插入排序，详见解法3。data 始终保持 Val 升序，Push 用插入排序定位后插入，Pop 取首元素即可。时间复杂度 O(N·k)，空间复杂度 O(k)。
- 堆（heap），详见解法4。Push 末尾追加后 siftUp，Pop 交换首尾后 siftDown。时间复杂度 O(N log k)，空间复杂度 O(k)。

**复杂度对比**

设链表个数为 `k`，所有节点总数为 `N`。

| 方法 | 时间 | 空间 |
|------|------|------|
| ① 顺序两两合并 | O(N·k) | O(1) |
| ② 分治归并 | O(N log k) | O(log k) |
| ③ 优先队列 + 插入排序 | O(N·k) | O(k) |
| ④ 堆（siftUp / siftDown） | O(N log k) | O(k) |

```go
// date 2023/10/17
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 解法1
// 最朴素的顺序两两合并
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    if len(lists) == 1 {
        return lists[0]
    }
    k := len(lists)
    l1, l2 := mergeKLists(lists[:k-1]), lists[k-1]
    dumy := &ListNode{}
    pre := dumy
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            pre.Next = l1
            l1 = l1.Next
        } else {
            pre.Next = l2
            l2 = l2.Next
        }
        pre = pre.Next
    }
    if l1 != nil {
        pre.Next = l1
    }
    if l2 != nil {
        pre.Next = l2
    }
    return dumy.Next
}

// 解法2
// 分治思想，归并排序思想的合并
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    
    var merge func(lists []*ListNode, left, right int) *ListNode
    merge = func(lists []*ListNode, left, right int) *ListNode {
        if left == right {
            return lists[left]
        }
        if left > right {
            return nil
        }
        mid := left + (right- left)/2
        return mergeTwoList(merge(lists, left, mid), merge(lists, mid+1, right))
    }

    return merge(lists, 0, len(lists)-1)
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    dummy := &ListNode{}
    cur := dummy
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            cur.Next = l1
            l1 = l1.Next
        } else {
            cur.Next = l2
            l2 = l2.Next
        }
        cur = cur.Next
    }
    if l1 != nil {
        cur.Next = l1
    }
    if l2 != nil {
        cur.Next = l2
    }
    return dummy.Next
}
```

```go
// date 2026/06/28
// 解法3
// 优先队列（有序数组实现）：Push 用插入排序维持 data 单调递增，Pop 取首元素
// 时间复杂度 O(N·k)，空间复杂度 O(k)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    dumy := &ListNode{}
    pq := newPriQueue()
    for _, v := range lists {
        if v != nil {
            pq.Push(v)
        }
    }

    cur := dumy
    for pq.Len() > 0 {
        x := pq.Pop()
        cur.Next = x
        cur = cur.Next

        if x.Next != nil {
            pq.Push(x.Next)
        }
    }

    return dumy.Next
}

// 优先队列，升序
type priQueue struct {
    data []*ListNode
}

func newPriQueue() *priQueue {
    return &priQueue{
        data: make([]*ListNode, 0, 16),
    }
}

func (p *priQueue) Len() int {
    return len(p.data)
}

// Push 用插入排序维持 data 单调递增（Val 升序）
// 时间复杂度 O(k)，空间复杂度 O(1)
func (p *priQueue) Push(v *ListNode) {
    // 1. 线性扫描找到插入位置：第一个大于 v.Val 的位置 i
    i := 0
    for i < len(p.data) && p.data[i].Val <= v.Val {
        i++
    }
    // 2. 在位置 i 插入 v：[i, n) 整体后移一位
    p.data = append(p.data, nil)
    copy(p.data[i+1:], p.data[i:])
    p.data[i] = v
}

// Pop 取出队首（最小）节点；O(1)，依赖 Push 维持的有序性
func (p *priQueue) Pop() *ListNode {
    v := p.data[0]
    p.data = p.data[1:]
    return v
}
```

```go
// date 2026/06/28
// 解法4
// 堆（heap）实现：Push 末尾追加后 siftUp，Pop 交换首尾后 siftDown
// 时间复杂度 O(N log k)，空间复杂度 O(k)
// 与解法3 共用同一接口（Push / Pop / Len），仅内部实现不同

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    dumy := &ListNode{}
    pq := newPriQueueHeap()
    for _, v := range lists {
        if v != nil {
            pq.Push(v)
        }
    }

    cur := dumy
    for pq.Len() > 0 {
        x := pq.Pop()
        cur.Next = x
        cur = cur.Next

        if x.Next != nil {
            pq.Push(x.Next)
        }
    }

    return dumy.Next
}

// 优先队列，最小堆（按 Val 升序）
type priQueueHeap struct {
    data []*ListNode
}

func newPriQueueHeap() *priQueueHeap {
    return &priQueueHeap{
        data: make([]*ListNode, 0, 16),
    }
}

func (p *priQueueHeap) Len() int {
    return len(p.data)
}

// Push 追加到末尾后沿父节点上浮，保持最小堆性质
// 时间复杂度 O(log k)
func (p *priQueueHeap) Push(v *ListNode) {
    p.data = append(p.data, v)
    // siftUp
    i := len(p.data) - 1
    for i > 0 {
        parent := (i - 1) / 2
        if p.data[parent].Val <= p.data[i].Val {
            break
        }
        p.data[parent], p.data[i] = p.data[i], p.data[parent]
        i = parent
    }
}

// Pop 取出堆顶（最小）节点；末尾元素提到堆顶后下沉
// 时间复杂度 O(log k)
func (p *priQueueHeap) Pop() *ListNode {
    if len(p.data) == 0 {
        return nil
    }
    top := p.data[0]
    n := len(p.data) - 1
    p.data[0] = p.data[n]
    p.data = p.data[:n]
    // siftDown
    for i := 0; ; {
        left, right := 2*i+1, 2*i+2
        if left >= len(p.data) {
            break
        }
        smallest := left
        if right < len(p.data) && p.data[right].Val < p.data[left].Val {
            smallest = right
        }
        if p.data[i].Val <= p.data[smallest].Val {
            break
        }
        p.data[i], p.data[smallest] = p.data[smallest], p.data[i]
        i = smallest
    }
    return top
}
```

