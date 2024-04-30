## 27 移除元素-简单

题目：

给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。


分析：

前后指针，前面的指针负责记录符合条件的元素，后面的指针遍历整个数组

```go
// date 2022-09-15
func removeElement(nums []int, val int) int {
    idx := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[idx] = nums[i]
            idx++
        }
    }
    return idx
}
```

算法2：交换的思想，将每个等于val的元素，交换至数组的尾部，维护尾部索引，返回新的尾部索引。时间复杂度O(n) 空间复杂度O(n)

```go
// date 2021-03-13
func removeElements(nums []int, val int) int {
  i, tail := 0, len(nums)
  for i < tail; {
    if nums[i] == val {
      tail--
      nums[i] = nums[tail]
      nums[tail] = val
    } else {
      i++
    }
  }
  return tail
}
```

