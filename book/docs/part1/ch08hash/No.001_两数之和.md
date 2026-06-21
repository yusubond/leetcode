## 001 两数之和

给定一个整数数组 `nums` 和一个整数目标值 `target`，请你在该数组中找出 **和为目标值** *`target`* 的那 **两个** 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。

你可以按任意顺序返回答案。



分析：

利用 map 存储遍历过的值，实现O(1)查找。

```go
func twoSum(nums []int, target int) []int {
	set := make(map[int]int, len(nums))
	ans := make([]int, 2)

	for i, v := range nums {
		j, ok := set[target-v]
		if ok {
			ans[0], ans[1] = i, j
			break
		}
		set[v] = i
	}

	return ans
}
```

