package leetcode

import (
	"math"
	"programs/internal/algorithmingo/algorithm"
	"programs/kit/utils"
	"strconv"
	"strings"
)

// leetcode162
func FindPeakElement(nums []int) int {
	n := len(nums)

	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}

	l, r := 0, n-1
	for {
		mid := (l + r) / 2
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
}

// leetcode165
func CompareVersion(version1 string, version2 string) int {
	res1 := strings.Split(version1, ".")
	res2 := strings.Split(version2, ".")
	i, j := 0, 0
	for ; i < len(res1) && j < len(res2); i, j = i+1, j+1 {
		num1, _ := strconv.Atoi(res1[i])
		num2, _ := strconv.Atoi(res2[j])
		if num1 < num2 {
			return -1
		} else if num1 > num2 {
			return 1
		}
	}
	for ; i < len(res1); i++ {
		num, _ := strconv.Atoi(res1[i])
		if num > 0 {
			return 1
		}
	}
	for ; j < len(res2); j++ {
		num, _ := strconv.Atoi(res2[j])
		if num > 0 {
			return -1
		}
	}
	return 0
}

// leetcode168
func ConvertToTitle(columnNumber int) string {
	var single byte
	var res string
	if columnNumber == 0 {
		return ""
	}
	for columnNumber != 0 {
		columnNumber--
		single = byte('A' + columnNumber%26)
		res = string(single) + res
		columnNumber /= 26
	}
	return res
}

// leetcode187
func FindRepeatedDnaSequences(s string) []string {
	root := algorithm.NewRuneTrie()
	res := []string{}
	for i := 0; i+10 <= len(s); i++ {
		if !root.Find(s[i : i+10]) {
			root.Insert(s[i : i+10])
		} else {
			res = append(res, s[i:i+10])
		}
	}
	return RemoveRepByMap(res)
}

func RemoveRepByMap(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

// leetcode188
func MaxProfit4(k int, prices []int) int {
	n := len(prices)
	if n == 0 || n == 1 || k == 0 {
		return 0
	}
	dp := make([][][2]int, n)
	for i := range dp {
		dp[i] = make([][2]int, k+1)
	}
	// j=1时，代表第一次交易
	dp[0][1][0] = 0          // 未持有股票
	dp[0][1][1] = -prices[0] // 持有股票
	for i := 1; i <= k; i++ {
		dp[0][i][1] = -prices[0]
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= k; j++ {
			// 第i天未持股票，之前一直没有或者今天把之前的卖了
			dp[i][j][0] = utils.MaxNum(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			// 第i天持有股票，之前一直有股票，或者是取消之前的买入操作，在股价更低的地方买入
			dp[i][j][1] = utils.MaxNum(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[n-1][k][0]
}

// leetcode198
func Rob(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]
	for i := 2; i <= len(nums); i++ {
		dp[i] = utils.MaxNum(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[len(nums)]
}

// leetcode199
func RightSideView(root *algorithm.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := [][]int{}

	queue := []*algorithm.TreeNode{root}
	for len(queue) != 0 {
		size := len(queue)
		cur := []int{}
		for i := 0; i < size; i++ {
			front := queue[0]
			queue = queue[1:]
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
			cur = append(cur, front.Val)
		}
		res = append(res, cur)
	}
	ans := []int{}
	for _, cur := range res {
		ans = append(ans, cur[len(cur)-1])
	}
	return ans
}
