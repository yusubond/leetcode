## 93 复原IP地址-中等

题目：

**有效 IP 地址** 正好由四个整数（每个整数位于 `0` 到 `255` 之间组成，且不能含有前导 `0`），整数之间用 `'.'` 分隔。

- 例如：`"0.1.2.201"` 和` "192.168.1.1"` 是 **有效** IP 地址，但是 `"0.011.255.245"`、`"192.168.1.312"` 和 `"192.168@1.1"` 是 **无效** IP 地址。

给定一个只包含数字的字符串 `s` ，用以表示一个 IP 地址，返回所有可能的**有效 IP 地址**，这些地址可以通过在 `s` 中插入 `'.'` 来形成。你 **不能** 重新排序或删除 `s` 中的任何数字。你可以按 **任何** 顺序返回答案。



> **示例 1：**
>
> ```
> 输入：s = "25525511135"
> 输出：["255.255.11.135","255.255.111.35"]
> ```
>
> **示例 2：**
>
> ```
> 输入：s = "0000"
> 输出：["0.0.0.0"]
> ```
>
> **示例 3：**
>
> ```
> 输入：s = "101023"
> 输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
> ```



分析：

这道题可以选用 DFS。递归深搜的时候需要注意两种情况：

第一种，多个值可以合并，只要不超过 255，且不为零

第二种，单个值，正常递归

中间值不可有前导零，通过判断 idx 来实现。

```go
// date 2023/12/26
func restoreIpAddresses(s string) []string {
    res := make([]string, 0, 16)

    var dfs func(s string, idx int, ip []int)
    dfs = func(s string, idx int, ip []int) {
        if idx == len(s) {
            if len(ip) == 4 {
                res = append(res, toString(ip))
            }
            return
        }
        if idx == 0 {
            // IP 地址的第一个值可以为零
            // 所以这里不需要判断 num == 0
            num, _ := strconv.Atoi(string(s[0]))
            ip = append(ip, num)
            dfs(s, idx+1, ip)
        } else {
            // 非 IP 地址的第一个值,都需要判断不为零
            num, _ := strconv.Atoi(string(s[idx]))
            next := ip[len(ip)-1]*10 + num
            // 如果多个可以当做一个
            if next <= 255 && ip[len(ip)-1] != 0 {
                ip[len(ip)-1] = next
                dfs(s, idx+1, ip)
                ip[len(ip)-1] /= 10  // 撤销 num 选择
            }
            if len(ip) < 4 {
                ip = append(ip, num)
                dfs(s, idx+1, ip)
                ip = ip[:len(ip)-1]
            }
        }
    }

    dfs(s, 0, []int{})

    return res
}

func toString(nums []int) string {
    res := strconv.Itoa(nums[0])

    for i := 1; i < len(nums); i++ {
        res += "." + strconv.Itoa(nums[i])
    }

    return res
}
```

