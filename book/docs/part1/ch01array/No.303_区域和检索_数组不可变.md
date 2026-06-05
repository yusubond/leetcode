## 303 区域和检索-数组不可变-简单

题目：

给定一个整数数组  `nums`，处理以下类型的多个查询:

1. 计算索引 `left` 和 `right` （包含 `left` 和 `right`）之间的 `nums` 元素的 **和** ，其中 `left <= right`

实现 `NumArray` 类：

- `NumArray(int[] nums)` 使用数组 `nums` 初始化对象
- `int sumRange(int i, int j)` 返回数组 `nums` 中索引 `left` 和 `right` 之间的元素的 **总和** ，包含 `left` 和 `right` 两点（也就是 `nums[left] + nums[left + 1] + ... + nums[right]` )



分析：

简单的前缀和。

```go
// date 2023/12/11
type NumArray struct {
    data []int
}


func Constructor(nums []int) NumArray {
    sums := make([]int, len(nums))
    sum := 0
    for i, v := range nums {
        sum += v
        sums[i] = sum
    }
    return NumArray{data: sums}
}


func (this *NumArray) SumRange(left int, right int) int {
    v1 := 0
    if left > 0 {
        v1 = this.data[left-1]
    }
    v2 := this.data[right]
    return v2-v1
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
```

