## 306 累加数-中等

题目：

**累加数** 是一个字符串，组成它的数字可以形成累加序列。

一个有效的 **累加序列** 必须 **至少** 包含 3 个数。除了最开始的两个数以外，序列中的每个后续数字必须是它之前两个数字之和。

给你一个只包含数字 `'0'-'9'` 的字符串，编写一个算法来判断给定输入是否是 **累加数** 。如果是，返回 `true` ；否则，返回 `false` 。

**说明：**累加序列里的数，除数字 0 之外，**不会** 以 0 开头，所以不会出现 `1, 2, 03` 或者 `1, 02, 3` 的情况。



> **示例 1：**
>
> ```
> 输入："112358"
> 输出：true 
> 解释：累加序列为: 1, 1, 2, 3, 5, 8 。1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8
> ```
>
> **示例 2：**
>
> ```
> 输入："199100199"
> 输出：true 
> 解释：累加序列为: 1, 99, 100, 199。1 + 99 = 100, 99 + 100 = 199
> ```



分析：

这道题的解题思路是判断字符串是否符合斐波那契形式。

- 每次判断需要维护两个数字，所以在 DFS 遍历的时候维护两个数的边界索引`firstEnd`和`secondEnd`。

- 每次移动`firstEnd`，`secondEnd`的时候，需要判断后面的字符串是否以两个数的和开头。

- ```
  strings.HasPrefix(nums[secondEnd+1:], strconv.Itoa(x1+x2))
  ```

- 其次，无论第一个数和第二个数，只要起始数字出现`0`，都算非法情况，直接返回 false。



```go
// date 2023/12/27
func isAdditiveNumber(num string) bool {
    if len(num) < 3 {
        return false
    }
    for firstEnd := 0; firstEnd < len(num)/2; firstEnd++ {
        if num[0] == '0' && firstEnd > 0 {
            break
        }
        first, _ := strconv.Atoi(num[:firstEnd+1])

        // find the second
        for secondEnd := firstEnd+1; max(firstEnd, secondEnd-firstEnd) <= len(num)-secondEnd; secondEnd++ {
            if num[firstEnd+1] == '0' && secondEnd - firstEnd > 1 {
                break
            }
            second, _ := strconv.Atoi(num[firstEnd+1:secondEnd+1])
            if recursiveCheck(num, first, second, secondEnd+1) {
                return true
            }
        }
    }
    return false
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func recursiveCheck(num string, x1, x2 int, left int) bool {
    if left == len(num) {
        return true
    }
    if strings.HasPrefix(num[left:], strconv.Itoa(x1+x2)) {
        return recursiveCheck(num, x2, x1+x2, left + len(strconv.Itoa(x1+x2)))
    }
    return false
}
```

