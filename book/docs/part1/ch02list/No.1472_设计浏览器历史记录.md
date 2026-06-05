## 1472 设计浏览器历史记录-中等

题目：

你有一个只支持单个标签页的 **浏览器** ，最开始你浏览的网页是 `homepage` ，你可以访问其他的网站 `url` ，也可以在浏览历史中后退 `steps` 步或前进 `steps` 步。

请你实现 `BrowserHistory` 类：

- `BrowserHistory(string homepage)` ，用 `homepage` 初始化浏览器类。
- `void visit(string url)` 从当前页跳转访问 `url` 对应的页面 。执行此操作会把浏览历史前进的记录全部删除。
- `string back(int steps)` 在浏览历史中后退 `steps` 步。如果你只能在浏览历史中后退至多 `x` 步且 `steps > x` ，那么你只后退 `x` 步。请返回后退 **至多** `steps` 步以后的 `url` 。
- `string forward(int steps)` 在浏览历史中前进 `steps` 步。如果你只能在浏览历史中前进至多 `x` 步且 `steps > x` ，那么你只前进 `x` 步。请返回前进 **至多** `steps`步以后的 `url` 。



分析：

这个问题重点在于理解 visit 操作，只要发生 visit 操作，**浏览历史中的前进记录全部删除**。

所以，需要一个当前指针 cur 指向当前浏览的页面。

然后，用 head, tail 辅助做头尾的哑结点。

设计 双向链表，方便从当前页面 cur 根据 steps 步数前进或后退。

前进、后退遍历的时候，要校验当前指针是否为头尾指针，以免越界。

```go 
// date 2023/10/17
type BrowserHistory struct {
    cur *MyNode
    head, tail *MyNode
}

type MyNode struct {
    key string
    prev, next *MyNode
}


func Constructor(homepage string) BrowserHistory {
    ss := BrowserHistory{
        head: &MyNode{},
        tail: &MyNode{},
    }
    node := &MyNode{key: homepage}
    node.next = ss.tail
    node.prev = ss.head
    ss.tail.prev = node
    ss.head.next = node
  
    ss.cur = node
    return ss
}


func (this *BrowserHistory) Visit(url string)  {
    node := &MyNode{key: url}
    this.cur.next = node
    node.prev = this.cur
  
    node.next = this.tail
    this.tail.prev = node
  
  	this.cur = node
}


func (this *BrowserHistory) Back(steps int) string {
    pre := this.cur
    for steps > 0 {
        if pre.prev == this.head {
            break
        }
        pre = pre.prev
        steps--
    }
    this.cur = pre
    return this.cur.key
}


func (this *BrowserHistory) Forward(steps int) string {
    pre := this.cur
    for steps > 0 {
        if pre.next == this.tail {
            break
        }
        pre = pre.next
        steps--
    }
    this.cur = pre
    return this.cur.key
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
```

