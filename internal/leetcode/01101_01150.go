package leetcode

import (
	"programs/kit/utils"
	"strings"
)

// leetcode1108
func DefangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

// leetcode1109
func CorpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n+1)
	for _, booking := range bookings {
		first, last, seats := booking[0], booking[1], booking[2]
		diff[first-1] += seats
		diff[last] -= seats
	}
	sum := make([]int, n+2)
	for i := 1; i <= len(diff); i++ {
		sum[i] = sum[i-1] + diff[i-1]
	}
	return sum[1 : len(sum)-1]
}

// leetcode1143
func LongestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = utils.MaxNum(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}
