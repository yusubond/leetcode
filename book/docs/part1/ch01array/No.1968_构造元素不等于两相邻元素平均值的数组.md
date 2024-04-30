## 1968 构造元素不等于两相邻元素平均值的数组-中等

题目：

给你一个 下标从 0 开始 的数组 nums ，数组由若干 互不相同的 整数组成。你打算重新排列数组中的元素以满足：重排后，数组中的每个元素都 不等于 其两侧相邻元素的 平均值 。

更公式化的说法是，重新排列的数组应当满足这一属性：对于范围 1 <= i < nums.length - 1 中的每个 i ，(nums[i-1] + nums[i+1]) / 2 不等于 nums[i] 均成立 。

返回满足题意的任一重排结果。


分析：

因为元素各不相同，直接升序排序。那么排序后的数组，前半部分的任意两个元素的平均值肯定不等于后半部分任意两个元素的平均值。

所以，再将前半部分和后半部分交叉存放即可。


```go
// date 2023/11/09
func rearrangeArray(nums []int) []int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })
    mid := (len(nums) + 1)/2
    res := make([]int, 0, len(nums))

    for i := 0; i < mid; i++ {
        res = append(res, nums[i])
        if i + mid < len(nums) {
            res = append(res, nums[i+mid])
        }
    }

    return res
}
```
