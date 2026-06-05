## 1299 将每个元素替换为右侧最大元素-简单

题目：

给你一个数组 `arr` ，请你将每个元素用它右边最大的元素替换，如果是最后一个元素，用 `-1` 替换。

完成所有替换操作后，请你返回这个数组。



算法：

```go
// date 2022/10/01
func replaceElements(arr []int) []int {
    cur, right_max := -1, -1
    for i := len(arr)-1; i >= 0; i-- {
        // 保存当前值
        cur = arr[i]
        // 将当前值赋值为右侧最大值
        arr[i] = right_max
        // 更新右侧最大值
        if cur > right_max {
            right_max = cur
        }
    }
    return arr
}
```

