## 对线性表进行归并排序

  * 1.[算法流程](#算法流程)
  * 2.[代码实现](#代码实现)

### 算法流程

归并排序通常是线性表(即，链表)排序的优先选择算法。链表的慢随机访问特性使得某些排序算法(比如，快速排序)的执行效率偏低，也使得某些算法(比如，堆排序)难以实现。

链表的归并排序伪代码如下：

```
// Head is the head node of a linked list
MergeSort(Head)
1. 如果头节点为空或只有一个节点，则返回
2. 其次，将链表分成两个子链表
  Split(Head, &a, &b); // a, b分别为两个子链表
3. 对两个子链表进行归并排序
  MergeSort(a);
  MergeSort(b);
4. 合并已经排序的a, b两个子链表，并更新头节点
  *Head = Merge(a, b);
```

时间复杂度为:O(nLogn)

### 代码实现

* C++

```
struct Node{
  int val;
  struct Node* next;
}
struct Node* Merge(struct Node* a, struct Node* b);
void Split(struct Node* head, struct Node** front, struct Node** back);
void MergeSort(struct Node** head) {
  struct Node* phead = *head;
  struct Node* a;
  struct Node* b;
  if(phead == NULL | phead->next == NULL)
    return;
  Split(head, &a, &b);
  MergeSort(&a);
  MergeSort(&b);
  *head = Merge(a, b);
}

struct Node* Merge(struct Node* a, struct Node* b) {
  struct Node* result = NULL;
  if(a == NULL)
    return b;
  else if (b == NULL)
    return a;
  if(a->val <= b->val) {
    result = a;
    result->next = Merge(a->next, b);
  } else {
    result = b;
    result->next = Merge(a, b->next);
  }
  return result;
}

void Split(struct Node* head, struct Node** front, struct Node** back) {
  struct Node* fast;
  struct Node* slow;
  if(head == NULL | head->next == NULL) {
    *front = head;
    *back = NULL;
  } else {
    slow = head;
    fast = head->next;
    while(fast != NULL) {
      fast = fast->next;
      if(fast != NULL) {
        slow = slow->next;
        fast = fast->next;
      }
    }
    *front = head;
    *back = slow->next;
    slow->next = NULL;
  }
}
```
