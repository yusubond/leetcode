## 11 盛最多水的容器-中等

题目：

给定一个长度为 `n` 的整数数组 `height` 。有 `n` 条垂线，第 `i` 条线的两个端点是 `(i, 0)` 和 `(i, height[i])` 。

找出其中的两条线，使得它们与 `x` 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

**说明：**你不能倾斜容器。



分析：

利用双指针从两头向中间逼近，计算中间结果，并保存最大结果。

```go
// date: 2022-09-14
func maxArea(height []int) int {
    left, right := 0, len(height)-1
    res, temp := 0, 0
    width := 0
    for left < right {
        width = right - left
        if height[left] < height[right] {
            temp = height[left] * width
            left++
        } else {
            temp = height[right] * width
            right--
        }
        if temp > res {
            res = temp
        }
    }
    return res
}
```

