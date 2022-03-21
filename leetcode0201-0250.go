package leetcode

import (
	"container/heap"
	"fmt"
	"programs/internal/algorithmingo/algorithm"
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
func SearchMatrix(matrix [][]int, target int) bool {
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
