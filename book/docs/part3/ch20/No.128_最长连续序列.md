## 128 最长连续序列-中等

题目：

给定一个未排序的整数数组 `nums` ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 `O(n)` 的算法解决此问题。



> **示例 1：**
>
> ```
> 输入：nums = [100,4,200,1,3,2]
> 输出：4
> 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
> ```
>
> **示例 2：**
>
> ```
> 输入：nums = [0,3,7,2,5,8,4,6,0,1]
> 输出：9
> ```



**解题思路**：

这道题就是求数组中连续元素的最大个数，要求时间复杂度`O(n)`。

- 可以先暴力求解，代码见解法一。思路是先把每个数存入map，然后对 map 进行清洗。

  - 遍历 map，把既没有前驱元素也没有后驱元素的节点全部删除。
  - 再次遍历 map，此时 map 中的元素肯定有连续的元素，但不确定连续有多少，那么依次往后查找并删除，在依次往前查找，并删除；过程中统计元素的个数，并更新最终的结果集。
  - 得到就是最长连续序列的个数

  需要注意 corner case 的处理。

- 解法二最优，思路是遍历数组，存入 map，存入的过程中做2件事情。第一件事，先查看`v-1`和`v+1`是否都存在 map 中，如果存在，表示存在连续序列，那么更新`left`和`right`边界，当前元素`v`对应的最小连续序列长度就是`sum = left + right + 1`。第二件事，更新左右边界对应元素的长度。

- 解法三，并查集思路。这个解法不是很好想，不过确实可以这样做，锻炼并查集思维，具体做法是：
  - 先把每个元素的下标当做一个集合。
  - 然后遍历数组，存入 map。如果元素的前一个值`v-1`在 map 中，或者后一个元素`v+1`在 map 中，则将它们的下标进行 `union()`。
  - 然后在`uf`中依次查询每个下标，找到元素最多的集合。



解法一：

```go
// date 2024/01/04
// 解法一
func longestConsecutive(nums []int) int {
    res := 0
    set := make(map[int]int, len(nums))
    for _, v := range nums {
        set[v] = 1
    }
		// case1 重复元素
    if len(set) < 2 {
        return len(set)
    }

    for k := range set {
        k0 := k-1
        k1 := k+1
        _, ok1 := set[k0]
        _, ok2 := set[k1]
        if !ok1 && !ok2 {
            delete(set, k)
        }
    }
    // case2 全部离散，没有连续的
    if len(set) == 0 {
        return 1
    }

    for k := range set {
        ans := 1
        kk := k+1
        _, ok := set[kk]
        for ok {
            delete(set, kk)
            kk++
            ans++
            _, ok = set[kk]
        }
        kk = k-1
        _, ok = set[kk]
        for ok {
            kk--
            ans++
            _, ok = set[kk]
        }
        res = max(res, ans)
    }

    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
```





解法二：

```go
// date 2024/01/04
// 解法二
func longestConsecutive(nums []int) int {
    res := 0
    // seqCt 存储包含num连续序列的长度
    seqCt := make(map[int]int)
    for _, num := range nums {
        if seqCt[num] == 0 {
            // left 表示 num 左边连续元素的个数
            // right 表示 num 右边连续元素的个数
            left, right := 0, 0
            if seqCt[num-1] > 0 {
                left = seqCt[num-1]
            }
            if seqCt[num+1] > 0 {
                right = seqCt[num+1]
            }

            // 左右加一起，再加上 num 本身，就是 num 所构成连续序列的个数
            sum := left + right + 1
            seqCt[num] = sum
            
            res = max(res, sum)

            // 因为序列是连续的，所以可以更新最左边和最右边的个数
            // 如果 left， right 为零，不影响结果
            seqCt[num-left] = sum
            seqCt[num+right] = sum
        }
    }

    return res
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
```



解法三：

```go
// date 2024/01/04
// 解法三
func longestConsecutive(nums []int) int {
    n := len(nums)
    idxMap := make(map[int]int, len(nums))
    // 初始化，每个idx都是一个集合
    for i, _ := range nums {
        idxMap[i] = 1
    }

    uf := newMyUnionFind(n)
    numMap := make(map[int]int, n)
    for i, v := range nums {
        // 跳过重复元素
        if _, ok := numMap[v]; ok {
            continue
        }
        // save idx to num map
        numMap[v] = i
        if idx2, ok1 := numMap[v-1]; ok1 {
            uf.union(i, idx2)
        }
        if idx3, ok2 := numMap[v+1]; ok2 {
            uf.union(i, idx3)
        }
    }

    res := 0
    for idx := range idxMap {
        p := uf.find(idx)
        if p != idx {
            idxMap[p]++
        }
        if idxMap[p] > res {
            res = idxMap[p]
        }
    }
    return res
}

type (
	MyUnionFind struct {
		parent []int // 存储每个节点的父结点
	}
)

// n 表示图中一共有多少个节点
func newMyUnionFind(n int) *MyUnionFind {
	u := &MyUnionFind{
		parent: make([]int, n),
	}
	// 初始化时, 每个节点的父结点都是自己
	for i := 0; i < n; i++ {
		u.parent[i] = i
	}
	return u
}

// 将两个节点x y 合并
func (u *MyUnionFind) union(x, y int) {
	xp, yp := u.find(x), u.find(y)
	if xp == yp {
		return
	}
	u.parent[yp] = xp
}

// 查找 x 的父结点
func (u *MyUnionFind) find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}
```

