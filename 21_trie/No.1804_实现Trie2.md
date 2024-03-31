## 1804 实现Trie2-中等

题目：

前缀树（**[trie](https://en.wikipedia.org/wiki/Trie)** ，发音为 "try"）是一个树状的数据结构，用于高效地存储和检索一系列字符串的前缀。前缀树有许多应用，如自动补全和拼写检查。

实现前缀树 Trie 类：

- `Trie()` 初始化前缀树对象。
- `void insert(String word)` 将字符串 `word` 插入前缀树中。
- `int countWordsEqualTo(String word)` 返回前缀树中字符串 `word` 的实例个数。
- `int countWordsStartingWith(String prefix)` 返回前缀树中以 `prefix` 为前缀的字符串个数。
- `void erase(String word)` 从前缀树中移除字符串 `word` 。



> **示例 1:**
>
> ```
> 输入
> ["Trie", "insert", "insert", "countWordsEqualTo", "countWordsStartingWith", "erase", "countWordsEqualTo", "countWordsStartingWith", "erase", "countWordsStartingWith"]
> [[], ["apple"], ["apple"], ["apple"], ["app"], ["apple"], ["apple"], ["app"], ["apple"], ["app"]]
> 输出
> [null, null, null, 2, 2, null, 1, 1, null, 0]
> 
> 解释
> Trie trie = new Trie();
> trie.insert("apple");               // 插入 "apple"。
> trie.insert("apple");               // 插入另一个 "apple"。
> trie.countWordsEqualTo("apple");    // 有两个 "apple" 实例，所以返回 2。
> trie.countWordsStartingWith("app"); // "app" 是 "apple" 的前缀，所以返回 2。
> trie.erase("apple");                // 移除一个 "apple"。
> trie.countWordsEqualTo("apple");    // 现在只有一个 "apple" 实例，所以返回 1。
> trie.countWordsStartingWith("app"); // 返回 1
> trie.erase("apple");                // 移除 "apple"。现在前缀树是空的。
> trie.countWordsStartingWith("app"); // 返回 0
> ```



**解题思路**

这道题也是字典树的标准解法，既要统计 key 的次数，又要统计前缀的总和。

```go
// date 2024/03/31
type Trie struct {
    child [26]*Trie
    isWord bool
    cnt int  // word count
    sum int   // prefix sum
}


func Constructor() Trie {
    t := Trie{
        child: [26]*Trie{},
        isWord: false,
        cnt: 0,
        sum: 0,
    }
    return t
}


func (this *Trie) Insert(word string)  {
    this.insertWithDelta(word, 1)
}


func (this *Trie) CountWordsEqualTo(word string) int {
    cur := this
    n := len(word)
    i := 0
    for i = 0; i < n; i++ {
        idx := word[i] - 'a'
        if cur.child[idx] != nil {
            cur = cur.child[idx]
        } else {
            break
        }
    }

    if i < n || cur == nil {
        return 0
    }

    return cur.cnt
}


func (this *Trie) CountWordsStartingWith(prefix string) int {
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
    return cur.sum
}


func (this *Trie) Erase(word string)  {
    this.insertWithDelta(word, -1)
}

func (this *Trie) insertWithDelta(word string, delt int) {
    cur := this
    i, n := 0, len(word)
    for i = 0; i < n; i++ {
        idx := word[i] - 'a'
        if cur.child[idx] == nil {
            cur.child[idx] = &Trie{}
        }
        cur.child[idx].sum += delt
        cur = cur.child[idx]
    }
    if delt > 0 {
        // to add
        cur.isWord = true
        cur.cnt += 1
    } else {
        // delete word
        cur.isWord = false
        cur.cnt -= 1
    }
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.CountWordsEqualTo(word);
 * param_3 := obj.CountWordsStartingWith(prefix);
 * obj.Erase(word);
 */
```

