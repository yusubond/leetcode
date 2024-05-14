package template

import (
	"fmt"
)

type (
	// SegmentTree define
	SegmentTree struct {
		data []int // 存储原始数据
		tree []int // 存储线段树数据
	}
)

func NewSegmentTree() *SegmentTree {
	return &SegmentTree{}
}

func (s *SegmentTree) Init(nums []int) {
	n := len(nums)
	s.data = make([]int, n)
	s.tree = make([]int, 4*n)
	for i := 0; i < n; i++ {
		s.data[i] = nums[i]
	}
	// build segment tree from nums
	left, right := 0, n-1
	s.buildSegmentTree(0, left, right)
}

func (s *SegmentTree) buildSegmentTree(treeIdx int, left, right int) {
	if left == right {
		// the leaf node
		s.tree[treeIdx] = s.data[left]
		return
	}
	mid := left + (right-left)/2
	leftTreeIdx := leftChild(treeIdx)
	rightTreeIdx := rightChild(treeIdx)
	// build left tree with data [left, mid] elem
	s.buildSegmentTree(leftTreeIdx, left, mid)
	// build right tree with data [mid+1, right] elem
	s.buildSegmentTree(rightTreeIdx, mid+1, right)
	// 注意，递归构建左右子树，别忘了最后求和
	s.tree[treeIdx] = s.tree[leftTreeIdx] + s.tree[rightTreeIdx]
}

// Query define
// 查询 [left, right] 区间和
func (s *SegmentTree) Query(left, right int) int {
	if len(s.data) > 0 {
		return s.queryInTree(0, 0, len(s.data)-1, left, right)
	}
	return 0
}

// 在以 root 为根的线段树中，[tl, tr] 范围内, 搜索 [left, right] 区间值
// 如果线段树的区间和查询区间一致，那么直接返回线段树的节点
// 否则分左子树，右子树，左右子树三个情况，分别查询
func (s *SegmentTree) queryInTree(root, tl, tr int, left, right int) int {
	// 如果区间一样，直接返回根节点数据
	if left == tl && right == tr {
		return s.tree[root]
	}
	mid := tl + (tr-tl)/2
	leftTree, rightTree := leftChild(root), rightChild(root)
	// query in the right tree
	// 注意，如果查询区间在某个子树里，直接返回查询子树的结构
	// 此时，查询子树，left, right 不能变，因为，区间[tl, tr] 包含 [left, right]
	if left > mid {
		return s.queryInTree(rightTree, mid+1, tr, left, right)
	} else if right <= mid {
		// query in the left tree
		return s.queryInTree(leftTree, tl, mid, left, right)
	}
	// 区间跨两个子树，那么分开查询
	// 此时，left, right 要跟线段树的左右区间边界重合
	lSum := s.queryInTree(leftTree, tl, mid, left, mid)
	rSum := s.queryInTree(rightTree, mid+1, tr, mid+1, right)
	return lSum + rSum
}

// Update define
// 数值全量更新，即原始数组中下标为 index 的值更新为 val
func (s *SegmentTree) Update(index, val int) {
	if len(s.data) > 0 {
		s.updateInTree(0, 0, len(s.data)-1, index, val)
	}
}

func (s *SegmentTree) updateInTree(root, tl, tr int, index, val int) {
	// find the leaf node, so update its value
	if tl == tr {
		s.tree[root] = val
		s.data[tl] = val
		return
	}

	mid := tl + (tr-tl)/2
	leftTree := leftChild(root)
	rightTree := rightChild(root)
	// if value index is at the right part, then update the right tree
	// otherwise update the left tree
	if index > mid {
		s.updateInTree(rightTree, mid+1, tr, index, val)
	} else if index <= mid {
		s.updateInTree(leftTree, tl, mid, index, val)
	}
	// update the parent's value
	s.tree[root] = s.tree[leftTree] + s.tree[rightTree]
}

func (s *SegmentTree) Show() {
	fmt.Printf("data = %d\n", s.data)
	fmt.Printf("tree = %d\n", s.tree)
}

func leftChild(treeIdx int) int {
	return treeIdx*2 + 1
}

func rightChild(treeIdx int) int {
	return treeIdx*2 + 2
}
