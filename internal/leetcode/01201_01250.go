package leetcode

// leetcode1218
func LongestSubsequence(arr []int, difference int) int {
	ans := 0
	dp := map[int]int{}
	for _, v := range arr {
		dp[v] = dp[v-difference] + 1
		if dp[v] > ans {
			ans = dp[v]
		}
	}
	return ans
}

// leetcode1221
func BalancedStringSplit(s string) int {
	lnum, rnum := 0, 0
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			lnum++
		} else {
			rnum++
		}
		if lnum == rnum {
			res++
		}
	}
	return res
}
