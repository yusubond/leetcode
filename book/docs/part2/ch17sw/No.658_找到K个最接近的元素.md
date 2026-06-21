## 658 找到K个最接近的元素-中等

题目：

给定一个 **排序好** 的数组 `arr` ，两个整数 `k` 和 `x` ，从数组中找到最靠近 `x`（两数之差最小）的 `k` 个数。返回的结果必须要是按升序排好的。

整数 `a` 比整数 `b` 更接近 `x` 需要满足：

- `|a - x| < |b - x|` 或者
- `|a - x| == |b - x|` 且 `a < b`



分析：

分情况讨论：

1. 如果 k 大于等于数组长度，那么整个数组就是答案
2. 如果 x 小于等于第一个元素，那么数组的前 K 个就是答案
3. 如果 x 大于等于数组的最后一个元素，那么数组的后 K 个就是答案
4. 其他情况，找到第一个大于等于 x 的元素，然后双向指针向两头查找；
5. 一旦一端指针达到临界后，且不满 K 个，那么从另一端开始凑



优化部分：第 4 步骤查找 x 可以使用二分查找。

```go
// date 2023/11/21
func findClosestElements(arr []int, k int, x int) []int {
    n := len(arr)
    if n == 0 || k >= n {
        return arr
    }

    if x <= arr[0] {
        return arr[0:k]
    } else if x >= arr[n-1] {
        return arr[n-k:n]
    }
    left, right := 0, 0
    for i := 0; i < n; i++ {
        if arr[i] >= x {
            right = i
            break
        }
    }
    left = right-1
    for k > 0 {
        if left >= 0 && right < n {
            r1 := abs(arr[right], x)
            l1 := abs(arr[left], x)
            if r1 < l1 {
                right++
            } else if l1 <= r1 {
                left--
            }
        } else if right < n {
            right++
        } else if left >= 0 {
            left--
        }
        k--
    }

    return arr[left+1:right]
}

func abs(x, y int) int {
    if x > y {
        return x-y
    }
    return y-x
}
```

