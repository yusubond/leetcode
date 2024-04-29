## 315 计算右侧小于当前元素的个数-困难

题目：

给你一个整数数组 `nums` ，按要求返回一个新数组 `counts` 。数组 `counts` 有该性质： `counts[i]` 的值是 `nums[i]` 右侧小于 `nums[i]` 的元素的数量。



> **示例 1：**
>
> ```
> 输入：nums = [5,2,6,1]
> 输出：[2,1,1,0] 
> 解释：
> 5 的右侧有 2 个更小的元素 (2 和 1)
> 2 的右侧仅有 1 个更小的元素 (1)
> 6 的右侧有 1 个更小的元素 (1)
> 1 的右侧有 0 个更小的元素
> ```
>
> **示例 2：**
>
> ```
> 输入：nums = [-1]
> 输出：[0]
> ```
>
> **示例 3：**
>
> ```
> 输入：nums = [-1,-1]
> 输出：[0,0]
> ```



**解题思路**

这道题可用树形数组来解决。

前提是先讲原始数组排序去重，并离散化到一段连续的区域。然后对这段连续的区域进行树形数组。

树形数组可求前缀和，在本题中树形数组初始为 0；增量更新每个位置，值为1，那么得到的前缀和就是小于当前元素的个数。

具体过程如下：

![image](images/image315.svg)



```go
// date 2024/01/11
type MyBit struct {
    tree []int
}

func NewMyBit(cap int) *MyBit {
    // 原数组中的值都为零，所以直接初始化tree就可以
    return &MyBit{tree: make([]int, cap+1)}
}

func (this *MyBit) Query(idx int) int {
    sum := 0
    for idx >= 1 {
        sum += this.tree[idx]
        idx -= lowbit(idx)
    }
    return sum
}

func (this *MyBit) Add(idx int, val int) {
    for idx < len(this.tree) {
        this.tree[idx] += val
        idx += lowbit(idx)
    }
}

func countSmaller(nums []int) []int {
    n := len(nums)
    // c
    one := make([]int, n)
    copy(one, nums)

    // 先排序,然后去重，并离散化
    // 这样做的目的是1)升序排序，这样树形数组查询的时候，前缀和就是小于当前元素的个数
    // 2）把原数组中的值映射到一个连续的区域，便于做树形数组
    sort.Slice(one, func(i, j int) bool {
        return one[i] < one[j]
    })
    k := 1
    kth := map[int]int{one[0]: 1}
    for i := 1; i < n; i++ {
        if one[i] != one[i-1] {
            k++
            kth[one[i]] = k
        }
    }

    bit := NewMyBit(k)
    res := make([]int, 0, n)
    for i := n-1; i >= 0; i-- {
        v := kth[nums[i]]
        // 求小于v的个数，就是求 v-1 前缀和
        res = append(res, bit.Query(v-1))
        bit.Add(v, 1)
    }
    reverseArr(res)
    return res
}

func reverseArr(arr []int) []int {
    left, right := 0, len(arr)-1
    for left < right {
        arr[left], arr[right] = arr[right], arr[left]
        left++
        right--
    }
    return arr
}

func lowbit(x int) int {
    return x & -x
}
```

