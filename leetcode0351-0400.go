package leetcode

import (
	"container/heap"
	"fmt"
	"math"
	"math/rand"
	"programs/internal/algorithmingo/algorithm"
	"programs/kit/common"
	"strconv"
	"time"
)

// leetcode352
type interval struct {
	lBorder int
	rBorder int
}
type SummaryRanges struct {
	intervals []interval
	length    int
}

func NewSummaryRanges() SummaryRanges {
	return SummaryRanges{intervals: []interval{}, length: 0}
}

func (t *SummaryRanges) AddNum(val int) {
	num := len(t.intervals)
	l, r := 0, num-1
	for l <= r {
		mid := (l + r) >> 1
		lBorder, rBorder := t.intervals[mid].lBorder, t.intervals[mid].rBorder
		if lBorder <= val && val <= rBorder {
			return
		}
		if val > rBorder {
			l = mid + 1
		} else if val < lBorder {
			r = mid - 1
		}
	}

	// 发现两个可以合并的区间
	if l-1 >= 0 && l < num {
		if t.intervals[l-1].rBorder+1 == val && t.intervals[l].lBorder-1 == val {
			t.intervals[l-1].rBorder = t.intervals[l].rBorder
			t.intervals = append(t.intervals[0:l], t.intervals[l+1:]...)
			return
		}
	}

	// 要和插入点两边的值进行检查 l,l-1
	if l-1 >= 0 && t.intervals[l-1].rBorder == val-1 {
		// 左边检查
		t.intervals[l-1].rBorder = val
		return
	}

	if l < num && t.intervals[l].lBorder == val+1 {
		// 右边检查
		t.intervals[l].lBorder = val
		return
	}

	remain := make([]interval, len(t.intervals)-l)
	copy(remain, t.intervals[l:])
	t.intervals = append(t.intervals[0:l], interval{lBorder: val, rBorder: val})
	t.intervals = append(t.intervals, remain...)
}

func (t *SummaryRanges) GetIntervals() [][]int {
	res := [][]int{}
	for i := 0; i < len(t.intervals); i++ {
		res = append(res, []int{t.intervals[i].lBorder, t.intervals[i].rBorder})
		//fmt.Println(t.intervals[i].lBorder, t.intervals[i].rBorder)
	}
	return res
}

// leetcode367
func IsPerfectSquare(num int) bool {
	cnt := 1
	for num > 0 {
		num -= cnt
		cnt += 2
	}
	return num == 0
}

// leetcode367
func IsPerfectSquare2(num int) bool {
	// 最后一个小于等于
	l, r := 0, num
	for l <= r {
		m := (l + r) / 2
		if m*m <= num {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return r*r == num
}

// leetcode372
func SuperPow(a int, b []int) int {
	modNumber, length, res := 1337, len(b), 1
	pow := func(x, cnt int) int {
		ans := 1
		for cnt > 0 {
			if cnt&1 > 0 {
				ans = ans * x % modNumber
			}
			x = (x * x) % modNumber
			cnt >>= 1
		}
		return ans
	}

	for i := length - 1; i >= 0; i-- {
		res = res * pow(a, b[i]) % modNumber
		a = pow(a, 10)
	}
	return res
}

// leetcode373
type idxPair struct{ i, j int }
type hp struct {
	data         []idxPair
	nums1, nums2 []int
}

func KSmallestPairs(nums1, nums2 []int, k int) (ans [][]int) {
	m, n := len(nums1), len(nums2)
	h := hp{nil, nums1, nums2}
	for i := 0; i < k && i < m; i++ {
		h.data = append(h.data, idxPair{i, 0})
	}
	for h.Len() > 0 && len(ans) < k {
		p := heap.Pop(&h).(idxPair)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(&h, idxPair{i, j + 1})
		}
	}
	return
}

func (h hp) Len() int { return len(h.data) }
func (h hp) Less(i, j int) bool {
	a, b := h.data[i], h.data[j]
	return h.nums1[a.i]+h.nums2[a.j] < h.nums1[b.i]+h.nums2[b.j]
}
func (h hp) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *hp) Push(v interface{}) { h.data = append(h.data, v.(idxPair)) }
func (h *hp) Pop() interface{}   { a := h.data; v := a[len(a)-1]; h.data = a[:len(a)-1]; return v }

// leetcode375
func GetMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 1; n-- {
		for j := i + 1; j <= n; j++ {
			minCost := math.MaxInt32
			for k := i; k < j; k++ {
				cost := k + common.LargerNumber(dp[i][k-1], dp[k+1][j])
				if cost < minCost {
					minCost = cost
				}
			}
			dp[i][j] = minCost
		}
	}
	return dp[1][n]
}

// leetcode382
// type Solution382 []int

// func Constructor382(head *algorithm.ListNode) Solution382 {
// 	var s Solution382
// 	for node := head; node != nil; node = node.Next {
// 		s = append(s, node.Val)
// 	}
// 	return s
// }

// func (s Solution382) GetRandom() int {
// 	return s[rand.Intn(len(s))]
// }

// 用蓄水池算法
type Solution382 struct {
	head *algorithm.ListNode
}

func Constructor(head *algorithm.ListNode) Solution382 {
	return Solution382{head}
}

func (s Solution382) GetRandom() int {
	if s.head == nil {
		return -1
	}
	ans := s.head.Val

	for node, i := s.head.Next, 1; node != nil; node, i = node.Next, i+1 {
		if rand.Intn(i+1) == 0 {
			ans = node.Val
		}
	}
	return ans
}

// leetcode383
func CanConstruct(ransomNote string, magazine string) bool {
	cnt := map[byte]int{}
	for i := 0; i < len(magazine); i++ {
		cnt[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		cnt[ransomNote[i]-'a']--
		if cnt[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}

// leetcode384
type Solution struct {
	src []int
	dst []int
}

func NewSolution(nums []int) Solution {
	return Solution{nums, []int{}}
}

func (s *Solution) Reset() []int {
	fmt.Println("src", s.src)
	return s.src
}

func (s *Solution) Shuffle() []int {
	s.dst = KnuthDurstenfeldShuffle(s.src)
	fmt.Println("src", s.src)
	fmt.Println(s.dst)
	return s.dst
}

func FisherYatesShuffle(arr []int) []int {
	res := []int{}
	rand.Seed(time.Now().UnixNano())
	for len(arr) != 0 {
		pos := rand.Intn(len(arr))
		res = append(res, arr[pos])
		arr = append(arr[:pos], arr[pos+1:]...)
	}
	fmt.Println(res)
	return res
}
func KnuthDurstenfeldShuffle(arr []int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	// copy arr, avoid changing the original data
	// this is a in-place algorithm
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(newArr); i++ {
		p := rand.Intn(i + 1)
		newArr[i], newArr[p] = newArr[p], newArr[i]
	}
	return newArr
}

// leetcode390
func LastRemaining(n int) int {
	k, step, a1, an := 0, 1, 1, n
	cnt := n

	for cnt > 1 {
		if k%2 == 0 { // 正向
			a1 += step
			if cnt%2 == 1 {
				an -= step
			}
		} else { // 反向
			an -= step
			if cnt%2 != 0 {
				a1 += step
			}
		}
		step <<= 1
		cnt >>= 1
		k++
	}
	return a1
}

// leetcode391
type Point struct {
	x, y int
}

func IsRectangleCover(rectangles [][]int) bool {
	totalArea := 0
	minX, minY, maxX, maxY := rectangles[0][0], rectangles[0][1], rectangles[0][2], rectangles[0][3]
	mp := make(map[Point]int)
	for _, rect := range rectangles {
		x1, y1, x2, y2 := rect[0], rect[1], rect[2], rect[3]
		totalArea += (x2 - x1) * (y2 - y1)
		mp[Point{x1, y1}]++
		mp[Point{x2, y2}]++
		mp[Point{x1, y2}]++
		mp[Point{x2, y1}]++
		minX = common.SmallerNumber(minX, x1)
		minY = common.SmallerNumber(minY, y1)
		maxX = common.LargerNumber(maxX, x2)
		maxY = common.LargerNumber(maxY, y2)
	}
	if totalArea != (maxX-minX)*(maxY-minY) {
		return false
	}
	if mp[Point{minX, minY}] != 1 || mp[Point{minX, maxY}] != 1 || mp[Point{maxX, minY}] != 1 || mp[Point{maxX, maxY}] != 1 {
		return false
	}
	delete(mp, Point{minX, minY})
	delete(mp, Point{minX, maxY})
	delete(mp, Point{maxX, minY})
	delete(mp, Point{maxX, maxY})
	for _, v := range mp {
		if v != 2 && v != 4 {
			return false
		}
	}
	return true
}

// leetcode397
// func IntegerReplacement(n int) int {
// 	dp := make([]int, n+1)
// 	for i := 2; i <= n; i++ {
// 		if i%2 == 0 {
// 			dp[i] = dp[i/2] + 1
// 		} else {
// 			dp[i] = algorithm.SmallerNumber(dp[i-1]+1, dp[(i+1)/2]+2)
// 			//dp = dp[i/2+1:]
// 		}
// 	}
// 	return dp[n]
// }

func IntegerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return IntegerReplacement(n/2) + 1
	}
	return 2 + common.SmallerNumber(IntegerReplacement(n/2), IntegerReplacement(n/2+1))
}

// leetcode400
func powInt(a, b int) int {
	ans := 1
	for i := 0; i < b; i++ {
		ans *= a
	}
	return ans
}
func FindNthDigit(n int) int {
	if n <= 9 {
		return n
	}
	length, place := 0, 0
	for length < n {
		place++
		length += 9 * powInt(10, place-1) * place
	}
	for i := place - 1; i >= 1; i-- {
		n -= (9 * powInt(10, i-1) * i)
	}
	num, remain := int(math.Ceil(float64(n)/float64(place))), n%place
	targetNum := strconv.Itoa(powInt(10, place-1) + num - 1)
	pos := -1
	if remain != 0 {
		pos = remain - 1
	} else {
		pos = len(targetNum) - 1
	}
	res, _ := strconv.Atoi(string(targetNum[pos]))
	return res
}
