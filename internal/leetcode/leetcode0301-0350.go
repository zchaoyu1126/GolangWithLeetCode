package leetcode

import (
	"container/heap"
	"fmt"
	"programs/internal/algorithmingo/algorithm"
	"programs/kit/utils"
	"sort"
)

// leetcode304
type NumMatrix struct {
	sum [][]int
}

func NewNumMatrix(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	sum := make([][]int, m+1)
	for i := range sum {
		sum[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{sum: sum}
}

func (n *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return n.sum[row2+1][col2+1] - n.sum[row1][col2+1] - n.sum[row2+1][col1] + n.sum[row1][col1]
}

// leetcode307
type NumArray struct {
	*algorithm.SegTree
}

func NewArray(nums []int) NumArray {
	return NumArray{algorithm.NewSegTree(nums)}
}

func (t *NumArray) Update(index int, val int) {
	t.Modify(index, val, 0)
}

func (t *NumArray) SumRange(left int, right int) int {
	return t.Query(left, right, 0)
}

// leetcode309
func MaxProfit(prices []int) int {
	n := len(prices)
	if n == 0 || n == 1 {
		return 0
	}
	dp := make([][3]int, n)
	dp[0][0] = -prices[0] // 持有一支股票
	dp[0][1] = 0          // 未持有股票，但处于冷冻期，第i天刚卖出
	dp[0][2] = 0          // 未持有股票，但不处于冷冻期

	for i := 1; i < n; i++ {
		// 持有股票
		dp[i][0] = utils.MaxNum(dp[i-1][0], dp[i-1][2]-prices[i])
		// 未持有股票, 因为今天刚卖出
		dp[i][1] = dp[i-1][0] + prices[i]
		// 未持有股票，昨天卖的，或者前几天卖的
		dp[i][2] = utils.MaxNum(dp[i-1][1], dp[i-1][2])
	}
	return utils.MaxNum(dp[n-1][1], dp[n-1][2])
}

// leetcode318
func MaxProduct(words []string) int {
	masks := make([]int, len(words))
	for i, word := range words {
		for _, ch := range word {
			masks[i] |= 1 << (ch - 'a')
		}
	}

	length := 0
	for i := 0; i < len(masks); i++ {
		for j := i + 1; j < len(masks); j++ {
			if masks[i]&masks[j] == 0 {
				if len(words[i])*len(words[j]) > length {
					length = len(words[j]) * len(words[i])
				}
			}
		}
	}
	return length
}

// leetcode322
func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = INT_MAX
	}
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != INT_MAX {
				dp[j] = utils.MinNum(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	return dp[amount]
}

// leetcode326
func IsPowerOfThree(n int) bool {
	l := 1
	if n == 1 {
		return true
	}
	for n > l {
		l *= 3
	}
	return n == l
}

// leetcode332
func FindItinerary(tickets [][]string) []string {
	cur := []string{"JFK"}

	mp := make(map[string][]string)
	used := make(map[string]map[string]int)

	for i := 0; i < len(tickets); i++ {
		from, to := tickets[i][0], tickets[i][1]
		if _, ok := mp[from]; !ok {
			mp[from] = []string{}
			used[from] = make(map[string]int)
		}
		mp[from] = append(mp[from], to)
		used[from][to]++
	}
	for key := range mp {
		sort.Strings(mp[key])
	}

	var backtrace func(startPos string) bool
	backtrace = func(startPos string) bool {
		// 找到之后应该给上一级报告说找到了
		if len(cur) == len(tickets)+1 {
			return true
		}

		for _, to := range mp[startPos] {
			if used[startPos][to] == 0 {
				continue
			}
			cur = append(cur, to)
			used[startPos][to]--
			if backtrace(to) {
				// 从to出发找到了 直接return,不能再cur = cur[:len(cur)-1]
				return true
			}
			cur = cur[:len(cur)-1]
			used[startPos][to]++
		}
		return false
	}
	backtrace("JFK")
	return cur
}

// leetcode334
func IncreasingTriplet(nums []int) bool {
	// 最长上升子序列
	stack := []int{}
	for i := 0; i < len(nums); i++ {
		if len(stack) == 0 {
			stack = append(stack, nums[i])
			continue
		}
		top := stack[len(stack)-1]
		if top >= nums[i] {
			// 找到第一个比nums[i]小的数进行修改
			cur := len(stack) - 1
			for cur >= 0 && stack[cur] > nums[i] {
				cur--
			}
			stack[cur+1] = nums[i]
		} else {
			stack = append(stack, nums[i])
		}
	}
	fmt.Println(stack)
	return len(stack) >= 3
}

// leetcode335
func IsSelfCrossing(distance []int) bool {
	for i := 3; i < len(distance); i++ {
		// 第 1 类路径交叉的情况
		if distance[i] >= distance[i-2] && distance[i-1] <= distance[i-3] {
			return true
		}

		// 第 2 类路径交叉的情况
		if i == 4 && distance[3] == distance[1] &&
			distance[4] >= distance[2]-distance[0] {
			return true
		}

		// 第 3 类路径交叉的情况
		if i >= 5 && distance[i-3]-distance[i-5] <= distance[i-1] &&
			distance[i-1] <= distance[i-3] &&
			distance[i] >= distance[i-2]-distance[i-4] &&
			distance[i-2] > distance[i-4] {
			return true
		}
	}
	return false
}

// leetcode337
var mp337 map[*algorithm.TreeNode]int

func Rob337_1(root *algorithm.TreeNode) int {
	mp337 = make(map[*algorithm.TreeNode]int)
	return dfs(root)
}
func dfs(root *algorithm.TreeNode) int {
	if root == nil {
		return 0
	}
	if val, has := mp337[root]; has {
		return val
	}
	// 决定抢根节点
	val1 := root.Val
	if root.Left != nil {
		val1 += dfs(root.Left.Left)
		val1 += dfs(root.Left.Right)
	}
	if root.Right != nil {
		val1 += dfs(root.Right.Left)
		val1 += dfs(root.Right.Right)
	}
	// 决定放弃根节点
	val2 := dfs(root.Left) + dfs(root.Right)
	res := utils.MaxNum(val1, val2)
	mp337[root] = res
	return res
}

func Rob337_2(root *algorithm.TreeNode) int {
	var dfs func(*algorithm.TreeNode) [2]int
	dfs = func(root *algorithm.TreeNode) [2]int {
		if root == nil {
			return [2]int{0, 0}
		}
		v1 := dfs(root.Left)
		v2 := dfs(root.Right)
		res := [2]int{}
		res[0] = v1[0] + v2[0] + root.Val
		res[1] = utils.MaxNum(v1[0], v1[1]) + utils.MaxNum(v2[0], v2[1])
		return res
	}
	res := dfs(root)
	return utils.MaxNum(res[0], res[1])
}

// leetcode343
func IntegerBreak(n int) int {
	dp := make([]int, n+1)
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	dp[1], dp[2] = 1, 1
	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			x := utils.MaxNum(dp[j], j)
			y := utils.MaxNum(dp[i-j], i-j)
			dp[i] = utils.MaxNum(dp[i], x*y)
		}
	}
	return dp[n]
}

// leetcode344
func ReverseString(s []byte) {
	n := len(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		tmp := s[j]
		s[j] = s[i]
		s[i] = tmp
	}
}

// leetcode345
func ReverseVowels(s string) string {
	i, j := 0, len(s)-1
	data := []byte(s)
	for i < j {
		for i <= len(s)-1 && !isAEIOU(data[i]) {
			i++
		}
		for j >= 0 && !isAEIOU(data[j]) {
			j--
		}
		if i < j {
			data[i], data[j] = data[j], data[i]
		}
		i++
		j--
	}
	fmt.Println(string(data))
	return string(data)
}

func isAEIOU(x byte) bool {
	if x == 'a' || x == 'A' || x == 'e' || x == 'E' || x == 'i' || x == 'I' || x == 'o' || x == 'O' || x == 'U' || x == 'u' {
		return true
	}
	return false
}

// leetcode347
type frequency struct {
	val int
	cnt int
}
type freHeap []frequency

func (fh freHeap) Len() int            { return len(fh) }
func (fh freHeap) Less(i, j int) bool  { return fh[i].cnt > fh[j].cnt }
func (fh freHeap) Swap(i, j int)       { fh[i], fh[j] = fh[j], fh[i] }
func (fh *freHeap) Push(x interface{}) { *fh = append(*fh, x.(frequency)) }
func (fh *freHeap) Pop() interface{}   { top := (*fh)[:len(*fh)-1]; *fh = (*fh)[:len(*fh)-1]; return top }
func TopKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}
	h := freHeap{}
	for key, val := range mp {
		h = append(h, frequency{key, val})
	}
	heap.Init(&h)
	ans := []int{}
	for i := 0; i < k; i++ {
		ans = append(ans, h[0].val)
		heap.Pop(&h)
	}
	return ans
}

// leetcode349
func Intersection(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		mp[nums1[i]] = 1
	}
	res := []int{}
	for i := 0; i < len(nums2); i++ {
		if val, ok := mp[nums2[i]]; ok && val > 0 {
			res = append(res, nums2[i])
			mp[nums2[i]]--
		}
	}
	return res
}

func Intersect(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		mp[nums1[i]]++
	}
	res := []int{}
	for i := 0; i < len(nums2); i++ {
		if cnt, ok := mp[nums2[i]]; ok && cnt > 0 {
			res = append(res, nums2[i])
			mp[nums2[i]]--
		}
	}
	return res
}

// leetcode350
func IsHappy(n int) bool {
	mp := make(map[int]int)
	for {
		sum := digitalSum(n)
		if sum == 1 {
			return true
		}
		if _, ok := mp[sum]; !ok {
			mp[sum]++
		} else {
			return false
		}
		n = sum
	}
}

func digitalSum(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
