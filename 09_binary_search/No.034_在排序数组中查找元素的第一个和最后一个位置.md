## 34 在排序数组中查找元素的第一个和最后一个位置-中等

题目：

给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。



分析：

复习二分查找，[详见](../binary_search.md)

```go
// date 2022/09/15
func searchRange(nums []int, target int) []int {
    v1 := searchFirstEqual(nums, target)
    v2 := searchFindLastEqual(nums, target)
    return []int{v1, v2}
}

func searchFirstEqual(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] >= target { // 舍弃掉的
            right = mid-1
        } else {
            left = mid+1
        }
    }
    if left < len(nums) && nums[left] == target {
        return left
    }

    return -1
}

func searchFindLastEqual(nums []int, target int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] <= target { // 舍弃掉的
            left = mid+1
        } else {
            right = mid - 1
        }
    }
    if right >= 0 && nums[right] == target {
        return right
    }
    return -1
}
```

