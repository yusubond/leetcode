## 211 添加与搜索单词-数据结构设计-中等

题目：

请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。

实现词典类 `WordDictionary` ：

- `WordDictionary()` 初始化词典对象
- `void addWord(word)` 将 `word` 添加到数据结构中，之后可以对它进行匹配
- `bool search(word)` 如果数据结构中存在字符串与 `word` 匹配，则返回 `true` ；否则，返回 `false` 。`word` 中可能包含一些 `'.'` ，每个 `.` 都可以表示任何一个字母。



> **示例：**
>
> ```
> 输入：
> ["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
> [[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
> 输出：
> [null,null,null,null,false,true,true,true]
> 
> 解释：
> WordDictionary wordDictionary = new WordDictionary();
> wordDictionary.addWord("bad");
> wordDictionary.addWord("dad");
> wordDictionary.addWord("mad");
> wordDictionary.search("pad"); // 返回 False
> wordDictionary.search("bad"); // 返回 True
> wordDictionary.search(".ad"); // 返回 True
> wordDictionary.search("b.."); // 返回 True
> wordDictionary.search("b.");  // 返回 false, 因为存在 b. 前缀，但不存在长度为2的单词
> ```



**解题思路**

这道题也是标准的字典树。需要注意的是：

- 尽管查找的时候，可用`'.'`进行模糊匹配，但是也要判断当前节点是不是单词节点，如果查找的 key 只是某个单词的前缀，需要返回 false。

```go
// date 2024/01/25
type WordDictionary struct {
    child [26]*WordDictionary
    isEndOfWord bool
}

func NewWordDictionary() *WordDictionary {
    return &WordDictionary{child: [26]*WordDictionary{}}
}

func Constructor() WordDictionary {
    return WordDictionary{child: [26]*WordDictionary{}}
}


func (this *WordDictionary) AddWord(word string)  {
    n := len(word)
    cur := this
    for i := 0; i < n; i++ {
        idx := word[i] - 'a'
        if cur.child[idx] == nil {
            cur.child[idx] = NewWordDictionary()
        }
        cur = cur.child[idx]
    }
    cur.isEndOfWord = true
}


func (this *WordDictionary) Search(word string) bool {
    var searchNode func(root *WordDictionary, word string, depth int) bool
    searchNode = func(root *WordDictionary, word string, depth int) bool {
        if root == nil {
            return false
        }
        // check the last character of word
        if depth == len(word) {
            return root.isEndOfWord
        }
        if word[depth] != '.' {
            idx := word[depth] - 'a'
            if root.child[idx] == nil {
                return false
            }
            return searchNode(root.child[idx], word, depth+1)
        } else {
            // '.'
            // 依次 dfs 搜索，并合并结果
            temp := false
            for i := 0; i < 26; i++ {
                if root.child[i] != nil {
                    r1 := searchNode(root.child[i], word, depth+1)
                    temp = temp || r1
                }
            }
            return temp
        }
        return false
    }


    return searchNode(this, word, 0)
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
```

