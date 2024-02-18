package leetcode

import (
	"fmt"
	"programs/kit/utils"
	"sort"
	"strconv"
	"strings"
)

// leetcode453
func MinMoves(nums []int) int {
	sum := 0
	min := 0xFFFFFFFF
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if nums[i] < min {
			min = nums[i]
		}
	}
	return sum - min*len(nums)
}

// leetcode454
func FourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	mp := make(map[int]int)
	for _, val1 := range nums1 {
		for _, val2 := range nums2 {
			mp[val1+val2]++
		}
	}
	cnt := 0
	for _, val1 := range nums3 {
		for _, val2 := range nums4 {
			if _, ok := mp[-(val1 + val2)]; ok {
				cnt++
			}
		}
	}
	return cnt
}

// leetcode455
func FindContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j, cnt := 0, 0, 0
	for i < len(s) && j < len(g) {
		if s[i] >= g[j] {
			cnt++
			i++
			j++
		} else {
			i++
		}
	}
	return cnt
}

// leetcode459
func RepeatedSubstringPattern(s string) bool {
	next := make([]int, len(s)+1)
	getNext := func(str string) {
		i, j := 0, -1
		next[0] = -1
		for i < len(str) {
			if j == -1 || str[i] == str[j] {
				i++
				j++
				next[i] = j
			} else {
				j = next[j]
			}
		}
	}
	getNext(s)
	last := next[len(s)]
	if last == 0 {
		return false
	}
	return len(s)%(len(s)-last) == 0
}

// leetcode467
func FindSubstringInWraproundString(p string) int {
	// dp[i]代表以字母'a'+i结尾的最大长度
	dp := make([]int, 26)
	k := 1 //连续的长度
	dp[int(p[0]-'a')] = 1
	for i := 1; i < len(p); i++ {
		if p[i]-p[i-1] == 1 || p[i-1]-p[i] == 25 {
			// 如果是连续的
			k++
		} else {
			// 如果不连续
			k = 1
		}
		dp[int(p[i]-'a')] = utils.MaxNum(k, dp[int(p[i]-'a')])
	}
	res := 0
	for i := 0; i < 26; i++ {
		res += dp[i]
	}
	return res
}

// leetcode468
func ValidIPAddress(queryIP string) string {
	// ipv4
	words := strings.Split(queryIP, ".")
	if len(res) == 4 {
		for _, word := range words {
			val, err := strconv.Atoi(word)
			if err != nil || val >= 256 || val < 0 {
				break
			}
		}
		return "IPv4"
	}
	words = strings.Split(queryIP, ":")
	check := func(ch byte) bool {
		if ch <= 'F' && ch >= 'A' || ch <= 'f' && ch >= 'a' || ch <= '9' && ch >= '0' {
			return true
		}
		return false
	}
	if len(res) == 8 {
		for _, word := range words {
			for i := 0; i < len(word); i++ {
				if !check(word[i]) {
					break
				}
			}
		}
		return "IPv6"
	}
	return "Neither"
}

// leetcode473
func Makesquare(matchsticks []int) bool {
	sum := 0
	for _, val := range matchsticks {
		sum += val
	}
	if sum%4 != 0 {
		return false
	}
	t := sum / 4
	edges := [4]int{}
	var dfs func(int) bool
	dfs = func(x int) bool {
		if x == len(matchsticks) {
			return true
		}
		for i := 0; i < 4; i++ {
			edges[i] += matchsticks[x]
			if edges[i] <= t && dfs(x+1) {
				return true
			}
			edges[i] -= matchsticks[x]
		}
		return false
	}
	return dfs(0)
}

// 状态压缩+记忆化搜索
func Makesquare2(matchsticks []int) bool {
	edgeLen, sum, n := 0, 0, len(matchsticks)
	dp := make([]int, 1<<n)

	for _, val := range matchsticks {
		sum += val
	}
	if sum%4 != 0 {
		return false
	}
	edgeLen = sum / 4

	var dfs func(int, int) bool
	dfs = func(status, sum int) bool {
		if status == 1<<n && sum == edgeLen {
			return true
		}
		if sum > edgeLen {
			return false
		}
		if dp[status] != 0 {
			return dp[status] == 1
		}

		for i, val := range matchsticks {
			if sum == edgeLen {
				sum = 0
			}
			if (status>>i)&i == 1 {
				continue
			}
			if dfs(status|(1<<i), sum+val) {
				dp[status] = 1
				return true
			}
		}
		dp[status] = -1
		return false
	}

	return dfs(0, 0)
}

// leetcode474
func FindMaxForm(strs []string, m int, n int) int {
	// calculate num of zero and one
	strsNum := func(str string) (int, int) {
		zeroNum, oneNum := 0, 0
		for _, ch := range str {
			if ch == '0' {
				zeroNum++
			} else {
				oneNum++
			}
		}
		return zeroNum, oneNum
	}
	w0, w1 := make([]int, len(strs)+1), make([]int, len(strs)+1)
	for i, str := range strs {
		w0[i+1], w1[i+1] = strsNum(str)
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= len(strs); i++ {
		for p := m; p >= w0[i]; p-- {
			for q := n; q >= w1[i]; q-- {
				dp[p][q] = utils.MaxNum(dp[p][q], dp[p-w0[i]][q-w1[i]]+1)
			}
		}
	}
	return dp[m][n]
}

// leetcode475
func FindRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	binarySearchFirst := func(nums []int, target int) int {
		low, high := 0, len(nums)-1
		for low <= high {
			mid := (low + high) / 2
			if nums[mid] >= target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		return low
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	res := 0
	for i := 0; i < len(houses); i++ {
		r := binarySearchFirst(heaters, houses[i])
		l := r - 1
		if l < 0 {
			if houses[i] < heaters[r]-res {
				res = heaters[r] - houses[i]
			}
			continue
		}
		if r >= len(heaters) {
			if houses[i] > heaters[l]+res {
				res = houses[i] - heaters[l]
			}
			continue
		}
		if houses[i] < heaters[r]-res && houses[i] > heaters[l]+res {
			//不在右边加热器的范围      并且         不在左边加热器的范围
			res = min(houses[i]-heaters[l], heaters[r]-houses[i])
		}
	}
	return res
}

// leetcode476
func FindComplement(num int) int {
	res := []int{}
	for num > 0 {
		res = append(res, num&1)
		num >>= 1
	}
	ans := 0
	fmt.Println(res)
	for i := len(res) - 1; i >= 0; i-- {
		if res[i] == 0 {
			ans |= 1
		}
		ans <<= 1
	}
	return ans >> 1
}

// leetcode488
func FindMinStep(board string, hand string) int {
	for i := 0; i < len(board); i++ {

	}
	return 1
}

// leetcode491
func FindSubsequences(nums []int) [][]int {
	var backtrace func(start int)
	res := [][]int{}
	cur := []int{}

	backtrace = func(start int) {
		if len(cur) >= 2 {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
		}
		if start == len(nums) {
			return
		}
		history := make([]int, 201)
		for i := start; i < len(nums); i++ {
			if len(cur) > 0 && nums[i] < cur[len(cur)-1] || history[nums[i]+100] == 1 {
				continue
			}
			history[nums[i]+100] = 1
			cur = append(cur, nums[i])
			backtrace(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	backtrace(0)
	return res
}

// leetcode492
func ConstructRectangle(area int) []int {
	num1 := 1
	for i := 1; i*i <= area; i++ {
		if area%i == 0 {
			num1 = i
		}
	}
	res := []int{area / num1, num1}
	return res
}

// leetcode495
func FindPoisonedDuration(timeSeries []int, duration int) int {
	begin, end := 0, -1
	res := 0
	for _, time := range timeSeries {
		if time > end {
			res += (end - begin + 1)
			end = time + duration - 1
			begin = time
		} else {
			end = time + duration - 1
		}
	}
	res += (end - begin + 1)
	return res
}

// leetcode494
func FindTargetSumWays1(nums []int, target int) int {
	var backtrace func(end, target int) int
	backtrace = func(end, target int) int {
		if end == 0 {
			if target == 0 && nums[end] == 0 {
				return 2
			} else if target == nums[end] || target == -nums[end] {
				return 1
			}
			return 0
		}
		res := 0
		res += backtrace(end-1, target+nums[end])
		res += backtrace(end-1, target-nums[end])
		return res
	}
	ans := backtrace(len(nums)-1, target)
	return ans
}

// leetcode494 剪枝版本
func FindTargetSumWays2(nums []int, target int) int {
	upperBound := make([]int, len(nums)+1)
	lowerBound := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		upperBound[i] = upperBound[i-1] + nums[i-1]
		lowerBound[i] = lowerBound[i-1] - nums[i-1]
	}

	var backtrace func(end, target int) int
	backtrace = func(end, target int) int {
		if end == 0 {
			if target == 0 && nums[end] == 0 {
				return 2
			} else if target == nums[end] || target == -nums[end] {
				return 1
			}
			return 0
		}
		if target > upperBound[end+1] || target < lowerBound[end+1] {
			return 0
		}
		res := 0
		res += backtrace(end-1, target+nums[end])
		res += backtrace(end-1, target-nums[end])
		return res
	}
	ans := backtrace(len(nums)-1, target)
	return ans
}

func FindTargetSumWays3(nums []int, target int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if abs(target) > sum {
		return 0
	}
	if (sum+target)/2%2 == 1 {
		return 0
	}

	leftSum := (sum + target) / 2
	dp := make([]int, leftSum+1)
	dp[0] = 1
	// dp[i][j] += dp[i-1][j-nums[i]]
	// 初始化dp[0]是1
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[leftSum]
}

// leetcode496
func NextGreaterElement1(nums []int) []int {
	// 给定一个数组，返回一个等长的数组，对应索引存储着下一个更大元素，如果没有更大的元素，就存 -1。
	// 输入：[2, 1, 2, 4, 3]
	// 输出：[4, 2, 4, -1, -1]

	res := make([]int, len(nums))
	stack := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = -1
		} else {
			res[i] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return res
}

func NextGreaterElement2(nums []int) []int {
	// 给定一个数组 T = [73, 74, 75, 71, 69, 72, 76, 73]
	// 该数组存放是近几天的天气气温（华氏度）
	// 返回一个数组，计算：对于每一天，至少等多少天才能等到一个更暖和的气温；如果等不到那一天，填 0 。
	// 输入 T = [73, 74, 75, 71, 69, 72, 76, 73]
	// 输出 R = [1, 1, 4, 2, 1, 1, 0, 0]。
	// 解释：第一天 73 华氏度，第二天 74 华氏度，比 73 大，所以对于第一天，只要等一天就能等到一个更暖和的气温。后面的同理。

	res := make([]int, len(nums))
	stack := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) != 0 && nums[stack[len(stack)-1]] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = 0
		} else {
			index := stack[len(stack)-1]
			res[i] = index - i
		}
		stack = append(stack, i)
	}
	return res
}

func NextGreaterElement3(nums1, nums2 []int) []int {
	res := make([]int, len(nums1))
	stack := []int{}
	mp := map[int]int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		for len(stack) > 0 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			mp[num] = stack[len(stack)-1]
		} else {
			mp[num] = -1
		}
		stack = append(stack, num)
	}
	for i, num := range nums1 {
		res[i] = mp[num]
	}
	return res
}

// leetcode498
func FindDiagonalOrder(mat [][]int) []int {
	curSum := 0
	n, m := len(mat), len(mat[0])
	res := make([]int, 0, n*m)
	for curSum < n+m-1 {
		if curSum%2 == 0 {
			x := utils.MinNum(curSum, n)
			for x > 0 && curSum-x < m {
				y := curSum - x
				res = append(res, mat[x][y])
				x--
			}

		} else {
			y := utils.MinNum(curSum, m)
			for y > 0 && curSum-y < n {
				x := curSum - y
				res = append(res, mat[x][y])
				y--
			}
		}
		curSum++
	}
	return res
}
