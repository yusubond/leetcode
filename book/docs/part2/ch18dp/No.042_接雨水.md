## 42 接雨水-困难

题目：

给定 `n` 个非负整数表示每个宽度为 `1` 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。



> **示例 1：**
>
> ![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/rainwatertrap.png)
>
> ```
> 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
> 输出：6
> 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 
> ```
>
> **示例 2：**
>
> ```
> 输入：height = [4,2,0,3,2,5]
> 输出：9
> ```



**解题思路**

这道题有多个思路可解。

- 动态规划，详见解法1
- 双指针，详见解法2
- 单调栈，详见解法3

对于每根柱子 i 来说，该位置上能够接多少雨水，取决于该柱子左右其他柱子高度的较小值，减去该柱子高度，就是该处的雨水量。

那么我们可以维护两个数组 leftMax 和 rightMax，`leftMax[i]`表示区间[0, i]的最大高度，递推公式如下：

```sh
i = 0, leftMax[i] = height[i]
0 < i < n, leftMax[i] = max(leftMax[i-1], height[i])
```

同理，对于 rightMax，表示右侧的最大值，有如下公式：

```sh
i = n-1, rightMax[i] = height[i]
0 <= i < n-1, rightMax[i] = max(rightMax[i+1], height[i])
```

有了这两个数组，我们就可以求每根柱子处的雨水，等于`min(leftMax[i], rightMax[i]) - height[i]`。

代码详见解法1。



解法1中，需要维护两个 max 数组，是不是可以用双指针代替呢？可以。

因为两个 max 数组的作用是为了求左右的较小值，那么我们可用左右指针来维护这个这两个变量。

只要有`height[left] < height[right]`，那么必有`leftMax < rightMax`；则此处雨水量为`leftMax - height[left]`

只要有`height[left] >= height[right]`，那么必有`leftMax > rightMax`，此处雨水量为`rightMax - height[right]`。

详见解法2。

解法1和解法2都是从"逐列"的角度来计算每根柱子的雨水量。换个角度，我们可以按"凹槽"逐层结算雨水，这就是单调栈解法。

维护一个单调递减栈（栈中存的是下标，对应高度从栈底到栈顶递减）。遍历每根柱子时，如果当前柱子比栈顶柱子高，说明出现了"右墙"，此时栈顶柱子即为凹槽的"底部"，将其弹出，新栈顶即为"左墙"，于是左墙、底部、右墙围成一个水洼，按层结算该水洼的雨水：

```sh
水洼雨水量 = (min(左墙, 右墙) - 底部) × 宽度
宽度 = 右墙下标 - 左墙下标 - 1
```

否则（当前柱子不高于栈顶），还未形成凹槽，将当前柱子入栈继续等待。

详见解法3。





```go
// date 2024/01/16
// 解法1
func trap(height []int) int {
    n := len(height)
    leftMax, rightMax := make([]int, n), make([]int, n)
    // fill left Max
    for i := 0; i < n; i++ {
        if i == 0 {
            leftMax[i] = height[i]
        } else {
            leftMax[i] = max(leftMax[i-1], height[i])
        }
    }

    // fill right Max
    for i := n-1; i >= 0; i-- {
        if i == n-1 {
            rightMax[i] = height[i]
        } else {
            rightMax[i] = max(rightMax[i+1], height[i])
        }
    }
    ans := 0

    for i := 0; i < n; i++ {
        ans += min(leftMax[i], rightMax[i]) - height[i]
    }

    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

// 解法2
// 双指针
func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	water := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}

	return water
}

// 解法3
// 单调栈，按"凹槽"逐层结算
func trap(height []int) int {
	stack := []int{} // 存下标，对应高度单调递减
	water := 0

	for i := 0; i < len(height); i++ {
		// 当前柱子比栈顶高 → 出现右墙，结算凹槽
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			bottom := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // 弹出凹槽底部

			if len(stack) == 0 {
				break // 左边没墙了，接不住水
			}

			leftWall := stack[len(stack)-1]
			distance := i - leftWall - 1
			water += (min(height[i], height[leftWall]) - height[bottom]) * distance
		}
		stack = append(stack, i)
	}

	return water
}
```

