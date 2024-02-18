package algorithm

import "github.com/zchaoyu1126/coding-practice/utils"

// ********************************************************* \\
// 平衡二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GetHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := GetHeight(root.Left)
	r := GetHeight(root.Right)
	if l == -1 || r == -1 {
		return -1
	}
	if l-r >= 2 || r-l >= 2 {
		return -1
	}
	return utils.MaxNum(l, r) + 1
}

func IsBalanced(root *TreeNode) bool {
	return GetHeight(root) != -1
}

// ********************************************************* \\
// 字典树
type TrieNode struct {
	Num       int
	Son       map[byte]*TrieNode
	IsEnd     bool
	Val       byte
	WordValue int
}

type Trie struct {
	root *TrieNode
}

func NewRuneTrie() *Trie {
	return &Trie{root: &TrieNode{Son: make(map[byte]*TrieNode)}}
}

func (t *Trie) InsertWithValue(str string, value int) bool {
	node := t.root
	if str == "" {
		return false
	}
	bytes := []byte(str)
	for _, val := range bytes {
		pos := val - 'a'
		if node.Son[pos] == nil {
			node.Son[pos] = &TrieNode{Son: make(map[byte]*TrieNode)}
			node.Son[pos].Val = val
		} else {
			node.Son[pos].Num++
		}
		node = node.Son[pos]
	}
	node.WordValue = value
	node.IsEnd = true
	return true
}

func (t *Trie) Insert(str string) bool {
	node := t.root
	if str == "" {
		return false
	}
	bytes := []byte(str)
	for _, val := range bytes {
		pos := val - 'a'
		if node.Son[pos] == nil {
			node.Son[pos] = &TrieNode{Son: make(map[byte]*TrieNode)}
			node.Son[pos].Val = val
		} else {
			node.Son[pos].Num++
		}
		node = node.Son[pos]
	}
	node.IsEnd = true
	return true
}

func (t *Trie) SumOfPrefix(prefix string) int {
	node := t.root
	if prefix == "" {
		return -1
	}
	bytes := []byte(prefix)
	for _, val := range bytes {
		pos := val - 'a'
		if node.Son[pos] != nil {
			node = node.Son[pos]
		} else {
			return -1
		}
	}

	sum := 0
	var traverse func(n *TrieNode)
	traverse = func(n *TrieNode) {
		if n.IsEnd {
			sum += n.WordValue
		}
		for _, next := range n.Son {
			traverse(next)
		}
	}
	traverse(node)
	return sum
}

func (t *Trie) Find(str string) bool {
	node := t.root
	if str == "" {
		return false
	}
	bytes := []byte(str)
	for _, val := range bytes {
		pos := val - 'a'
		if node.Son[pos] != nil {
			node = node.Son[pos]
		} else {
			return false
		}
	}
	return true
}
