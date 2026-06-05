## 56 合并区间-中等

题目：

以数组 `intervals` 表示若干个区间的集合，其中单个区间为 `intervals[i] = [starti, endi]` 。请你合并所有重叠的区间，并返回 *一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间* 。



> **示例 1：**
>
> ```
> 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
> 输出：[[1,6],[8,10],[15,18]]
> 解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
> ```
>
> **示例 2：**
>
> ```
> 输入：intervals = [[1,4],[4,5]]
> 输出：[[1,5]]
> 解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
> ```



**解题思路**

这道题最朴素的做法就是按照题意依次合并。具体为，先对集合按每个区间的 start 大小排序，排序后每个区间是依次展开的；然后把第一个区间保存为 one，从第二个开始进行合并，合并规则如下：

1. 如果当前区间的 start 和 end，均在上一个区间内部，即跳过该区间
2. 如果当前区间的 start 小于等于上一个区间（one） 的 end，那么上个区间（one）可修改为`[start, end2]`
3. 如果 1和 2 都不满足，则表示得到一个新的区间，加入结果集即可；同时将当前区间保存为 one，合并

```go
// date 2024/01/16
func merge(intervals [][]int) [][]int {
    n := len(intervals)
    if n < 2 {
        return intervals
    }
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    ans := make([][]int, 0, 16)
    one := intervals[0]
    for i := 1; i < n; i++ {
        if intervals[i][0] >= one[0] && intervals[i][1] <= one[1] {
            continue
        }
        if intervals[i][0] <= one[1] {
            one[1] = intervals[i][1]
        } else {
            ans = append(ans, one)
            one = intervals[i]
        }
    }
    ans = append(ans, one)
    return ans
}
```

