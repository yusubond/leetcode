## 1477 找两个和为目标值且不重叠的子数组-中等

题目：

给你一个整数数组 `arr` 和一个整数值 `target` 。

请你在 `arr` 中找 **两个互不重叠的子数组** 且它们的和都等于 `target` 。可能会有多种方案，请你返回满足要求的两个子数组长度和的 **最小值** 。

请返回满足要求的最小长度和，如果无法找到这样的两个子数组，请返回 **-1** 。



分析：



1. 利用滑动窗口，找出所有满足条件的子数组【这些子数组可能重叠，也可能不重叠，需要剪枝操作】
2. 对所有的子数组按 长度 排序，以便找到第一个答案后即可退出循环
3. 剪枝1：判断子数组的长度，如果超过原数组的长度的一半，那么肯定无解，直接退出。
4. 剪枝2：对有包含关系的子数组直接跳过

```go
// date 2023/11/24
type resNum struct {
    start int
    length int
}

func minSumOfLengths(arr []int, target int) int {
    ans := len(arr) * 2
    res := make([]*resNum, 0, 16)
    left, right := 0, 0
    sum := 0
    n := len(arr)

    for right < n {
        sum += arr[right]
        right++

        for sum >= target {
            if sum == target {
                // one result
                // res = append(res, right-left)
                res = append(res, &resNum{start: left, length: right - left})
            }
            sum -= arr[left]
            left++
        }
    }

    // 防止超时
    // 按子数组长度排序，后面剪枝之后，只要找到答案，即可 break
    sort.Slice(res, func(i, j int) bool {
        return res[i].length < res[j].length
    })

    m := len(res)
    for i := 0; i < m; i++ {
        // 防止超时
        // 如果单个子数组的长度已经超过一半，那么肯定无解，直接跳出
        if res[i].length * 2 > n {
            break
        }
        for j := i+1; j < m; j++ {
            if res[i].start < res[j].start && res[j].start < res[i].start + res[i].length {
                continue
            }
            if res[j].start < res[i].start && res[i].start < res[j].start + res[j].length {
                continue
            }
            ans = min(ans, res[i].length + res[j].length)
            break
        }
    }

    if ans == len(arr) * 2 {
        return -1
    }
    return ans
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
```

