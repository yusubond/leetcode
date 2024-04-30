## 全排列-中等

题目：

给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。


分析：

回溯算法

```go
// date 2023/11/06
func permute(nums []int) [][]int {
    // res 存放所有的结果
    res := make([][]int, 0, 1024)

    // nums 表示待添加的元素
    // temp 表示临时结果
    var backtrack func(nums []int, temp []int)
    backtrack = func(nums []int, temp []int) {
        // 说明都已经添加过了，那么 temp 就是其中的一个结果，直接追加到结果集中
        if len(nums) == 0 {
            res = append(res, temp)
            return
        }
        for i := range nums {
            // 依次添加到 temp 中，形成新的切片
            // 并对未使用的元素，回溯添加
            newtemp := append(temp, nums[i])
            unused := make([]int, 0, len(nums)-1)
            unused = append(unused, nums[:i]...)
            unused = append(unused, nums[i+1:]...)
            backtrack(unused, newtemp)
        }
    }

    backtrack(nums, make([]int, 0))

    return res
}
```
