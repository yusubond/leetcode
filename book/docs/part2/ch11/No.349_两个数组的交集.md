## 349 两个数组的交集-中等

题目：

给定两个数组 `nums1` 和 `nums2` ，返回 *它们的交集* 。输出结果中的每个元素一定是 **唯一** 的。我们可以 **不考虑输出结果的顺序** 。



分析：

哈希表存储，如果存在第一个数组中，值为1，如果存在第二个数组，值为2。

最后遍历哈希表即可。

```go
// date 2023/12/05
func intersection(nums1 []int, nums2 []int) []int {
    set := make(map[int]int, 32)
    for _, v := range nums1 {
        _, ok := set[v]
        if !ok {
            set[v] = 1
        }
    }
    for _, v := range nums2 {
        _, ok := set[v]
        if ok {
            set[v] = 2
        }
    }
    res := make([]int, 0, 4)
    for k, v := range set {
        if v == 2 {
            res = append(res, k)
        }
    }
    return res
}
```

思考：如果是已经排序的数组，如何求交集。

如果是有序数组，直接在另一个数组中二分查找。算法如下：

```go
// date 2023/12/05
func intersection(nums1 []int, nums2 []int) []int {
    res := make([]int, 0, 16)
    sort.Slice(nums1, func(i, j int) bool {
        return nums1[i] < nums1[j]
    })
    sort.Slice(nums2, func(i, j int) bool {
        return nums2[i] < nums2[j]
    })

    for i := 0; i < len(nums1); i++ {
        if i > 0 && nums1[i] == nums1[i-1] {
            continue
        }
        if findEqual(nums2, nums1[i]) {
            res = append(res, nums1[i])
        }
    }

    return res
}

func findEqual(nums []int, key int) bool {
    n := len(nums)

    if nums[0] > key || nums[n-1] < key {
        return false
    }

    left, right := 0, n-1
    for left <= right {
        mid := (left + right) / 2
        if nums[mid] == key {
            return true
        } else if nums[mid] < key {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return false
}
```

