## 933 最近的请求次数-简单

题目：

写一个 `RecentCounter` 类来计算特定时间范围内最近的请求。

请你实现 `RecentCounter` 类：

- `RecentCounter()` 初始化计数器，请求数为 0 。
- `int ping(int t)` 在时间 `t` 添加一个新请求，其中 `t` 表示以毫秒为单位的某个时间，并返回过去 `3000` 毫秒内发生的所有请求数（包括新请求）。确切地说，返回在 `[t-3000, t]` 内发生的请求数。

**保证** 每次对 `ping` 的调用都使用比之前更大的 `t` 值。



分析：

先直接入队，入队后，把历史上超过时间的出队即可。

```go
// date 2023/11/30
type RecentCounter struct {
    ct []int
}


func Constructor() RecentCounter {
    return RecentCounter{
        ct: make([]int, 0, 16),
    }
}


func (this *RecentCounter) Ping(t int) int {
    this.ct = append(this.ct, t)
    start := t - 3000
    for i := 0; i < len(this.ct); i++ {
        if this.ct[i] >= start {
            this.ct = this.ct[i:]
            break
        }
    }
    return len(this.ct)
}
```