package template

import (
	"fmt"
	"testing"
)

func TestSegmentTree_Init(t *testing.T) {
	arr := []int{0, 1, 3, 5, -2, 3}
	sg := NewSegmentTree()
	sg.Init(arr)
	sg.Show()
	res := sg.Query(0, 4)
	fmt.Printf("sum [0, 4] = %d\n", res)

	res2 := sg.Query(3, 4)
	fmt.Printf("sum [3, 4] = %d\n", res2)

	sg.Update(3, 12)
	fmt.Printf("update index 3 to 12\n")
	sg.Show()
	res2 = sg.Query(3, 4)
	fmt.Printf("sum [3, 4] = %d\n", res2)
}
