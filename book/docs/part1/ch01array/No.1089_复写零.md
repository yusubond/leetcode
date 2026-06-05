## 1089 复写零-简单

题目：

给你一个长度固定的整数数组 `arr` ，请你将该数组中出现的每个零都复写一遍，并将其余的元素向右平移。

注意：请不要在超过该数组长度的位置写入元素。请对输入的数组 **就地** 进行上述修改，不要从函数返回任何东西。



分析：

直接遍历，遇到 零 时，将后续的元素向右移动一格。

```go
// date 2020/04/19
func duplicateZeros(arr []int)  {
    n := len(arr)
    j := 0

    for i := 0; i < n; i++ {
        if arr[i] == 0 {
            // copy and remove the last one
            j = n-2
            for j > i {
                arr[j+1] = arr[j]
                j--
            }
            arr[j+1] = 0
            i++
        }
    }
}
```

