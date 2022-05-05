package leetcode

import (
	"container/heap"
	"fmt"
	"programs/internal/algorithmingo/algorithm"
	"programs/kit/common"
)

// leetcode203
func RemoveElements(head *algorithm.ListNode, val int) *algorithm.ListNode {
	dummyHead := &algorithm.ListNode{}
	dummyHead.Next = head
	cur := dummyHead

	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}

// leetcode206
func ReverseList1(head *algorithm.ListNode) *algorithm.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	others := head.Next             // 得到尾节点
	newHead := ReverseList1(others) // 返回头节点
	others.Next = head
	head.Next = nil
	return newHead
}

func ReverseList2(head *algorithm.ListNode) *algorithm.ListNode {
	if head == nil {
		return nil
	}

	has := head         // 已经反转好的
	remain := head.Next // 没有反转的
	has.Next = nil      // 断开连接

	for remain != nil {
		top := remain        // 取出第一个
		remain = remain.Next // 取出第一个top
		top.Next = has       // 将top加到has链表头部
		has = top            // has更新
	}
	return has
}

// leetcode207
func CanFinish(numCourses int, prerequisites [][]int) bool {
	mp := make(map[int][]int)
	cnt := make([]int, numCourses) // 入度

	outCnt := 0
	// Create Directed Acyclic Graph
	for i := 0; i < len(prerequisites); i++ {
		x, y := prerequisites[i][0], prerequisites[i][1]
		if _, ok := mp[y]; !ok {
			mp[y] = []int{x}
		} else {
			mp[y] = append(mp[y], x)
		}
		cnt[x]++ // x的入度加一  学习x的前提是先学y  所以是y-->x
	}
	queue := []int{}
	for idx, value := range cnt {
		if value == 0 {
			queue = append(queue, idx)
		}
	}
	for len(queue) != 0 {
		top := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		outCnt++
		for i := 0; i < len(mp[top]); i++ {
			node := mp[top][i]
			cnt[node]--
			if cnt[node] == 0 {
				queue = append(queue, node)
			}
		}
	}
	return outCnt == numCourses
}

// leetcode209
func MinSubArrayLen(target int, nums []int) int {
	l, r, sum, res := 0, 0, 0, 0xFFFFFFFF
	for ; r < len(nums); r++ {
		sum += nums[r]
		for sum >= target {
			if res > r-l+1 {
				res = r - l + 1
			}
			sum -= nums[l]
			l++
		}
	}
	if res == 0xFFFFFFFF {
		return 0
	}
	return res
}

func FindOrder(numCourses int, prerequisites [][]int) []int {
	mp := make(map[int][]int)
	cnt := make([]int, numCourses) // 入度
	res := []int{}
	outCnt := 0
	// Create Directed Acyclic Graph
	for i := 0; i < len(prerequisites); i++ {
		x, y := prerequisites[i][0], prerequisites[i][1]
		if _, ok := mp[y]; !ok {
			mp[y] = []int{x}
		} else {
			mp[y] = append(mp[y], x)
		}
		cnt[x]++ // x的入度加一  学习x的前提是先学y  所以是y-->x
	}
	queue := []int{}
	for idx, value := range cnt {
		if value == 0 {
			queue = append(queue, idx)
		}
	}
	for len(queue) != 0 {
		top := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		res = append(res, top)
		outCnt++
		for i := 0; i < len(mp[top]); i++ {
			node := mp[top][i]
			cnt[node]--
			if cnt[node] == 0 {
				queue = append(queue, node)
			}
		}
	}
	if outCnt == numCourses {
		return res
	}
	return []int{}
}

// leetcode 211
type TrieNode struct {
	Num   int
	Son   map[byte]*TrieNode
	IsEnd bool
	Val   byte
}
type WordDictionary struct {
	root *TrieNode
}

func NewWordDictionary() WordDictionary {
	return WordDictionary{root: &TrieNode{Son: make(map[byte]*TrieNode)}}
}

func (t *WordDictionary) AddWord(word string) {
	node := t.root
	if word == "" {
		return
	}
	bytes := []byte(word)
	for _, val := range bytes {
		pos := val - 'a'
		if node.Son[pos] == nil {
			node.Son[pos] = &TrieNode{Son: make(map[byte]*TrieNode)}
			node.Son[pos].Val = val
		} else {
			node.Son[pos].Num++
		}
		node = node.Son[pos]
	}
	node.IsEnd = true
}

func (t *WordDictionary) Search(word string) bool {
	var dfs func(int, *TrieNode) bool
	dfs = func(index int, node *TrieNode) bool {
		if index == len(word) {
			return node.IsEnd
		}
		ch := word[index]
		if ch != '.' {
			child := node.Son[ch-'a']
			if child != nil && dfs(index+1, child) {
				return true
			}
		} else {
			for i := range node.Son {
				child := node.Son[i]
				if child != nil && dfs(index+1, child) {
					return true
				}
			}
		}
		return false
	}
	return dfs(0, t.root)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

// leetcode212
type Cord struct {
	x, y int
}

func FindWords(board [][]byte, words []string) []string {
	// timeout at the last test case
	res := []string{}
	var dfs func(int, int, int, string) bool
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}

	dfs = func(x, y int, cur int, target string) bool {
		findRes := false
		if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || board[x][y] == '#' {
			return false
		}
		if cur == len(target)-1 && target[cur] == board[x][y] {
			res = append(res, target)
			return true
		}
		if target[cur] == board[x][y] {
			origin := board[x][y]
			board[x][y] = '#'
			for i := 0; i < 4; i++ {
				if x+dx[i] < 0 || x+dx[i] >= len(board) || y+dy[i] < 0 || y+dy[i] >= len(board[0]) || board[x+dx[i]][y+dy[i]] == '#' {
					continue
				}
				tmp := dfs(x+dx[i], y+dy[i], cur+1, target)
				if tmp {
					findRes = tmp
					break
				}
			}
			board[x][y] = origin
		}
		return findRes
	}

	mp := map[byte][]Cord{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			mp[board[i][j]] = append(mp[board[i][j]], Cord{i, j})
		}
	}

	for _, word := range words {
		for i := 0; i < len(mp[word[0]]); i++ {
			start := mp[word[0]][i]
			findRes := dfs(start.x, start.y, 0, word)
			if findRes {
				break
			}
		}
	}
	return res
}

// leetcode213
func RobII(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	} else if n == 2 {
		return common.LargerNumber(nums[0], nums[1])
	}
	return common.LargerNumber(Rob(nums[0:n-1]), Rob(nums[1:]))
}

// leetcode216
func CombinationSum3(k int, n int) [][]int {
	res := [][]int{}
	cur := []int{}
	var backtrace func(start, target int)
	backtrace = func(start, target int) {
		if target == 0 {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
		} else if target < 0 {
			return
		}
		for i := start; i <= 9; i++ {
			cur = append(cur, i)
			backtrace(i+1, target-i)
			cur = cur[:len(cur)-1]
		}
	}
	return res
}

// leetcode217
func ContainsDuplicate(nums []int) bool {
	mp := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, has := mp[nums[i]]; has {
			return false
		} else {
			mp[nums[i]] = true
		}
	}
	return true
}

// leetcode222
func CountNodes(root *algorithm.TreeNode) int {
	if root == nil {
		return 0
	}
	leftNum, rightNum := 0, 0
	left, right := root.Left, root.Right
	for left != nil {
		left = left.Left
		leftNum++
	}
	for right != nil {
		right = right.Right
		rightNum++
	}
	if leftNum == rightNum {
		return (2 << leftNum) - 1
	}
	return CountNodes(root.Left) + CountNodes(root.Right) + 1
}

// leetcode225
type MyStack struct {
	queue []int
	cnt   int
}

func NewMyStack() MyStack {
	return MyStack{[]int{}, 0}
}

func (ms *MyStack) Push(x int) {
	ms.cnt++
	ms.queue = append(ms.queue, x)
}

func (ms *MyStack) Pop() int {
	tmp := []int{}
	for i := 0; i < cnt-1; i++ {
		tmp = append(tmp, ms.queue[0])
		ms.queue = ms.queue[1:]
	}
	top := ms.queue[0]
	ms.cnt--
	ms.queue = tmp
	return top
}

func (ms *MyStack) Top() int {
	return ms.queue[cnt-1]
}

func (ms *MyStack) Empty() bool {
	return ms.cnt == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// leetcode226
func InvertTree(root *algorithm.TreeNode) *algorithm.TreeNode {
	if root == nil {
		return nil
	}
	left := root.Left
	right := root.Right
	root.Left = InvertTree(right)
	root.Right = InvertTree(left)
	return root
}

// leetcode229
func MajorityElement2(nums []int) []int {
	cnt1, cnt2 := 0, 0
	num1, num2 := 0, 0
	for i := 0; i < len(nums); i++ {
		if cnt1 > 0 && nums[i] == num1 {
			cnt1++
		} else if cnt2 > 0 && nums[i] == num2 {
			cnt2++
		} else if cnt1 == 0 {
			num1 = nums[i]
			cnt1++
		} else if cnt2 == 0 {
			num2 = nums[i]
			cnt2++
		} else {
			cnt1--
			cnt2--
		}
	}
	vote1, vote2 := 0, 0
	for i := 0; i < len(nums); i++ {
		if cnt1 > 0 && nums[i] == num1 {
			vote1++
		}
		if cnt2 > 0 && nums[i] == num2 {
			vote2++
		}
	}
	res := []int{}
	if vote1 >= len(nums)/3 {
		res = append(res, num1)
	}
	if vote2 >= len(nums)/3 {
		res = append(res, num2)
	}
	return res
}

// leetcode230
var cnt int
var resKthSamllest int

func KthSmallest(root *algorithm.TreeNode, k int) int {
	cnt = 0
	InOrder(root, k)
	return resKthSamllest
}

func InOrder(root *algorithm.TreeNode, k int) {
	if root == nil {
		return
	}
	InOrder(root.Left, k)
	fmt.Println(root.Val)
	cnt++
	if cnt == k {
		resKthSamllest = root.Val
	}
	InOrder(root.Right, k)
}

// leetcode232
// 用栈模拟队列
type MyQueue struct {
	stackIn  []int
	stackOut []int
	cnt      int
}

func NewMyQueue() MyQueue {
	return MyQueue{[]int{}, []int{}, 0}
}

func (mq *MyQueue) Push(x int) {
	mq.stackIn = append(mq.stackIn, x)
	mq.cnt++
}

func (mq *MyQueue) Pop() int {
	if len(mq.stackIn) == 0 && len(mq.stackOut) == 0 {
		return -1
	}
	if len(mq.stackOut) == 0 {
		for len(mq.stackIn) != 0 {
			top := mq.stackIn[len(mq.stackIn)-1]
			mq.stackOut = append(mq.stackOut, top)
			mq.stackIn = mq.stackIn[:len(mq.stackIn)-1]
		}
	}
	top := mq.stackOut[len(mq.stackOut)-1]
	mq.stackOut = mq.stackOut[:len(mq.stackOut)-1]
	mq.cnt--
	return top
}

func (mq *MyQueue) Peek() int {
	if len(mq.stackIn) == 0 && len(mq.stackOut) == 0 {
		return -1
	}
	if len(mq.stackOut) == 0 {
		for len(mq.stackIn) != 0 {
			top := mq.stackIn[len(mq.stackIn)-1]
			mq.stackOut = append(mq.stackOut, top)
			mq.stackIn = mq.stackIn[:len(mq.stackIn)-1]
		}
	}
	top := mq.stackOut[len(mq.stackOut)-1]
	return top
}

func (mq *MyQueue) Empty() bool {
	return mq.cnt == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

// leetcode233
func pow10(num int) (res int) {
	res = 1
	for i := 0; i < num; i++ {
		res *= 10
	}
	return
}

func CountDigitOne(n int) int {
	mp := map[int]int{}
	res := 0
	for i := 1; i <= 9; i++ {
		key := pow10(i)
		value := i * pow10(i-1)
		mp[key] = value
	}
	for i := 9; i >= 1; i-- {
		t := pow10(i)
		remain := n / t
		if remain > 0 {
			res += (remain * mp[t])
		} else {
			continue
		}

		if remain == 1 {
			res += n%t + 1
		} else if remain > 1 {
			res += t
		}
		n %= t
	}
	if n >= 1 {
		res += 1
	}
	return res
}

// leetcode235
func LowestCommonAncestor1(root, p, q *algorithm.TreeNode) *algorithm.TreeNode {
	var postTraversal func(node, p, q *algorithm.TreeNode) *algorithm.TreeNode
	postTraversal = func(node, p, q *algorithm.TreeNode) *algorithm.TreeNode {
		if node == nil || node.Val == p.Val || node.Val == q.Val {
			return node
		}
		l := postTraversal(node.Left, p, q)
		r := postTraversal(node.Right, p, q)
		if l == nil && r == nil {
			return nil
		}
		if l != nil && r != nil {
			return node
		}
		if l != nil {
			return l
		}

		return r

	}
	return postTraversal(root, p, q)
}

func LowestCommonAncestor235(root, p, q *algorithm.TreeNode) *algorithm.TreeNode {
	cur := root
	for cur != nil {
		if cur.Val > p.Val && cur.Val > q.Val {
			cur = cur.Left
		} else if cur.Val < p.Val && cur.Val < q.Val {
			cur = cur.Right
		} else {
			// else的情况中，说明有p.Val==cur.Val 或者q.Val == cur.Val
			// 或者 p.Val < cur.Val < q.Val
			// 或者 q.Val < cur.Val < p.Val
			return cur
		}
	}
	return nil
}

// leetcode236
func LowestCommonAncestor(root, p, q *algorithm.TreeNode) *algorithm.TreeNode {
	//  使用DFS遍历，时间复杂度O(N)，空间复杂度O(N)
	//              1
	//         2          3
	//      4    5     *6   7
	//    8  9 *10 11 12 13 14 15
	// 返回不为空(nil) 一定是说明找了某个节点
	// 在此，5的左子树找到了，所以其返回值非空，2的返回值非空
	// 3的返回值非空，由于2，3的返回值非空，所以其最近公共祖先为1
	var dfs func(root, p, q *algorithm.TreeNode) *algorithm.TreeNode
	dfs = func(root, p, q *algorithm.TreeNode) *algorithm.TreeNode {
		if root == nil || root.Val == p.Val || root.Val == q.Val {
			return root
		}
		left := dfs(root.Left, p, q)
		right := dfs(root.Right, p, q)
		// 一共四种情况
		// 在左、右子树中找到了p或q，一边一个的情况
		// 说明当前节点就是其最近公共祖先
		if left != nil && right != nil {
			return root
		}
		// left无right有，那么返回right，return到上一层
		if left == nil {
			// 右子树中含有目标节点
			return right
		}
		// left有right无，那么返回left，返回到上一层
		if right == nil {
			return left
		}
		// 都没找到直接返回nil
		return nil
	}
	return dfs(root, p, q)
}
func LowestCommonAncestor2(root, p, q *algorithm.TreeNode) *algorithm.TreeNode {
	// 使用hash表存储父节点信息
	parent := make(map[int]*algorithm.TreeNode)
	visited := make(map[int]bool)
	var dfs func(root *algorithm.TreeNode)
	dfs = func(root *algorithm.TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			parent[root.Left.Val] = root
		}
		if root.Right != nil {
			parent[root.Right.Val] = root
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	for p != nil {
		visited[p.Val] = true
		p = parent[p.Val]
	}

	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}
	return nil
}

func LowestCommonAncestor3(root, p, q *algorithm.TreeNode) *algorithm.TreeNode {
	// 倍增思想前篇
	// 首先需要遍历一遍，获得father和depth
	depth := make(map[int]int)
	father := make(map[int]*algorithm.TreeNode)
	var dfs func(root *algorithm.TreeNode, d int)
	dfs = func(root *algorithm.TreeNode, d int) {
		if root == nil {
			return
		}
		if root.Left != nil {
			father[root.Left.Val] = root
			depth[root.Left.Val] = d + 1
			dfs(root.Left, d+1)
		}
		if root.Right != nil {
			father[root.Right.Val] = root
			depth[root.Right.Val] = d + 1
			dfs(root.Right, d+1)
		}
	}
	dfs(root, 0)
	// 始终保持p的深度小
	if depth[p.Val] > depth[q.Val] {
		p, q = q, p
	}
	for depth[q.Val] > depth[p.Val] {
		q = father[q.Val]
	}
	//现在p,q深度相同， 同时往上
	for p.Val != q.Val {
		p = father[p.Val]
		q = father[q.Val]
	}
	return p
}

//倍增法 https://www.cnblogs.com/darlingroot/p/10597611.html
func LowestCommonAncestor4(root, p, q *algorithm.TreeNode) *algorithm.TreeNode {
	// 节点数目2~10^5
	// 步长设置为20
	var dfs func(root *algorithm.TreeNode, d int)
	father := make(map[int][]*algorithm.TreeNode)

	depth := make(map[int]int)
	// 在dfs中进行预处理
	dfs = func(root *algorithm.TreeNode, d int) {
		if root == nil {
			return
		}
		if root.Left != nil {
			// 如果没分配就分配
			if _, ok := father[root.Left.Val]; !ok {
				father[root.Left.Val] = make([]*algorithm.TreeNode, 20)
			}
			// 初始化
			father[root.Left.Val][0] = root
			for i := 0; i < 19; i++ {
				if father[root.Left.Val][i] == nil {
					father[root.Left.Val][i+1] = nil
					continue
				}
				father[root.Left.Val][i+1] = father[father[root.Left.Val][i].Val][i]
			}

			depth[root.Left.Val] = d + 1
			dfs(root.Left, d+1)
		}
		if root.Right != nil {
			if _, ok := father[root.Right.Val]; !ok {
				father[root.Right.Val] = make([]*algorithm.TreeNode, 20)
			}
			father[root.Right.Val][0] = root
			for i := 0; i < 19; i++ {
				if father[root.Right.Val][i] == nil {
					father[root.Right.Val][i+1] = nil
					continue
				}
				father[root.Right.Val][i+1] = father[father[root.Right.Val][i].Val][i]
			}
			depth[root.Right.Val] = d + 1
			dfs(root.Right, d+1)
		}
	}
	// father[i][j] 这个节点2^(j)次能跳到哪？

	father[root.Val] = make([]*algorithm.TreeNode, 20)
	for i := 0; i < 20; i++ {
		father[root.Val][i] = nil
	}

	dfs(root, 0)
	// 保证p的深度永远是小的
	if depth[p.Val] > depth[q.Val] {
		p, q = q, p
	}
	// 将q往上
	for i := 19; i >= 0; i-- {
		if father[q.Val][i] == nil {
			continue
		}
		if depth[father[q.Val][i].Val] >= depth[p.Val] {
			q = father[q.Val][i]
		}
		if p.Val == q.Val {
			return p
		}
	}
	// 现在深度一致了
	for i := 19; i >= 0; i-- {
		if father[p.Val][i] == nil && father[q.Val][i] == nil {
			continue
		}
		// 如果前面计算的没错，这里是不会出现一个为nil，另一个不为nil的情况
		if father[p.Val][i] != father[q.Val][i] {
			p = father[p.Val][i]
			q = father[q.Val][i]
		}
	}
	return father[p.Val][0]
}

// leetcode237

func DeleteNode(node *algorithm.ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// leetcode239
type info struct {
	val int
	idx int
}
type hp239 []info

func (h hp239) Len() int            { return len(h) }
func (h hp239) Less(i, j int) bool  { return h[i].val < h[j].val }
func (h hp239) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp239) Push(x interface{}) { *h = append(*h, x.(info)) }
func (h *hp239) Pop() interface{}   { top := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return top }
func (h *hp239) Top() interface{}   { return (*h)[len(*h)-1] }

func MaxSlidingWindow(nums []int, k int) []int {
	h := hp239{}
	for i := 0; i < k-1; i++ {
		h = append(h, info{nums[i], i})
	}
	heap.Init(&h)
	res := []int{}

	for i := k - 1; i < len(nums); i++ {
		heap.Push(&h, info{nums[i], i})
		for h.Top().(info).idx <= i-k {
			heap.Pop(&h)
		}
		res = append(res, h.Top().(info).val)
	}
	return res
}

func MaxSlidingWindow2(nums []int, k int) []int {
	queue, ans := []int{}, []int{}
	push := func(x int) {
		for len(queue) > 0 && queue[len(queue)-1] < nums[x] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, x)
	}
	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			push(i)
		} else if i >= k-1 {
			for len(queue) > 0 && queue[0] <= i-k {
				queue = queue[1:]
			}
			ans = append(ans, queue[0])
			push(nums[i])
		}
	}
	return ans
}

// leetcide240
func SearchMatrixII(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	lastRow := matrix[rows-1]
	index, isfind := binarySearch(lastRow, target)
	if !isfind {
		if index == cols {
			// 全部都比target小
			return false
		}
		for i := index; i < cols; i++ {
			for j := 0; j < rows; j++ {
				if matrix[j][i] == target {
					return true
				}
			}
		}
		return false
	}
	return true
}

func binarySearch(arr []int, num int) (int, bool) {
	l, r := 0, len(arr)-1
	for l <= r {
		m := (l + r) / 2
		if arr[m] < num {
			l = m + 1
		} else if arr[m] > num {
			r = m - 1
		} else {
			return m, true
		}
	}
	return l, false
}

func SearchMatrix2(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	x, y := 0, cols-1
	for x < rows-1 && y >= 0 {
		if target < matrix[x][y] {
			y--
		} else if target > matrix[x][y] {
			x++
		} else {
			return true
		}
	}
	return false
}

// leetcode242
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mp := make([]int, 26)
	for i := 0; i < len(s); i++ {
		mp[int(s[i]-'a')]++
		mp[int(t[i]-'a')]--
	}
	for i := 0; i < 26; i++ {
		if mp[i] != 0 {
			return false
		}
	}
	return true
}
