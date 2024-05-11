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
