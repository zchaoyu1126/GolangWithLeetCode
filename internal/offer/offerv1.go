package offer

import (
	"programs/internal/algorithmingo/algorithm"
)

// offerv2 64 脑筋急转弯
// 求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）
func SumNums(n int) int {
	res := 0
	var sum func(int) bool
	sum = func(x int) bool {
		res += x
		return x > 0 && sum(x-1)
	}
	sum(n)
	return res
}

// offerv2 65 不用加减乘除做加法 位运算

// offerv2 66 构建乘积数组 脑筋急转弯
// 没意思，从左从右分别构造两个数组
func ConstructArr(a []int) []int {
	l, r := make([]int, len(a)), make([]int, len(a))
	ans := make([]int, len(a))
	l[0] = 1
	for i := 1; i < len(a); i++ {
		l[i] = l[i-1] * a[i]
	}
	r[len(a)-1] = 1
	for i := len(a) - 2; i >= 0; i-- {
		r[i] = r[i+1] * a[i]
	}
	for i := 0; i < len(a); i++ {
		ans[i] = l[i] * r[i]
	}
	return ans
}

// offerv2 67 atoi转换

// offerv2 68 I 二叉搜索树地最近公共祖先
func LowestCommonAncestor1(root, p, q *algorithm.TreeNode) *algorithm.TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	if p.Val > q.Val {
		p, q = q, p
	}
	if p.Val < root.Val && root.Val < q.Val {
		return root
	} else if p.Val < root.Val && q.Val < root.Val {
		return LowestCommonAncestor1(root.Left, p, q)
	}
	return LowestCommonAncestor1(root.Right, p, q)
}

// offerv2 68 II 二叉树的最近公共祖先
func LowestCommonAncestor2(root, p, q *algorithm.TreeNode) *algorithm.TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	v1 := LowestCommonAncestor2(root.Left, p, q)
	v2 := LowestCommonAncestor2(root.Right, p, q)
	if v1 != nil && v2 != nil {
		return root
	}
	if v1 == nil {
		return v2
	}
	return v1
}
