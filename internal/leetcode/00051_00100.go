package leetcode

// import (
// 	"fmt"
// 	"programs/internal/algorithmingo/algorithm"
// 	"programs/kit/utils"
// 	"sort"
// )

// // leetcode51
// func SolveNQueens(n int) [][]string {
// 	col := make([]int, n)
// 	l2r := make([]int, 2*n)
// 	r2l := make([]int, 2*n)

// 	res := [][]string{}
// 	cur := make([][]byte, n)
// 	for i := range cur {
// 		cur[i] = make([]byte, n)
// 		for j := range cur[i] {
// 			cur[i][j] = '.'
// 		}
// 	}

// 	var backtrace func(ridx int)
// 	backtrace = func(ridx int) {
// 		if ridx == n {
// 			ans := make([]string, n)
// 			for i := range cur {
// 				ans[i] = string(cur[i])
// 			}
// 			res = append(res, ans)
// 			return
// 		}

// 		i := ridx
// 		for j := 0; j < n; j++ {
// 			if col[j] == 1 || r2l[i+j] == 1 || l2r[i-j+n] == 1 {
// 				continue
// 			}
// 			col[j], r2l[i+j], l2r[i-j+n] = 1, 1, 1
// 			cur[i][j] = 'Q'
// 			backtrace(ridx + 1)
// 			col[j], r2l[i+j], l2r[i-j+n] = 0, 0, 0
// 			cur[i][j] = '.'
// 		}
// 	}
// 	backtrace(0)
// 	return res
// }

// // leetcode53
// // 最大子序和
// // 法一：暴力求解，求前缀和，然后使用二重循环遍历各个子区间
// func MaxSubArray(nums []int) int {
// 	sum := make([]int, len(nums)+1)
// 	for i := 1; i <= len(nums); i++ {
// 		sum[i] = sum[i-1] + nums[i-1]
// 	}
// 	max := -0xFFFFFFFF
// 	for i := 1; i <= len(nums); i++ {
// 		for j := 1; j <= i; j++ {
// 			if sum[i]-sum[j-1] > max {
// 				max = sum[i] - sum[j-1]
// 			}
// 		}
// 	}
// 	fmt.Println(sum)
// 	return max
// }

// // 法二：DP，状态转移方程为f[i] = largerNumber(f[i-1]+nums[i], nums[i])
// // 其中f[i]表示以第i个数结尾时，得到的最大子序列和, 在下述代码中i的取值范围为1~len(nums)
// // 所求结果为 max(f[1], f[2], f[3]... f[len(nums)])
// func MaxSubArray2(nums []int) int {
// 	dp := make([]int, len(nums)+1)
// 	res := -0xFFFFFFFF
// 	for i := 1; i <= len(nums); i++ {
// 		dp[i] = utils.MinNum(dp[i-1]+nums[i-1], nums[i-1])
// 		res = utils.MaxNum(res, dp[i])
// 	}
// 	fmt.Println(dp)
// 	return dp[len(nums)]
// }

// // 法三：分治 线段树的思想

// func MaxSubArray3(nums []int) int {
// 	return get(nums, 0, len(nums)-1).mSum
// }

// func pushUp(l, r Status) Status {
// 	iSum := l.iSum + r.iSum
// 	lSum := utils.MaxNum(l.iSum+r.lSum, l.lSum)
// 	rSum := utils.MaxNum(l.rSum+r.iSum, r.rSum)
// 	mSum := utils.MaxNum(l.rSum+r.lSum, utils.MaxNum(l.mSum, r.mSum))
// 	return Status{lSum: lSum, rSum: rSum, mSum: mSum, iSum: iSum}
// }

// func get(nums []int, l, r int) Status {
// 	if l == r {
// 		return Status{nums[l], nums[l], nums[l], nums[l]}
// 	}
// 	m := (l + r) >> 1
// 	lSub := get(nums, l, m)
// 	rSub := get(nums, m+1, r)
// 	return pushUp(lSub, rSub)
// }

// type Status struct {
// 	lSum, rSum, mSum, iSum int
// }

// // leetcode54
// func SpiralOrder(matrix [][]int) []int {
// 	if len(matrix) == 0 || len(matrix[0]) == 0 {
// 		return []int{}
// 	}
// 	m, n := len(matrix), len(matrix[0])
// 	top, bottom := 0, m-1
// 	left, right := 0, n-1
// 	num, total := 0, m*n
// 	res := make([]int, 0)

// 	for num < total {
// 		for i := left; i <= right; i++ {
// 			res = append(res, matrix[top][i])
// 			num++
// 		}
// 		top++

// 		for i := top; i <= bottom; i++ {
// 			res = append(res, matrix[i][right])
// 			num++
// 		}
// 		right--

// 		fmt.Println(num, total)
// 		if num == total {
// 			break
// 		}

// 		for i := right; i >= left; i-- {
// 			res = append(res, matrix[bottom][i])
// 			num++
// 		}
// 		bottom--

// 		for i := bottom; i >= top; i-- {
// 			res = append(res, matrix[i][left])
// 			num++
// 		}
// 		left++
// 	}
// 	return res
// }

// // leetcode58
// func LengthOfLastWord(s string) int {
// 	start := len(s) - 1
// 	for i := len(s) - 1; i >= 0; i-- {
// 		if s[i] == ' ' {
// 			continue
// 		} else {
// 			start = i
// 			break
// 		}
// 	}
// 	cnt := 0
// 	for i := start; i >= 0; i-- {
// 		if s[i] == ' ' {
// 			break
// 		} else {
// 			cnt++
// 		}
// 	}
// 	return cnt
// }

// // leetcode62
// func UniquePaths(m int, n int) int {
// 	dp := make([][]int, m)
// 	for i := range dp {
// 		dp[i] = make([]int, n)
// 	}
// 	dp[0][0] = 1
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if i == 0 && j == 0 {
// 				continue
// 			}
// 			if i-1 >= 0 && j-1 >= 0 {
// 				dp[i][j] = dp[i-1][j] + dp[i][j-1]
// 			} else if i-1 < 0 {
// 				dp[i][j] = dp[i][j-1]
// 			} else if j-1 < 0 {
// 				dp[i][j] = dp[i-1][j]
// 			}
// 		}
// 	}
// 	fmt.Println(dp)
// 	return dp[m-1][n-1]
// }

// // leetcode63
// func UniquePathsWithObstacles(obstacleGrid [][]int) int {
// 	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
// 		return 0
// 	}
// 	m, n := len(obstacleGrid), len(obstacleGrid[0])
// 	dp := make([][]int, m)
// 	for i := range dp {
// 		dp[i] = make([]int, n)
// 	}
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if obstacleGrid[i][j] == 1 {
// 				continue
// 			}
// 			if i == 0 && j == 0 && obstacleGrid[i][j] == 0 {
// 				dp[i][j] = 1
// 				continue
// 			}

// 			if i-1 >= 0 && j-1 >= 0 {
// 				dp[i][j] = dp[i-1][j] + dp[i][j-1]
// 			} else if i-1 < 0 {
// 				dp[i][j] = dp[i][j-1]
// 			} else if j-1 < 0 {
// 				dp[i][j] = dp[i-1][j]
// 			}
// 		}
// 	}
// 	return dp[m-1][n-1]
// }

// // leetcode66
// func PlusOne(digits []int) []int {
// 	for i := len(digits) - 1; i >= 0; i-- {
// 		if digits[i] != 9 {
// 			digits[i] += 1
// 			return digits
// 		}
// 		digits[i] = 0
// 	}
// 	res := make([]int, len(digits)+1)
// 	res = append(res, 1)
// 	res = append(res, digits...)
// 	return res
// 	//  4 5 6  459  460
// }

// // leetcode67
// func AddBinary(a string, b string) string {
// 	n, m := len(a), len(b)
// 	res := ""
// 	flag := 0
// 	i, j := n-1, m-1
// 	for i >= 0 && j >= 0 {
// 		tmp := int(a[i]-'0') + int(b[j]-'0') + flag
// 		if tmp >= 2 {
// 			flag = 1
// 			res = string('0'+byte(tmp-2)) + res
// 		} else {
// 			res = string('0'+byte(tmp)) + res
// 			flag = 0
// 		}
// 		i--
// 		j--
// 	}

// 	for i >= 0 {
// 		tmp := int(a[i]-'0') + flag
// 		if tmp >= 2 {
// 			flag = 1
// 			res = string('0'+byte(tmp-2)) + res
// 		} else {
// 			res = string('0'+byte(tmp)) + res
// 			flag = 0
// 		}
// 		i--
// 	}
// 	for j >= 0 {
// 		tmp := int(b[j]-'0') + flag
// 		if tmp >= 2 {
// 			flag = 1
// 			res = string('0'+byte(tmp-2)) + res
// 		} else {
// 			res = string('0'+byte(tmp)) + res
// 			flag = 0
// 		}
// 		j--
// 	}
// 	if flag == 1 {
// 		res = "1" + res
// 	}
// 	return res
// }

// // leetcode68
// func FullJustify(words []string, maxWidth int) []string {
// 	oneline := []string{}
// 	length := 0
// 	res := []string{}
// 	for i := 0; i < len(words); i++ {
// 		if len(words[i])+length < maxWidth {
// 			fmt.Println(words[i], "test1")
// 			oneline = append(oneline, words[i])
// 			length += len(words[i])
// 			length++
// 		} else {
// 			fmt.Println(words[i], "test2", len(words[i])+length, length)
// 			tmp := ""
// 			for j := 0; j < len(oneline); j++ {
// 				tmp += oneline[j]
// 				tmp += " "
// 			}
// 			res = append(res, tmp)
// 			oneline = []string{}
// 			oneline = append(oneline, words[i])
// 			length = len(words)
// 		}
// 	}
// 	tmp := ""
// 	for j := 0; j < len(oneline); j++ {
// 		tmp += oneline[j]
// 		tmp += " "
// 	}
// 	res = append(res, tmp)
// 	return res
// }

// // leetcode69
// func MySqrt(x int) int {
// 	l, r := 0, x
// 	// 找最后一个小于等于target的目标
// 	for l <= r {
// 		m := (l + r) / 2
// 		if m*m <= x {
// 			l = m + 1
// 		} else {
// 			r = m - 1
// 		}
// 	}
// 	return r
// }

// // leetcode70
// func ClimbStairs(n int) int {
// 	if n == 1 {
// 		return 1
// 	} else if n == 2 {
// 		return 2
// 	}
// 	pre1, pre2 := 1, 2
// 	for i := 3; i <= n; i++ {
// 		tmp := pre2
// 		pre2 = pre1 + pre2
// 		pre1 = tmp
// 	}
// 	return pre2
// }

// // leetcode72
// func MinDistance(word1 string, word2 string) int {
// 	// dp[i][j] 表示word1[i]到word2[j]需要多少步操作
// 	n, m := len(word1), len(word2)
// 	dp := make([][]int, n+1)
// 	for i := range dp {
// 		dp[i] = make([]int, m+1)
// 	}
// 	// 边界条件
// 	for i := 1; i <= n; i++ {
// 		dp[i][0] = i
// 	}
// 	for j := 1; j <= m; j++ {
// 		dp[0][j] = j
// 	}
// 	for i := 1; i <= n; i++ {
// 		for j := 1; j <= m; j++ {
// 			if word1[i-1] == word2[j-1] {
// 				dp[i][j] = utils.MinNum(dp[i-1][j], dp[i][j-1]) + 1
// 				dp[i][j] = utils.MinNum(dp[i][j], dp[i-1][j-1])
// 			} else {
// 				dp[i][j] = utils.MinNum(dp[i-1][j], dp[i][j-1]) + 1
// 				dp[i][j] = utils.MinNum(dp[i][j], dp[i-1][j-1]+1)
// 			}
// 		}
// 	}
// 	return dp[n][m]
// }

// // leetcode74
// func SearchMatrixI(matrix [][]int, target int) bool {
// 	// 此题中的matrix可转化为一维矩阵
// 	m, n := len(matrix), len(matrix[0])
// 	low, high := 0, m*n-1
// 	for low <= high {
// 		mid := (low + high) / 2
// 		x, y := mid/m, mid%n
// 		if target == matrix[x][y] {
// 			return true
// 		} else if target < matrix[x][y] {
// 			low = mid + 1
// 		} else if target > matrix[x][y] {
// 			high = mid - 1
// 		}
// 	}
// 	return false
// }

// // leetcode76
// func MinWindow1(s string, t string) string {
// 	res := []byte{}
// 	ans := s
// 	mp := make(map[byte]int)
// 	for i := 0; i < len(t); i++ {
// 		mp[t[i]]++
// 	}

// 	check := func() bool {
// 		for _, val := range mp {
// 			if val > 0 {
// 				return false
// 			}
// 		}
// 		return true
// 	}

// 	for l, r := 0, 0; r < len(s); r++ {
// 		res = append(res, s[r])
// 		if _, ok := mp[s[r]]; ok {
// 			mp[s[r]]--
// 			for check() {
// 				if len(string(res)) < len(ans) {
// 					ans = string(res)
// 				}
// 				res = res[1:]
// 				if _, ok := mp[s[l]]; ok {
// 					mp[s[l]]++
// 				}
// 				l++
// 			}
// 		}
// 	}
// 	return ans
// }

// // 更好的写法
// func MinWindow2(s string, t string) string {
// 	window := make(map[byte]int, 0)
// 	need := make(map[byte]int, 0)

// 	left, match := -1, 0
// 	start, end, min := 0, 0, len(s)+1

// 	for i := range t {
// 		need[t[i]]++
// 	}

// 	for right := 0; right < len(s); right++ {
// 		// 1. 直接将s[right]加入到区间，形成（left, right]
// 		ch1 := s[right]
// 		window[ch1]++

// 		//  2. 更新状态
// 		if window[ch1] == need[ch1] {
// 			match++
// 		}

// 		// 3. 超出区间，或者满足条件
// 		for match == len(need) {
// 			if right-left < min {
// 				start, end = left, right
// 				min = right - left
// 			}

// 			// 4. 移除s[++left]，更新状态
// 			left++
// 			ch2 := s[left]
// 			if window[ch2] == need[ch2] {
// 				match--
// 			}
// 			window[ch2]--
// 		}
// 	}

// 	return s[start+1 : end+1]
// }

// // leetcode77
// func Combine(n int, k int) [][]int {
// 	if k == 0 {
// 		return [][]int{}
// 	}
// 	res := [][]int{}
// 	cur := []int{}
// 	var backtrace func(n int, k int)
// 	backtrace = func(n int, k int) {
// 		// 简单的进行减枝操作
// 		if n < k {
// 			return
// 		} else if k == 0 {
// 			tmp := make([]int, len(cur))
// 			copy(tmp, cur)
// 			res = append(res, tmp)
// 		} else if n == 0 {
// 			return
// 		}
// 		for i := n; i >= 1; i-- {
// 			cur = append(cur, i)
// 			backtrace(i-1, k-1)
// 			cur = cur[:len(cur)-1]
// 		}
// 	}
// 	backtrace(n, k)
// 	return res
// }

// // leetcode78
// func Subsets(nums []int) [][]int {
// 	var backtrace func(start int)
// 	res := [][]int{}
// 	cur := []int{}
// 	backtrace = func(start int) {
// 		if start == len(nums) {
// 			return
// 		}
// 		for i := start; i < len(nums); i++ {
// 			cur = append(cur, nums[i])
// 			tmp := make([]int, len(cur))
// 			copy(tmp, cur)
// 			res = append(res, tmp)
// 			backtrace(i + 1)
// 			cur = cur[:len(cur)-1]
// 		}
// 	}
// 	backtrace(0)
// 	res = append(res, []int{})
// 	return res
// }

// // leetcode86
// func Partition86(head *algorithm.ListNode, x int) *algorithm.ListNode {
// 	cur := head
// 	l := &algorithm.ListNode{}
// 	r := &algorithm.ListNode{}
// 	n1, n2 := l, r
// 	for cur != nil {
// 		if cur.Val >= x {
// 			r.Next = cur
// 			r = r.Next
// 		} else {
// 			l.Next = cur
// 			l = l.Next
// 		}
// 		cur = cur.Next
// 	}
// 	l.Next = n2.Next
// 	r.Next = nil
// 	return n1.Next
// }

// // leetcode88
// func Merge(nums1 []int, m int, nums2 []int, n int) {
// 	i, j := m-1, n-1
// 	cur := m + n - 1
// 	for i >= 0 && j >= 0 {
// 		if nums2[j] > nums1[i] {
// 			nums1[cur] = nums2[j]
// 			j--
// 		} else {
// 			nums1[cur] = nums1[i]
// 			i--
// 		}
// 		cur--
// 	}
// 	for i >= 0 {
// 		nums1[cur] = nums1[i]
// 		cur--
// 		i--
// 	}
// 	for j >= 0 {
// 		nums1[cur] = nums2[j]
// 		j--
// 		cur--
// 	}
// 	fmt.Println(nums1)
// }

// // leetcode90
// func SubsetsWithDup(nums []int) [][]int {
// 	sort.Ints(nums)
// 	var backtrace func(start int)
// 	res := [][]int{}
// 	cur := []int{}
// 	backtrace = func(start int) {
// 		if start == len(nums) {
// 			return
// 		}
// 		for i := start; i < len(nums); i++ {
// 			cur = append(cur, nums[i])
// 			tmp := make([]int, len(cur))
// 			copy(tmp, cur)
// 			res = append(res, tmp)
// 			backtrace(i + 1)
// 			cur = cur[:len(cur)-1]
// 			for i+1 < len(nums) && nums[i+1] == nums[i] {
// 				i++
// 			}
// 		}
// 	}
// 	backtrace(0)
// 	res = append(res, []int{})
// 	return res
// }

// // leetcode92
// func ReverseBetween(head *algorithm.ListNode, left int, right int) *algorithm.ListNode {
// 	// 注：1<=left<=right<=n
// 	dummy := &algorithm.ListNode{}
// 	dummy.Next = head
// 	cur := dummy
// 	guard := dummy
// 	for i := 0; i < left-1; i++ {
// 		cur = cur.Next
// 		guard = guard.Next
// 	}
// 	// guard用于连接反转后的头节点
// 	// left节点前一个 cur
// 	// 开始反转 cur.Next是要反转的目标

// 	// cur.Next不可能为nil, cur.Next.Next成立不会越界
// 	last, has := cur.Next, cur.Next
// 	remain := cur.Next.Next
// 	has.Next = nil
// 	// 极端情况right=left=n, 不会执行这个循环
// 	// 此时has = last node, remain = nil
// 	for i := 0; i < right-left; i++ {
// 		// 将第一个节点放入has中
// 		tmp := remain.Next
// 		remain.Next = has
// 		has = remain
// 		// 更新remain
// 		remain = tmp
// 	}
// 	// 剩下的节点
// 	guard.Next = has
// 	remain.Next = last
// 	return dummy.Next
// }

// // leetcode93
// func RestoreIpAddresses(s string) []string {
// 	var backtrace func(start int)
// 	res := []string{}
// 	cur := []string{}
// 	backtrace = func(start int) {
// 		if start == len(s) {
// 			fmt.Println(cur)
// 		}
// 		for i := start; i < len(s); i++ {
// 			cur = append(cur, s[start:i+1])
// 			backtrace(i + 1)
// 			cur = cur[:len(cur)-1]
// 		}
// 	}

// 	backtrace(0)
// 	return res
// }

// // leetcode94
// func InorderTraversal1(root *algorithm.TreeNode) []int {
// 	var traversal func(root *algorithm.TreeNode)
// 	res := []int{}
// 	traversal = func(root *algorithm.TreeNode) {
// 		if root == nil {
// 			return
// 		}
// 		res = append(res, root.Val)
// 		traversal(root.Left)
// 		traversal(root.Right)
// 	}
// 	traversal(root)
// 	return res
// }

// func InorderTraversal2(root *algorithm.TreeNode) []int {
// 	stack := []*algorithm.TreeNode{}
// 	res := []int{}
// 	cur := root
// 	// 先找到最左边的节点
// 	for cur != nil || len(stack) != 0 {
// 		if cur != nil {
// 			stack = append(stack, cur)
// 			cur = cur.Left
// 		} else {
// 			top := stack[len(stack)-1]
// 			res = append(res, top.Val)
// 			stack = stack[:len(stack)-1]
// 			cur = top.Right
// 		}
// 	}
// 	return res
// }

// // leetcode95
// func GenerateTrees(n int) []*algorithm.TreeNode {

// 	var generate func(start, end int) []*algorithm.TreeNode
// 	generate = func(start, end int) []*algorithm.TreeNode {
// 		if start > end {
// 			return []*algorithm.TreeNode{nil}
// 		}

// 		res := []*algorithm.TreeNode{}
// 		for i := start; i <= end; i++ {
// 			left := generate(start, i-1)
// 			right := generate(i+1, end)

// 			for _, leftTree := range left {
// 				for _, rightTree := range right {
// 					root := algorithm.TreeNode{Val: i}
// 					root.Left = leftTree
// 					root.Right = rightTree
// 					res = append(res, &root)
// 				}
// 			}
// 		}
// 		return res
// 	}

// 	return generate(1, n)
// }

// // leetcode96
// func NumTrees1(n int) int {
// 	ans := 1
// 	for i := 2; i <= n; i++ {
// 		ans = ans * (4*i - 2) / (i + 1)
// 	}
// 	return ans
// }

// func NumTrees2(n int) int {
// 	dp := make([]int, n+5)
// 	dp[0] = 1
// 	for i := 1; i <= n; i++ {
// 		for j := 0; j < i; j++ {
// 			// 左边j个节点，右边i-j-1个节点
// 			dp[i] += dp[j] * dp[i-j-1]
// 		}
// 	}
// 	return dp[n]
// }

// // leetcode98
// func IsValidBST(root *algorithm.TreeNode) bool {
// 	var prev *algorithm.TreeNode
// 	cur := root
// 	stack := []*algorithm.TreeNode{}
// 	for cur != nil || len(stack) != 0 {
// 		if cur != nil {
// 			stack = append(stack, cur)
// 			cur = cur.Left
// 		} else {
// 			top := stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]
// 			if prev != nil && cur.Val < prev.Val {
// 				return false
// 			}
// 			prev = cur
// 			cur = top.Right
// 		}
// 	}
// 	return true
// }

// // leetcode100
// func IsSameTree(p *algorithm.TreeNode, q *algorithm.TreeNode) bool {
// 	if p == nil && q != nil {
// 		return false
// 	}
// 	if p != nil && q == nil {
// 		return false
// 	}
// 	if p == nil && q == nil {
// 		return true
// 	}
// 	l := IsSameTree(p.Left, q.Left)
// 	r := IsSameTree(p.Right, q.Right)
// 	return p.Val == q.Val && l && r
// }
