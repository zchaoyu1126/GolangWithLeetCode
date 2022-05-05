package leetcode

import (
	"math/rand"
	"programs/internal/algorithmingo/algorithm"
	"programs/kit/common"
	"sort"
	"strconv"
	"strings"
)

// Leetcode501
func FindMode(root *algorithm.TreeNode) []int {
	var prev *algorithm.TreeNode
	var inorder func(node *algorithm.TreeNode)
	mp := make(map[int][]int)
	curReNum := 1
	maxReNum := 0
	inorder = func(node *algorithm.TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if prev != nil {
			if prev.Val == node.Val {
				curReNum++
			} else {
				if curReNum >= maxReNum {
					if _, ok := mp[curReNum]; !ok {
						mp[curReNum] = []int{}
					}
					mp[curReNum] = append(mp[curReNum], prev.Val)
					maxReNum = curReNum
				}
				curReNum = 1
			}
		}
		prev = node
		inorder(node.Right)
	}
	inorder(root)
	if curReNum >= maxReNum {
		if _, ok := mp[curReNum]; !ok {
			mp[curReNum] = []int{}
		}
		mp[curReNum] = append(mp[curReNum], prev.Val)
		maxReNum = curReNum
	}
	return mp[maxReNum]
}

// leetcode503
func NextGreaterElements4(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	stack := []int{}
	for i := 2*len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i%length] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i%length] = -1
		} else {
			res[i%length] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i%length])
	}
	return res
}

// leetcode506
func FindRelativeRanks(score []int) []string {
	mp := make(map[int]string)
	tmp := make([]int, len(score))
	copy(tmp, score)
	sort.Sort(sort.Reverse(sort.IntSlice(tmp)))
	for i := 0; i < len(tmp); i++ {
		if i == 0 {
			mp[tmp[i]] = "Gold Medal"
		} else if i == 1 {
			mp[tmp[i]] = "Silver Medal"
		} else if i == 2 {
			mp[tmp[i]] = "Bronze Medal"
		} else {
			mp[tmp[i]] = strconv.Itoa(i)
		}
	}
	res := []string{}
	for i := 0; i < len(score); i++ {
		res = append(res, mp[score[i]])
	}
	return res
}

// leetcode507
func CheckPerfectNumber(num int) bool {
	sum := 0
	for i := 0; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			sum += (num / i)
		}
	}
	return sum == num
}

// leetcode509
func Fib509(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	prev1, prev2 := 0, 1
	ans := 0
	for i := 2; i <= n; i++ {
		ans = prev1 + prev2
		prev1, prev2 = prev2, ans
	}
	return ans
}

// leetcode513
func FindBottomLeftValue(root *algorithm.TreeNode) int {
	res, curDepth := 0, -1
	var dfs func(node *algorithm.TreeNode, d int)
	dfs = func(node *algorithm.TreeNode, d int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if curDepth < d {
				curDepth = d
				res = node.Val
			}
		}
		dfs(node.Left, d+1)
		dfs(node.Right, d+1)
	}
	dfs(root, 0)
	return res
}

// leetcode515
func LargestValues(root *algorithm.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	queue := []*algorithm.TreeNode{root}
	for len(queue) != 0 {
		size := len(queue)
		cur := INT_MIN
		for i := 0; i < size; i++ {
			front := queue[0]
			queue = queue[1:]
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
			if front.Val > cur {
				cur = front.Val
			}
		}
		res = append(res, cur)
	}
	return res
}

// leetcode516 最长回文子序列
func LongestPalindromeSubseq(s string) int {
	length := len(s)
	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, length)
	}
	for i := length - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < length; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = common.LargerNumber(dp[i+1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[0][length-1]
}

// 问题可以转化为求最长公共子序列
func LongestPalindromeSubseq2(s string) int {
	t := reverseString(s)
	res := longestCommonSubseq(s, t)
	return res
}

func longestCommonSubseq(s, t string) int {
	n, m := len(s), len(t)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = common.LargerNumber(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}

func reverseString(s string) string {
	bytes := []byte(s)
	for from, to := 0, len(bytes)-1; from < to; from, to = from+1, to-1 {
		bytes[from], bytes[to] = bytes[to], bytes[from]
	}
	return string(bytes)
}

// leetcode518
func Change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}

// leetcode519
type Leetcode519Solution struct {
	mp          map[int]int
	m, n, total int
}

func Leetcode519Constructor(m int, n int) Leetcode519Solution {
	return Leetcode519Solution{map[int]int{}, m, n, m * n}
}

func (t *Leetcode519Solution) Flip() (ans []int) {
	pos := rand.Intn(t.total)
	t.total--
	if newPos, taken := t.mp[pos]; taken {
		ans = []int{newPos / t.n, newPos % t.n}
	} else {
		ans = []int{pos / t.n, pos % t.n}
	}
	if newPos, taken := t.mp[t.total]; taken {
		t.mp[pos] = newPos
	} else {
		t.mp[pos] = t.total
	}
	return
}

func (t *Leetcode519Solution) Leetcode519Reset() {
	t.mp = map[int]int{}
	t.total = t.m * t.n
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(m, n);
 * param_1 := obj.Flip();
 * obj.Reset();
 */

// leetcode520
func DetectCapitalUse(word string) bool {
	sum, lower := 0, strings.ToLower(word)
	for i := 0; i < len(word); i++ {
		sum += int(lower[i] - word[i])
	}
	if sum == len(word)*32 || sum == 0 {
		return true
	}
	if sum == 32 && word[0] >= 'Z' && word[0] <= 'A' {
		return true
	}
	return false
}

// leetcode524
type StringSlice []string

func (s StringSlice) Len() int      { return len(s) }
func (s StringSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s StringSlice) Less(i, j int) bool {
	if len(s[i]) == len(s[j]) {
		return s[i] < s[j]
	}
	return len(s[i]) > len(s[j])
}

func FindLongestWord(s string, dictionary []string) string {
	mp := []string{}
	for _, str := range dictionary {
		i, j := 0, 0
		for i < len(s) && j < len(str) {
			if s[i] == str[j] {
				i++
				j++
			} else {
				i++
			}
		}
		if j == len(str) {
			mp = append(mp, str)
		}
	}
	if len(mp) == 0 {
		return ""
	}
	sort.Sort(StringSlice(mp))
	return mp[0]
}

// leetcode530
func GetMinimumDifference(root *algorithm.TreeNode) int {
	res := INT_MAX
	var prev *algorithm.TreeNode
	var inorderTraversal func(node *algorithm.TreeNode)
	inorderTraversal = func(node *algorithm.TreeNode) {
		if node == nil {
			return
		}
		inorderTraversal(node.Left)
		if prev != nil && node.Val-prev.Val < res {
			res = node.Val - prev.Val
		}
		prev = node
		inorderTraversal(node.Right)
	}
	inorderTraversal(root)
	return res
}

// leetcode538
func ConvertBST(root *algorithm.TreeNode) *algorithm.TreeNode {
	var traversal func(root *algorithm.TreeNode)
	sum := 0
	traversal = func(root *algorithm.TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Right)
		sum += root.Val
		root.Val = sum
		traversal(root.Left)
	}
	traversal(root)
	return root
}

// leetcode539
func FindMinDifference(timePoints []string) int {
	data := []int{}
	if len(timePoints) > 1440 {
		return 0
	}
	for i := 0; i < len(timePoints); i++ {
		hour, _ := strconv.Atoi(timePoints[i][0:2])
		minute, _ := strconv.Atoi(timePoints[i][3:5])
		data = append(data, hour*60+minute)
	}
	sort.Ints(data)
	res := 1440
	for i := 0; i < len(data)-1; i++ {
		gap := data[i+1] - data[i]
		gap = common.SmallerNumber(gap, 1440-gap)
		if gap < res {
			res = gap
		}
	}
	return res
}
