## 283 移动零-简单

> 给定一个数组 `nums`，编写一个函数将所有 `0` 移动到数组的末尾，同时保持非零元素的相对顺序。
>
> **请注意** ，必须在不复制数组的情况下原地对数组进行操作。



算法：

维护好非零元素应该使用的索引，遍历数组，将非零元素放在其索引。

```go
// date 2023/11/21
func moveZeroes(nums []int)  {
    idx := 0
    i, n := 0, len(nums)
    for i < n {
        if nums[i] != 0 {
          	// 有可能整个数组全是非零元素，所以要判断
          	// 因为只要有零存在，idx 比小于 i
            if idx < i {
                nums[idx] = nums[i]
                nums[i] = 0
            }
            idx++
        }
        i++
    }
}
```

推荐下面的写法，简洁粗暴。

先直接复制非零元素，然后尾部全部置零。

```go
func moveZeroes(nums []int)  {
    idx := 0
    i, n := 0, len(nums)
    for i < n {
        if nums[i] != 0 {
            nums[idx] = nums[i]
            idx++
        }
        i++
    }
    for idx < n {
        nums[idx] = 0
        idx++
    }
}
```

