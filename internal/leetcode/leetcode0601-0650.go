package leetcode

import (
	"programs/internal/algorithmingo/algorithm"
	"programs/kit/utils"
)

// leetcode617
func MergeTrees(root1 *algorithm.TreeNode, root2 *algorithm.TreeNode) *algorithm.TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil || root2 == nil {
		if root1 == nil {
			return root2
		}
		return root1
	}
	root1.Val += root2.Val
	root1.Left = MergeTrees(root1.Left, root2.Left)
	root1.Right = MergeTrees(root1.Right, root2.Right)
	return root1
}

// leetcode629
func KInversePairs(n, k int) int {
	const mod int = 1e9 + 7
	f := [2][]int{make([]int, k+1), make([]int, k+1)}
	f[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= k; j++ {
			cur := i & 1
			prev := cur ^ 1
			f[cur][j] = 0
			if j > 0 {
				f[cur][j] = f[cur][j-1]
			}
			if j >= i {
				f[cur][j] -= f[prev][j-i]
			}
			f[cur][j] += f[prev][j]
			if f[cur][j] >= mod {
				f[cur][j] -= mod
			} else if f[cur][j] < 0 {
				f[cur][j] += mod
			}
		}
	}
	return f[n&1][k]
}

// leetcode630
func ScheduleCourse(courses [][]int) int {
	cnt := 0
	for i := 0; i < len(courses); i++ {

	}
	return cnt
}

// leetcode637
func AverageOfLevels(root *algorithm.TreeNode) []float64 {
	res := []float64{}
	queue := []*algorithm.TreeNode{root}
	for len(queue) != 0 {
		size := len(queue)
		sum := 0
		for i := 0; i < size; i++ {
			front := queue[0]
			queue = queue[1:]
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
			sum += front.Val
		}
		res = append(res, float64(sum)/float64(size))
	}
	return res
}

// leetcode638
func ShoppingOffers(price []int, special [][]int, needs []int) int {
	n := len(price)

	// 过滤不需要计算的大礼包，只保留需要计算的大礼包
	filterSpecial := [][]int{}
	for _, s := range special {
		totalCount, totalPrice := 0, 0
		for i, c := range s[:n] {
			totalCount += c
			totalPrice += c * price[i]
		}
		if totalCount > 0 && totalPrice > s[n] {
			filterSpecial = append(filterSpecial, s)
		}
	}

	// 记忆化搜索计算满足购物清单所需花费的最低价格
	dp := map[string]int{}

	var dfs func([]byte) int
	dfs = func(curNeeds []byte) (minPrice int) {
		if res, has := dp[string(curNeeds)]; has {
			return res
		}
		for i, p := range price {
			minPrice += int(curNeeds[i]) * p // 不购买任何大礼包，原价购买购物清单中的所有物品
		}
		nextNeeds := make([]byte, n)
	outer:
		for _, s := range filterSpecial {
			for i, need := range curNeeds {
				if need < byte(s[i]) { // 不能购买超出购物清单指定数量的物品
					continue outer
				}
				nextNeeds[i] = need - byte(s[i])
			}
			minPrice = utils.MinNum(minPrice, dfs(nextNeeds)+s[n])
		}
		dp[string(curNeeds)] = minPrice
		return
	}

	curNeeds := make([]byte, n)
	for i, need := range needs {
		curNeeds[i] = byte(need)
	}
	return dfs(curNeeds)
}

func CountSubstrings1dp(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	res := 0
	// s[i]==s[j]
	// dp[i+1][j-1] ij相差大于1的时候
	// true      i+1=j
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					res++
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					res++
					dp[i][j] = true
				}
			}
		}
	}
	return res
}

func CountSubstrings2(s string) int {
	// 枚举回文中心起点
	// a b c 回文中心 n+n-1 = 2n-1
	n := len(s)
	res := 0
	for i := 0; i < 2*n-1; i++ {
		var l, r int
		if i%2 == 0 {
			l, r = i/2-1, i/2+1
			res++

		} else if i%2 == 1 {
			l, r = i/2, i/2+1
		}
		for l >= 0 && r < n && s[l] == s[r] {
			res++
			l--
			r++
		}
	}
	return res
}
