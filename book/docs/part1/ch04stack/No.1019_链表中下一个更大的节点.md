## 1019 链表中的下一个更大的节点-中等

题目：

给定一个长度为 `n` 的链表 `head`

对于列表中的每个节点，查找下一个 **更大节点** 的值。也就是说，对于每个节点，找到它旁边的第一个节点的值，这个节点的值 **严格大于** 它的值。

返回一个整数数组 `answer` ，其中 `answer[i]` 是第 `i` 个节点( **从1开始** )的下一个更大的节点的值。如果第 `i` 个节点没有下一个更大的节点，设置 `answer[i] = 0` 。



> **示例 1：**
>
> ![img](https://assets.leetcode.com/uploads/2021/08/05/linkedlistnext1.jpg)
>
> ```
> 输入：head = [2,1,5]
> 输出：[5,5,0]
> ```
>
> **示例 2：**
>
> ![img](https://assets.leetcode.com/uploads/2021/08/05/linkedlistnext2.jpg)
>
> ```
> 输入：head = [2,7,4,3,5]
> 输出：[7,0,5,5,0]
> ```



分析：

解法1：最朴素的解法就是两层循环，依次查找，找到就退出

```go
// date 2023/12/21
// 两层查找
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
    ans := make([]int, 0, 16)
    
    pre := head

    for pre != nil {
        res := 0
        next := pre.Next
        for next != nil {
            if next.Val > pre.Val {
                res = next.Val
                break
            }
            next = next.Next
        }
        ans = append(ans, res)
        pre = pre.Next
    }

    return ans
}
```



解法2：单调栈

跟数组的题目类似，只不过链表的要自己保存索引。

```go
// date 2023/12/21
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
    type node struct {
        idx, val int
    }
    ans := make([]int, 0, 16)
    stack := make([]node, 0, 16)

    pre := head
    i := 0

    for pre != nil {
        v := pre.Val

        for len(stack) > 0 && v > stack[len(stack)-1].val {
            ans[stack[len(stack)-1].idx] = v
            stack = stack[:len(stack)-1]
        }


        stack = append(stack, node{idx: i, val: v})
        ans = append(ans, 0)
        pre = pre.Next
        i++
    }

    return ans
}
```

