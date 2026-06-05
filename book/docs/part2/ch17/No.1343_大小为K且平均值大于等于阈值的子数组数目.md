## 1343 大小为K且平均值大于等于阈值的子数组数目-中等

题目：

给你一个整数数组 `arr` 和两个整数 `k` 和 `threshold` 。

请你返回长度为 `k` 且平均值大于等于 `threshold` 的子数组数目。



分析：

最基本的滑动窗口。

```go
// date 2023/11/22
func numOfSubarrays(arr []int, k int, threshold int) int {
    var ans int
    left, right := 0, 0
    n := len(arr)
    sum := 0

    for right < n {
        sum += arr[right]
        right++
        if right - left == k {
            fm := float64(sum) / float64(k)
            if fm >= float64(threshold) {
                ans++
            }
            sum -= arr[left]
            left++
        }
    }

    return ans
}
```

