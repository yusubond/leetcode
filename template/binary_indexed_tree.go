package template

import "fmt"

type (
	// BinaryIndexedTree define
	BinaryIndexedTree struct {
		capacity int
		data     []int // array A
		tree     []int // array B
	}
)

// NewBinaryIndexedTree define
func NewBinaryIndexedTree(arr []int) *BinaryIndexedTree {
	n := len(arr)
	b := &BinaryIndexedTree{
		capacity: n + 1,
		data:     make([]int, n+1),
		tree:     make([]int, n+1), // 为了计算方便，初始化为n+1
	}
	for i := 1; i <= n; i++ {
		b.data[i] = arr[i-1]
		b.tree[i] = arr[i-1]
		for j := i - 2; j >= i-lowbit(i); j-- {
			b.tree[i] += arr[j]
		}
	}
	return b
}

// Query define
// 求区间和，返回原始数组中[0, index] 的区间和
// 输入index 为原始数组的下标，需要+1
func (b *BinaryIndexedTree) Query(i int) int {
	return b.queryWithBIT(i + 1)
}

// SumRange define
// 对任意区间求和，返回原始数组 A [start, end]区间和
func (b *BinaryIndexedTree) SumRange(left, right int) int {
	s1 := b.Query(left - 1)
	s2 := b.Query(right)
	return s2 - s1
}

// Add define
// 增量更新，即对原始数组A index下标的值 增加 val
func (b *BinaryIndexedTree) Add(i, val int) {
	b.addWithBit(i+1, val)
}

// Set define
func (b *BinaryIndexedTree) Set(i, val int) {
	// old := b.data[i+1]
	// dt := val - old
	b.Add(i, val-b.data[i+1])
}

func (b *BinaryIndexedTree) queryWithBIT(index int) int {
	sum := 0
	for index >= 1 {
		sum += b.tree[index]
		index -= lowbit(index)
	}
	return sum
}

func (b *BinaryIndexedTree) addWithBit(idx int, val int) {
	b.data[idx] += val
	for idx <= b.capacity {
		b.tree[idx] += val
		idx += lowbit(idx)
	}
}

func (b *BinaryIndexedTree) Show() {
	fmt.Printf("data = %d\n", b.data)
	fmt.Printf("tree = %d\n", b.tree)
}

// 返回 x 二进制形式中，末尾最后一个1所代表的数值，即 2^k
// k 为 x 二进制中末尾零的的个数
func lowbit(x int) int {
	return x & -x
}
