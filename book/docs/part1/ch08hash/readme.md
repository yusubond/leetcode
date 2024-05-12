### HashTable

哈希表是一种数据结构，使用哈希函数组织数据，支持快速插入和搜索。



#### 哈希表原理

哈希表的关键思想是使用**哈希函数**将**键映射到存储桶**。具体来说就是：

1.当我们插入一个新的键时，哈希函数将决定该键应该分配到哪个桶中，并将该键存储在相应的桶中；

2.当我们想要搜索一个键时，哈希表将使用同样的哈希函数来查找相应的桶，并只在特定的桶中进行搜索。

#### 哈希集合的设计

```go
const MAX_LEN = 10000
type MyHashSet struct {
    set [MAX_LEN][]int
}
/** Initialize your data structure here. */
func Constructor() MyHashSet {
    return MyHashSet{
        set: [MAX_LEN][]int{},
    }
}
func (this *MyHashSet) getIndex(key int) int {
    return key % MAX_LEN  // 这就是关键的哈希函数
}
func (this *MyHashSet) getPos(key, index int) int {
    for i := 0; i < len(this.set[index]); i++ {
        if this.set[index][i] == key {
            return i
        }
    }
    return -1
}

func (this *MyHashSet) Add(key int)  {
    index := this.getIndex(key)
    pos := this.getPos(key, index)
    if pos < 0 {
        this.set[index] = append(this.set[index], key)
    }
}

func (this *MyHashSet) Remove(key int)  {
    index := this.getIndex(key)
    pos := this.getPos(key, index)
    if pos >= 0 {
        for i := 0; i < len(this.set[index]); i++ {
            if this.set[index][i] == key {
                // 优化，利用交换的思想
                ns := this.set[index][:i]
                ns = append(ns, this.set[index][i+1:]...)
                this.set[index] = ns
                return
            }
        }
    }
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
    index := this.getIndex(key)
    pos := this.getPos(key, index)
    return pos >= 0
}
```

remove元素中可以使用交换的思想，代码如下

```go
for i := 0; i < len(this.set[index]); i++ {
  if this.set[index][i] == key {
    this.set[index][i] = this.set[index][len(this.set[index]) - 1]
    break
  }
}
this.set[index] = this.set[index][:n-1]
```



#### 哈希映射的设计

```go
const MAX_LEN = 1000000
type MyHashMap struct {
    hashmap [MAX_LEN]map[int]int
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
    return MyHashMap{
        hashmap: [MAX_LEN]map[int]int{},
    }
}

func (this *MyHashMap) getIndex(key int) int {
    return key % MAX_LEN  // 关键哈希函数
}
func (this *MyHashMap) getPos(key, index int) int {
    for k := range this.hashmap[index] {
        if k == key {
            return k
        }
    }
    return -1
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
    index := this.getIndex(key)
    pos := this.getPos(key, index)
    if pos < 0 {
        if len(this.hashmap[index]) == 0 {
            this.hashmap[index] = make(map[int]int)
        }
        this.hashmap[index][key] = value
    } else {
        this.hashmap[index][pos] = value
    }
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
    index := this.getIndex(key)
    pos := this.getPos(key, index)
    if pos >= 0 {
        return this.hashmap[index][pos]
    }
    return -1
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
    index := this.getIndex(key)
    pos := this.getPos(key, index)
    if pos >= 0 {
        nm := make(map[int]int, len(this.hashmap[index]) - 1)
        for k, v := range this.hashmap[index] {
            if k != key {
                nm[k] = v
            }
        }
        this.hashmap[index] = nm
    }
}
```

### 相关题目

- 1 两数之和
- 202 快乐数
- 36 有效的数独
- 170 两数之和III - 数据结构设计
- 349 两个数组的交集
- 652 寻找重复的树【中等】
- 49 字母异位词分组【M】
- 四数相加II
- 常数时间插入，删除和获取
- 146 LRU缓存机制【M】

#### 1 两数之和

思路分析：

利用哈希映射原理。先将数组中所有的元素存入map中(相同的元素值，取index较大的一个，后序判断需要)，然后遍历数组，并判断target-value是否在map中。如果map存在，并且index不一样则找到结果。

```go
// date
func twoSum(nums []int, target int) []int {
  res := make([]int, 2)
  m := make(m[int]int, len(nums))
  for index, v := range nums {
    m[v] = index
  }
  for i, v := range nums {
    if j, ok := m[target-v]; ok && i != j {
       res[0], res[1] = i, j
       break
    }
  }
  return res
}
// date 2019/12/30
// 上一版的优化，上一版中需要遍历数组两次，而实际上可以只遍历一次，如下
func twoSum(nums []int, target int) []int {
  res := make([]int, 2)
  m := make(map[int]int)
  for i, v := range nums {
    if j, ok := m[target-v]; ok {
      res[0], res[1] = j, i
      break
    }
    m[v] = i
  }
  return res
}
```

#### 202 快乐数

思路分析

算法1：利用快慢指针的思想判断是否有循环。

```go
func isHappy(n int) bool {
  slow := bitSquareSum(n)
  fast := bitSquareSum(bitSquareSum(n))
  for slow != fast {
    slow = bitSquareSum(slow)
    fast = bitSquareSum(bitSquareSum(fast))
  }
  return slow == 1
}
func bitSquareSum(n int) int {
  sum := 0
  for n > 0 {
    bit := n % 10
    sum += bit * bit
    n /= 10
  }
  return sum
}
```

算法2：利用set，判断之前出现的数是否会再次出现，如果出现，则不是快乐数

```go
// date 2019/12/30
// 算法：set检查是否重复
func isHappy(n int) bool {
  sum, m := 0, make(map[int]struct{})
  for n != 1 {
    for n > 0 {
      sum += (n%10) * (n%10)
      n /= 10
    }
    if _, ok := m[sum]; ok {
      return sum == 1
    }
    m[sum] = struct{}{}
    n, sum = sum, 
  }
  return true
}
```





#### 36 有效的数据 【M】

思路分析

如何获得3x3的子数独，boxes = (rows / 3) * 3 + (columns / 3)。然后保证行，列，子数独每个元素只能出现一次。

```go
func isValidSudoku(board [][]byte) bool {
  rows := [9]map[byte]int{}
  cols := [9]map[byte]int{}
  boxes := [9]map[byte]int{}
  for i := 0; i < 9; i++ {
    rows[i] = make(map[byte]int)
    cols[i] = make(map[byte]int)
    boxes[i] = make(map[byte]int)
  }
  for i := 0; i < 9; i++ {
    for j := 0; j < 9; j++ {
      k := board[i][j]
      if k == byte('.') {continue}
      if _, ok := rows[i][k]; ok {
        return false
      } else {rows[i][k] = 1}
      
      if _, ok := cols[j][k]; ok {
        return false
      } else {cols[j][k] = 1}
      
      b := (i/3)*3 + j/3
      if _, ok := boxes[b][k]; ok {
        return false
      } else {boxes[b][k] = 1}
    }
  }
  return true
}
```



#### 170 两数之和III - 数据结构设计

思路分析:



```go
type TwoSum struct {
    set map[int]int
}

func Constructor() TwoSum {
    return TwoSum{
        set: make(map[int]int),
    }
}

func (this *TwoSum) Add(number int)  {
    if _, ok := this.set[number]; ok {
        this.set[number]++
    } else {
        this.set[number] = 1
    }
}

func (this *TwoSum) Find(value int) bool {
    for k, _ := range this.set {
      if c, ok := this.set[value - k]; ok && (k == value - k && c > 1 || k != value - k) {
          return true
        }
    }
    return false
}
```



#### 349 两个数组的交集

思路分析：两个数组中均可能存在重复元素，所以在第二个数组中找到后需要删除key。

```go
func intersection(nums1, nums2 []int) []int {
  res := make([]int, 0)
  m := make(map[int]struct{})
  for _, v := range nums1 {
    if _, ok := m[v]; !ok {
      m[v] = struct{}{}
    }
  }
  for _, v := range nums2 {
    if _, ok := m[v]; ok {
      res = append(res, v)
      delete(m, v)
    }
  }
  return res
}
```

#### 350 两个数组的交集 II

题目：

![截屏2019-12-0722.20.47](/Users/subond/Documents/MyLife/leetcode/截屏2019-12-0722.20.47.png)

```go
func intersect(nums1, nums2 []int) []int {
  res := make([]int, 0)
  m := make(map[int]int)
  for _, v := range nums1 {
    if _, ok := m[v]; ok {
      m[v]++
    } else {
      m[v] = 1
    }
  }
  for _, v := range nums2 {
    if c, ok := m[v]; ok && c >= 1 {
      res = append(res, v)
      m[v]--
    }
  }
  return res
}
```

#### 存在重复元素II

```go
func containsNearByDuplicate(nums []int, k int) bool {
  m := make(map[int]int)
  for index, v := range nums {
    if j, ok := m[v]; ok && index - j <= k {
      return true
    }
    m[v] = index
  }
  return false
}
```

#### 652 寻找重复的树

思路分析：

将树的节点进行先序遍历，并存储。【即序列化】，然后判断是否出现过。

```go
func 
```



#### 49 字母异位词分组【M】

思路分析：

将每个单词中，每个字符出现的频率映射到一个map中，然后比较每个单词的map是否一致。

```go
func groupAnagrams(strs []string) [][]string {
    n := len(strs)
    grouped := make(map[string]int, n)
    res := make([][]string, 0)
    strM := make(map[string]map[rune]int)
    for i, str := range strs {
        // 已经归属于某个组，直接跳过
        if _, ok := grouped[str]; ok {continue}
        grouped[str] = i
        gstr := make([]string, 0)
        gstr = append(gstr, str)
        im := strMap(str)
        strM[str] = im
        for j := i + 1; j < n; j++ {
            // 长度不一样，肯定不在一组，跳过
            if len(strs[j]) != len(str) {continue}
            // 两个单词一样，直接加入当前组
            if strs[j] == str {
                grouped[strs[j]] = j
                gstr = append(gstr, strs[j])
                continue
            }
            // 如果被分组过，且是当前元素，跳过
            if v, ok := grouped[strs[j]]; ok && v == j {continue}
            // 如果被分过组，且不是当前元素，可能是重复元素
            jm := strMap(strs[j])
            strM[strs[j]] = jm
            if isSame(im, jm) {
                grouped[strs[j]] = j
                gstr = append(gstr, strs[j])
            }
        }
        res = append(res, gstr)
    }
    return res
}

func isSame(m1, m2 map[rune]int) bool {
    if len(m1) != len(m2) {return false}
    for k, v := range m1 {
        if v2, ok := m2[k]; !ok || v2 != v {
            return false
        }
    }
    return true
}

func strMap(s string) map[rune]int {
    res := make(map[rune]int)
    for _, v := range s {
        if _, ok := res[v]; ok {
            res[v]++
        } else {
            res[v] = 1
        }
    }
    return res
}
```

思路分析

算法优化:将每个单词按字符出现的次数形成key，作为映射的基础，同一个key，可以对应多个单词，即分组。【golang数组是可比较类型，可以用作map的key】

例如aab, aba, baa, 这三个单词的key都是a2b1。

```go
func groupAnagrams(strs []string) [][]string {
    m := make(map[[26]int][]string)
    for _, str := range strs {
        k := strArray(str)
        s, ok := m[k]
        if !ok {
            s = make([]string, 0)
        }
        s = append(s, str)
        m[k] = s
    }
    res := make([][]string, 0, len(m))
    for _, v := range m {
        res = append(res, v)
    }
    return res
}

func strArray(s string) [26]int {
    res := [26]int{}
    for _, v := range s {
        res[v - 'a']++
    }
    return res
}
```

#### 249 移位字符串分组

思路分析

异位字符串分组中对字符出现的次数作为key，移位字符串的key可以视为在一个具有26个槽位的哈希环上，每个字符的相对位置是一样的。

```go
func groupStrings(strs []string) [][]string {
  m := make(map[[26]int][]string)
  for _, str := range strs {
    k := strArray(str)
    s, ok := m[k]
    if !ok {
      s = make([]string, 0)
    }
    s = append(s, str)
    m[k] = s
  }
  res := make([][]string, 0, len(m))
  for _, v := range m {
    res = append(res, v)
  }
  return res
}

func strArray(s string) [26]int {
  res := [26]int{}
  if len(s) == 0 {return res}
  for _, v := range s {
    res[(v - s[0] + 26) % 26]++
  }
  return res
}
```

#### 四数相加II

思路分析

正常情况下需要四层循环，比较耗时；可以通过map优化，即先将两个数组合并，记录其出现的值及可能的组合次数；然后用map查看结果，并将组合数相乘。

```go
func fourSumCount(A []int, B []int, C []int, D []int) int {
    res := 0
    abm, cdm := make(map[int]int, len(A)*2), make(map[int]int, len(A)*2)
    for _, a := range A {
        for _, b := range B {
            if _, ok := abm[a+b]; ok {
                abm[a+b]++
            } else {
                abm[a+b] = 1
            }
        }
    }
    for _, c := range C {
        for _, d := range D {
            if _, ok := cdm[c+d]; ok {
                cdm[c+d]++
            } else {
                cdm[c+d] = 1
            }
        }
    }
    for k, v1 := range abm {
        if v2, ok := cdm[0-k]; ok {
            res += v1 * v2
        }
    }
    return res
}
```

#### 常数时间插入、删除和获取

```go
type RandomizedSet struct {
    set  map[int]int  // 存放元素及其对应的下标
    data []int        // 存放元素
}

func Constructor() RandomizedSet {
    return RandomizedSet{
        set: make(map[int]int),
        data: make([]int, 0),
    }
}
func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.set[val]; ok {
        return false
    }
    this.set[val] = len(this.data)
    this.data = append(this.data, val)
    return true
}

func (this *RandomizedSet) Remove(val int) bool {
    if _, ok := this.set[val]; !ok {
        return false
    }
    // 将欲删除的元素从data中删除，利用和最后一个元素交换的方式删除
    // 将set中data中最后一个元素的下标更新为想要删除元素的下标
    this.set[this.data[len(this.data) - 1]] = this.set[val]
    // data中要删除的元素的位置上放上最后一个元素
    this.data[this.set[val]] = this.data[len(this.data) - 1]
    // 删除set
    delete(this.set, val)
    // 删除data最后一个元素
    this.data = this.data[:len(this.data) - 1]
    return true
}

func (this *RandomizedSet) GetRandom() int {
    return this.data[rand.Intn(len(this.data))]
}
```

#### 146 LRU缓存机制

题目解析：

LRU是Least Recently Used，即最近最少使用。

算法：哈希表+双向链表

```go
// data 2020/01/29
// 算法 哈希表+双向链表，用头节点保存最近操作过的节点，尾巴节点即是最近最少使用的节点。
type LRUCache struct {
    cache map[int]*DLinkNode
    size int
    head, tail *DLinkNode
}
type (
    DLinkNode struct {
        key, value int
        pre, next *DLinkNode
    }
)
func newLinkNode(k, v int) *DLinkNode {
    return &DLinkNode{
        key: k,
        value: v,
    }
}
// alway add node next to head
func (this *LRUCache) addNode(n *DLinkNode) {
    n.pre = this.head
    n.next = this.head.next
    this.head.next.pre = n
    this.head.next = n
}
// delete node from the link list
func (this *LRUCache) removeNode(n *DLinkNode) {
    p1, p2 := n.pre, n.next
    p1.next = p2
    p2.pre = p1
}
// move certain node to the head
func (this *LRUCache) move2Head(n *DLinkNode) {
    this.removeNode(n)
    this.addNode(n)
}
// delete tail node
func (this *LRUCache) removeTail() int {
    t := this.tail.pre
    this.removeNode(t)
    return t.key
}

func Constructor(capacity int) LRUCache {
    l :=  LRUCache{
        cache: make(map[int]*DLinkNode),
        size: capacity,
        head: newLinkNode(0,0),
        tail: newLinkNode(0,0),
    }
    l.head.next = l.tail
    l.tail.pre = l.head
    return l
}


func (this *LRUCache) Get(key int) int {
    v, ok := this.cache[key]
    if !ok { return -1 }
    this.move2Head(v)  // head总是保存最新的记录
    return v.value
}


func (this *LRUCache) Put(key int, value int)  {
    v, ok := this.cache[key]
    if ok {
        v.value = value
        this.move2Head(v)  // head总是保存最新的记录
        return
    }
    v = newLinkNode(key, value)
    this.cache[key] = v
    this.addNode(v)
    if len(this.cache) > this.size {
        tkey := this.removeTail()
        delete(this.cache, tkey)
    }
}
```

