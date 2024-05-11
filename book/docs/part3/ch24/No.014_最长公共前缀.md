## 14 最长公共前缀-简单

题目：

编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 `""`。



**解题思路**

字典树的词频应用。

具体为，先用输入构造字典树，然后在字典树中查找出现次数等于字符串总数的字符。

```go
// date 2024/01/25
type Trie struct {
	child [26]*Trie
	cnt   int
}

func NewTrie() *Trie {
	return &Trie{
		child: [26]*Trie{},
		cnt:   0,
	}
}

func (this *Trie) Insert(word string) {
	n := len(word)
	cur := this
	for i := 0; i < n; i++ {
		idx := word[i] - 'a'
		if cur.child[idx] == nil {
			cur.child[idx] = &Trie{child: [26]*Trie{}, cnt: 0}
		}
		cur.child[idx].cnt++
		cur = cur.child[idx]
	}
}

func (this *Trie) findLongestCommonPrefix(totalTime int) string {
	ans := ""
	idx := -1
	cur := this
	for cur != nil {
		idx = -1
		for i := 0; i < 26; i++ {
			node := cur.child[i]
			if node != nil && node.cnt == totalTime {
				idx = i
			}
		}
		if idx == -1 {
			break
		}
		ans += string(idx + 'a')
		cur = cur.child[idx]
	}
	return ans
}

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}
	tie := NewTrie()
	for _, v := range strs {
		tie.Insert(v)
	}
	res := tie.findLongestCommonPrefix(n)
	return res
}
```

