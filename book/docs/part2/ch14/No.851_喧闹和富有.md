## 851 喧闹和富有-中等

题目：

有一组 `n` 个人作为实验对象，从 `0` 到 `n - 1` 编号，其中每个人都有不同数目的钱，以及不同程度的安静值（quietness）。为了方便起见，我们将编号为 `x` 的人简称为 "person `x` "。

给你一个数组 `richer` ，其中 `richer[i] = [ai, bi]` 表示 person `ai` 比 person `bi` 更有钱。另给你一个整数数组 `quiet` ，其中 `quiet[i]` 是 person `i` 的安静值。`richer` 中所给出的数据 **逻辑自洽**（也就是说，在 person `x` 比 person `y` 更有钱的同时，不会出现 person `y` 比 person `x` 更有钱的情况 ）。

现在，返回一个整数数组 `answer` 作为答案，其中 `answer[x] = y` 的前提是，在所有拥有的钱肯定不少于 person `x` 的人中，person `y` 是最安静的人（也就是安静值 `quiet[y]` 最小的人）。



分析：

这道题的整体思路是 AOV 网的拓扑排序。

由题意可知，`richer`关系构成了一张有向无环图，则据此构建图的关系，记录每个节点的入度和出度节点。

利用 AOV 网的拓扑排序，从入度为零的节点开始广度优先遍历；不断地使用更富有的人，但安静值更小的人。

当所有入度为零的序列为零时，则终止搜索。

初始化`ans[i] = i`，拓扑排序时，如果`quiet[ans[i]] > quiet[ans[j]]`，那么`ans[i] = ans[j]`。

```go
// date 2024/01/01
func loudAndRich(richer [][]int, quiet []int) []int {
    n := len(quiet)
    ans := make([]int, n)
    for i := 0; i < n; i++ {
        ans[i] = i
    }

    in := make([]int, n)
    out := make(map[int][]int, 4)
    for _, v := range richer {
        in[v[1]]++
        _, ok := out[v[0]]
        if !ok {
            out[v[0]] = make([]int, 0, 16)
        }
        out[v[0]] = append(out[v[0]], v[1])
    }

    // aov 网 拓扑排序
    queue := make([]int, 0, 16)
    for i, v := range in {
        if v == 0 {
            queue = append(queue, i)
        }
    }

    for len(queue) != 0 {
        n := len(queue)
        for i := 0; i < n; i++ {
            u := queue[i]
            if ov, ok := out[u]; ok && len(ov) != 0 {
                for _, v := range ov {
                    in[v]--
                    if quiet[ans[v]] > quiet[ans[u]] {
                        ans[v] = ans[u]
                    }
                    if in[v] == 0 {
                        queue = append(queue, v)
                    }
                }
            }
        }

        queue = queue[n:]
    }

    return ans
}
```

