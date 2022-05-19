package algorithm

import "programs/kit/utils"

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

// ********************************************************* \\
// 线段树
type SegTree struct {
	nums   []int
	tree   []int
	length int
}

func NewSegTree(nums []int) *SegTree {
	n := len(nums)
	// 开四倍的空间 n << 2
	segTree := &SegTree{nums: nums, tree: make([]int, n<<2)}
	segTree.buildTree(0, len(nums)-1, 0)
	segTree.length = n
	return segTree
}

func (s *SegTree) buildTree(l, r int, pos int) {
	if l == r {
		s.tree[pos] = s.nums[l]
		return
	}
	m := (l + r) >> 1
	s.buildTree(l, m, pos<<1+1)
	s.buildTree(m+1, r, pos<<1+2)
	s.pushUp(pos)
}

func (s *SegTree) pushUp(pos int) {
	s.tree[pos] = s.tree[pos<<1+1] + s.tree[pos<<1+2]
}

func (s *SegTree) update(index, val int, l, r int, pos int) {
	if l == r {
		s.tree[pos] = val
		return
	}
	m := (l + r) >> 1
	if index <= m {
		s.update(index, val, l, m, pos<<1+1)
	} else {
		s.update(index, val, m+1, r, pos<<1+2)
	}
	s.pushUp(pos)
}

func (s *SegTree) query(L, R int, l, r int, pos int) int {
	if L <= l && r <= R {
		return s.tree[pos]
	}
	m := (l + r) >> 1
	ans := 0
	if L <= m {
		// 左子区间有重合
		ans += s.query(L, R, l, m, pos<<1+1)
	}
	if R > m {
		// 右子区间有重合
		ans += s.query(L, R, m+1, r, pos<<1+2)
	}
	return ans
}

func (s *SegTree) GetTree() []int {
	return s.tree
}

func (s *SegTree) UpdateValByIndex(index int, val int) {
	s.update(index, val, 0, s.length-1, 0)
}

func (s *SegTree) SumRangeBetweenLR(left int, right int) int {
	return s.query(left, right, 0, s.length-1, 0)
}
