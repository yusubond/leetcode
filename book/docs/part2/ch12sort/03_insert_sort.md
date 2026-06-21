## 插入排序

插入排序的思想是：

依次选择未排序的元素，将其插入到已排序的合适位置。

```sh
Loop from i = 1 to n-1.
  a) Pick element arr[i] and insert it into sorted sequence arr[0…i-1]
```



```go
// 插入排序
// 每次选择未排序的元素，将其插入到已排序的合适位置。
func insertSort(nums []int) []int {
	n := len(nums)

	for i := 1; i < n; i++ {
		v := nums[i]  // the unsort elem
		j := i - 1
    // find the last euqal or small than v
		for j >= 0 && nums[j] > v {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = v
	}

	return nums
}
```

