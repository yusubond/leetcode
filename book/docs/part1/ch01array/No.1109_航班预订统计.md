## 1109 航班预订统计-中等

题目：

这里有 `n` 个航班，它们分别从 `1` 到 `n` 进行编号。

有一份航班预订表 `bookings` ，表中第 `i` 条预订记录 `bookings[i] = [firsti, lasti, seatsi]` 意味着在从 `firsti` 到 `lasti` （**包含** `firsti` 和 `lasti` ）的 **每个航班** 上预订了 `seatsi` 个座位。

请你返回一个长度为 `n` 的数组 `answer`，里面的元素是每个航班预定的座位总数。



分析：

差分数组。

```go
// date 2023/12/12
func corpFlightBookings(bookings [][]int, n int) []int {
    // diff nums
    ans := make([]int, n)
    for _, num := range bookings {
        if num[0] <= n {
            ans[num[0]-1] += num[2]
        }
        if num[1] < n {
            ans[num[1]] -= num[2]
        }
    }

    // 复原差分数组
    for i := 1; i < n; i++ {
        ans[i] += ans[i-1]
    }

    return ans
}
```

