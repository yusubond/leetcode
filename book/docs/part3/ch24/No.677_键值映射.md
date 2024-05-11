## 677 键值映射-中等

题目：

设计一个 map ，满足以下几点:

- 字符串表示键，整数表示值
- 返回具有前缀等于给定字符串的键的值的总和

实现一个 `MapSum` 类：

- `MapSum()` 初始化 `MapSum` 对象
- `void insert(String key, int val)` 插入 `key-val` 键值对，字符串表示键 `key` ，整数表示值 `val` 。如果键 `key` 已经存在，那么原来的键值对 `key-value` 将被替代成新的键值对。
- `int sum(string prefix)` 返回所有以该前缀 `prefix` 开头的键 `key` 的值的总和。



> **示例 1：**
>
> ```
> 输入：
> ["MapSum", "insert", "sum", "insert", "sum"]
> [[], ["apple", 3], ["ap"], ["app", 2], ["ap"]]
> 输出：
> [null, null, 3, null, 5]
> 
> 解释：
> MapSum mapSum = new MapSum();
> mapSum.insert("apple", 3);  
> mapSum.sum("ap");           // 返回 3 (apple = 3)
> mapSum.insert("app", 2);    
> mapSum.sum("ap");           // 返回 5 (apple + app = 3 + 2 = 5)
> ```



**解题思路**

这道题可以有多种解法。

解法1：标准的map和strings库函数，求和的时候通过strings.HasPrefix函数来判断是否包含前缀。

解法2：字典树，将key存放在字典树中，key对应的值存在key结束的地方。这样做的好处是，插入的时候可以无脑的可操作，求和时深搜即可。

解法3：字典树，增量更新。这样做的好处是插入的时候使用增量更新，求和的时候，在字典树中查找即可。



```go
// date 2024/03/30
// 解法1
type MapSum struct {
    cache map[string]int
}


func Constructor() MapSum {
    return MapSum{
        cache: make(map[string]int),
    }
}


func (this *MapSum) Insert(key string, val int)  {
    this.cache[key] = val
}


func (this *MapSum) Sum(prefix string) int {
    ans := 0

    for k, v := range this.cache {
        if strings.HasPrefix(k, prefix) {
            ans += v
        }
    }

    return ans
}


/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */

// 解法2
type MapSum struct {
    child [26]*MapSum
    isEnd bool
    size int
}


func Constructor() MapSum {
    m := MapSum{
        child: [26]*MapSum{},
        isEnd: false,
        size: 0,
    }
    return m
}


func (this *MapSum) Insert(key string, val int)  {
    cur := this
    for i := 0; i < len(key); i++ {
        idx := key[i] - 'a'
        if cur.child[idx] == nil {
            cur.child[idx] = &MapSum{child: [26]*MapSum{}}
        }
        cur = cur.child[idx]
    }
    cur.isEnd = true
    cur.size = val
}


func (this *MapSum) Sum(prefix string) int {
    ans := 0
    cur := this

    i, n := 0, len(prefix)
    for i = 0; i < n; i++ {
        idx := prefix[i] - 'a'
        if cur.child[idx] != nil {
            cur = cur.child[idx]
        } else {
            break
        }
    }

    if i < n || cur == nil {
        return 0
    }

    // dfs cur
    var dfs func(root *MapSum)
    dfs = func(root *MapSum) {
        if root.isEnd {
            ans += root.size
        }
        for i := 0; i < 26; i++ {
            if root.child[i] != nil {
                dfs(root.child[i])
            }
        }
    }

    dfs(cur)

    return ans
}


/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */

// 解法3
type TrieSumNode struct {
	child [26]*TrieSumNode
	sum   int
}

type MapSum struct {
	cache map[string]int
	root  *TrieSumNode
}


func Constructor() MapSum {
    m := MapSum{
        cache: make(map[string]int),
        root: newTrieSumNode(),
    }
    return m
}

func newTrieSumNode() *TrieSumNode {
	t := &TrieSumNode{
		child: [26]*TrieSumNode{},
		sum:   0,
	}
	return t
}


func (this *MapSum) Insert(key string, val int)  {
    old, ok := this.cache[key]
    if !ok {
        this.root.insertKey(key, val)
    } else {
        this.root.insertKey(key, val - old)
    }
    this.cache[key] = val
}


func (this *MapSum) Sum(prefix string) int {
    return this.root.searchPrefixSum(prefix)
}

func (t *TrieSumNode) insertKey(key string, delt int) {
	cur, n := t, len(key)
	for i := 0; i < n; i++ {
		idx := key[i] - 'a'
		if cur.child[idx] == nil {
			cur.child[idx] = newTrieSumNode()
		}
		cur.child[idx].sum += delt
		cur = cur.child[idx]
	}
}

func (t *TrieSumNode) searchPrefixSum(s string) int {
	i, n := 0, len(s)
	cur := t
	for i = 0; i < n; i++ {
		idx := s[i] - 'a'
		if cur.child[idx] != nil {
			cur = cur.child[idx]
		} else {
			break
		}
	}
	if i < n || cur == nil {
		return 0
	}
	return cur.sum
}


/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
```

