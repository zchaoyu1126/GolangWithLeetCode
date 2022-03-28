package leetcode

import (
	"fmt"
	"programs/internal/algorithmingo/algorithm"
	"programs/kit/common"
)

// leetcode53
// 最大子序和
// 法一：暴力求解，求前缀和，然后使用二重循环遍历各个子区间
func MaxSubArray(nums []int) int {
	sum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	max := -0xFFFFFFFF
	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= i; j++ {
			if sum[i]-sum[j-1] > max {
				max = sum[i] - sum[j-1]
			}
		}
	}
	fmt.Println(sum)
	return max
}

// 法二：DP，状态转移方程为f[i] = largerNumber(f[i-1]+nums[i], nums[i])
// 其中f[i]表示以第i个数结尾时，得到的最大子序列和, 在下述代码中i的取值范围为1~len(nums)
// 所求结果为 max(f[1], f[2], f[3]... f[len(nums)])
func MaxSubArray2(nums []int) int {
	dp := make([]int, len(nums)+1)
	res := -0xFFFFFFFF
	for i := 1; i <= len(nums); i++ {
		dp[i] = common.LargerNumber(dp[i-1]+nums[i-1], nums[i-1])
		res = common.LargerNumber(res, dp[i])
	}
	fmt.Println(dp)
	return dp[len(nums)]
}

// 法三：分治 线段树的思想

func MaxSubArray3(nums []int) int {
	return get(nums, 0, len(nums)-1).mSum
}

func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := common.LargerNumber(l.iSum+r.lSum, l.lSum)
	rSum := common.LargerNumber(l.rSum+r.iSum, r.rSum)
	mSum := common.LargerNumber(l.rSum+r.lSum, common.LargerNumber(l.mSum, r.mSum))
	return Status{lSum: lSum, rSum: rSum, mSum: mSum, iSum: iSum}
}

func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m+1, r)
	return pushUp(lSub, rSub)
}

type Status struct {
	lSum, rSum, mSum, iSum int
}

// leetcode54
func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	top, bottom := 0, m-1
	left, right := 0, n-1
	num, total := 0, m*n
	res := make([]int, 0)

	for num < total {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
			num++
		}
		top++

		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
			num++
		}
		right--

		fmt.Println(num, total)
		if num == total {
			break
		}

		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
			num++
		}
		bottom--

		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
			num++
		}
		left++
	}
	return res
}

// leetcode58
func LengthOfLastWord(s string) int {
	start := len(s) - 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			continue
		} else {
			start = i
			break
		}
	}
	cnt := 0
	for i := start; i >= 0; i-- {
		if s[i] == ' ' {
			break
		} else {
			cnt++
		}
	}
	return cnt
}

// leetcode66
func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] += 1
			return digits
		}
		digits[i] = 0
	}
	res := make([]int, len(digits)+1)
	res = append(res, 1)
	res = append(res, digits...)
	return res
	//  4 5 6  459  460
}

// leetcode67
func AddBinary(a string, b string) string {
	n, m := len(a), len(b)
	res := ""
	flag := 0
	i, j := n-1, m-1
	for i >= 0 && j >= 0 {
		tmp := int(a[i]-'0') + int(b[j]-'0') + flag
		if tmp >= 2 {
			flag = 1
			res = string('0'+byte(tmp-2)) + res
		} else {
			res = string('0'+byte(tmp)) + res
			flag = 0
		}
		i--
		j--
	}

	for i >= 0 {
		tmp := int(a[i]-'0') + flag
		if tmp >= 2 {
			flag = 1
			res = string('0'+byte(tmp-2)) + res
		} else {
			res = string('0'+byte(tmp)) + res
			flag = 0
		}
		i--
	}
	for j >= 0 {
		tmp := int(b[j]-'0') + flag
		if tmp >= 2 {
			flag = 1
			res = string('0'+byte(tmp-2)) + res
		} else {
			res = string('0'+byte(tmp)) + res
			flag = 0
		}
		j--
	}
	if flag == 1 {
		res = "1" + res
	}
	return res
}

// leetcode68
func FullJustify(words []string, maxWidth int) []string {
	oneline := []string{}
	length := 0
	res := []string{}
	for i := 0; i < len(words); i++ {
		if len(words[i])+length < maxWidth {
			fmt.Println(words[i], "test1")
			oneline = append(oneline, words[i])
			length += len(words[i])
			length++
		} else {
			fmt.Println(words[i], "test2", len(words[i])+length, length)
			tmp := ""
			for j := 0; j < len(oneline); j++ {
				tmp += oneline[j]
				tmp += " "
			}
			res = append(res, tmp)
			oneline = []string{}
			oneline = append(oneline, words[i])
			length = len(words)
		}
	}
	tmp := ""
	for j := 0; j < len(oneline); j++ {
		tmp += oneline[j]
		tmp += " "
	}
	res = append(res, tmp)
	return res
}

// leetcode69
func MySqrt(x int) int {
	l, r := 0, x
	// 找最后一个小于等于target的目标
	for l <= r {
		m := (l + r) / 2
		if m*m <= x {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return r
}

// leetcode70
func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	pre1, pre2 := 1, 2
	for i := 3; i <= n; i++ {
		tmp := pre2
		pre2 = pre1 + pre2
		pre1 = tmp
	}
	return pre2
}

// leetcode76
func MinWindow1(s string, t string) string {
	res := []byte{}
	ans := s
	mp := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		mp[t[i]]++
	}

	check := func() bool {
		for _, val := range mp {
			if val > 0 {
				return false
			}
		}
		return true
	}

	for l, r := 0, 0; r < len(s); r++ {
		res = append(res, s[r])
		if _, ok := mp[s[r]]; ok {
			mp[s[r]]--
			for check() {
				if len(string(res)) < len(ans) {
					ans = string(res)
				}
				res = res[1:]
				if _, ok := mp[s[l]]; ok {
					mp[s[l]]++
				}
				l++
			}
		}
	}
	return ans
}

// 更好的写法
func MinWindow2(s string, t string) string {
	window := make(map[byte]int, 0)
	need := make(map[byte]int, 0)

	left, match := -1, 0
	start, end, min := 0, 0, len(s)+1

	for i := range t {
		need[t[i]]++
	}

	for right := 0; right < len(s); right++ {
		// 1. 直接将s[right]加入到区间，形成（left, right]
		ch1 := s[right]
		window[ch1]++

		//  2. 更新状态
		if window[ch1] == need[ch1] {
			match++
		}

		// 3. 超出区间，或者满足条件
		for match == len(need) {
			if right-left < min {
				start, end = left, right
				min = right - left
			}

			// 4. 移除s[++left]，更新状态
			left++
			ch2 := s[left]
			if window[ch2] == need[ch2] {
				match--
			}
			window[ch2]--
		}
	}

	return s[start+1 : end+1]
}

// leetcode88
func Merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	cur := m + n - 1
	for i >= 0 && j >= 0 {
		if nums2[j] > nums1[i] {
			nums1[cur] = nums2[j]
			j--
		} else {
			nums1[cur] = nums1[i]
			i--
		}
		cur--
	}
	for i >= 0 {
		nums1[cur] = nums1[i]
		cur--
		i--
	}
	for j >= 0 {
		nums1[cur] = nums2[j]
		j--
		cur--
	}
	fmt.Println(nums1)
}

// leetcode92
func ReverseBetween(head *algorithm.ListNode, left int, right int) *algorithm.ListNode {
	// 注：1<=left<=right<=n
	dummy := &algorithm.ListNode{}
	dummy.Next = head
	cur := dummy
	guard := dummy
	for i := 0; i < left-1; i++ {
		cur = cur.Next
		guard = guard.Next
	}
	// guard用于连接反转后的头节点
	// left节点前一个 cur
	// 开始反转 cur.Next是要反转的目标

	// cur.Next不可能为nil, cur.Next.Next成立不会越界
	last, has := cur.Next, cur.Next
	remain := cur.Next.Next
	has.Next = nil
	// 极端情况right=left=n, 不会执行这个循环
	// 此时has = last node, remain = nil
	for i := 0; i < right-left; i++ {
		// 将第一个节点放入has中
		tmp := remain.Next
		remain.Next = has
		has = remain
		// 更新remain
		remain = tmp
	}
	// 剩下的节点
	guard.Next = has
	remain.Next = last
	return dummy.Next
}

// leetcode94
func InorderTraversal1(root *algorithm.TreeNode) []int {
	var traversal func(root *algorithm.TreeNode)
	res := []int{}
	traversal = func(root *algorithm.TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		traversal(root.Left)
		traversal(root.Right)
	}
	traversal(root)
	return res
}

func InorderTraversal2(root *algorithm.TreeNode) []int {
	stack := []*algorithm.TreeNode{}
	res := []int{}
	cur := root
	// 先找到最左边的节点
	for cur != nil || len(stack) != 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			top := stack[len(stack)-1]
			res = append(res, top.Val)
			stack = stack[:len(stack)-1]
			cur = top.Right
		}
	}
	return res
}

// leetcode95
func GenerateTrees(n int) []*algorithm.TreeNode {

	var generate func(start, end int) []*algorithm.TreeNode
	generate = func(start, end int) []*algorithm.TreeNode {
		if start > end {
			return []*algorithm.TreeNode{nil}
		}

		res := []*algorithm.TreeNode{}
		for i := start; i <= end; i++ {
			left := generate(start, i-1)
			right := generate(i+1, end)

			for _, leftTree := range left {
				for _, rightTree := range right {
					root := algorithm.TreeNode{Val: i}
					root.Left = leftTree
					root.Right = rightTree
					res = append(res, &root)
				}
			}
		}
		return res
	}

	return generate(1, n)
}

// leetcode96
func NumTrees(n int) int {
	ans := 1
	for i := 2; i <= n; i++ {
		ans = ans * (4*i - 2) / (i + 1)
	}
	return ans
}

// leetcode98
func IsValidBST(root *algorithm.TreeNode) bool {
	var prev *algorithm.TreeNode
	cur := root
	stack := []*algorithm.TreeNode{}
	for cur != nil || len(stack) != 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if prev != nil && cur.Val < prev.Val {
				return false
			}
			prev = cur
			cur = top.Right
		}
	}
	return true
}

// leetcode100
func IsSameTree(p *algorithm.TreeNode, q *algorithm.TreeNode) bool {
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	if p == nil && q == nil {
		return true
	}
	l := IsSameTree(p.Left, q.Left)
	r := IsSameTree(p.Right, q.Right)
	return p.Val == q.Val && l && r
}
