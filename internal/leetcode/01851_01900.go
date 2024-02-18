package leetcode

import "programs/kit/utils"

// leetcode1854
func MaximumPopulation(logs [][]int) int {
	diff := make([]int, 105)
	sum := make([]int, 105)
	for _, log := range logs {
		birth, death := log[0], log[1]
		diff[birth-1950] += 1
		diff[death-1950] -= 1
	}
	for i := 1; i < 105; i++ {
		sum[i] = sum[i-1] + diff[i-1]
	}
	// sum[i]代表的年份i-1+1950
	res, year := 0, -1
	for i := 1; i < 105; i++ {
		if sum[i] > res {
			res = sum[i]
			year = i - 1 + 1950
		}
	}
	return year
}

// leetcode1883
func MaxIceCream(costs []int, coins int) int {
	dp := make([]int, coins+1)

	for i := 0; i < len(costs); i++ {
		for j := coins; j >= 0; j-- {
			if j >= costs[i] {
				dp[j] = utils.MaxNum(dp[j], dp[j-costs[i]]+1)
			}
		}
	}
	return dp[coins]
}
