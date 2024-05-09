## 621 任务调度器

题目：

给定一个用字符数组表示的 CPU 需要执行的任务列表。其中包含使用大写的 A - Z 字母表示的26 种不同种类的任务。任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。CPU 在任何一个单位时间内都可以执行一个任务，或者在待命状态。

然而，两个相同种类的任务之间必须有长度为 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。

你需要计算完成所有任务所需要的最短时间。

示例1：

```sh
输入: tasks = ["A","A","A","B","B","B"], n = 2
输出: 8
执行顺序: A -> B -> (待命) -> A -> B -> (待命) -> A -> B.
```



**解题思路**



```go
// 算法1：优先安排出现次数最多的任务
func leastInterval(tasks []byte, n int) int {
    m := make([]int, 26)
    for _, t := range tasks {
        m[t - 'A']++
    }
    sort.Slice(m, func(i, j int) bool {
        return m[i] < m[j]
    })
    var res int
    for m[25] > 0 {
        i := 0
        for i <= n {
            if m[25] == 0 {
                break
            }
            if i < 16 && m[25-i] > 0 {
                m[25-i]--
            }
            res++
            i++
        }
        sort.Slice(m, func(i, j int) bool {
            return m[i] < m[j]
        })
    }
    return res
}
// 算法二：设计，利用空闲时间[leetcode-cn]
func leastInterval(tasks []byte, n int) int {
    m := make([]int, 26)
    for _, t := range tasks {
        m[t - 'A']++
    }
    sort.Slice(m, func(i, j int) bool {
        return m[i] < m[j]
    })
    max_val := m[25] - 1
    idle := n * max_val
    for i := 24; i >= 0 && m[i] > 0; i-- {
        if m[i] <= max_val {
            idle -= m[i]
        } else {
            idle -= max_val
        }
    }
    if idle > 0 {
        return idle + len(tasks)
    }
    return len(tasks)
}
```

