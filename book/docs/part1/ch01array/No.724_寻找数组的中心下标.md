## 724 寻找数组的中心下标-简单

给你一个整数数组 nums ，请计算数组的 中心下标 。

数组 中心下标 是数组的一个下标，其左侧所有元素相加的和等于右侧所有元素相加的和。

如果中心下标位于数组最左端，那么左侧数之和视为 0 ，因为在下标的左侧不存在元素。这一点对于中心下标位于数组最右端同样适用。

如果数组有多个中心下标，应该返回 最靠近左边 的那一个。如果数组不存在中心下标，返回 -1 。



分析：

两次遍历，先求总和，二次遍历的时候判断左右和是否相等。

```go
// date 2022/10/07
func pivotIndex(nums []int) int {
    leftSum, rightSum := 0, 0
    for _, v := range nums {
        rightSum += v
    }

    for i, v := range nums {
        rightSum -= v
        if leftSum == rightSum {
            return i
        }
        leftSum += v
    }
    return -1
}
```

