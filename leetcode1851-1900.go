package leetcode

import "programs/kit/common"

// leetcode1883
func MaxIceCream(costs []int, coins int) int {
	dp := make([]int, coins+1)

	for i := 0; i < len(costs); i++ {
		for j := coins; j >= 0; j-- {
			if j >= costs[i] {
				dp[j] = common.LargerNumber(dp[j], dp[j-costs[i]]+1)
			}
		}
	}
	return dp[coins]
}
