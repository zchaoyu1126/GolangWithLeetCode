package leetcode

import (
	"fmt"
	"programs/internal/algorithmingo/algorithm"
	"programs/kit/utils"
	"sort"
	"strconv"
	"strings"
)

// leetcode404
func SumOfLeftLeaves(root *algorithm.TreeNode) int {
	sum := 0
	var dfs func(node *algorithm.TreeNode, flag int)
	dfs = func(node *algorithm.TreeNode, flag int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil && flag == 1 {
			sum += node.Val
			return
		}
		dfs(node.Left, 1)
		dfs(node.Right, 2)
	}
	dfs(root, 0)
	return sum
}

// leetcode412
func FizzBuzz(n int) []string {
	res := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			res = append(res, "FizzBuzz")
		} else if i%3 == 0 {
			res = append(res, "Fizz")
		} else if i%5 == 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}

// leetcode413
func NumberOfArithmeticSlices(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}
	pre, sum := 0, 0
	if nums[2]-nums[1] == nums[1]-nums[0] {
		pre += 1
		sum += 1
	}
	for i := 3; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			pre += pre + 1
			sum += pre
		} else {
			pre = 0
		}
	}
	return sum
}

// leetcode416
func CanPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	V := int(sum / 2)
	dp := make([]int, V+1)
	for i := range nums {
		for j := V; j >= nums[i]; j-- {
			dp[j] = utils.MaxNum(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[V] == V
}

// leetcode419
func CountBattleships(board [][]byte) int {
	m, n := len(board), len(board[0])
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'X' && !(i > 0 && board[i-1][j] == 'X' || j > 0 && board[i][j-1] == 'X') {
				ans++
			}
		}
	}
	return ans
}

// leetcode423
func OriginalDigits(s string) string {
	cntChar := map[byte]int{}
	var cntNum [10]int
	for i := 0; i < len(s); i++ {
		cntChar[s[i]]++
	}
	cntNum[0] = cntChar['z']
	cntNum[2] = cntChar['w']
	cntNum[4] = cntChar['u']
	cntNum[6] = cntChar['x']
	cntNum[8] = cntChar['g']

	cntNum[3] = cntChar['h'] - cntChar['g']
	cntNum[5] = cntChar['f'] - cntChar['u']
	cntNum[7] = cntChar['s'] - cntChar['x']

	cntNum[1] = cntChar['o'] - cntNum[0] - cntNum[2] - cntNum[4]
	cntNum[9] = cntChar['i'] - cntNum[5] - cntNum[6] - cntNum[8]

	res := ""
	for i := 0; i < 10; i++ {
		for j := 0; j < cntNum[i]; j++ {
			res += strconv.Itoa(i)
		}
	}
	fmt.Println(res)
	return res
}

// leetcode429
type Node429 struct {
	Val      int
	Children []*Node429
}

func LevelOrder429(root *Node429) [][]int {
	queue := []*Node429{}
	queue = append(queue, root)
	res := [][]int{}
	for len(queue) != 0 {
		size := len(queue)
		cur := []int{}
		for i := 0; i < size; i++ {
			front := queue[0]
			queue = queue[1:]
			cur = append(cur, front.Val)
			queue = append(queue, front.Children...)
		}
		res = append(res, cur)
	}
	return res
}

// leetcode430
func Flatten(root *algorithm.MultiListNode) *algorithm.MultiListNode {
	head, _ := flat(root)
	return head
}

func flat(root *algorithm.MultiListNode) (head, tail *algorithm.MultiListNode) {
	if root == nil {
		return nil, nil
	}
	cur := root
	h, t := root, root
	for cur != nil {
		if cur.Child != nil {
			head, tail := flat(cur.Child)
			cur.Child = nil
			originNext := cur.Next
			cur.Next = head
			head.Prev = cur
			tail.Next = originNext
			if originNext != nil {
				originNext.Prev = tail
			}
			t = tail
			cur = originNext
		} else {
			t = cur
			cur = cur.Next
		}
	}
	return h, t
}

// leetcode433
func MinMutation(start string, end string, bank []string) int {
	n := len(bank) + 1
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	bank = append(bank, start)
	diff := func(x, y string) int {
		res := 0
		for i := 0; i < len(x); i++ {
			if x[i] != y[i] {
				res++
			}
		}
		return res
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if diff(bank[i], bank[j]) == 1 {
				matrix[i][j] = 1
				matrix[j][i] = 1
			}
		}
	}

	queue := []int{n - 1}
	visited := make([]int, n)
	visited[n-1] = 1

	curSize, nextSize, step, find := 1, 0, -1, false
	for len(queue) != 0 && !find {
		step++
		for curSize != 0 {
			front := queue[0]
			visited[front] = 1
			if bank[front] == end {
				find = true
				break
			}
			queue = queue[1:]
			curSize--

			for i := 0; i < n; i++ {
				if matrix[front][i] == 1 && visited[i] == 0 {
					queue = append(queue, i)
					nextSize++
				}
			}
		}
		curSize = nextSize
		nextSize = 0
	}
	if step == 0 || !find {
		return -1
	}
	return step
}

// leetcode436
// 第一维是区间起点
// 第二维是下标信息
type IntervalInfo [][2]int

func (it IntervalInfo) Len() int      { return len(it) }
func (it IntervalInfo) Swap(i, j int) { it[i], it[j] = it[j], it[i] }
func (it IntervalInfo) Less(i, j int) bool {
	if it[i][0] == it[j][0] {
		return it[i][1] < it[j][1]
	}
	return it[i][0] < it[j][0]
}

func FindRightInterval(intervals [][]int) []int {
	n := len(intervals)
	info := make([][2]int, n)
	res := make([]int, n)
	for i := 0; i < len(intervals); i++ {
		info[i][0] = intervals[i][0]
		info[i][1] = i
	}
	sort.Sort(IntervalInfo(info))
	for i, interval := range intervals {
		r := interval[1]
		// 返回第一个使得f[i]=true的值，没有则返回n
		// 注意一个区间的右区间可以是它自己，这个题目里面没有说明
		idx := sort.Search(n, func(j int) bool {
			// return i != info[j][1] && info[j][0] >= r
			return info[j][0] >= r
		})
		if idx == n {
			res[i] = -1
		} else {
			res[i] = idx
		}
	}
	return res
}

// leetcode438
func FindAnagrams(s string, p string) []int {
	mp1, mp2 := [26]int{}, [26]int{}
	for i := 0; i < len(p); i++ {
		mp1[s[i]]++
		mp2[p[i]]++
	}
	res := []int{}
	if mp1 == mp2 {
		res = append(res, 0)
	}
	for i := len(p); i < len(s); i++ {
		mp1[s[i-1]]--
		mp1[s[i]]++
		if mp1 == mp2 {
			res = append(res, i-len(p)+1)
		}
	}
	return res
}

// leetcode441
func ArrangeCoins(n int) int {
	res := 1
	for (res*res+res)/2 < n {
		res <<= 1
	}
	if (res*res+res)/2 > n {
		l, r := res>>1, res
		for l < r {
			m := (l + r + 1) / 2
			if (m*m+m)/2 <= n {
				l = m
			} else {
				r = m - 1
			}
		}
		res = l
	}
	return res
}

// leetcode447
func NumberOfBoomerangs(points [][]int) int {
	ans := 0
	for _, p := range points {
		cnt := map[int]int{}
		for _, q := range points {
			dis := (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])
			cnt[dis]++
		}
		for _, m := range cnt {
			ans += m * (m - 1)
		}
	}
	return ans
}

// leetcode449
type Codec struct {
}

func NewCodec() Codec {
	return Codec{}
}
func (c *Codec) PrevOrder(root *algorithm.TreeNode, builder *strings.Builder) {
	if root == nil {
		return
	}
	str := strconv.Itoa(root.Val) + " "
	builder.WriteString(str)
	c.PrevOrder(root.Left, builder)
	c.PrevOrder(root.Right, builder)
}

func (c *Codec) BuildTree(preorder []int) *algorithm.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	val := preorder[0]
	node := &algorithm.TreeNode{Val: val}
	idx := sort.Search(len(preorder), func(x int) bool {
		return preorder[x] > val
	})
	node.Left = c.BuildTree(preorder[1:idx])
	node.Right = c.BuildTree(preorder[idx:])
	return node
}

// Serializes a tree to a single string.
func (c *Codec) Serialize(root *algorithm.TreeNode) string {
	var builder strings.Builder
	c.PrevOrder(root, &builder)
	return builder.String()
}

// Deserializes your encoded data to tree.
func (c *Codec) Deserialize(data string) *algorithm.TreeNode {
	if data == "" {
		return nil
	}
	data = strings.Trim(data, " ")
	vals := strings.Split(data, " ")
	nums := make([]int, 0, len(vals))
	for _, val := range vals {
		num, _ := strconv.Atoi(val)
		nums = append(nums, num)
	}
	return c.BuildTree(nums)
}

// leetcode450
func DeleteNode450(root *algorithm.TreeNode, key int) *algorithm.TreeNode {
	if root == nil {
		return root
	}
	if key > root.Val {
		root.Right = DeleteNode450(root.Right, key)
	} else if key < root.Val {
		root.Left = DeleteNode450(root.Left, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = root.Left
		}
		root = root.Right
	}
	return root
}
