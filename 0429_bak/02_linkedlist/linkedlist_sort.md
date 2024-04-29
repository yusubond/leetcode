## 单链表的排序算法

### 1、插入排序

插入排序可以通过直接交换节点得到。

```go
/*
功能：插入排序
参数：传入链表头指针作为参数，返回排序后的头指针
说明：时间复杂度O(n^2)，空间复杂度O(1)
第一步：选择插入的位置;
第二步：如果在已排序的链表的尾部，则直接加入到尾部，否则，插入到选好的位置。注意：需要保存前驱节点。
 */

func InsertSort(head *ListNode) *ListNode {
  if head == nil || head.Next == nil { return head }
  var pre, p, start, end *ListNode
  
  p, start, end, pre = head.Next, head, head, head
  var next *ListNode
  
  for p != nil {
    next, pre = p.Next, start
    for next != p && p.Val >= next.Val {
      next = next.Next
      pre = pre.Next
    }
    // 追加到已经排序的队尾
    if next == p {
      pend = p
    } else {
      pend.Next = p.Next
      p.Next = next
      pre.Next = p
    }
    p = pend.Next
  }
  return head
}
```

### 2、选择排序

```
/*
功能：选择排序(选择未排序序列中的最值，然后放到已排序的最前面或最后面，只交换节点的值)
参数：输入链表的头指针，返回排序后的头指针
说明：时间复杂度O(n^2)，空间复杂度O(1)
 */
Node *SelectSort(Node *phead) {
  if(phead == NULL || phead->next == NULL) return phead;
  Node *pend = phead;
  while(pend->next != NULL) {
    Node *minNode = pend->next;
    Node *p = minNode->next;
    while(p != NULL) {
      if(p->val < minNode->val)
        minNode = p;
      p = p->next;
    }
    swap(minNode->val, pend->val);
    pend = pend->next;
  }
  return phead;
}
```

### 3、快速排序

```
/*
功能：快速排序，链表指向下一个元素的特性，partition是选择左闭合区间
第一步，partiton，因为链表不支持随机访问元素，因此partiton中选取第一个节点值作为基准
第二步，排序
参数：输入链表头指针，输出排序后的头指针
说明：平均时间复杂度O(nlogn)，空间复杂度O(1)
 */
Node *Partition(Node *low, Node *high) {
  //左闭合区间[low, high)
  int base = low->val;
  Node *location = low;
  for(Node *i = low->next; i != high; i = i->next) {
    if(i->val > base) {
      location = location->next;
      swap(location->val, i->val);
    }
  }
  swap(low->val, location->val);
  return location;
}
void QuickList(Node *phead, Node *tail) {
  //左闭合区间[phead, tail)
  if(phead != tail && phead->next != tail) {
    Node *mid = Partition(phead, tail);
    QuickList(phead, mid);
    QuickList(mid->next, tail);
  }
}
Node *QuickSort(Node *phead) {
  if(phead == NULL || phead->next == NULL) return phead;
  QuickList(phead, NULL);
  return phead;
}
```

### 4、归并排序

```
/*
功能：归并排序，交换链表节点
第一步：先写归并函数；第二步：利用快慢指针找到链表中点，然后递归对子链进行排序
参数：输出链表的头指针，输出排序后的头指针
说明：时间复杂度O(nlogn)，空间复杂度O(1)。归并排序应该算是链表排序中最佳的选择，保证最好和最坏的时间复杂度都是O(nlogn)，而且在将空间复杂度由数组中O(n)，降到链表中的O(1)
 */
Node *merge(Node *phead1, Node *phead2) {
  if(phead1 == NULL) return phead2;
  if(phead2 == NULL) return phead1;
  Node *res, *p;
  if(phead1->val < phead2->val) {
    res = phead1;
    phead1 = phead1->next;
  } else {
    res = phead2;
    phead2 = phead2->next;
  }
  p = res;
  while(phead1 != NULL && phead2 != NULL) {
    if(phead1->val < phead2->val) {
      p->next = phead1;
      phead1 = phead1->next;
    } else {
      p->next = phead2;
      phead2 = phead2->next;
    }
    p = p->next;
  }
  if(phead1 != NULL) p->next = phead1;
  else if(phead2 != NULL) p->next = phead2;
  return res;
}
Node *MergeSort(Node *phead) {
  if(phead == NULL || phead->next == NULL) return phead;
  Node *fast = phead;
  Node *slow = phead;
  while(fast->next != NULL && fast->next->next != NULL) {
    fast = fast->next->next;
    slow = slow->next;
  }
  fast = slow;
  slow = slow->next;
  fast->next = NULL;
  fast = MergeSort(phead);
  slow = MergeSort(slow);
  return merge(fast, slow);
}
```
