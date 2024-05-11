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

这道题就是标准的线段树解法，有几个注意点：

- 查询的时候，如果查询区间在某个子树里面，直接查询子树，且查询区间直接透传不变。如果查询区间跨根节点，那么需要左右一起查，并合并结果，这时查询区间要跟线段树的区间边界保持一致。
- 更新的时候。无论更新左子树还是右子树，最后都要记得更新根节点。

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


func (this *NumArray) Update(index int, val int)  {
    if len(this.data) > 0 {
        this.updateInTree(0, 0, len(this.data)-1, index, val)
    }
}


func (this *NumArray) SumRange(left int, right int) int {
    if len(this.data) > 0 {
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
    ltree := 2*root+1
    rtree := 2*root+2
    this.buildSegmentTree(ltree, left, mid)
    this.buildSegmentTree(rtree, mid+1, right)
    this.tree[root] = this.tree[ltree] + this.tree[rtree]
}

func (this *NumArray) queryInTree(root, tl, tr, left, right int) int {
    if left == tl && right == tr {
        return this.tree[root]
    }
    mid := tl + (tr-tl)/2
    ltree := root*2+1
    rtree := root*2+2
    if left > mid {
        return this.queryInTree(rtree, mid+1, tr, left, right)
    } else if right <= mid {
        return this.queryInTree(ltree, tl, mid, left, right)
    }
    lsum := this.queryInTree(ltree, tl, mid, left, mid)
    rsum := this.queryInTree(rtree, mid+1, tr, mid+1, right)
    return lsum + rsum
}

func (this *NumArray) updateInTree(root, tl, tr int, index, val int) {
    if tl == tr {
        this.tree[root] = val
        this.data[index] = val
        return
    }
    mid := tl + (tr-tl)/2
    ltree := 2*root+1
    rtree := 2*root+2
    if index > mid {
        this.updateInTree(rtree, mid+1, tr, index, val)
    } else {
        this.updateInTree(ltree, tl, mid, index, val)
    }
    this.tree[root] = this.tree[ltree]+this.tree[rtree]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
```

