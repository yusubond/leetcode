## 303 区域和检索-数组不可变-中等

题目：

给定一个整数数组  `nums`，处理以下类型的多个查询:

1. 计算索引 `left` 和 `right` （包含 `left` 和 `right`）之间的 `nums` 元素的 **和** ，其中 `left <= right`

实现 `NumArray` 类：

- `NumArray(int[] nums)` 使用数组 `nums` 初始化对象
- `int sumRange(int i, int j)` 返回数组 `nums` 中索引 `left` 和 `right` 之间的元素的 **总和** ，包含 `left` 和 `right` 两点（也就是 `nums[left] + nums[left + 1] + ... + nums[right]` )



**解题思路**

这道题可有两个思路。第一个，既然题目中已经说明元素不可变，那么**前缀和**可以实现。具体为遍历数组，计算前缀和，并将前缀和存储另一个数组。查找的时候直接在前缀和数组中做减法即可。

第二个思路是线段树，只要实现 Query 方法即可，详见下面代码。

```go
// date 2024/01/09
type NumArray struct {
    data []int
    tree []int
}


func Constructor(nums []int) NumArray {
    n := len(nums)
    res := NumArray{
        data: make([]int, n),
        tree: make([]int, 4*n),
    }
    for i := 0; i < n; i++ {
        res.data[i] = nums[i]
    }
    res.buildSegmentTree(0, 0, n-1)
    return res
}


func (this *NumArray) SumRange(left int, right int) int {
    if left >= 0 && right < len(this.data) {
        return this.queryInTree(0, 0, len(this.data)-1, left, right)
    }
    return 0
}

func (this *NumArray) buildSegmentTree(root, left, right int) {
    if left == right {
        this.tree[root] = this.data[left]
        return
    }
    mid := left + (right-left)/2
    tl := root*2+1
    tr := root*2+2
    this.buildSegmentTree(tl, left, mid)
    this.buildSegmentTree(tr, mid+1, right)
    this.tree[root] = this.tree[tl] + this.tree[tr]
}

func (this *NumArray) queryInTree(root, tl, tr, left, right int) int {
    if left == tl && right == tr {
        return this.tree[root]
    }
    mid := tl + (tr - tl)/2
    if left > mid {
        return this.queryInTree(root*2+2, mid+1, tr, left, right)
    }
    if right <= mid {
        return this.queryInTree(root*2+1, tl, mid, left, right)
    }
    lsum := this.queryInTree(root*2+1, tl, mid, left, mid)
    rsum := this.queryInTree(root*2+2, mid+1, tr, mid+1, right)
    return lsum+rsum
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
```

