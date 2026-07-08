## 166 分数到小数-中等

题目：

给定两个整数，分别表示分数的分子 `numerator` 和分母 `denominator`，以 **字符串形式** 返回小数。

如果小数部分为循环小数，则将循环的部分括在括号内。

如果存在多个答案，只需返回 **任意一个**。

对于所有给定的输入，保证 **答案字符串的长度小于 10^4**。

注意，如果分数可以表示为有限长度的字符串，则 **必须** 返回它。

> **示例 1：**
>
> ```
> 输入：numerator = 1, denominator = 2
> 输出："0.5"
> ```
>
> **示例 2：**
>
> ```
> 输入：numerator = 2, denominator = 1
> 输出："2"
> ```
>
> **示例 3：**
>
> ```
> 输入：numerator = 4, denominator = 333
> 输出："0.(012)"
> ```
>
> **示例 4：**
>
> ```
> 输入：numerator = 1, denominator = 6
> 输出："0.1(6)"
> ```

**关键观察：长除法 + 哈希表检测余数循环**

分数转小数本质是模拟**长除法（long division）**：

1. 先处理**符号**：同号为正，异号为负。
2. 算出**整数部分**：`|numerator| / |denominator|`。
3. 如有余数，开始小数部分：余数 × 10 ÷ 分母，商追加到结果，余数更新为 `余数 × 10 % 分母`。
4. **核心**：小数部分的每一步产生一个新余数。如果某个余数**重复出现**，则之后的商也会重复——这就是循环节。用哈希表 `map[余数]→位置` 记录每个余数首次出现在结果字符串中的位置，当余数再次出现时，在该位置前插入 `(`，末尾加 `)`，结束。

**边缘情况**：

- `numerator = 0`：直接返回 `"0"`。
- `denominator = 1`：只有整数部分，没有小数点。
- 负数取绝对值时，`INT_MIN` 取绝对值会溢出——Go 中需要先转为 `int64` 避免越界，或直接用 `int64` 做所有计算。

**解法 1：长除法 + 哈希表**

```go
// date 2026/07/08
func fractionToDecimal(numerator int, denominator int) string {
    if numerator == 0 {
        return "0"
    }

    res := make([]byte, 0)

    // 1. 处理符号
    if (numerator < 0) != (denominator < 0) {
        res = append(res, '-')
    }

    // 2. 转为绝对值（使用 int64 防止 INT_MIN 越界）
    n := abs64(int64(numerator))
    d := abs64(int64(denominator))

    // 3. 整数部分
    res = append(res, []byte(strconv.FormatInt(n/d, 10))...)
    remainder := n % d
    if remainder == 0 {
        return string(res)
    }

    // 4. 小数部分
    res = append(res, '.')
    // 哈希表：余数 → 在 res 中的位置
    posMap := make(map[int64]int)
    for remainder != 0 {
        // 余数重复 → 找到循环节
        if idx, ok := posMap[remainder]; ok {
            // 在 idx 前插入 '('，末尾加 ')'
            prefix := string(res[:idx])
            repeat := string(res[idx:])
            return prefix + "(" + repeat + ")"
        }
        posMap[remainder] = len(res)
        remainder *= 10
        res = append(res, byte('0'+remainder/d))
        remainder %= d
    }

    return string(res)
}

func abs64(x int64) int64 {
    if x < 0 {
        return -x
    }
    return x
}
```

**解法 1 变体：直接用 `int64` 参数**

为规避 `int` 取绝对值时的溢出风险，计算全程在 `int64` 上进行。

```go
// date 2026/07/08
func fractionToDecimal(numerator int, denominator int) string {
    if numerator == 0 {
        return "0"
    }
    n, d := int64(numerator), int64(denominator)

    var res strings.Builder
    // 符号
    if (n < 0) != (d < 0) {
        res.WriteByte('-')
    }
    if n < 0 { n = -n }
    if d < 0 { d = -d }

    // 整数部分
    res.WriteString(strconv.FormatInt(n/d, 10))
    n %= d
    if n == 0 {
        return res.String()
    }
    res.WriteByte('.')

    // 小数部分，哈希表记录余数→位置
    seen := make(map[int64]int)
    for n != 0 {
        if pos, ok := seen[n]; ok {
            s := res.String()
            return s[:pos] + "(" + s[pos:] + ")"
        }
        seen[n] = res.Len()
        n *= 10
        res.WriteByte(byte('0' + n/d))
        n %= d
    }
    return res.String()
}
```

**复杂度**

| 解法 | 时间 | 空间 | 备注 |
|---|---|---|---|
| 长除法 + 哈希表 | O(L) | O(L) | L 为答案字符串长度，最大 10^4；每个余数最多出现一次 |

> 余数范围是 `[0, denominator-1]`，所以最多循环 `denominator` 次。循环节长度最大为 `denominator-1`，符合题目 10^4 的上限保证。

**关键细节：为什么余数重复就意味着循环**

长除法每一步：`当前余数 × 10 → 商 → 新余数`。这是一个**确定性过程**——给定余数 r，下一步的商和新余数是唯一确定的。因此如果余数 r 重复出现，从该点往后的所有商都会完全重复上一次出现 r 时的序列，这就是循环节。哈希表正是利用这一确定性来提前终止。

**相邻题**

- [No.029 两数相除](https://leetcode.cn/problems/divide-two-integers/)：只用加减法实现整数除法，同属"不用内置除法"的数字模拟题。
- [No.202 快乐数](../ch08hash/readme.md)：同为"检测数字循环"模式——用哈希表（或快慢指针）判断是否会重复，与本题的余数循环检测异曲同工。（收录于章节 readme）
