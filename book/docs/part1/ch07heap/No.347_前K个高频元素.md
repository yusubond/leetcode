## 347 前K个高频元素-中等

题目：

给你一个整数数组 `nums` 和一个整数 `k` ，请你返回其中出现频率前 `k` 高的元素。你可以按 **任意顺序** 返回答案。



> **示例 1:**
>
> ```
> 输入: nums = [1,1,1,2,2,3], k = 2
> 输出: [1,2]
> ```
>
> **示例 2:**
>
> ```
> 输入: nums = [1], k = 1
> 输出: [1]
> ```



分析：

这道题依然是小顶堆的解题思路。

自己构造一个结构体，保存元素值和元素出现的次数。

先遍历数组，统计元素的出现次数，保存在 map 中，然后对 map 遍历，构造小顶堆。

```go
// date 2024/01/03
func topKFrequent(nums []int, k int) []int {
    set := make(map[int]int, 16)
    arr := make([]*MyNode, 0, 16)
    for _, v := range nums {
        set[v]++
    }
    idx := 0
    for num, count := range set {
        re := &MyNode{val: num, time: count}
        if idx < k {
            arr = append(arr, re)
            idx++
            if idx == k {
                makeMinHeap(arr)
            }
        } else if count > arr[0].time {
            arr[0] = re
            minHeapify(arr, 0)
        }
    }
    ans := make([]int, 0, k)
    for _, v := range arr {
        ans = append(ans, v.val)
    }
    return ans
}

type MyNode struct {
    val int
    time int
}

func makeMinHeap(arr []*MyNode) {
    for i := len(arr)/2 - 1; i >= 0; i-- {
        minHeapify(arr, i)
    }
}

func minHeapify(arr []*MyNode, start int) {
    n := len(arr)
    if start > n {
        return
    }
    temp := arr[start]
    left := start * 2 + 1
    for left < n {
        right := left + 1
        if right < n && arr[right].time < arr[left].time {
            left++
        }
        if arr[start].time < arr[left].time {
            break
        }
        arr[start] = arr[left]
        arr[left] = temp

        start = left
        left = 2*start + 1
    }
}
```

