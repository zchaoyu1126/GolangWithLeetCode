package leetcode

import (
	"fmt"
	"programs/internal/algorithmingo/algorithm"
	"programs/kit/utils"
	"sort"
)

type Node559 struct {
	Val      int
	Children []*Node559
}

// leetcode559
func MaxDepth559(root *Node559) int {
	if root == nil {
		return 0
	}
	res := 0
	for i := 0; i < len(root.Children); i++ {
		tmp := MaxDepth559(root.Children[i])
		if tmp > res {
			res = tmp
		}
	}
	return res + 1
}

// leetcode563
func FindTilt(root *algorithm.TreeNode) int {
	var traverse func(root *algorithm.TreeNode, sum *int) int
	sum := 0
	traverse = func(root *algorithm.TreeNode, sum *int) int {
		if root == nil {
			return 0
		}
		lValue := traverse(root.Left, sum)
		rValue := traverse(root.Right, sum)
		(*sum) += abs_int(rValue - lValue)
		return lValue + rValue + root.Val
	}
	traverse(root, &sum)
	return sum
}

// leetcode572
func IsSubtree(root *algorithm.TreeNode, subRoot *algorithm.TreeNode) bool {
	var isSame func(p, q *algorithm.TreeNode) bool
	isSame = func(p, q *algorithm.TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		l := isSame(p.Left, q.Left)
		r := isSame(p.Right, q.Right)
		return p.Val == q.Val && l && r
	}
	cur := root
	stack := []*algorithm.TreeNode{}
	for cur != nil || len(stack) != 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			top := stack[len(stack)-1]
			if isSame(top, subRoot) {
				return true
			}
			stack = stack[:len(stack)-1]
			cur = top.Right
		}
	}
	return false
}

func IsSubtree2(root *algorithm.TreeNode, subRoot *algorithm.TreeNode) bool {
	var isSame func(p, q *algorithm.TreeNode) bool
	isSame = func(p, q *algorithm.TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		l := isSame(p.Left, q.Left)
		r := isSame(p.Right, q.Right)
		return p.Val == q.Val && l && r
	}
	var dfs func(p *algorithm.TreeNode) bool
	dfs = func(p *algorithm.TreeNode) bool {
		if isSame(p, subRoot) {
			return true
		}
		return dfs(p.Left) || dfs(p.Right)
	}

	return dfs(root)
}

// leetcode573
func Fib(n int) int {
	preOne, preTwo := 0, 1
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	res := 0
	for n-2 >= 0 {
		n--
		res = preOne + preTwo
		preOne = preTwo
		preTwo = res
	}
	return res
}

// leetcode575
func DistributeCandies(candyType []int) int {
	mp := make(map[int]int)
	for i := 0; i < len(candyType); i++ {
		mp[candyType[i]] = i
	}
	res := len(mp)
	if res > len(candyType)/2 {
		return len(candyType) / 2
	}
	return res
}

// leetcode583
// 求最长公共子序列
func MinDistance1(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = utils.MaxNum(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return n + m - 2*dp[n][m]
}

func MinDistance2(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	// dp[i][j] 表示 word[:i], word[:j]相同时需要删除多少
	// 初始化条件dp[0][0] = 0
	for i := 1; i <= n; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= m; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = utils.MinNum(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[n][m]
}

// leetcode594
func FindLHS(nums []int) int {
	sort.Ints(nums)
	cnt1, cnt2, res := 0, 0, 0
	start, end := nums[0], nums[0]-1
	for i := 0; i < len(nums); i++ {
		if nums[i] == start {
			cnt1++
			//fmt.Println(start, end, nums[i], cnt1, cnt2, res, 1)
		} else if nums[i] == end {
			cnt2++
			//fmt.Println(start, end, nums[i], cnt1, cnt2, res, 2)
		} else if nums[i] == start+1 {
			end = nums[i]
			cnt2++
			//fmt.Println(start, end, nums[i], cnt1, cnt2, res, 3)
		} else if nums[i] == end+1 {
			if cnt1+cnt2 > res {
				res = cnt1 + cnt2
			}
			cnt1 = cnt2
			cnt2 = 1
			start = end
			end = nums[i]
			//fmt.Println(start, end, nums[i], cnt1, cnt2, res, 4)
		} else if nums[i] > end+1 {
			if cnt1+cnt2 > res {
				res = cnt1 + cnt2
			}
			cnt1 = 1
			cnt2 = 0
			start = nums[i]
			end = nums[i] - 1
			//fmt.Println(start, end, nums[i], cnt1, cnt2, res, 5)
		}
	}
	fmt.Println(start, end, res)
	if start == end-1 {
		if cnt1+cnt2 > res {
			res = cnt1 + cnt2
		}
	}
	fmt.Println(res)
	return res
}
