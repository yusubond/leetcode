### 贪心算法

贪心算法一般用来解决需要"找到要做某事的最小数量"或”找到在某些情况下适合的最大物品数量的问题”，且提供无序的输入。

贪心算法的思想是每一步都选择最佳解决方案，最终获得全局最佳的解决方案。

### 相关题目

- 55 跳跃游戏
- 45 跳远游戏II



#### 55 跳跃游戏

思路分析

第一种：正向，检查当前的每一个值，看是否可能跳跃到更远的地方。

第二种：反向检查，看是否能够从最远处到达起点。

```go
// 正向检查
func canJump(nums []int) bool {
  if len(nums) <= 1 {return true}
  reach := nums[0]
  for i := 0; i <= reach && i < len(nums); i++ {
    if reach < i + nums[i] {
      reach = i + nums[i]
    }
  }
  return reach >= len(nums) - 1
}
// 反向检查
func canJump(nums []int) bool {
  if len(nums) <= 1 {return true}
  left := len(nums)-1
  for i := len(nums) - 2; i >= 0; i-- {
    if i + nums[i] >= left {
      left = i
    }
  }
  return left == 0
}
```

#### 45 跳跃游戏II



```go
func jump(nums[] int) int {
  if len(nums) <= 1 {return 0}
  reach := nums[0]
  for i := 0; i < len(nums) - 1; i++ {
    for j := i; j < reach; j++ {
      if j + nums[j] > reach {
        reach = j + nums[j]
      }
    }
  }
}
```

#### 452 用最少数量的箭引爆气球

思路分析

1 先对气球的结束end坐标进行排序

2 凡是气球的起始坐标大于当前能够引爆的气球的end坐标的均需要增加一支箭，并更新end坐标

```go
func findMinArrowsShots(points [][]int) int {
  if len(points) == 0 {return 0}
  sort.Slice(points, func(i, j int) bool {
    return s[i][1] < s[j][1]
  })
  res := 1
  end := points[0][1]
  for _, p := range points {
    if end < p[0] {
      res++
      end = p[1]
    }
  }
  return res
}
```

