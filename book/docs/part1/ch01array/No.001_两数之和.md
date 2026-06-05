## 1 两数之和-简单

题目：

给定一个整数数组 `nums` 和一个整数目标值 `target`，请你在该数组中找出 **和为目标值** *`target`* 的那 **两个** 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。



**解题思路**

题目中已经说明每种输入只会对应一种输出，所以可以有多种解法。

- 两层循环。第一层遍历当前 x，然后从元素的 x 的下一个元素 y 开始遍历，如果 x + y 等于目标值，则返回两者的坐标，详见解法1。

- 哈希表。利用map存储已经遍历过的元素和下标，达到O(1)查找，详见解法2。时间和空间复杂度都是`O(N)`。

```go
// date 2020/09/14
// 解法1
// 两层循环
// 时间复杂度O(N^2)
func twoSum(nums []int, target int) []int {
    res := make([]int, 0, 2)
    n := len(nums)

    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            if nums[i] + nums[j] == target {
                res = append(res, i, j)
                return res
            }
        }
    }
    
    return res
}
// 解法2
// map 查找
func twoSum(nums []int, target int) []int {
    set := make(map[int]int, len(nums))
    res := []int{}
    for j, v := range nums {
        d := target - v
        if i, ok := set[d]; ok && i != j {
            res = append(res, i, j)
            break
        } else {
            set[v] = j
        }
    }
    return res
}
```

扩展题目：两数之和

**问题延伸：给定一个无序且元素可重复的数组a[],以及一个数sum，求a[]中是否存在两个元素的和等于sum，并输出这两个元素的下标。答案可能不止一种，请输出所有可能的答案结果集。**

这是VMware面试中的一道题，以下是个人的解法。这个算法的时间复杂度为O(n)，但是需要辅助空间。

```go
// data 2021/02/21
// 解法一：两层循环
func twoSum(nums []int, target int) [][]int {
  res := make([][]int, 0, 16)
  idx := make(map[int][]int, 16)
  for i, v := range nums {
    d := target - v
    if oldIdx, ok := idx[d]; ok && len(oldIdx) != 0 {
      for _, j := range oldIdx {
        res = append(res, []int{j, i})
      }
    }
    _, ok := idx[v]
    if !ok {
      idx[v] = []int{}
    }
    idx[v] = append(idx[v], i)
  }
  return res
}
```

