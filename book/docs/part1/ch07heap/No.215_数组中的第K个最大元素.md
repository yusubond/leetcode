## 215 数组中的第K个最大元素-中等

题目：

给定整数数组 `nums` 和整数 `k`，请返回数组中第 `**k**` 个最大的元素。

请注意，你需要找的是数组排序后的第 `k` 个最大的元素，而不是第 `k` 个不同的元素。

你必须设计并实现时间复杂度为 `O(n)` 的算法解决此问题。



> **示例 1:**
>
> ```
> 输入: [3,2,1,5,6,4], k = 2
> 输出: 5
> ```
>
> **示例 2:**
>
> ```
> 输入: [3,2,3,1,2,4,5,5,6], k = 4
> 输出: 4
> ```



分析：

算法1：直接排序，返回第k个

算法2：小顶堆

```go
// date 2024/01/03
func findKthLargest(nums []int, k int) int {
    res := make([]int, k)
    idx := 0
    for _, v := range nums {
        if idx < k {
            res[idx] = v
            idx++
            if idx == k {
                makeMinHeap(res)
            }
        } else if v > res[0] {
            res[0] = v
            minHeapify(res, 0)
        }
    }
    return res[0]
}

// 把 arr 数组变成最小堆
func makeMinHeap(arr []int) {
    for i := len(arr)/2 - 1; i >= 0; i-- {
        minHeapify(arr, i)
    }
}

// sink
// 把较小值浮到上面
func minHeapify(arr []int, i int) {
    n := len(arr)
    if i > n {
        return
    }
    temp := arr[i]

    // left = 2*i + 1
    // right = left + 1
    l := i << 1 + 1
    for l < n {
        // 取左右节点中的较小值
        if l+1 < n && arr[l+1] < arr[l] {
            l++
        }
        // 如果 父结点小于左右节点，那么最小堆已成立，直接跳出
        if arr[i] < arr[l] {
            break
        }
        // swap
        arr[i] = arr[l]
        arr[l] = temp

        i = l
        l = i << 1 + 1
    }
}
```

