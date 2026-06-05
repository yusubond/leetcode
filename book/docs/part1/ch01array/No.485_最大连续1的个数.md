## 485 最大连续1的个数-简单

题目：

给定一个二进制数组 `nums` ， 计算其中最大连续 `1` 的个数。

数组中只包含1和0。



分析：

前后指针

```go
// date 2022/10/07
func findMaxConsecutiveOnes(nums []int) int {
    res, n := 0, len(nums)
    start, end := 0, 0
    for start < n && nums[start] == 0 {
        start++
    }

    end = start

    for end < n {
        // find the first 0
        if nums[end] == 0 {
            if res < (end - start) {
                res = end - start
            }
            for end < n && nums[end] == 0 {
                end++
            }
            start = end
        }
        end++
    }
    if res < end - start {
        res = end - start
    }
    return res
}
// 解法2 直接数数
func findMaxConsecutiveOnes(nums []int) int {
	ans, cnt := 0, 0

	for _, v := range nums {
		if v == 1 {
			cnt++
		} else {
			cnt = 0
		}
		if cnt > ans {
			ans = cnt
		}
	}

	return ans
}
```

