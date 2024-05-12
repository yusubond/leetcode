## 480 滑动窗口中位数-困难

题目：

中位数是有序序列中最中间的那个数。如果序列长度为偶数，那么则是最中间两个数的平均数。

给定一个数组 `nums` 和 整数 `k`，`k` 表示滑动窗口大小，从最左边滑动到最右边，每次窗口向右移动1位，请你找出滑动窗口中的中位数。



分析：

因为原始数组是无序的，所以形成的窗口也是无序的。需要借助辅助数组，对窗口内的元素进行排序。

另外，如果形成窗口之后进行排序，会超时，所以构造窗口时就要保证辅助数组的有序性。

为此，对辅助数组实现三个函数：

```
1. add(x int) 插入排序，保证插入元素后辅助数组有序
2. del(x int) 在有序数组中删除一个元素
3. mid() 返回辅助数组的中位数
```

因为 k 可能为奇数或偶数，所以分别设计。

```go
func medianSlidingWindow(nums []int, k int) []float64 {
    left, right := 0, 0
    ans := make([]float64, 0, 64)
    f := k & 0x1 == 0x1
    s := k >> 1

    lt := &winList{one: f, step: s, data: make([]int, 0, 16)}

    for right < len(nums) {
        lt.add(nums[right])
        right++
        if right - left >= k {
            ans = append(ans, lt.mid())
            lt.del(nums[left])
            left++
        }
    }

    return ans
}

type winList struct {
    one bool
    step int
    data []int
}

func (w *winList) add(x int) {
    n := len(w.data)
    w.data = append(w.data, 0)
    i := n-1
    for i >= 0 {
        if x > w.data[i] {
            break
        } else {
            w.data[i+1] = w.data[i]
        }
        i--
    }
    w.data[i+1] = x
}

func (w *winList) del(x int) {
    i := 0
    n := len(w.data)
    s1 := n-1
    for i < n {
        if w.data[i] == x {
            copy(w.data[i:], w.data[i+1:])
            break
        } else {
            i++
        }
    }
    w.data = w.data[:s1]
}

func (w *winList) mid() float64 {
    if w.one {
        return float64(w.data[w.step])
    }
    return (float64(w.data[w.step-1]) + float64(w.data[w.step])) / 2.0
}
```
