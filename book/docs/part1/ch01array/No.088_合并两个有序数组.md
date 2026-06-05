## 88 合并两个有序数组-简单

题目：

将两个非递减数组nums1和nums2，按非递减顺序合并到nums1



分析：

从后往前充填，可以节省nums1中元素的移动。

```go
// date 2022/09/24
func merge(nums1 []int, m int, nums2 []int, n int)  {
    i, j, cur := m-1, n-1, m+n-1
    for i >= 0 && j >= 0 {
        if nums1[i] > nums2[j] {
            nums1[cur] = nums1[i]
            i--
        } else {
            nums1[cur] = nums2[j]
            j--
        }
        cur--
    }
    for i >= 0 {
        nums1[cur] = nums1[i]
        cur--
        i--
    }
    for j >= 0 {
        nums1[cur] = nums2[j]
        cur--
        j--
    }
}
```

