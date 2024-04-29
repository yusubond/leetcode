package template

import (
	"fmt"
	"testing"
)

var (
	unSortArr = []int{5, 4, 3, 3, 2, 2, 1}
)

func TestBubbleSort(t *testing.T) {
	// deep copy from un sort array
	n := len(unSortArr)
	one := make([]int, n)
	copy(one, unSortArr)

	fmt.Printf("unSortArr = %d\n", one)
	res := BubbleSort(one)
	checkSortArr(t, res)
	fmt.Printf("sort res = %d\n", res)
}

func TestSelectSort(t *testing.T) {
	// deep copy from un sort array
	n := len(unSortArr)
	one := make([]int, n)
	copy(one, unSortArr)

	fmt.Printf("unSortArr = %d\n", one)
	res := SelectSort(one)
	checkSortArr(t, res)
	fmt.Printf("sort res = %d\n", res)
}

func TestInsertSort(t *testing.T) {
	// deep copy from un sort array
	n := len(unSortArr)
	one := make([]int, n)
	copy(one, unSortArr)

	fmt.Printf("unSortArr = %d\n", one)
	res := InsertSort(one)
	checkSortArr(t, res)
	fmt.Printf("sort res = %d\n", res)
}

func TestQuickSort(t *testing.T) {
	// deep copy from un sort array
	n := len(unSortArr)
	one := make([]int, n)
	copy(one, unSortArr)

	fmt.Printf("unSortArr = %d\n", one)
	res := QuickSort(one)
	checkSortArr(t, res)
	fmt.Printf("sort res = %d\n", res)
}

func TestMergeSort(t *testing.T) {
	// deep copy from un sort array
	n := len(unSortArr)
	one := make([]int, n)
	copy(one, unSortArr)

	fmt.Printf("unSortArr = %d\n", one)
	res := MergeSort(one)
	checkSortArr(t, res)
	fmt.Printf("sort res = %d\n", res)
}

func checkSortArr(t *testing.T, sortArr []int) {
	expectLen := len(unSortArr)
	if len(sortArr) != expectLen {
		t.Fatalf("expect arr len %d\n", expectLen)
	}
	for i := 1; i < expectLen; i++ {
		if sortArr[i-1] > sortArr[i] {
			t.Fatalf("sort fail res = %d\n", sortArr)
		}
	}
}
