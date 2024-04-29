package template

import (
	"testing"
)

func TestNewBinaryIndexedTree(t *testing.T) {
	arr := []int{2, 1, 1, 3, 2, 3, 4, 5, 6, 7, 8, 9}
	preSum := getPreSum(arr)

	bit := NewBinaryIndexedTree(arr)
	bit.Show()

	// define test query func
	testQuery := func(arr []int) {
		n := len(arr)
		sum := 0
		for i := 0; i < n; i++ {
			sum += arr[i]
			res := bit.Query(i)
			if res != sum {
				t.Fatalf("fail, query [0, %d] = %d, expect %d\n", i, res, sum)
			}
		}
	}

	// test query
	testQuery(arr)

	// test sum range
	left, right := 2, 6
	res := bit.SumRange(left, right)
	expect := preSum[right] - preSum[left-1]
	if res != expect {
		t.Fatalf("SumRange fail, get sum range of [%d, %d] = %d, expect %d\n", left, right, res, expect)
	}

	// test add
	arr[3] += 3
	preSum = getPreSum(arr)

	bit.Add(3, 3)
	bit.Show()

	testQuery(arr)
	res = bit.SumRange(left, right)
	expect = preSum[right] - preSum[left-1]
	if res != expect {
		t.Fatalf("SumRange fail, get sum range of [%d, %d] = %d, expect %d\n", left, right, res, expect)
	}
}

func getPreSum(arr []int) []int {
	n := len(arr)
	preSum := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
		preSum[i] = sum
	}
	return preSum
}
