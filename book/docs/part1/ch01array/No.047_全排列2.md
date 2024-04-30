## 47 全排列2-中等

题目：

给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。


分析：

剪枝，回溯

```go
// date 2023/11/07
func permuteUnique(nums []int) [][]int {
    res := make([][]int, 0, 1024)
    vis := make(map[int]bool)
    n := len(nums)

    temp := make([]int, 0, 16)

    var backtrack func(idx int)
    backtrack = func(idx int) {
        if idx == n {
            res = append(res, append([]int{}, temp...))
            return
        }
        for i, v := range nums {
            // vis[i] 已经填充过的，不再填充
            // i > 0 && nums[i-1] == && !vis[i-1]
            // 相同元素，及时没填充过，也不需要填充
            if i > 0 && nums[i-1] == v && !vis[i-1] || vis[i] {
                continue
            }
            temp = append(temp, v)
            vis[i] = true
            backtrack(idx+1)
            vis[i] = false
            temp = temp[:len(temp)-1]
        }
    }

    sort.Slice(nums, func(i,j int) bool {
        return nums[i] < nums[j]
    })

    backtrack(0)

    return res
}
```
