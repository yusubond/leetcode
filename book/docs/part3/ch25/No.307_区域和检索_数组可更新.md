## 307 区域和检索-数组可修改-中等

题目：

给你一个数组 `nums` ，请你完成两类查询。

1. 其中一类查询要求 **更新** 数组 `nums` 下标对应的值
2. 另一类查询要求返回数组 `nums` 中索引 `left` 和索引 `right` 之间（ **包含** ）的nums元素的 **和** ，其中 `left <= right`

实现 `NumArray` 类：

- `NumArray(int[] nums)` 用整数数组 `nums` 初始化对象
- `void update(int index, int val)` 将 `nums[index]` 的值 **更新** 为 `val`
- `int sumRange(int left, int right)` 返回数组 `nums` 中索引 `left` 和索引 `right` 之间（ **包含** ）的nums元素的 **和** （即，`nums[left] + nums[left + 1], ..., nums[right]`）



**解题思路**

这道题既可以用标准的线段数解法，也可以用树形数组（即二叉索引树）解决。线段树解法详见线段树部分，下面给出树形数组的解法。

注意，这里的树形数组更新使用的是增量更新，所以维护原始数组的数据。

```go
// 2024/01/10
// binary indexed tree
type NumArray struct {
    data []int
    tree []int
}

func Constructor(nums []int) NumArray {
    n := len(nums)
    res := NumArray{
        data: make([]int, n+1),
        tree: make([]int, n+1),
    }
    // B[i] = sum(j, i) of A
    // j = i - 2^k - 1
    for i := 1; i <= n; i++ {
        res.data[i] = nums[i-1]
        res.tree[i] = nums[i-1]
        for j := i-2; j >= i-lowbit(i); j-- {
            res.tree[i] += nums[j]
        }
    }
    return res
}


func (this *NumArray) Update(index int, val int)  {
    // convert to bit index
    index++
    dt := val - this.data[index]
    this.data[index] = val
    this.update(index, dt)
}


func (this *NumArray) SumRange(left int, right int) int {
    s1 := this.query(left-1)
    s2 := this.query(right)
    return s2-s1
}

// 增量更新
func (this *NumArray) update(index, val int) {
    // son to parent
    for index < len(this.tree) {
        this.tree[index] += val
        index += lowbit(index)
    }
}

func (this *NumArray) query(index int) int {
    // sum is [1, index]
    sum := 0
    index += 1
    for index >= 1 {
        sum += this.tree[index]
        index -= lowbit(index)
    }
    return sum
}

func lowbit(x int) int {
    return x & -x
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
```

