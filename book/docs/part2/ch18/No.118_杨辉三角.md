## 118 杨辉三角-中等

题目：

给定一个非负整数 *`numRows`，*生成「杨辉三角」的前 *`numRows`* 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。



**解题思路**

这道题也是动态规划的一种，直接模拟杨辉三角的计算过程即可。

```go
// date 2024/10/30
func generate(numRows int) [][]int {
    ans := make([][]int, numRows)
    for i := 0; i < numRows; i++ {
        ans[i] = make([]int, 0, 16)
    }

    for i := 0; i < numRows; i++ {
        if i == 0 {
            ans[i] = append(ans[i], 1)
        } else {
            ans[i] = append(ans[i], 1)
            for j := 1; j < i; j++ {
                ans[i] = append(ans[i], ans[i-1][j]+ans[i-1][j-1])
            }
            ans[i] = append(ans[i], 1)
        }
    }

    return ans
}
```

