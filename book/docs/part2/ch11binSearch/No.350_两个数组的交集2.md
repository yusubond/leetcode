## 350 两个数组的交集2-中等

题目：

给你两个整数数组 `nums1` 和 `nums2` ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。



分析：

两个哈希表分别记录两个数组中各元素出现的次数，然后对哈希表进行遍历查找，取重复元素在数组中出现的最小次数，追加到结果中。

```go
// date 2023/12/05
func intersect(nums1 []int, nums2 []int) []int {
    n1, n2 := len(nums1), len(nums2)
    set1 := make(map[int]int, n1)
    set2 := make(map[int]int, n2)
    for _, v := range nums1 {
        set1[v]++
    }
    for _, v := range nums2 {
        set2[v]++
    }
    
    res := make([]int, 0, 16)
    for k, v := range set1 {
        v2, ok := set2[k]
        if ok {
            ct := min(v, v2)
            for ct > 0 {
                res = append(res, k)
                ct--
            }
        }
    }
    return res
}

func min(x, y int) int {
    if x <= y {
        return x
    }
    return y
}
```

进阶思考：

1. 如果给定的数组已经排好序呢？你将如何优化你的算法？

【可用二分查找，返回第一个和最后一个等于 key 的下标】

2. 如果 `nums1` 的大小比 `nums2` 小，哪种方法更优？



3. 如果 `nums2` 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？