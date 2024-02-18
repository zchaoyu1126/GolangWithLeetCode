package leetcode

import (
	"container/heap"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/zchaoyu1126/coding-practice/internal/algorithmingo/algorithm"
	"github.com/zchaoyu1126/coding-practice/utils"
)

// leetcode1
func TwoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, has := mp[target-nums[i]]; has {
			return []int{i, val}
		} else {
			mp[nums[i]] = i
		}
	}
	return []int{}
}

// leetcode7
func Reverse(x int) int {
	tmp, res := x, 0
	for tmp != 0 {
		if res < math.MinInt32/10 || res > (math.MaxInt32-1)/10 {
			return 0
		}
		res = res*10 + tmp%10
		tmp /= 10
	}
	return res
}

// leetcode8
func MyAtoi(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	bytes := []byte(s)
	var sign int32 = 1
	if bytes[0] == '-' {
		sign = -1
		bytes = bytes[1:]
	} else if bytes[0] == '+' {
		bytes = bytes[1:]
	} else if !(bytes[0] <= '9' && bytes[0] >= '0') {
		return 0
	}

	var res int32 = 0
	for i := 0; i < len(bytes); i++ {
		if bytes[i] <= '9' && bytes[i] >= '0' {
			if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && int(bytes[i]-'0') > math.MaxInt32%10) {
				return math.MaxInt32
			}
			if res < math.MinInt32/10 || (res == math.MinInt32/10 && int(bytes[i]-'0') > -(math.MinInt32%10)) {
				return math.MinInt32
			}
			res = res*10 + sign*int32(bytes[i]-'0')
		} else {
			break
		}
	}
	return int(res)
}

// leetcode9
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

// leetcode15
func ThreeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		// 正确的去重方式
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			// 如果把去重逻辑放在这里，可能会直接导致right<=left 例如0,0,0这种数据
			// for k > j && nums[k] == nums[k-1] {
			//     k--
			// }
			// for k > j && nums[j] == nums[j+1] {
			//     j++
			// }
			if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else {
				// 找到一个三元组后再去重
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for k > j && nums[k] == nums[k-1] {
					k--
				}
				for k > j && nums[j] == nums[j+1] {
					j++
				}
				k--
				j++
			}
		}
	}
	return res
}

// leetcode17
func LetterCombinations(digits string) []string {
	mp := make(map[byte][]byte)
	mp['2'] = []byte{'a', 'b', 'c'}
	mp['3'] = []byte{'d', 'e', 'f'}
	mp['4'] = []byte{'g', 'h', 'i'}
	mp['5'] = []byte{'j', 'k', 'l'}
	mp['6'] = []byte{'m', 'n', 'o'}
	mp['7'] = []byte{'p', 'q', 'r', 's'}
	mp['8'] = []byte{'t', 'u', 'v'}
	mp['9'] = []byte{'w', 'x', 'y', 'z'}
	res := []string{}
	cur := []byte{}
	var backtrace func(digits string)
	backtrace = func(digits string) {
		if len(digits) <= 0 {
			res = append(res, string(cur))
			return
		}
		for i := 0; i < len(digits); i++ {
			for _, ch := range mp[digits[i]] {
				cur = append(cur, ch)
				backtrace(digits[1:])
				cur = cur[:len(cur)-1]
			}
		}
	}
	backtrace(digits)
	return res
}

// leetcode18
func FourSum(nums []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, len(nums)-1
			for left < right {

				if nums[i]+nums[j]+nums[left]+nums[right] > target {
					right--
				} else if nums[i]+nums[j]+nums[left]+nums[right] < target {
					left++
				} else {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for right > left && nums[right] == nums[right-1] {
						right--
					}
					for right > left && nums[left] == nums[left+1] {
						left++
					}
					right--
					left++
				}
			}
		}
	}
	return res
}

// leetcode19
func RemoveNthFromEnd(head *algorithm.ListNode, n int) *algorithm.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	dummy := &algorithm.ListNode{}
	dummy.Next = head
	slow, fast := dummy, dummy
	for fast.Next != nil {
		for n > 0 {
			fast = fast.Next
			n--
		}
		if fast.Next == nil {
			break
		}
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return dummy.Next
}

// leetcode22
func GenerateParenthesis(n int) []string {
	var generate func(str string, l, r int)
	res := []string{}
	generate = func(str string, l, r int) {
		if l == 0 && r == 0 {
			res = append(res, str)
			return
		}
		if l == r {
			generate(str+"(", l-1, r)
		} else {
			if l > 0 {
				generate(str+"(", l-1, r)
			}
			generate(str+")", l, r-1)
		}
	}
	generate("", n, n)
	return res
}

// leetcode21
func MergeTwoLists(l1 *algorithm.ListNode, l2 *algorithm.ListNode) *algorithm.ListNode {
	i, j := l1, l2
	newHead := &algorithm.ListNode{}
	cur := newHead
	for i != nil && j != nil {
		if i.Val < j.Val {
			cur.Next = i
			i = i.Next
		} else {
			cur.Next = j
			j = j.Next
		}
		cur = cur.Next
	}
	for i != nil {
		cur.Next = i
		cur = cur.Next
		i = i.Next
	}
	for j != nil {
		cur.Next = j
		cur = cur.Next
		j = j.Next
	}
	return newHead.Next
}

// leetcode23
// 使用堆来合并k个列表
type ListNodeHeap []*algorithm.ListNode

func (h ListNodeHeap) Len() int           { return len(h) }
func (h ListNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h ListNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*algorithm.ListNode))
}
func (h *ListNodeHeap) Pop() interface{} {
	length := len(*h)
	x := (*h)[length-1]
	*h = (*h)[:length-1]
	return x
}

func MergeKLists(lists []*algorithm.ListNode) *algorithm.ListNode {
	h := &ListNodeHeap{}
	heap.Init(h)
	for _, list := range lists {
		cur := list
		for cur != nil {
			//fmt.Println(cur.Val)
			heap.Push(h, cur)
			cur = cur.Next
		}
	}
	if len(*h) == 0 {
		return nil
	}
	newList := heap.Pop(h).(*algorithm.ListNode)
	cur := newList
	for len(*h) != 0 {
		cur.Next = heap.Pop(h).(*algorithm.ListNode)
		//fmt.Println(cur.Next.Val)
		cur = cur.Next
	}
	cur.Next = nil
	return newList
}

// leetcode24
func ReverseLinkedList(root *algorithm.ListNode) *algorithm.ListNode {
	if root == nil || root.Next == nil {
		return root
	}
	others := root.Next
	newHead := ReverseLinkedList(others)
	others.Next = root
	root.Next = nil
	return newHead
}

func SwapPairs1(head *algorithm.ListNode) *algorithm.ListNode {
	return ReverseKGroup3(head, 2)
}

func SwapPairs2(head *algorithm.ListNode) *algorithm.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	others := newHead.Next
	newHead.Next = head
	head.Next = SwapPairs2(others)
	return newHead
}

// leetcode25
func ReverseKGroup1(head *algorithm.ListNode, k int) *algorithm.ListNode {
	if k == 1 {
		return head
	}
	stack := []*algorithm.ListNode{}
	cnt := 0
	newHead := &algorithm.ListNode{}
	preEnd := newHead

	for cur := head; cur != nil; cur = cur.Next {
		if cnt < k {
			stack = append(stack, cur)
			cnt++
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			top.Next = nil
			preEnd.Next = top
			prev := top
			for len(stack) != 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				top.Next = nil
				prev.Next = top
				prev = top
			}
			preEnd = prev
			cnt = 1
			stack = append(stack, cur)
		}
	}
	if len(stack) == k {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		top.Next = nil
		preEnd.Next = top
		prev := top
		for len(stack) != 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			top.Next = nil
			prev.Next = top
			prev = top
		}
	} else if len(stack) != 0 {
		for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
			stack[i], stack[j] = stack[j], stack[i]
		}
		for len(stack) != 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			top.Next = nil
			preEnd.Next = top
			preEnd = top
		}
	}
	return newHead.Next
}

func ReverseKGroup2(head *algorithm.ListNode, k int) *algorithm.ListNode {
	cur := head
	// 前进k-1次，此时cur为结束的节点
	// cur.Next就是新一组要反转的开头

	for i := 0; i < k-1; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}
	// 这里的cur也有可能是nil
	// 比如[1,2,3,4,5] 2
	if cur == nil {
		return head
	}

	// 断开cur后面的节点
	next := cur.Next
	cur.Next = nil

	// head成为尾巴
	newHead := ReverseLinkedList(head)
	nextHead := ReverseKGroup2(next, k)
	head.Next = nextHead
	return newHead
}

func ReverseKGroup3(head *algorithm.ListNode, k int) *algorithm.ListNode {
	if k == 1 {
		return head
	}

	dummy := &algorithm.ListNode{}
	cur := dummy

	for {
		stack := []*algorithm.ListNode{}
		cnt := k
		// 移动过去，看一下是否有k个
		tmp := head
		for tmp != nil && cnt > 0 {
			stack = append(stack, tmp)
			tmp = tmp.Next
			cnt++
		}
		if cnt > 0 {
			cur.Next = head
			break
		}

		for len(stack) != 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cur.Next = top
			cur = top

		}

		head = tmp
	}
	return dummy.Next
}

// leetcode26
func RemoveDuplicates(nums []int) int {
	slow, fast := 0, 1

	for fast = 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			nums[slow+1] = nums[fast]
			slow++
		}
	}
	return slow
}

// leetcode27
func RemoveElement1(nums []int, val int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l] == val {
			for r >= l && nums[r] == val {
				r--
			}
			if r < l {
				break
			}
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		} else {
			l++
		}
	}
	return l
}

func RemoveElement2(nums []int, val int) int {
	slow, fast := 0, 0
	for fast = 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// leetcode28
func StrStr1(s string, t string) int {
	if s == "" && t == "" {
		return 0
	}
	n, m := len(s), len(t)
	for i := 0; i <= n-m; i++ {
		match := 0
		for j := 0; j < m; j++ {
			if s[i+j] == t[j] {
				match++
			} else {
				break
			}
		}
		if match == m {
			return i
		}
	}
	return -1
}

func StrStr2(s string, t string) int {
	if s == "" && t == "" {
		return 0
	}
	n, m := len(s), len(t)
	if n < m {
		return -1
	}
	var p, k int = 23, 1e7 + 7
	targetVal := HashForStr(t, p, k)
	val := HashForStr(s[:m], p, k)
	tmp := QuickPow(p, m-1, k)
	if val == targetVal {
		if Check(s[:m], t) {
			return 0
		}
	}
	for i := 1; i <= n-m; i++ {
		val = (((val-int(s[i-1])*tmp%k+k)%k*p)%k + int(s[i+m-1])%k) % k
		if val == targetVal {
			if Check(s[i:i+m], t) {
				return i
			}
		}
	}
	return -1
}

func HashForStr(str string, prime, k int) int {
	hashVal := 0
	n := len(str)
	for i := 0; i < n; i++ {
		hashVal = (hashVal + int(str[i])*QuickPow(prime, n-i-1, k)%k) % k
	}
	return hashVal
}

func QuickPow(x, cnt, k int) int {
	res := 1
	for cnt > 0 {
		if cnt&1 == 1 {
			res = (res * x) % k
		}
		x = (x * x) % k
		cnt >>= 1
	}
	return res
}

func Check(s, t string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			return false
		}
	}
	return true
}

// leetcode29
func DivideLeetCode(dividend int, divisor int) int {
	// 首先考虑溢出的情况
	// dividend = 2147483647,这种情况不可能溢出
	// dividend = -2147483648, 除-1或者除1
	if dividend == minNUmber && divisor == -1 {
		return maxNumber
	} else if dividend == minNUmber && divisor == 1 {
		return dividend
	}
	res := 0
	dividend, isMinus1 := utils.AbsInt(dividend)
	divisor, isMinus2 := utils.AbsInt(divisor)
	for dividend >= divisor {
		n := divisor
		cnt := 1
		for n<<1 < dividend {
			n <<= 1
			cnt <<= 1
		}
		res += cnt
		dividend -= n
	}
	if isMinus1 || isMinus2 {
		return -res
	}
	return res
}

// leetcode35
func SearchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		}
	}
	return l
}

// leetcode36
func IsValidSudoku(board [][]byte) bool {
	row, col, area := [10]int{}, [10]int{}, [10]int{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				continue
			}
			cnt := board[i][j] - '0'
			if (row[i]>>cnt&1 == 1) || (col[j]>>cnt&1 == 1) || (area[(i/3)*3+j/3]>>cnt&1 == 1) {
				return false
			}
			row[i] |= 1 << cnt
			col[j] |= 1 << cnt
			area[(i/3)*3+j/3] |= 1 << cnt
		}
	}
	return true
}

// leetcode37
func SolveSudoku(board [][]byte) {
	m, n := 9, 9
	position := []int{}
	row, col := [9][9]int{}, [9][9]int{}
	block := [3][3][9]int{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				position = append(position, i*9+j)
			} else {
				digit := int(board[i][j] - '0')
				row[i][digit] = 1
				col[j][digit] = 1
				block[i/3][j/3][digit] = 1
			}
		}
	}

	var dfs func(n int)
	dfs = func(n int) {
		if n == len(position) {
			return
		}
		x, y := position[n]/9, position[n]%9
		for i := 1; i <= 9; i++ {
			if row[x][i] == 0 && col[y][i] == 0 && block[x/3][y/3][i] == 0 {
				board[x][y] = byte('0' + i)
				row[x][i], col[y][i], block[x/3][y/3][i] = 1, 1, 1
				dfs(n + 1)
				board[x][y] = '.'
				row[x][i], col[y][i], block[x/3][y/3][i] = 0, 0, 0
			}
		}
	}
	dfs(0)
}

// leetcode38
func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	prev := "1"
	cur := ""
	for i := 1; i < n; i++ {
		cnt := 1
		lastNumber := prev[0]
		for j := 1; j < len(prev); j++ {
			if prev[j] != lastNumber {
				cur += string(byte(cnt) + '0')
				cur += string(lastNumber)
				cnt = 1
				lastNumber = prev[j]
			} else {
				cnt++
			}
		}
		cur += string(byte(cnt) + '0')
		cur += string(lastNumber)
		prev = cur
		cur = ""
	}
	return cur
}

// leetcode39
func CombinationSum(candidates []int, target int) [][]int {

	sort.Ints(candidates)
	minN := candidates[0]

	cur := []int{}
	res := [][]int{}
	var backtrace func(start, target int)
	backtrace = func(start, target int) {
		if target == 0 {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
		} else if target < minN {
			return
		}
		for i := start; i < len(candidates); i++ {
			cur = append(cur, candidates[i])
			backtrace(i, target-candidates[i])
			cur = cur[:len(cur)-1]
		}
	}
	backtrace(0, target)
	return res
}

// leetcode40
func CombinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	minN := candidates[0]

	cur := []int{}
	res := [][]int{}
	var backtrace func(start, target int)
	prev := minN - 1
	backtrace = func(start, target int) {
		if target == 0 {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
		} else if target < minN {
			return
		}
		i := start
		for i < len(candidates) && candidates[i] == prev {
			i++
		}
		for ; i < len(candidates); i++ {
			cur = append(cur, candidates[i])
			backtrace(i+1, target-candidates[i])
			cur = cur[:len(cur)-1]
		}
	}
	backtrace(0, target)
	return res
}

// leetcode42
func Trap(height []int) int {
	res := 0
	stack := []int{}

	for i := 0; i < len(height); i++ {

		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			curWidth := i - left - 1
			curHeight := utils.MinNum(height[i], height[left]) - height[top]
			res += curHeight * curWidth
		}
		stack = append(stack, i)
	}
	return res
}

// leetcode49
func GroupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		if _, ok := mp[sortedStr]; !ok {
			mp[sortedStr] = []string{}
		}
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	res, i := [][]string{}, 0
	for _, val := range mp {
		res[i] = []string{}
		res[i] = append(res[i], val...)

	}
	return res
}

// leetcode45
func Jump(nums []int) int {
	end, maxPosition, step := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = utils.MaxNum(maxPosition, nums[i]+i)
		if i == end {
			end = maxPosition
			step++
		}
	}
	return step
}

// leetcode46
func Permute(nums []int) [][]int {
	res := [][]int{}
	cur := []int{}
	used := make([]int, len(nums))
	var backtrace func()
	backtrace = func() {
		if len(cur) == len(nums) {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
		}

		for i := 0; i < len(nums); i++ {
			if used[i] == 1 {
				continue
			}
			used[i] = 1
			cur = append(cur, nums[i])
			backtrace()
			used[i] = 0
			cur = cur[:len(cur)-1]
		}
	}
	backtrace()
	return res
}

// leetcode47
func PermuteUnique(nums []int) [][]int {
	res := [][]int{}
	cur := []int{}
	used := make([]int, len(nums))

	var backtrace func()
	backtrace = func() {
		if len(cur) == len(nums) {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
		}
		history := make([]int, 21)
		for i := 0; i < len(nums); i++ {
			if used[i] == 1 || history[nums[i]] == 1 {
				continue
			}
			used[i] = 1
			history[nums[i]] = 1
			cur = append(cur, nums[i])
			backtrace()
			used[i] = 0
			cur = cur[:len(cur)-1]
		}
	}
	backtrace()
	return res
}

// leetcode50
func MyPow(x float64, n int) float64 {
	cnt := 0
	if n < 0 {
		cnt = -n
	} else {
		cnt = n
	}
	res := qPow(x, cnt)
	if n < 0 {
		return 1.0 / res
	}
	return res
}

func qPow(x float64, n int) float64 {
	res := 1.0
	for n > 0 {
		if n&1 == 1 {
			res = res * x
		}
		x = x * x
		n >>= 1
	}
	return res
}
