package leetcode

import (
	"fmt"
	"programs/internal/algorithmingo/algorithm"
	"programs/kit/utils"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// leetcode701
func InsertIntoBST(root *algorithm.TreeNode, val int) *algorithm.TreeNode {
	if root == nil {
		return &algorithm.TreeNode{Val: val}
	}
	cur := root
	prev := cur
	for cur != nil {
		if cur.Val < val {
			prev = cur
			cur = cur.Right
		} else if cur.Val > val {
			prev = cur
			cur = cur.Left
		}
	}
	if prev.Val < val {
		prev.Right = &algorithm.TreeNode{Val: val}
	} else if prev.Val > val {
		prev.Left = &algorithm.TreeNode{Val: val}
	}
	return root
}

// leetcode704
func Search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

// leetcode707
// type MyLinkedList struct {
//     Size		int
//     DummyHead 	*ListNode
// }

// type ListNode struct {
//     Val		int
//     Next 	*ListNode
// }

// func NewMyLinkedList() MyLinkedList {
//     return MyLinkedList{0, &ListNode{-1, nil}}
// }

// func (l *MyLinkedList) AddAtHead(val int) {
//     node := &ListNode{Val:val, l.DummyHead.Next}
//     l.DummyHead.Next = node
//     l.Size++
// }

// func (l *MyLinkedList) AddAtTail(val int) {
//     cur := l.DummyHead
//     for cur.Next != nil {
//         cur = cur.Next
//     }
//     cur.Next = &ListNode{Val:val}
//     l.Size++
// }

// func (l *MyLinkedList) AddAtIndex(idx, val int) {
//     cur := l.DummyHead
//     if idx <= 0 {
//         l.AddAtHead(val)
//     } else if idx > l.Size {
//         return
//     } else if idx == l.Size {
//         l.AddAtTail(val)
//     } else {
//         for cur.Next != nil && idx > 0 {
//             cur = cur.Next
//             idx--
//         }
//         node := &ListNode{val, cur.Next}
//         cur.Next = node
//         l.Size++
//     }
// }

// func (l *MyLinkedList) Get(idx int) int {
//     if idx < 0 || idx >= l.Size {
//         return -1
//     }
//     cur := l.DummyHead
//     for cur.Next != nil && idx >= 0{
//         cur = cur.Next
//         idx--
//     }
//     return cur.Val
// }

// func (l *MyLinkedList) Len() int {
//     cur := l.DummyHead
//     res := 0
//     for cur.Next != nil {
//         cur = cur.Next
//         res++
//     }
//     return res
// }

// func (l *MyLinkedList) DeleteAtIndex(idx int) {
//     if idx < 0 || idx >= l.Size {
//         return
//     }
//     cur := l.DummyHead
//     for cur.Next != nil && idx > 0 {
//         cur = cur.Next
//         idx--
//     }
//     cur.Next = cur.Next.Next
//     l.Size--
// }

// func (l *MyLinkedList) DeleteVal(val int) {
//     cur := l.DummyHead
//     for cur.Next != nil {
//         if cur.Next.Val == val {
//             cur.Next = cur.Next.Next
//         } else {
//             cur = cur.Next
//         }
//     }
// }

type MyLinkedList struct {
	Size      int
	DummyHead *MyListNode
	DummyTail *MyListNode
}

type MyListNode struct {
	Val  int
	Prev *MyListNode
	Next *MyListNode
}

func NewMyLinkedList() MyLinkedList {
	node1, node2 := &MyListNode{Val: -1}, &MyListNode{Val: -1}
	node1.Next = node2
	node2.Prev = node1
	return MyLinkedList{0, node1, node2}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.Size {
		return -1
	}
	if index <= l.Size/2 {
		cur := l.DummyHead
		for cur.Next != nil && index >= 0 {
			cur = cur.Next
			index--
		}
		return cur.Val
	} else {
		cur := l.DummyTail
		cnt := l.Size - index - 1
		for cur.Prev != nil && cnt >= 0 {
			cur = cur.Prev
			cnt--
		}
		return cur.Val
	}
}

func (l *MyLinkedList) AddAtHead(val int) {
	node := &MyListNode{Val: val}
	node.Prev = l.DummyHead
	node.Next = l.DummyHead.Next
	l.DummyHead.Next.Prev = node
	l.DummyHead.Next = node
	l.Size++
}

func (l *MyLinkedList) AddAtTail(val int) {
	node := &MyListNode{Val: val}
	node.Next = l.DummyTail
	node.Prev = l.DummyTail.Prev
	l.DummyTail.Prev.Next = node
	l.DummyTail.Prev = node
	l.Size++
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		l.AddAtHead(val)
	} else if index > l.Size {
		return
	} else if index == l.Size {
		l.AddAtTail(val)
	} else {
		if index <= l.Size/2 {
			cur := l.DummyHead
			for cur.Next != nil && index > 0 {
				cur = cur.Next
				index--
			}
			node := &MyListNode{Val: val}
			node.Prev = cur
			node.Next = cur.Next.Next
			cur.Next.Prev = node
			cur.Next = node
		} else {
			cur := l.DummyTail
			cnt := l.Size - 1 - index
			for cur.Prev != nil && cnt > 0 {
				cur = cur.Prev
				cnt--
			}
			node := &MyListNode{val, cur.Prev, cur}
			node.Prev = cur.Prev
			node.Next = cur
			cur.Prev.Next = node
			cur.Prev = node
		}
		l.Size++
	}
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	//l.Output()
	if index < 0 || index >= l.Size {
		return
	}
	if index <= l.Size/2 {
		cur := l.DummyHead
		for cur.Next != nil && index > 0 {
			cur = cur.Next
			index--
		}
		cur.Next = cur.Next.Next
		cur.Next.Prev = cur
	} else {
		cur := l.DummyTail
		cnt := l.Size - index - 1
		for cur.Prev != nil && cnt > 0 {
			cur = cur.Prev
			cnt--
		}
		cur.Prev = cur.Prev.Prev
		cur.Prev.Next = cur
	}
	l.Size--
}

func (l *MyLinkedList) Output() {
	cur := l.DummyHead
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

// leetcode713
func NumSubarrayProductLessThanK713I(nums []int, k int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	res := 0
	for i := 0; i < n; i++ {
		mulV := 1
		for j := i; j < n; j++ {
			mulV *= nums[j]
			if mulV < k {
				res++
			} else {
				break
			}
		}
	}
	return res
}

func NumSubarrayProductLessThanK713II(nums []int, k int) int {
	n := len(nums)
	l := 0
	res, mulV := 0, 1
	for r := 0; r < n; r++ {
		mulV *= nums[r]
		for mulV >= k {
			mulV /= nums[l]
			l++
		}
		res += r - l + 1
	}
	return res
}

// leetcode714
func MaxProfit6(prices []int, fee int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0], dp[0][1] = 0, -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = utils.MaxNum(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = utils.MaxNum(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

// leetcode715
// 线段树动态开点
// 假设不会出现还没跟踪就删除的情况
type dsegTreeNode struct {
	l, m, r    int
	val        int
	add        int
	lson, rson *dsegTreeNode
}

type DsegTree struct {
	root *dsegTreeNode
}

func NewDsegTreeNode(l, r int) *dsegTreeNode {
	return &dsegTreeNode{
		l:   l,
		r:   r,
		m:   (l + r) >> 1,
		val: -1,
	}
}

func NewDsegTree(l, r int) *DsegTree {
	return &DsegTree{root: NewDsegTreeNode(l, r)}
}

func (t *DsegTree) pushUp(node *dsegTreeNode) {
	node.val = utils.MinNum(node.lson.val, node.rson.val)
}

func (t *DsegTree) pushDown(node *dsegTreeNode) {
	if node.lson == nil {
		node.lson = NewDsegTreeNode(node.l, node.m)
	}
	if node.rson == nil {
		node.rson = NewDsegTreeNode(node.m+1, node.r)
	}
	if node.add != 0 {
		node.lson.val = node.add
		node.rson.val = node.add
		node.lson.add = node.add
		node.rson.add = node.add
		node.add = 0
	}
}

func (t *DsegTree) modify(l, r, val int, node *dsegTreeNode) {
	if l > r {
		return
	}
	if l <= node.l && node.r <= r {
		node.val = val
		node.add = val
		return
	}
	t.pushDown(node)
	if l <= node.m {
		t.modify(l, r, val, node.lson)
	}
	if r > node.m {
		t.modify(l, r, val, node.rson)
	}
	t.pushUp(node)
}

func (t *DsegTree) query(l, r int, node *dsegTreeNode) int {
	if l > r {
		return 1e9
	}
	if l <= node.l && node.r <= r {
		return node.val
	}
	t.pushDown(node)
	// 和左子区间有交集
	var val int = 1e9
	if l <= node.m {
		val = utils.MinNum(val, t.query(l, r, node.lson))
	}
	// 和右子区间有交集
	if r > node.m {
		val = utils.MinNum(val, t.query(l, r, node.rson))
	}
	return val
}

type RangeModule struct {
	tree *DsegTree
}

func NewRangeModule() RangeModule {
	return RangeModule{tree: NewDsegTree(1, 1e9+1)}
}

func (r *RangeModule) AddRange(left int, right int) {
	r.tree.modify(left, right-1, 1, r.tree.root)
}

func (r *RangeModule) QueryRange(left int, right int) bool {
	ans := r.tree.query(left, right-1, r.tree.root)
	return ans > 0
}

func (r *RangeModule) RemoveRange(left int, right int) {
	r.tree.modify(left, right-1, -1, r.tree.root)
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */

// leetcode718
func FindLength(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	res := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > res {
					res = dp[i][j]
				}
			}
		}
	}
	return res
}

// leetcode719
func SmallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	l, r := 0, nums[n-1]-nums[0]

	getCount := func(nums []int, m int) int {
		res := 0
		l, r := 0, 0
		for r < len(nums) {
			for nums[r]-nums[l] > m {
				l++
			}
			res += r - l
			r++
		}
		return res
	}
	for l <= r {
		m := (l + r) >> 1
		// 找大于等于k的左边界
		if getCount(nums, m) >= k {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

// leetcode725
func SplitListToParts(head *algorithm.ListNode, k int) []*algorithm.ListNode {
	p := head
	listLen := 0
	for p != nil {
		listLen++
		p = p.Next
	}
	remain := listLen % k
	average := listLen / k
	res := []*algorithm.ListNode{}
	curLen := 0
	prev := head
	for p = head; p != nil; p = p.Next {
		if curLen == 0 {
			curLen++
			res = append(res, p)
			continue
		}
		if curLen < average {
			curLen++
		} else if curLen == average && remain != 0 {
			remain--
			curLen++
		} else {
			curLen = 1
			prev.Next = nil
			res = append(res, p)
		}
		prev = p
	}
	prevListLen := len(res)
	for prevListLen != k {
		res = append(res, nil)
		prevListLen++
	}
	return res
}

// leetcode726
func CountOfAtoms(formula string) string {
	stk := []map[string]int{{}}
	i, n := 0, len(formula)
	getCurAtom := func() string {
		start := i
		i++
		for i < n && unicode.IsLower(rune(formula[i])) {
			i++
		}
		return formula[start:i]
	}

	getNum := func() int {
		if i == n || !unicode.IsDigit(rune(formula[i])) {
			return 1
		}
		num := 0
		for ; i < n && unicode.IsDigit(rune(formula[i])); i++ {
			num = num*10 + int(formula[i]-'0')
		}
		return num
	}

	for i < n {
		if ch := formula[i]; ch == '(' {
			i++
			stk = append(stk, map[string]int{})
		} else if ch == ')' {
			i++
			num := getNum()
			curLevel := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			for atom, v := range curLevel {
				stk[len(stk)-1][atom] += v * num
			}
		} else {
			atom := getCurAtom()
			num := getNum()
			stk[len(stk)-1][atom] += num
		}
	}
	atoms := []string{}
	for atom := range stk[0] {
		atoms = append(atoms, atom)
	}
	sort.Strings(atoms)
	res := ""
	for i := 0; i < len(atoms); i++ {
		res += atoms[i]
		if num := stk[0][atoms[i]]; num != 1 {
			res += strconv.Itoa(num)
		}
	}
	return res
}

// leetcode729
type Events [][2]int

func (e Events) Len() int           { return len(e) }
func (e Events) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e Events) Less(i, j int) bool { return e[i][0] < e[j][0] }

type MyCalendar struct {
	events [][2]int
}

func NewMyCalendar() MyCalendar {
	return MyCalendar{make([][2]int, 0, 1000)}
}

func (m *MyCalendar) Book(start int, end int) bool {
	n := len(m.events)
	if n == 0 {
		m.events = append(m.events, [2]int{start, end})
		return true
	}

	idx := sort.Search(n, func(i int) bool {
		return m.events[i][0] > start
	})
	// fmt.Println(idx)
	// fmt.Println(idx, m.events)
	if idx == n {
		// m.events[i][0]全都比start小
		if m.events[n-1][1] > start {
			return false
		}
		m.events = append(m.events, [2]int{start, end})
		return true
	} else if idx == 0 {
		// m.events[i][0]全都比start大
		if m.events[0][0] < end {
			return false
		}
		m.events = append([][2]int{{start, end}}, m.events...)
		return true
	}
	// m.events[idx][0] > start
	// m.events[idx-1][0] <= start
	if m.events[idx-1][0] == start {
		return false
	}
	if m.events[idx-1][1] <= start && m.events[idx][0] >= end {
		m.events = append(m.events[:idx], append([][2]int{{start, end}}, m.events[idx:]...)...)
		return true
	}
	return false
}

// leetcode732
type node struct {
	left      *node
	right     *node
	l, mid, r int
	v, add    int
}

func newNode(l, r int) *node {
	return &node{
		l:   l,
		r:   r,
		mid: int(uint(l+r) >> 1),
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type segmentTree struct {
	root *node
}

func newSegmentTree() *segmentTree {
	return &segmentTree{
		root: newNode(1, 1e9+1),
	}
}

func (t *segmentTree) modify(l, r, v int, n *node) {
	if l > r {
		return
	}
	if n.l >= l && n.r <= r {
		n.v += v
		n.add += v
		return
	}
	t.pushdown(n)
	if l <= n.mid {
		t.modify(l, r, v, n.left)
	}
	if r > n.mid {
		t.modify(l, r, v, n.right)
	}
	t.pushup(n)
}

func (t *segmentTree) query(l, r int, n *node) int {
	if l > r {
		return 0
	}
	if n.l >= l && n.r <= r {
		return n.v
	}
	t.pushdown(n)
	v := 0
	if l <= n.mid {
		v = max(v, t.query(l, r, n.left))
	}
	if r > n.mid {
		v = max(v, t.query(l, r, n.right))
	}
	return v
}

func (t *segmentTree) pushup(n *node) {
	n.v = max(n.left.v, n.right.v)
}

func (t *segmentTree) pushdown(n *node) {
	if n.left == nil {
		n.left = newNode(n.l, n.mid)
	}
	if n.right == nil {
		n.right = newNode(n.mid+1, n.r)
	}
	if n.add != 0 {
		n.left.add += n.add
		n.right.add += n.add
		n.left.v += n.add
		n.right.v += n.add
		n.add = 0
	}
}

type MyCalendarThree struct {
	tree *segmentTree
}

func NewMyCalendarThree() MyCalendarThree {
	return MyCalendarThree{newSegmentTree()}
}

func (m *MyCalendarThree) BookThree(start int, end int) int {
	m.tree.modify(start+1, end, 1, m.tree.root)
	return m.tree.query(1, int(1e9)+1, m.tree.root)
}

// leetcode743
func NetworkDelayTime(times [][]int, n int, k int) int {
	mp := make(map[int]map[int]int)
	minTime := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minTime[i] = 0xFFFFFFFF
	}

	for i := 0; i < len(times); i++ {
		from, to, dis := times[i][0], times[i][1], times[i][2]
		if mp[from] == nil {
			mp[from] = make(map[int]int)
		}
		mp[from][to] = dis
	}

	var dfs func(start int, curTime int)
	dfs = func(start, curTime int) {
		if minTime[start] > curTime {
			minTime[start] = curTime
		}
		for key, value := range mp[start] {

			dfs(key, curTime+value)
		}
	}

	dfs(k, 0)

	maxTime := 0
	for i := 1; i <= n; i++ {
		if minTime[i] > maxTime && len(mp[i]) == 0 {
			maxTime = minTime[i]
		}
	}
	fmt.Println(minTime)
	if maxTime == 0xFFFFFFFF {
		return -1
	}
	return maxTime
}

// leetcode746 看不懂题目系列
func MinCostClimbingStairs1(cost []int) int {
	// 网友热评:
	// 我觉得这个题的描述应该改改：每个阶梯都有一定数量坨屎，
	// 一次只能跨一个或者两个阶梯，走到一个阶梯就要吃光上面的屎，问怎么走才能吃最少的屎？
	// 开局你选前两个阶梯的其中一个作为开头点，并吃光该阶梯的屎。
	// 最后一个阶梯后还存在一个无屎的平台
	cost = append(cost, 0)
	length := len(cost)
	dp := make([][]int, length+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0], dp[0][1] = 0, 0
	dp[1][0], dp[1][1] = cost[0], cost[0]
	for i := 1; i < length; i++ {
		dp[i+1][0] = utils.MinNum(dp[i][0], dp[i][1]) + cost[i]
		dp[i+1][1] = utils.MinNum(dp[i-1][0], dp[i-1][1]) + cost[i]
	}
	return utils.MinNum(dp[length][0], dp[length][1])
}

func MinCostClimbingStairs3(cost []int) int {
	dp := make([]int, len(cost)+1)
	if len(cost) <= 1 {
		return 0
	}
	dp[0], dp[1] = 0, 0
	for i := 2; i <= len(cost); i++ {
		dp[i] = utils.MinNum(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}

// leetcode747
func DominantIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	} else if len(nums) == 1 {
		return 0
	}
	var a, b, idx int
	if a < b {
		a, b, idx = nums[0], nums[1], 1
	} else {
		a, b, idx = nums[1], nums[0], 0
	}
	for i := 2; i < len(nums); i++ {
		if nums[i] > b {
			a = b
			b = nums[i]
			idx = i
		} else if nums[i] > a {
			a = nums[i]
		}
	}
	fmt.Println(a, b, idx)
	if b >= 2*a {
		return idx
	} else {
		return -1
	}
}

// leetcode748
func ShortestCompletingWord(licensePlate string, words []string) string {
	licensePlate = strings.ToLower(licensePlate)
	minLen := 1005
	var num [26]int
	for i := 0; i < len(licensePlate); i++ {
		if licensePlate[i] <= 'z' && licensePlate[i] >= 'a' {
			num[licensePlate[i]-'a']++
		}
	}
	ans := ""
	for i := 0; i < len(words); i++ {
		cnt := [26]int{}
		flag := true
		for j := 0; j < len(words[i]); j++ {
			cnt[words[i][j]-'a']++
		}
		for i := 0; i < 26; i++ {
			if cnt[i] < num[i] {
				flag = false
				break
			}
		}
		if flag && len(words[i]) < minLen {
			minLen = len(words[i])
			ans = words[i]
		}
	}
	return ans
}
