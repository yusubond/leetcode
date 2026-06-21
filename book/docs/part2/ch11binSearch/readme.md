## 二分查找

二分查找，是指在已排序的数组中查找目标值。

感觉很简单，其实二分查找并不简单，尤其二分法的各种变种。

首先，我们从最简单的开始：

问题1：从已经排序好的数组中找出与 key 相等的元素，如果找到，请返回元素下标，否则返回 -1。

如果存在答案，你可以假设只有唯一的答案。



### 1.找出与 key 相等的元素

请在已排序的数组中找出与 key 相等的元素，且数组中的元素不重复。如果能够找到，请返回元素下标，否则返回 -1。

因为答案唯一，所以判断 `nums[mid] == key`，只要相等直接返回结果。

```go
func search(nums []int, n int, key int) int {
  left, right := 0, n-1
  // 因为要判断所有元素，所以循环区间为 left <= right
  // left == right 就是判断最后一个元素
  for left <= right {
    // 避免下标溢出
    mid := left + (right - left) / 2
    if nums[mid] == key {
      return mid
    } else if nums[mid] > key {
      right = mid -1
    } else if nums[mid] < key {
      left = mid + 1
    }
  }
  return -1
}
```

程序不复杂，但是也要注意一点，因为已经判断了`nums[mid] == key`，所以每次移动 left 和 right 指针的时候，要在 mid 的基础上 +1 或者 -1。



### 2.找出第一个与 key 相等的元素

已知排序数组中的元素有重复，请找到第一个。

```go
func searchFirstEqual(nums []int, n int, key int) int {
  left, right := 0, n-1
  for left <= right {
    mid := (left + right) / 2
    if nums[mid] >= key {
      right = mid - 1
    } else if nums[mid] < key {
      left = mid + 1
    }
  }
  if left < n && nums[left] == key {
    return left
  }
  return -1
}
```

注意，只要是查找第一个的，一定是返回 left 指针。

另外，在二分查找的循环体中，left 指针的判断条件是`nums[mid] < key`，所以有可能不存在第一个与 key 相等的元素，所以二分查找结束后，要校验 left 的 合法性，以及 nums[left] == key。



### 3.找出最后一个与 key 相等的元素

已知排序数组中的元素有重复，请找出最后一个。

```go
func searchLastEqual(nums []int, n int, key int) int {
  left, right := 0, n-1
  // left <= right, 
  for left <= right {
    mid := (left + right) / 2
    if nums[mid] > key {
      right = mid - 1
    } else if nums[mid] <= key {
      left = mid + 1
    }
  }
  if right >= 0 && nums[right] == key {
    return right
  }
  return -1
}
```

注意：只要是查找最后一个，一定是返回 right 指针。



### 4.找出第一个等于或大于 key 的元素

这个变种，不在要求一定是等于，而是等于或者大于都可以，且第一个。

那么，返回 left 指针，且去掉最后 left 合法性的校验。

```go
func searchFirstEqualOrLarger(nums []int, n int, key int) int {
  left, right := 0, n-1
  for left <= right {
    mid := (left + right) / 2
    if nums[mid] >= key {
      right = mid - 1
    } else if nums[mid] < key {
      left = mid + 1
    }
  }
  if left > n-1 {
    return -1
  }
  return left
}
```



### 5.找出第一个大于 key 的元素

这个变种，去掉了相当，寻找第一个大于的。

```go
func searchFirstLarger(nums []int, n int, key int) int {
  left, right := 0, n-1
  for left <= right {
    mid := (left + right) / 2
    if nums[mid] > key {
      right = mid - 1
    } else if nums[mid] <= key {  // 不满足条件的
      left = mid + 1
    }
  }
  if left > n-1 {
    return -1
  }
  return left
}
```



### 6.找出最后一个等于或小于 key 的元素

返回右侧。

```go
func searchLastEqualOrSmaller(nums []int, n int, key int) int {
  left, right := 0, n-1
  for left <= right {
    mid := (left + right) / 2
    if nums[mid] > key {
      right = mid - 1
    } else if nums[mid] <= key {
      left = mid + 1
    }
  }
  return right
}
```



### 7.找出最后一个小于 key 的元素



```go
func searchLastSmaller(nums []int, n int, key int) int {
  left, right := 0, n-1
  for left <= right {
    mid := (left + right) / 2
    if nums[mid] >= key { // 这些是不要的
      right = mid - 1
    } else if nums[mid] < key {
      left = mid + 1
    }
  }
  return right
}
```



## 小结

总结下，这类问题的要点是：

- 根据返回元素的位置（第一个还是最后一个），决定返回 left 指针，还是 right 指针
- 在二分查找判断的过程中，根据要求排除不符合的元素，缩小 left 或者 right 的范围。



以下是测试用例。

```go
// date 2023/12/04
package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3, 3, 3, 5, 5, 7, 7}

	fmt.Printf("nums = %d\n", nums)
	fmt.Printf("expect = 6, First Equal 5: %d\n", searchFirstEqual(nums, 5))
	fmt.Printf("expect = 7, Last Equal 5: %d\n", searchLastEqual(nums, 5))
	fmt.Printf("expect = 6, First Equal Or Larger 4: %d\n", searchFirstEqualOrLarger(nums, 4))
	fmt.Printf("expect = 6, First Equal Or Larger 5: %d\n", searchFirstEqualOrLarger(nums, 5))
	fmt.Printf("expect = 3, First Larger 2: %d\n", searchFirstLarger(nums, 2))
	fmt.Printf("expect = 6, First Larger 4: %d\n", searchFirstLarger(nums, 4))
	fmt.Printf("expect = 5, Last Equal or Smaller 3: %d\n", searchLastEqualOrSmaller(nums, 3))
	fmt.Printf("expect = 5, Last Equal or Smaller 4: %d\n", searchLastEqualOrSmaller(nums, 4))
	fmt.Printf("expect = 7, Last Smaller 6: %d\n", searchLastSmaller(nums, 6))
}

// 1. 查找第一个与 key 相等的元素
func searchFirstEqual(nums []int, key int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= key {
			right = mid - 1
		} else if nums[mid] < key {
			left = mid + 1
		}
	}
	if left < n && nums[left] == key {
		return left
	}
	return -1
}

// 2. 查找最后一个与 key 相等的元素
func searchLastEqual(nums []int, key int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > key {
			right = mid - 1
		} else if nums[mid] <= key {
			left = mid + 1
		}
	}
	if right >= 0 && nums[right] == key {
		return right
	}
	return -1
}

// 3. 查找第一个等于或大于 key 的元素
func searchFirstEqualOrLarger(nums []int, key int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < key {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left >= n {
		return -1
	}
	return left
}

// 4. 查找第一个大于 key 的元素
func searchFirstLarger(nums []int, key int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= key {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left >= n {
		return -1
	}
	return left
}

// 5. 查找最后一个等于或小于 key 的元素
func searchLastEqualOrSmaller(nums []int, key int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > key {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if right >= 0 {
		return right
	}
	return -1
}

// 6. 查找最后一个小于 key 的元素
func searchLastSmaller(nums []int, key int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= key {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if right >= 0 {
		return right
	}
	return -1
}
```



## 二分查找的模板

**模板1**

1. 查找等于目标值 key 的索引

```go
func searchEqual(nums []int, key int) int {
  n := len(nums)
  if n == 0 {
    return -1
  }
  left, right := 0, n-1
  for left <= right {
    // 为了避免 left + right 溢出，使用 left + (right - left) / 2 的方式求中心点
    mid := left + (right - left) / 2
    if nums[mid] == key {
      return mid
    } else if nums[mid] < key {
      left = mid +1
    } else {
      right = mid - 1
    }
  }
  return -1
}
```

