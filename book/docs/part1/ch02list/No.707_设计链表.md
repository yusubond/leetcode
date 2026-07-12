## 707 设计链表-中等

题目：

你可以选择使用单链表或者双链表，设计并实现自己的链表。

单链表中的节点应该具备两个属性：`val` 和 `next` 。`val` 是当前节点的值，`next` 是指向下一个节点的指针/引用。

如果是双向链表，则还需要属性 `prev` 以指示链表中的上一个节点。假设链表中的所有节点下标从 **0** 开始。

实现 `MyLinkedList` 类：

- `MyLinkedList()` 初始化 `MyLinkedList` 对象。
- `int get(int index)` 获取链表中下标为 `index` 的节点的值。如果下标无效，则返回 `-1` 。
- `void addAtHead(int val)` 将一个值为 `val` 的节点插入到链表中第一个元素之前。在插入完成后，新节点会成为链表的第一个节点。
- `void addAtTail(int val)` 将一个值为 `val` 的节点追加到链表中作为链表的最后一个元素。
- `void addAtIndex(int index, int val)` 将一个值为 `val` 的节点插入到链表中下标为 `index` 的节点之前。如果 `index` 等于链表的长度，那么该节点会被追加到链表的末尾。如果 `index` 比长度更大，该节点将 **不会插入** 到链表中。
- `void deleteAtIndex(int index)` 如果下标有效，则删除链表中下标为 `index` 的节点。



**关键观察**

1. **哑结点（Dummy Node）消除边界判断**：使用哑 head 和哑 tail，所有插入/删除都在"中间"完成，无需特殊处理头尾插入。这是链表题的经典技巧——哨兵节点让代码逻辑统一。

2. **双向链表优于单向**：`addAtTail` 在单链表中需 O(n) 遍历到末尾；双向链表借助 tail.prev 直接定位，O(1) 即可。同理，`deleteAtIndex` 删除最后一个节点时也需要前驱，双向链表天然可回溯。

3. **`size` 字段统一校验**：用 `size` 记录节点数，`index` 有效性判断只需一行 `index < 0 || index >= size`。



**解法一：双向链表 + 哑结点**

```go
// date 2023/10/17
type MyListNode struct {
    Val        int
    prev, next *MyListNode
}

type MyLinkedList struct {
    size       int
    head, tail *MyListNode
}

func Constructor() MyLinkedList {
    list := MyLinkedList{
        size: 0,
        head: &MyListNode{},
        tail: &MyListNode{},
    }
    list.head.next = list.tail
    list.tail.prev = list.head
    return list
}

func (this *MyLinkedList) Get(index int) int {
    if index >= 0 && index < this.size {
        cur := this.head.next
        for index > 0 {
            cur = cur.next
            index--
        }
        return cur.Val
    }
    return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
    node := &MyListNode{Val: val}
    node.next = this.head.next
    node.prev = this.head

    this.head.next.prev = node
    this.head.next = node
    this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
    node := &MyListNode{Val: val}
    node.prev = this.tail.prev
    node.next = this.tail

    this.tail.prev.next = node
    this.tail.prev = node
    this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
    if index > this.size {
        return
    }
    if index == this.size {
        this.AddAtTail(val)
        return
    }
    if index == 0 {
        this.AddAtHead(val)
        return
    }
    node := &MyListNode{Val: val}
    pre := this.head
    for index > 0 {
        pre = pre.next
        index--
    }
    node.next = pre.next
    node.prev = pre
    pre.next.prev = node
    pre.next = node
    this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
    if index < 0 || index >= this.size {
        return
    }
    pre := this.head
    for index > 0 {
        pre = pre.next
        index--
    }

    pre.next = pre.next.next
    pre.next.prev = pre
    this.size--
}
```



**Get 遍历优化**

双向链表的优势：根据 index 位置选择从 head 或 tail 出发，将遍历次数从 O(index) 降到 O(min(index, size-index))。

```go
// date 2026/07/12
func (this *MyLinkedList) Get(index int) int {
    if index < 0 || index >= this.size {
        return -1
    }
    var cur *MyListNode
    if index < this.size/2 {
        // 从前向后
        cur = this.head.next
        for i := 0; i < index; i++ {
            cur = cur.next
        }
    } else {
        // 从后向前
        cur = this.tail.prev
        for i := this.size - 1; i > index; i-- {
            cur = cur.prev
        }
    }
    return cur.Val
}
```

`AddAtIndex` 和 `DeleteAtIndex` 中寻找前驱的循环也可同样优化。

> 这个优化来自 `template/double_linkedlist.go` 中的实现思路——双向链表的核心价值不仅是 O(1) 尾操作，还在于两端可达带来的遍历减半。



**解法二：单链表 + 哑头结点**

单链表没有 `prev` 指针，`addAtTail` 需要 O(n) 遍历到底，但结构更轻量。适合只需头部操作的场景。

```go
// date 2026/07/12
type SinglyNode struct {
    Val  int
    Next *SinglyNode
}

type MyLinkedList2 struct {
    size int
    head *SinglyNode // 哑头结点
}

func Constructor2() MyLinkedList2 {
    return MyLinkedList2{head: &SinglyNode{}}
}

func (this *MyLinkedList2) Get(index int) int {
    if index < 0 || index >= this.size {
        return -1
    }
    cur := this.head.Next
    for i := 0; i < index; i++ {
        cur = cur.Next
    }
    return cur.Val
}

func (this *MyLinkedList2) AddAtHead(val int) {
    node := &SinglyNode{Val: val}
    node.Next = this.head.Next
    this.head.Next = node
    this.size++
}

func (this *MyLinkedList2) AddAtTail(val int) {
    pre := this.head
    for pre.Next != nil {
        pre = pre.Next
    }
    pre.Next = &SinglyNode{Val: val}
    this.size++
}

func (this *MyLinkedList2) AddAtIndex(index int, val int) {
    if index > this.size {
        return
    }
    pre := this.head
    for i := 0; i < index; i++ {
        pre = pre.Next
    }
    node := &SinglyNode{Val: val}
    node.Next = pre.Next
    pre.Next = node
    this.size++
}

func (this *MyLinkedList2) DeleteAtIndex(index int) {
    if index < 0 || index >= this.size {
        return
    }
    pre := this.head
    for i := 0; i < index; i++ {
        pre = pre.Next
    }
    pre.Next = pre.Next.Next
    this.size--
}
```

注意单链表 `DeleteAtIndex` 不需要 `prev` 回指操作，代码更简洁——但代价是 `addAtTail` 必须 O(n) 遍历。



**复杂度对比**

| 操作 | 双向链表 | 双向链表（优化遍历） | 单链表 |
|------|---------|-------------------|--------|
| `get` | O(index) | O(min(index, n-index)) | O(index) |
| `addAtHead` | O(1) | O(1) | O(1) |
| `addAtTail` | O(1) | O(1) | O(n) |
| `addAtIndex` | O(index) | O(min(index, n-index)) | O(index) |
| `deleteAtIndex` | O(index) | O(min(index, n-index)) | O(index) |
| 空间 | 每个节点 2 个指针 | 同左 | 每个节点 1 个指针 |



**与相邻题的关系**

| 题目 | 关系 |
|------|------|
| No.146 LRU 缓存 | 双向链表 + 哈希表的经典组合，本质就是本题的双向链表结构 |
| No.460 LFU 缓存 | 在 No.146 基础上增加了频率维度 |
| No.206 反转链表 | 单链表基本操作，本题 `template/linkedlist.go` 中 `Reverse()` 的实现 |
| No.92 反转链表 II | 区间反转，需要精确定位前驱和后继，与本题的 `addAtIndex` 定位模式相同 |

> 本题是设计题，更完整的链表操作集合（翻转、排序、删除特定节点等）见 `template/double_linkedlist.go` 和 `template/linkedlist.go`。
