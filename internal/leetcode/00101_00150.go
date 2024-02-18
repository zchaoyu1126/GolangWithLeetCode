package leetcode

import (
	"container/list"
	"programs/internal/algorithmingo/algorithm"
	"programs/kit/utils"
	"strconv"
)

// leetcode101
func CheckSymmetric(root1, root2 *algorithm.TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	return root1.Val == root2.Val && CheckSymmetric(root1.Left, root2.Right) && CheckSymmetric(root1.Right, root2.Left)
}

func IsSymmetric(root *algorithm.TreeNode) bool {
	if root == nil {
		return true
	}
	return CheckSymmetric(root.Left, root.Right)
}

// leetcode102
func LevelOrder(root *algorithm.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	cur, next := []*algorithm.TreeNode{root}, []*algorithm.TreeNode{}
	curVal := []int{}
	for len(cur) != 0 || len(next) != 0 {
		if len(cur) == 0 {
			res = append(res, curVal)
			curVal = []int{}
			cur = next
			next = []*algorithm.TreeNode{}
			continue
		}
		front := cur[0]
		cur = cur[1:]
		curVal = append(curVal, front.Val)
		if front.Left != nil {
			next = append(next, front.Left)
		}
		if front.Right != nil {
			next = append(next, front.Right)
		}
	}
	res = append(res, curVal)
	return res
}

// leetcode104
func MaxDepth104(root *algorithm.TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := MaxDepth104(root.Left)
	rightDepth := MaxDepth104(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func BuildTree105(preorder []int, inorder []int) *algorithm.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	target := preorder[0]
	idx := -1
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == target {
			idx = i
		}
	}
	node := &algorithm.TreeNode{Val: target}
	node.Left = BuildTree(preorder[1:1+idx], inorder[:idx])
	node.Right = BuildTree(preorder[1+idx:], inorder[idx+1:])
	return node
}

// leetcode106
func BuildTree106(inorder []int, postorder []int) *algorithm.TreeNode {
	//后序遍历，最后一个一定是顶点
	//找postorder[last]在inorder中的位置
	if len(postorder) <= 0 {
		return nil
	}
	target := postorder[len(postorder)-1]
	idx := -1
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == target {
			idx = i
		}
	}
	node := &algorithm.TreeNode{Val: target}
	node.Left = BuildTree(inorder[:idx], postorder[:idx])
	node.Right = BuildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])
	return node
}

// leetcode107
func LevelOrderBottom(root *algorithm.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
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
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

// leetcode108
func SortedArrayToBST(nums []int) *algorithm.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &algorithm.TreeNode{Val: nums[mid]}
	root.Left = SortedArrayToBST(nums[:mid])
	root.Right = SortedArrayToBST(nums[mid+1:])
	return root
}

// leetcode110
func GetHeight(root *algorithm.TreeNode) int {
	if root == nil {
		return 0
	}
	l := GetHeight(root.Left)
	r := GetHeight(root.Right)
	if l == -1 || r == -1 {
		return -1
	}
	if l-r >= 2 || r-l >= 2 {
		return -1
	}
	return utils.MaxNum(l, r) + 1
}

func IsBalanced(root *algorithm.TreeNode) bool {
	return GetHeight(root) != -1
}

// leetcode111
func MinDepth(root *algorithm.TreeNode) int {

	if root == nil {
		return 0
	}
	//res := INT_MAX
	left := MinDepth(root.Left)
	right := MinDepth(root.Right)
	if left != 0 && right == 0 {
		return left + 1
	}
	if right != 0 && left == 0 {
		return right + 1
	}
	if left < right {
		return left + 1
	}
	return right + 1
}

// leetcode112
func HasPathSum(root *algorithm.TreeNode, targetSum int) bool {
	res := false

	var dfs func(root *algorithm.TreeNode, sum int)
	dfs = func(root *algorithm.TreeNode, sum int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil && !res {
			res = targetSum == sum+root.Val
		}
		dfs(root.Left, sum+root.Val)
		dfs(root.Right, sum+root.Val)
	}
	dfs(root, 0)
	return res
}

// leetcode113
func PathSum(root *algorithm.TreeNode, targetSum int) [][]int {
	var dfs func(root *algorithm.TreeNode)
	curPath := []int{}
	curSum := 0
	res := [][]int{}
	dfs = func(root *algorithm.TreeNode) {
		if root == nil {
			return
		}
		curPath = append(curPath, root.Val)
		curSum += root.Val
		// 是叶子节点
		if root.Left == nil && root.Right == nil {
			if curSum == targetSum {
				aPath := make([]int, len(curPath))
				copy(aPath, curPath)
				res = append(res, aPath)
			}
		}
		dfs(root.Left)
		dfs(root.Right)
		curPath = curPath[:len(curPath)-1]
		curSum -= root.Val
	}
	dfs(root)
	return res
}

// leetcode115
func NumDistinct(s string, t string) int {
	// dp
	n, m := len(s), len(t)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][m]
}

// leetcode116
type Node116 struct {
	Val   int
	Left  *Node116
	Right *Node116
	Next  *Node116
}

func Connect(root *Node116) *Node116 {
	if root == nil {
		return nil
	}
	queue := []*Node116{root}
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			front := queue[0]
			queue = queue[1:]
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
			if i == size-1 {
				front.Next = nil
			} else {
				front.Next = queue[0]
			}
		}
	}
	return root
}

// leetcode117
func Connect117(root *Node116) *Node116 {
	if root == nil {
		return nil
	}
	cur := root
	for cur != nil {
		dummy := &Node116{}
		prev := dummy
		for cur != nil {
			if cur.Left != nil {
				prev.Next = cur.Left
				prev = cur.Left
			}
			if cur.Right != nil {
				prev.Next = cur.Right
				prev = cur.Right
			}
			cur = cur.Next
		}
		cur = dummy.Next
	}
	return root
}

// leetcode118
func Generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	res := make([][]int, 2)
	res[0] = []int{1}
	res[1] = []int{1, 1}

	for i := 3; i <= numRows; i++ {
		row := make([]int, i)
		row[0] = 1
		row[i-1] = 1
		for j := 1; j < i-1; j++ {
			row[j] = res[i-2][j] + res[i-2][j-1]
		}
		res = append(res, row)
	}
	return res
}

// leetcode119
func GetRow(rowIndex int) []int {
	cmn := func(m, n int) int {
		if n == 0 || m == 1 {
			return 1
		}
		if n > m/2 {
			n = m - n
		}
		x, y := 1, 1
		for i := 0; i < n; i++ {
			x *= (m - i)
			y *= (1 + i)
			if x%y == 0 {
				x /= y
				y = 1
			}
		}
		return x / y
	}
	res := []int{}
	for i := 0; i <= rowIndex; i++ {
		res = append(res, cmn(rowIndex, i))
	}
	return res
}

// 我觉得我是个呆逼 cao
// func getRow(rowIndex int) []int {
//     row := make([]int, rowIndex+1)
//     row[0] = 1
//     for i := 1; i <= rowIndex; i++ {
//         row[i] = row[i-1] * (rowIndex - i + 1) / i
//     }
//     return row
// }

// leetcode120
func MinimumTotal(triangle [][]int) int {
	prev, cur := []int{}, []int{}
	prev = append(prev, triangle[0][0])
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				cur = append(cur, triangle[i][j]+prev[j])
				continue
			} else if j == len(triangle[i])-1 {
				cur = append(cur, triangle[i][j]+prev[j-1])
				continue
			}
			cur = append(cur, triangle[i][j]+utils.MinNum(prev[j], prev[j-1]))
		}
		prev = cur
		cur = []int{}
	}
	res := 0xFFFFFFFF
	for i := 0; i < len(prev); i++ {
		if prev[i] < res {
			res = prev[i]
		}
	}
	return res
}

// leetcode121
func MaxProfit1BF(prices []int) int {
	res := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			res = utils.MaxNum(prices[j]-prices[i], res)
		}
	}
	return res
}

func MaxProfit1DP(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < n; i++ {
		dp[i][0] = utils.MaxNum(dp[i-1][0], -prices[i])
		dp[i][1] = utils.MaxNum(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[n-1][1]
}

// leetcode122
func MaxProfit2(prices []int) int {
	prices = append(prices, 0)
	curMinPrice := prices[0]
	res := 0
	for i := 1; i <= len(prices); i++ {
		if prices[i] >= prices[i-1] {
			// 单调递增的过程
			continue
		} else {
			res += prices[i-1] - curMinPrice
			curMinPrice = prices[i]
		}
	}
	return res
}

// leetcode123
func MaxProfit3(prices []int) int {
	// dp[i][j][0] 第j次交易 买入
	// dp[i][j][1] 第j次交易 卖出
	// 买卖股票的最佳时机Ⅲ的泛化版
	// 1.定义 dp[i][0/1][k]表示第i天(非)持有股票并且交易次数为k的利润最大值
	// 2.递推公式
	// 	2.1第i天没有持有股票
	// 		2.1.1第i-1天没有持股 dp[i][0][k] = dp[i-1][0][k]
	// 		2.1.2第i-1天持股    dp[i][0][k] = dp[i-1][1][k] + prices[i]
	// 	2.2第i天持有股票
	// 		2.2.1第i-1天持股    dp[i][1][k] = dp[i-1][1][k]
	// 		2.2.2第i-1天没有持股 dp[i][1][k] = dp[i-1][0][k-1] - prices[i]
	// 			 进行了一次新的交易买入股票 因此第k次交易的状态需要由k-1次交易的状态得到
	// 3.初始化
	// 	dp[0][0][1...k] = 0             //第1天没有买入股票利润值为0
	// 	dp[0][1][1...k] = -prices[0]    //第1天进行k次交易买入股票利润值为-prices[0]
	// 4.遍历顺序
	// 	外层循环遍历天数 从前往后遍历
	// 	内层循环穷举交易次数状态
	n := len(prices)
	dp := make([][][2]int, n)
	for i := range dp {
		dp[i] = make([][2]int, 3)
	}
	// j=1时，代表第一次交易
	dp[0][1][0] = 0          // 未持有股票
	dp[0][1][1] = -prices[0] // 持有股票
	dp[0][2][0] = 0
	dp[0][2][1] = -prices[0]
	for i := 1; i < n; i++ {
		for j := 1; j <= 2; j++ {
			// 第i天未持股票，之前一直没有或者今天把之前的卖了
			dp[i][j][0] = utils.MaxNum(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			// 第i天持有股票，之前一直有股票，或者是取消之前的买入操作，在股价更低的地方买入
			dp[i][j][1] = utils.MaxNum(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[n-1][2][0]
}

// leetcode131
func Partition(s string) [][]string {
	var backtrace func(start int, s string)
	isPalindrome := func(str string) bool {
		for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
			if str[i] != str[j] {
				return false
			}
		}
		return true
	}
	res := [][]string{}
	cur := []string{}
	backtrace = func(start int, s string) {
		if s == "" {
			tmp := make([]string, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
		}
		for i := start; i < len(s); i++ {
			if !isPalindrome(s[start : i+1]) {
				continue
			}
			cur = append(cur, s[start:i+1])
			backtrace(i+1, s[i+1:])
			cur = cur[:len(cur)-1]
		}
	}
	backtrace(0, s)
	return res
}

// leetcode138
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var cacheNode map[*Node]*Node

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	} else if value, ok := cacheNode[node]; ok {
		return value
	}
	newNode := &Node{Val: node.Val}
	cacheNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

func CopyRandomList2(head *Node) *Node {
	cacheNode = map[*Node]*Node{}
	return deepCopy(head)
}
func CopyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	mp1 := make(map[*Node]int)
	mp2 := make(map[int]*Node)

	n := head
	var pre, cur *Node
	newHead := &Node{}
	newHead.Val = head.Val

	cnt := 1
	mp1[head] = cnt
	mp2[cnt] = newHead

	pre = newHead
	n = n.Next
	for n != nil {
		cur = &Node{}
		cur.Val = n.Val
		cur.Next = nil
		cnt++
		mp1[n] = cnt
		mp2[cnt] = cur

		pre.Next = cur
		pre = cur
		n = n.Next
	}

	n = head
	for t := newHead; t != nil; t = t.Next {
		if n.Random != nil {
			pos := mp1[n.Random]
			t.Random = mp2[pos]
		} else {
			t.Random = nil
		}
		n = n.Next
	}
	return newHead
}

// leetcode139
func WordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < len(wordDict); j++ {
			if len(wordDict[j]) < i {
				// word的长度小于i 直接continue
				continue
			} else {
				// 是否相等？前面是否能被表示？
				length := len(wordDict[j])
				if dp[i-length] && s[i-length:i] == wordDict[j] {
					dp[i] = true
				}
			}
		}
	}
	return dp[len(s)]
}

// leetcode142
func DetectCycle(head *algorithm.ListNode) *algorithm.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

// leetcode144
func PreorderTraversal(root *algorithm.TreeNode) []int {
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

// leetcode145
func PostorderTraversal(root *algorithm.TreeNode) []int {
	var traversal func(root *algorithm.TreeNode)
	res := []int{}
	traversal = func(root *algorithm.TreeNode) {
		if root == nil {
			return
		}
		traversal(root.Left)
		traversal(root.Right)
		res = append(res, root.Val)
	}
	traversal(root)
	return res
}

// leetcode146

type LRUCache struct {
	l          *list.List
	curSize    int
	capability int
	mp         map[int]*list.Element
}

type Info struct {
	key int
	val int
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{list.New(), 0, capacity, make(map[int]*list.Element)}
}

func (cache *LRUCache) Get(key int) int {
	if node, has := cache.mp[key]; !has {
		return -1
	} else {
		cache.l.MoveToFront(node)
		return node.Value.(Info).val
	}
}

func (cache *LRUCache) Put(key int, value int) {
	var node *list.Element
	if _, has := cache.mp[key]; !has {
		// 如果不存在
		node = cache.l.PushFront(Info{key, value})
		cache.mp[key] = node
		if cache.curSize < cache.capability {
			cache.curSize++
		} else {
			tail := cache.l.Back()
			delete(cache.mp, tail.Value.(Info).key)
			cache.l.Remove(tail)
		}
	} else {
		// 如果已存在
		// 那么就不用考虑大小的问题了
		node = cache.mp[key]
		node.Value = Info{key, value}
		cache.l.MoveToFront(node)
	}
}

// leetcode150
func EvalRPN(tokens []string) int {
	stack := []int{}
	for i := 0; i < len(tokens); i++ {
		if val, ok := strconv.Atoi(tokens[i]); ok == nil {
			stack = append(stack, val)
		} else {
			num1 := stack[len(stack)-2]
			num2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch tokens[i] {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}
