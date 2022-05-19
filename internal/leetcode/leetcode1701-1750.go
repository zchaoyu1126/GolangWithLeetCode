package leetcode

import (
	"programs/kit/utils"
	"sort"
)

// leetcode1711
func CountPairs(deliciousness []int) int {
	maxVal := 0
	for i := 0; i < len(deliciousness); i++ {
		if maxVal < deliciousness[i] {
			maxVal = deliciousness[i]
		}
	}
	maxVal *= 2
	res := 0
	mp := make(map[int]int)
	for i := 0; i < len(deliciousness); i++ {
		for sum := 1; sum <= maxVal; sum <<= 1 {
			res += mp[sum-deliciousness[i]]
		}
		mp[deliciousness[i]]++
	}
	return res % (1e9 + 7)
}

// leetcode1713
// LCS 最长公共子序列
func MinOperations(target []int, arr []int) int {
	dp := make([]int, len(arr)+1)

	for i := 1; i <= len(target); i++ {
		pre := 0
		for j := 1; j <= len(arr); j++ {
			tmp := dp[j]
			if target[i-1] == arr[j-1] {
				dp[j] = pre + 1
			} else {
				dp[j] = utils.MaxNum(dp[j], dp[j-1])
			}
			pre = tmp
		}
	}
	return len(target) - dp[len(arr)]
}

func LCS_MemoryON2(target []int, arr []int) int {
	n, m := len(target), len(arr)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if target[i-1] == arr[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = utils.MaxNum(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}

func LCS_MemoryON(target []int, arr []int) int {
	n, m := len(target), len(arr)
	dp := make([]int, m+1)
	// pre用于保存dp[i-1][j-1]的值
	// dp[i-1][j]就是上一行的值  dp[i][j-1]已经计算好了
	for i := 1; i <= n; i++ {
		pre := 0
		for j := 1; j <= m; j++ {
			tmp := dp[j]
			if target[i-1] == arr[j-1] {
				dp[j] = pre + 1
			} else {
				dp[j] = utils.MaxNum(dp[j], dp[j-1])
			}
			pre = tmp
		}
	}
	return dp[m]
}

func LSC_TimeONlogN(target []int, arr []int) int {
	n := len(target)
	pos := make(map[int]int, n)
	for i, val := range target {
		pos[val] = i
	}

	d := []int{}
	for _, val := range arr {
		if idx, has := pos[val]; has {
			// if arr has val
			if p := sort.SearchInts(d, idx); p < len(d) {
				d[p] = idx
			} else {
				d = append(d, idx)
			}
		}
	}
	return n - len(d)
}

// leetcode1716
func TotalMoney(n int) int {
	cnt := n / 7
	res := 0
	if cnt == 0 {
		return (1 + n%7) * (n % 7) / 2
	}
	// cnt 必须大于1
	res += (28 + 28 + 7*(cnt-1)) * cnt / 2
	// 如果cnt是被整除的
	if n%7 != 0 {
		// 首项cnt+1  末项cnt+1+n%7
		res += (cnt + 1 + cnt + n%7) * (n % 7) / 2
	}
	return res
}

// leetcode1728
// rows == grid.length
// cols = grid[i].length
// 1 <= rows, cols <= 8
// grid[i][j] 只包含字符 'C' ，'M' ，'F' ，'.' 和 '#' 。
// grid 中只包含一个 'C' ，'M' 和 'F' 。
// 1 <= catJump, mouseJump <= 8

// func CanMouseWin(grid []string, catJump int, mouseJump int) bool {
// 	n, m := len(grid), len(grid[0])
// 	var fx, fy int
// 	var cx, cy int
// 	var mx, my int
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < m; j++ {
// 			if grid[i][j] == 'F' {
// 				fx, fy = i, j
// 			} else if grid[i][j] == 'M' {
// 				mx, my = i, j
// 			} else if grid[i][j] == 'C' {
// 				cx, cy = i, j
// 			}
// 		}
// 	}

// }
