package leetcode

import (
	"container/heap"
	"errors"
	"fmt"
	"programs/internal/algorithmingo/algorithm"
	"programs/kit/utils"
	"sort"
	"strconv"
	"strings"
)

// leetcode257
func BinaryTreePaths(root *algorithm.TreeNode) []string {
	var dfs func(*algorithm.TreeNode, string)
	res := []string{}
	dfs = func(root *algorithm.TreeNode, cur string) {
		if root != nil {
			cur += strconv.Itoa(root.Val)
		}
		if root.Left == nil && root.Right == nil {
			res = append(res, cur)
			return
		}
		cur += "->"
		if root.Left != nil {
			dfs(root.Left, cur)
		}
		if root.Right != nil {
			dfs(root.Right, cur)
		}
	}
	dfs(root, "")
	fmt.Println(res)
	return res
}

// leetcode258
func AddDigits(num int) int {
	tmp := num
	sum := 0
	for tmp > 10 {
		for tmp != 0 {
			sum += tmp % 10
			tmp /= 10
		}
		tmp = sum
	}
	return sum
}

// leetcode268
func MissingNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res += i
		res -= nums[i]
	}
	res += len(nums)
	return res
}

// leetcode273
var trans map[int]string

func NumberToWords(num int) (res string) {
	trans = map[int]string{
		0: "Zero", 1: "One", 2: "Two", 3: "Three", 4: "Four", 5: "Five", 6: "Six", 7: "Seven", 8: "Eight", 9: "Nine",
		10: "Ten", 11: "Eleven", 12: "Twelve", 13: "Thirteen", 14: "Fourteen", 15: "Fifteen", 16: "Sixteen", 17: "Seventeen", 18: "Eighteen", 19: "Nineteen",
		20: "Twenty", 30: "Thirty", 40: "Forty", 50: "Fifty", 60: "Sixty", 70: "Seventy", 80: "Eighty", 90: "Ninety", 100: "Hundred",
	}

	fourth := num / 1000000000
	num %= 1000000000
	if fourth != 0 {
		tmp, _ := SmallNumberToWord(fourth)
		res += strings.TrimSpace(tmp)
		res += " Billon"
	}

	third := num / 1000000
	num %= 1000000
	if third != 0 {
		tmp, _ := SmallNumberToWord(third)
		res += " "
		res += strings.TrimSpace(tmp)

		res += " Millon"
	}

	second := num / 1000
	num %= 1000
	if second != 0 {
		tmp, _ := SmallNumberToWord(second)
		res += " "
		res += strings.TrimSpace(tmp)
		res += " Thousand"
	}

	if num != 0 {
		tmp, _ := SmallNumberToWord(num)
		res += " "
		res += strings.TrimSpace(tmp)
	}

	return strings.TrimSpace(res)
}

func SmallNumberToWord(num int) (res string, err error) {
	if num > 999 || num < 0 {
		return "", errors.New("wrong number")
	}
	high := num / 100
	if high != 0 {
		res += " "
		res += trans[high]
		res += " "
		res += trans[100]
	}
	remain := num % 100
	if remain <= 19 && remain >= 10 {
		res += " "
		res += trans[remain]
	} else {
		mid := remain / 10
		if mid != 0 {
			res += " "
			res += trans[mid*10]
		}
		low := remain % 10
		if low != 0 {
			res += " "
			res += trans[low]
		}
	}
	return res, nil
}

// leetcode274

func HIndex(citations []int) int {
	sort.Ints(citations)
	mp := make(map[int]int)
	cnt := 1
	for i := len(citations) - 1; i >= 0; i-- {
		mp[citations[i]] = cnt
		cnt++
	}

	max := 1
	for key, value := range mp {
		x := utils.MinNum(key, value)
		if x > max {
			max = x
		}
	}
	return max
}

// leetcode275
func HIndex2(citations []int) int {
	h := 0
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] > h {
			h++
		}
	}
	return h
}

// leetcode279
func NumSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i*i <= n; i++ {
		for j := i * i; j <= n; j++ {
			if dp[j] == 0 {
				dp[j] = dp[j-i*i] + 1
			} else {
				dp[j] = utils.MinNum(dp[j], dp[j-i*i]+1)
			}
		}
	}
	return dp[n]
}

// leetcode283
func MoveZeroes(nums []int) {
	slow, fast := 0, 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			// if slow == fast {
			//     slow++
			//     continue
			// }
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}

// leetcode295
// 超时的做法 o(N) 移动数组元素时间复杂度太高
// type MedianFinder struct {
// 	arr  []int
// 	size int
// }

// func NewMedianFinder() MedianFinder {
// 	return MedianFinder{arr: []int{}, size: 0}
// }

// func (mf *MedianFinder) AddNum(num int) {
// 	l, r := 0, len(mf.arr)-1
// 	for l <= r {
// 		m := (l + r) / 2
// 		if mf.arr[m] <= num {
// 			l = m + 1
// 		} else {
// 			r = m - 1
// 		}
// 	}
// 	mf.arr = append(mf.arr[:l], append([]int{num}, mf.arr[l:]...)...)
// 	mf.size++
// }

// func (mf *MedianFinder) FindMedian() float64 {
// 	if mf.size == 0 {
// 		return 0.0
// 	}
// 	if mf.size%2 == 0 {
// 		return float64(mf.arr[mf.size/2]+mf.arr[mf.size/2-1]) / 2.0
// 	} else {
// 		return float64(mf.arr[mf.size/2])
// 	}
// }

type hp295 struct {
	sort.IntSlice
}

func (h *hp295) Push(x interface{}) { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h *hp295) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

type MedianFinder struct {
	hpMin, hpMax hp295
	cnt          int
}

func NewMedianFinder() MedianFinder {
	return MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {
	mf.cnt++
	// 奇数时hpMin始终多一个
	if mf.hpMin.Len() == 0 || num >= mf.hpMin.IntSlice[0] {
		heap.Push(&mf.hpMin, num)
		if len(mf.hpMax.IntSlice)+1 < len(mf.hpMin.IntSlice) {
			heap.Push(&mf.hpMax, -mf.hpMin.IntSlice[0])
			heap.Pop(&mf.hpMin)
		}
	} else {
		heap.Push(&(mf.hpMax), -num)
		if len(mf.hpMax.IntSlice) > len(mf.hpMin.IntSlice) {
			heap.Push(&mf.hpMin, -mf.hpMax.IntSlice[0])
			heap.Pop(&mf.hpMax)
		}
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	minQ, maxQ := mf.hpMin, mf.hpMax
	if minQ.Len() > maxQ.Len() {
		return float64(minQ.IntSlice[0])
	}
	return float64(-maxQ.IntSlice[0]+minQ.IntSlice[0]) / 2

}

// leetcode299
func GetHint(secret string, guess string) string {
	mp1 := map[int]int{}
	mp2 := map[int]int{}
	aNum := 0
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			aNum++
		} else {
			index1 := int(secret[i] - '0')
			index2 := int(guess[i] - '0')
			mp1[index1]++
			mp2[index2]++
		}
	}
	bNum := 0
	for k, v := range mp2 {
		bNum += utils.MinNum(mp1[k], v)
	}
	str := fmt.Sprintf("%dA%dB", aNum, bNum)
	fmt.Println(str)
	return str
}

// leetcode300

func LengthOfLIS(nums []int) int {
	// Time O(N²)
	dp := make([]int, len(nums))
	maxV := -1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = utils.MaxNum(dp[i], dp[j]+1)
			}
		}
		if dp[i] > maxV {
			maxV = dp[i]
		}
	}
	return maxV
}

func LengthOfLIS2(nums []int) int {
	// Time O(NlogN)
	ceil := []int{}
	for i := 0; i < len(nums); i++ {
		if len(ceil) == 0 {
			ceil = append(ceil, nums[i])
		}
		l, r := 0, len(ceil)-1
		for l <= r {
			mid := (l + r) / 2
			if ceil[mid] < nums[i] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		if l == len(ceil) {
			ceil = append(ceil, nums[i])
		} else {
			ceil[l] = nums[i]
		}
	}
	return len(ceil)
}

// leetcode300
func LengthOfLIS300(nums []int) int {
	// 最长递增子序列
	binarySearch := func(nums []int, target int) int {
		l, r := 0, len(nums)-1
		for l <= r {
			m := (l + r) / 2
			if nums[m] <= target {
				l = m + 1
			} else {
				r = m - 1
			}
		}
		return r
	}
	stack := []int{}
	for i := 0; i < len(nums); i++ {
		if len(stack) == 0 {
			stack = append(stack, nums[i])
		} else {
			top := stack[len(stack)-1]
			if nums[i] > top {
				// 严格递增
				stack = append(stack, nums[i])
			} else {
				// 找最后一个小于等于nums[i]
				// 小于等于nums[i]的右边界
				pos := binarySearch(stack, nums[i])
				stack[pos] = nums[i]
			}
		}
	}
	return len(stack)
}

// leetcode300
func LengthOfLIS300II(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	// dp[i] 表示到i位置最长上升子序列的长度
	for i := 0; i < n; i++ {
		// 状态转移方程
		for j := i - 1; j >= 0; j-- {
			// 找比i小的
			if nums[j] < nums[i] {
				dp[i] = utils.MaxNum(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n-1]
}

func LengthOfLISIII(nums []int) int {
	var maxN int = 3000
	n := len(nums)

	// a是原来的数组，b是用于离散化
	a := make([]int, n+1)
	b := make([]int, n+1)
	copy(a[1:], nums)
	copy(b[1:], nums)
	// c是树状数组
	c := make([]int, maxN)

	// 离散化操作，先排序
	sort.Ints(b[1:])
	// slow 永远代表下一个要存储的地方
	slow := 2
	for fast := 2; fast <= n; fast++ {
		if b[fast] != b[slow-1] {
			b[slow] = b[fast]
			slow++
		}
	}
	// slow同时也是去重之后数组元素的长度
	b = b[:slow]
	cnt := slow
	lsh := func(b []int, target int) int {
		l, r := 0, len(b)-1
		for l <= r {
			m := (l + r) / 2
			if b[m] == target {
				return m
			} else if b[m] < target {
				l = m + 1
			} else {
				r = m - 1
			}
		}
		// 必定能找到，所以return -1无用
		return -1
	}

	lowbit := func(x int) int {
		return x & -x
	}

	modify := func(x, t int) {
		for x <= cnt {
			c[x] = utils.MaxNum(c[x], t)
			x += lowbit(x)
		}
	}

	query := func(x int) int {
		res := 0
		for x > 0 {
			res = utils.MaxNum(res, c[x])
			x -= lowbit(x)
		}
		return res
	}

	res := 0
	for i := 1; i <= n; i++ {
		// a[i]经过了离散化，得到的是一个下标
		a[i] = lsh(b[1:], a[i]) + 1
		val := query(a[i]-1) + 1
		res = utils.MaxNum(res, val)
		// 更新下标i处的最大值
		modify(a[i], val)
	}

	return res
}
