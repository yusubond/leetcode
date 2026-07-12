## 148 排序链表-中等

题目：

给你链表的头节点 `head` ，请你将其按 **升序** 排列并返回排序后的链表。

**进阶**：你可以在 `O(n log n)` 时间复杂度和**常数级空间复杂度**下，对链表进行排序吗？



> **示例 1：**
>
> ```
> 输入：head = [4,2,1,3]
> 输出：[1,2,3,4]
> ```
>
> **示例 2：**
>
> ```
> 输入：head = [-1,5,3,4,0]
> 输出：[-1,0,3,4,5]
> ```
>
> **示例 3：**
>
> ```
> 输入：head = []
> 输出：[]
> ```



**解题思路**

要在 `O(n log n)` 内排序，可选的有快排、归并、堆排。链表不支持随机访问，**快排的划分效率低、堆排需要下标**，而**归并排序**天然契合链表——它只需顺序遍历，且合并操作就是 [No.021 合并两个有序链表](./No.021_合并两个有序链表.md)。

下面给出两种归并实现：自顶向下（递归，最易写）和自底向上（迭代，**真正满足进阶的常数空间**）。

### 方法一：自顶向下归并（递归）✅ 易实现

三步：

1. **找中点**：快慢指针，`fast` 走两步、`slow` 走一步；为方便断开，让 `fast` 先出发一步（`fast = head.Next`），这样 `slow` 最终落在左半段的最后一个节点；
2. **从中点断开**为两条子链，分别递归排序；
3. **合并**两条有序子链。

```go
// date 2026/07/12
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val  int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    // 快慢指针找中点；fast 先走一步，使 slow 停在左半段末尾，便于断开
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    mid := slow.Next
    slow.Next = nil // 断开成 [head..slow] 和 [mid..]

    left := sortList(head)
    right := sortList(mid)
    return mergeTwoList(left, right)
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    cur := dummy
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
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
    } else {
        cur.Next = l2
    }
    return dummy.Next
}
```

**手动跑示例 1** `4→2→1→3`：

```
sortList(4→2→1→3)  断成 4→2 | 1→3
├─ sortList(4→2)   断成 4 | 2     → merge 得 2→4
├─ sortList(1→3)   断成 1 | 3     → merge 得 1→3
└─ merge(2→4, 1→3)               → 1→2→3→4
```

> **空间**：递归深度 `O(log n)`，**不满足**进阶的常数空间要求。下面方法二解决这个问题。

### 方法二：自底向上归并（迭代）⭐ 满足 O(1) 空间

不用递归，改为按**子段长度**逐轮翻倍地两两合并：

1. 子段长度从 `size = 1` 开始；
2. 每轮从左到右：取两段长度至多为 `size` 的子链，合并后接在已排序部分末尾；
3. 本轮结束后 `size *= 2`，直到 `size ≥ n`。

全程用 dummy 节点串接，只有几个指针变量，**额外空间 O(1)**。

```go
func sortListBottomUp(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    // 求链表长度
    n := 0
    for cur := head; cur != nil; cur = cur.Next {
        n++
    }

    dummy := &ListNode{Next: head}
    for size := 1; size < n; size <<= 1 { // 每轮子段长度翻倍
        tail := dummy     // 已合并部分的尾
        cur := dummy.Next // 未处理部分的头
        for cur != nil {
            left := cur
            right := cut(left, size) // 切下 left（至多 size 个），返回剩余头
            cur = cut(right, size)   // 切下 right（至多 size 个），返回剩余头
            tail = mergeAppend(left, right, tail) // 合并并接在 tail 后，更新尾
        }
    }
    return dummy.Next
}

// cut 从 head 切下至多 n 个节点：把切下段尾部 Next 置 nil，返回剩余链表头
func cut(head *ListNode, n int) *ListNode {
    for n > 1 && head != nil {
        head = head.Next
        n--
    }
    if head == nil {
        return nil
    }
    next := head.Next
    head.Next = nil
    return next
}

// mergeAppend 合并 l1、l2，接在 prev 之后，返回合并段的尾节点
func mergeAppend(l1, l2, prev *ListNode) *ListNode {
    cur := prev
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
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
    } else {
        cur.Next = l2
    }
    for cur.Next != nil { // 移到本段末尾，作为下一对合并的前驱
        cur = cur.Next
    }
    return cur
}
```

**手动跑示例 1** `4→2→1→3`（n=4）：

| 轮次 size | 合并过程 | 结果 |
|----------|---------|------|
| size=1 | (4)(2)→2→4，(1)(3)→1→3 | 2→4→1→3 |
| size=2 | (2→4)(1→3)→1→2→3→4 | 1→2→3→4 |
| size=4 | `4 < 4` 不成立，结束 | — |

### 复杂度分析

| 方法 | 时间 | 空间 | 备注 |
|------|------|------|------|
| 自顶向下（递归） | O(n log n) | O(log n) 递归栈 | 代码最简，面试首选 |
| 自底向上（迭代） | O(n log n) | **O(1)** | 满足进阶的常数空间 |

每轮合并总工作量是 O(n)（所有子段加起来遍历一遍），共 `log n` 轮，故时间 O(n log n)。

**如何选择**：面试中写出自顶向下即可；若面试官追问「常数空间」，再给出自底向上版本。

### 相邻题

- [No.147 对链表进行插入排序](./No.147_对链表进行插入排序.md)：同为链表排序，但是 O(n²) 的插入排序，可作为对照——体会归并把时间降到 O(n log n) 的意义。
- [No.021 合并两个有序链表](./No.021_合并两个有序链表.md)：本题的核心子过程。
- [No.023 合并 K 个升序链表](./No.023_合并K个升序链表.md)：合并的进阶——多条链表用最小堆或分治归并。
- [No.876 链表的中间节点](./No.876_链表的中间节点.md)：快慢指针找中点的基础题。
