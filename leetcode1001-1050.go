package leetcode

import (
	"container/heap"
	"math"
	"math/rand"
	"programs/kit/common"
	"sort"
	"time"
)

// leetcode1005
func arrSum(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res += nums[i]
	}
	return res
}
func LargestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	cnt := k
	for i := 0; i < len(nums) && cnt > 0; i++ {
		if nums[i] < 0 {
			nums[i] = -nums[i]
			cnt--
		}
	}
	if cnt%2 == 0 {
		return arrSum(nums)
	} else {
		hasZero, min := false, 0xffffffff
		for i := 0; i < len(nums); i++ {
			if nums[i] == 0 {
				hasZero = true
			}
			if nums[i] < min {
				min = nums[i]
			}
		}
		if hasZero {
			return arrSum(nums)
		} else {
			return arrSum(nums) - 2*min
		}
	}
}

// leetcode1034
type Coordinate struct {
	x int
	y int
}

func ColorBorder(grid [][]int, row int, col int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	tag := make([][]int, m)
	for i := range tag {
		tag[i] = make([]int, n)
	}
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
		copy(res[i], grid[i])
	}
	queue := []Coordinate{}
	queue = append(queue, Coordinate{row, col})
	tag[row][col] = 1
	dx := []int{0, 0, 1, -1}
	dy := []int{-1, 1, 0, 0}
	for len(queue) != 0 {
		top := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		x, y := top.x, top.y
		cnt := 0
		for i := 0; i < 4; i++ {
			newx, newy := x+dx[i], y+dy[i]
			if newx >= 0 && newx < m && newy >= 0 && newy < n {
				if grid[newx][newy] == grid[x][y] {
					cnt++
					if tag[newx][newy] == 0 {
						tag[newx][newy] = 1
						queue = append(queue, Coordinate{newx, newy})
					}
				}
			}
		}
		if cnt < 4 {
			res[x][y] = color
		}
	}
	return res
}

// leetcode1035
func MaxUncrossedLines(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = common.LargerNumber(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}

// leetcode1044
func randInt(a, b int) int {
	return a + rand.Intn(b-a)
}

func pow(x, n, mod int) int {
	res := 1
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func check(arr []byte, length int, a1, a2 int, mod1, mod2 int) int {
	aL1, aL2 := pow(a1, length, mod1), pow(a2, length, mod2)
	h1, h2 := 0, 0
	for _, c := range arr[:length] {
		h1 = (h1*a1 + int(c)) % mod1
		h2 = (h2*a2 + int(c)) % mod2
	}
	seen := map[[2]int]bool{{h1, h2}: true}
	for start := 1; start <= len(arr)-length; start++ {
		h1 = (h1*a1 - int(arr[start-1])*aL1 + int(arr[start+length-1])) % mod1
		h2 = (h2*a2 - int(arr[start-1])*aL2 + int(arr[start+length-1])) % mod2
		if h1 < 0 {
			h1 += mod1
		}
		if h2 < 0 {
			h2 += mod2
		}
		if seen[[2]int{h1, h2}] {
			return start
		}
		seen[[2]int{h1, h2}] = true
	}
	return -1
}
func LongestDupSubstring(s string) string {
	rand.Seed(time.Now().UnixNano())
	a1, a2 := randInt(26, 100), randInt(26, 100)
	mod1, mod2 := randInt(1e9+7, math.MaxInt32), randInt(1e9+7, math.MaxInt32)
	arr := []byte(s)
	for i := range arr {
		arr[i] -= 'a'
	}
	l, r := 1, len(s)-1
	length, start := 0, -1
	for l <= r {
		m := (l + r) / 2
		idx := check(arr, m, a1, a2, mod1, mod2)
		if idx != -1 {
			l = m + 1
			length = m
			start = idx
		} else {
			r = m - 1
		}
	}
	if start == -1 {
		return ""
	}
	return s[start : start+length]
}

// leetcode1046
type IntSliceHeap []int

func (ls IntSliceHeap) Len() int {
	return len(ls)
}
func (ls IntSliceHeap) Swap(i, j int) {
	ls[i], ls[j] = ls[j], ls[i]
}
func (ls IntSliceHeap) Less(i, j int) bool {
	return ls[i] > ls[j]
}
func (ls *IntSliceHeap) Push(x interface{}) {
	*ls = append(*ls, x.(int))
}
func (ls *IntSliceHeap) Pop() interface{} {
	old := *ls
	n := len(old)
	x := old[n-1]
	*ls = old[0 : n-1]
	return x
}
func LastStoneWeight(stones []int) int {
	h := IntSliceHeap(stones)
	heap.Init(&h)
	for len(h) >= 2 {
		y := heap.Pop(&h).(int)
		x := heap.Pop(&h).(int)
		if x == y {
			continue
		} else {
			heap.Push(&h, y-x)
		}
	}
	if len(h) == 0 {
		return 0
	} else {
		return h[0]
	}
}

// leetcode1047
func RemoveDuplicates1047(s string) string {
	bytes := []byte(s)
	stack := []byte{}

	for i := 0; i < len(bytes); i++ {
		// 如果由重复的那就全部清除
		if len(stack) != 0 && stack[len(stack)-1] == bytes[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, bytes[i])
		}
	}
	return string(stack)
}
