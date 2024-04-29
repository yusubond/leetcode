## 421 数组中两个数的最大异或值

题目：

给你一个整数数组 `nums` ，返回 `nums[i] XOR nums[j]` 的最大运算结果，其中 `0 ≤ i ≤ j < n` 。

 **提示：**

- `1 <= nums.length <= 2 * 105`
- `0 <= nums[i] <= 231 - 1`



> **示例 1：**
>
> ```
> 输入：nums = [3,10,5,25,2,8]
> 输出：28
> 解释：最大运算结果是 5 XOR 25 = 28.
> ```
>
> **示例 2：**
>
> ```
> 输入：nums = [14,70,53,83,49,91,36,80,92,51,66,70]
> 输出：127
> ```



**解题思路**

异或，就是二进制位中不同为 1，相同为 0。要求两个数的异或值最大，那么这两个数的二进制位应该尽可能的不一样。

最朴素的解法是两层循环，依次计算。

更快的方法是字典树。

题目中已经说明数组中的值在区间`[0, 2^31-1]`之间，所以我们可以按位处理每个值，存储在字典树中。

考虑到二进制中每位只可能是0或者1，所以字典树的节点可以用左右表示：

```go
type Trie struct {
  left, right *Trie  // left for 0, right for 1
}
```

具体为：

- 把数组第一个元素加入字典树
- 从第二个元素x开始，先从字典树中查找与 x 异或的最大值，然后再把 x 也加入到字典树中。
- 遍历的时候，顺便保留最大的异或值即可。

如何从字典树中查找最大的异或值呢？

从字典树的根节点开始，如果元素 x 的当前位为 0，那么就往字典树的 right（即该位 为1）继续查找；否则，往字典树的 left 中查找。



```go
// date 2024/01/26
type Trie struct {
    left, right *Trie // left 0, right 1
}

func (t *Trie) Add(num int) {
    cur := t
    for i := 30; i >= 0; i-- {
        bit := num >> i & 0x1
        if bit == 0 {
            // left 0
            if cur.left == nil {
                cur.left = &Trie{}
            }
            cur = cur.left
        } else {
            if cur.right == nil {
                cur.right = &Trie{}
            }
            cur = cur.right
        }
    }
}

func (t *Trie) findMaxXORForNum(num int) int {
    cur := t
    ans := 0
    for i := 30; i >= 0; i-- {
        bit := num >> i & 0x1
        ans = ans << 1
        if bit == 0 {
            // use the right of exist num
            if cur.right != nil {
                cur = cur.right
                ans += 1
            } else {
                cur = cur.left
            }
        } else {
            // use the left
            if cur.left != nil {
                cur = cur.left
                ans += 1
            } else {
                cur = cur.right
            }
        }
    }
    return ans
}

func findMaximumXOR(nums []int) int {
    ans := 0
    tie := &Trie{}
    for i, v := range nums {
        if i == 0 {
            tie.Add(v)
            continue
        }
        res := tie.findMaxXORForNum(nums[i])
        if res > ans {
            ans = res
        }
        tie.Add(nums[i])
    }
    return ans
}
```



为什么一次遍历就可以？

因为异或满足交换律。