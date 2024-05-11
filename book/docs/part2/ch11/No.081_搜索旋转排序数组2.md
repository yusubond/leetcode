## 81 搜索旋转排序数组2-中等

题目：

已知存在一个按非降序排列的整数数组 `nums` ，数组中的值不必互不相同。

在传递给函数之前，`nums` 在预先未知的某个下标 `k`（`0 <= k < nums.length`）上进行了 **旋转** ，使数组变为 `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]`（下标 **从 0 开始** 计数）。例如， `[0,1,2,4,4,4,5,6,6,7]` 在下标 `5` 处经旋转后可能变为 `[4,5,6,6,7,0,1,2,4,4]` 。

给你 **旋转后** 的数组 `nums` 和一个整数 `target` ，请你编写一个函数来判断给定的目标值是否存在于数组中。如果 `nums` 中存在这个目标值 `target` ，则返回 `true` ，否则返回 `false` 。

你必须尽可能减少整个操作步骤。



分析：

先根据第一个元素与目标值的大小关系，去掉不可能的部分，然后二分查找。

```go
// date 2023/12/04
func search(nums []int, target int) bool {
    n := len(nums)
    if target == nums[0] {
        return true
    } else if target > nums[0] {
        left, right := 1, n-1
        for right >= left && nums[right] <= nums[0] {
            right--
        }
        for left <= right {
            mid := (left + right) / 2
            if nums[mid] == target {
                return true
            } else if nums[mid] < target {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
        return false
    } else if target < nums[0] {
        left, right := 0, n-1
        for left <= right && nums[left] >= nums[0] {
            left++
        }
        for left <= right {
            mid := (left + right) / 2
            if nums[mid] == target {
                return true
            } else if nums[mid] < target {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
        return false
    }
    return false
}
```

