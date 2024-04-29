package template

const ALPHABET_SIZE = 26

// TrieNode define
type TrieNode struct {
	child [ALPHABET_SIZE]*TrieNode
	// isEndOfWord is true if the node represents
	// end of a word
	isEndOfWord bool
}

// NewTrieNode define
func NewTrieNode() *TrieNode {
	node := &TrieNode{
		child:       [26]*TrieNode{},
		isEndOfWord: false,
	}
	return node
}

// InsertKey define
// If not present, inserts key into trie
// If the key is prefix of trie node, just
// marks leaf node
func (t *TrieNode) InsertKey(key string) {
	n := len(key)
	if n <= 0 {
		return
	}
	cur := t
	for i := 0; i < n; i++ {
		idx := key[i] - 'a'
		if cur.child[idx] == nil {
			cur.child[idx] = NewTrieNode()
		}
		cur = cur.child[idx]
	}
	cur.isEndOfWord = true
}

// SearchKey define
// Returns true if key presents in trie, else
// false
func (t *TrieNode) SearchKey(key string) bool {
	n := len(key)
	if n == 0 {
		return false
	}
	cur := t
	for i := 0; i < n; i++ {
		idx := key[i] - 'a'
		if cur.child[idx] == nil {
			return false
		}
		cur = cur.child[idx]
	}
	return cur.isEndOfWord
}

// DeleteKey define
func (t *TrieNode) DeleteKey(key string) {
	n := len(key)
	if n == 0 {
		return
	}
	// dfs
	var dfs func(root *TrieNode, key string, depth int) *TrieNode
	dfs = func(root *TrieNode, key string, depth int) *TrieNode {
		// if tree is empty
		if root == nil {
			return root
		}
		// process the last character of key
		if depth == len(key) {
			// unmark
			// this node is no more end of word after delete of given key
			if root.isEndOfWord {
				root.isEndOfWord = false
			}
			// if the given key is not prefix of any other word
			if root.isEmpty() {
				root = nil
			}
			return root
		}
		// if not last character, recur for the child
		idx := key[depth] - 'a'
		// if not find key, just return
		if root.child[idx] == nil {
			return root
		}
		// recur the child
		root.child[idx] = dfs(root.child[idx], key, depth+1)
		// if current node is not the prefix of given key, remote it
		if root.isEmpty() && !root.isEndOfWord {
			root = nil
		}
		return root
	}
	dfs(t, key, 0)
}

func (t *TrieNode) isEmpty() bool {
	for _, v := range t.child {
		if v != nil {
			return false
		}
	}
	return true
}
