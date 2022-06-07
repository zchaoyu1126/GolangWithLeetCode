package algorithm

type segTreeNode struct {
	l, r int
	sum  int
	lazy int
}
type SegTree struct {
	tree []segTreeNode
}

func NewSegTree(nums []int) *SegTree {
	n := len(nums)
	// tres需要四倍的空间 n << 2
	segTree := &SegTree{tree: make([]segTreeNode, n<<2)}
	segTree.buildTree(nums, 0, n-1, 0)
	return segTree
}

func (s *SegTree) buildTree(nums []int, l, r int, pos int) {
	s.tree[pos].l = l
	s.tree[pos].r = r
	if l == r {
		s.tree[pos].sum = nums[l]
		return
	}
	m := (l + r) >> 1
	s.buildTree(nums, l, m, pos<<1+1)
	s.buildTree(nums, m+1, r, pos<<1+2)
	s.pushUp(pos)
}

func (s *SegTree) pushUp(pos int) {
	s.tree[pos].sum = s.tree[pos<<1+1].sum + s.tree[pos<<1+2].sum
}

// 线段树单点修改
func (s *SegTree) Modify(idx, val, pos int) {
	if s.tree[pos].l == s.tree[pos].r {
		s.tree[pos].sum = val
		return
	}
	m := (s.tree[pos].l + s.tree[pos].r) >> 1
	if idx <= m {
		s.Modify(idx, val, pos<<1+1)
	} else {
		s.Modify(idx, val, pos<<1+2)
	}
	s.pushUp(pos)
}

// 线段树区间修改
func (s *SegTree) ModifyInterval(l, r, val, pos int) {
	if s.tree[pos].l == l && s.tree[pos].r == r {
		s.tree[pos].sum += (r - l + 1) * val
		s.tree[pos].lazy += val
		return
	}
	m := (s.tree[pos].l + s.tree[pos].r) / 2
	if r <= m {
		// 全部在左子区间
		s.ModifyInterval(l, r, val, pos<<1+1)
	} else if l > m {
		// 全部在右子区间
		s.ModifyInterval(l, r, val, pos<<1+2)
	} else {
		s.ModifyInterval(l, m, val, pos<<1+1)
		s.ModifyInterval(m+1, r, val, pos<<1+2)
	}
	s.pushUp(pos)
}

func (s *SegTree) pushDown(pos int) {
	if s.tree[pos].l == s.tree[pos].r {
		s.tree[pos].lazy = 0
		return
	}
	s.tree[pos<<1+1].lazy += s.tree[pos].lazy
	s.tree[pos<<1+2].lazy += s.tree[pos].lazy
	val := s.tree[pos].lazy
	s.tree[pos<<1+1].sum += (s.tree[pos<<1+1].r - s.tree[pos<<1+2].l + 1) * val
	s.tree[pos<<1+2].sum += (s.tree[pos<<1+2].r - s.tree[pos<<1+2].l + 1) * val
	s.tree[pos].lazy = 0
}

func (s *SegTree) Query(l, r int, pos int) int {
	if s.tree[pos].lazy != 0 {
		s.pushDown(pos)
	}
	if s.tree[pos].l == l && s.tree[pos].r == r {
		return s.tree[pos].sum
	}
	m := (s.tree[pos].l + s.tree[pos].r) / 2
	if r <= m {
		// 说明都在左子区间
		return s.Query(l, r, pos<<1+1)
	} else if l > m {
		return s.Query(l, r, pos<<1+2)
	}
	lsum := s.Query(l, m, pos<<1+1)
	rsum := s.Query(m+1, r, pos<<1+2)
	return lsum + rsum
}
