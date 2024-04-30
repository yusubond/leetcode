## 941 有效的山脉数组-简单

给定一个整数数组 arr，如果它是有效的山脉数组就返回 true，否则返回 false。

让我们回顾一下，如果 arr 满足下述条件，那么它是一个山脉数组：

arr.length >= 3
在 0 < i < arr.length - 1 条件下，存在 i 使得：
arr[0] < arr[1] < ... arr[i-1] < arr[i]
arr[i] > arr[i+1] > ... > arr[arr.length - 1]



算法分析：

```go
// date 2020/05/20
func validMountainArray(arr []int) bool {
    n := len(arr)
    if n < 3 { return false }
    top, index := arr[0], 0
    // 找到最高峰
    for i := 1; i < n; i++ {
        if arr[i] == arr[i-1] { return false } // 不能有平的
        if arr[i] > top {
            top = arr[i]
            index = i
        }
    }
    // 判断是否是斜坡
    if index == 0 || index == n-1 {
        return false
    }
    // 判断上坡
    for i := 1; i < index; i++ {
        if arr[i-1] >= arr[i] {
            return false
        }
    }
    // 判断下坡
    for i := index+1; i < n; i++ {
        if arr[i-1] <= arr[i] {
            return false
        }
    }
    return true
}
```

