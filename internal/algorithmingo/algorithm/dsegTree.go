package algorithm

type dsegTreeNode struct {
	ls, rs  *dsegTreeNode
	l, r, m int
	val     int
	add     int
}

type DsegTree struct {
	root *dsegTreeNode
}

func newDsegTreeNode(l, r int) *dsegTreeNode {
	return &dsegTreeNode{
		l: l,
		r: r,
		m: (l + r) >> 1,
	}
}

func NewDsegTree(l, r int) *DsegTree {
	return &DsegTree{
		root: newDsegTreeNode(l, r),
	}
}

func (s *DsegTree) pushUp(n *dsegTreeNode) {
	n.val = n.ls.val + n.rs.val
}

func (s *DsegTree) pushDown(n *dsegTreeNode) {
	if n.ls == nil {
		n.ls = newDsegTreeNode(n.l, n.m)
	}
	if n.rs == nil {
		n.rs = newDsegTreeNode(n.m+1, n.r)
	}
	if n.add != 0 {
		n.ls.add += n.add
		n.rs.add += n.add
		n.ls.val += n.add * (n.ls.r - n.ls.l + 1)
		n.rs.val += n.add * (n.rs.r - n.rs.l + 1)
		n.add = 0
	}
}

func (s *DsegTree) ModifyInterval(l, r, v int, n *dsegTreeNode) {
	if l <= n.l && n.r <= r {
		n.val += (r - l + 1) * v
		n.add += v
		return
	}
	s.pushDown(n)
	// 全部都在左区间
	if l <= n.m {

		s.ModifyInterval(l, r, v, n.ls)
	}
	// 全部都在右区间
	if r > n.m {
		s.ModifyInterval(l, r, v, n.rs)
	}
	s.pushUp(n)
}

func (s *DsegTree) Query(l, r int, n *dsegTreeNode) int {
	if l > r {
		return 0
	}
	if n.l >= l && n.r <= r {
		return n.val
	}
	s.pushDown(n)

	if l <= n.m {
		return s.Query(l, r, n.ls)
	}
	if r > n.m {
		return s.Query(l, r, n.rs)
	}
	return 0
}
