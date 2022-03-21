package leetcode

import (
	"container/heap"
	"fmt"
	"programs/internal/algorithmingo/algorithm"
)

// leetcode307
type NumArray struct {
	*algorithm.SegTree
}

func NewArray(nums []int) NumArray {
	return NumArray{algorithm.NewSegTree(nums)}
}

func (t *NumArray) Update(index int, val int) {
	t.UpdateValByIndex(index, val)
}

func (t *NumArray) SumRange(left int, right int) int {
	return t.SumRangeBetweenLR(left, right)
}

// leetcode318
func MaxProduct(words []string) int {
	masks := make([]int, len(words))
	for i, word := range words {
		for _, ch := range word {
			masks[i] |= 1 << (ch - 'a')
		}
	}

	length := 0
	for i := 0; i < len(masks); i++ {
		for j := i + 1; j < len(masks); j++ {
			if masks[i]&masks[j] == 0 {
				if len(words[i])*len(words[j]) > length {
					length = len(words[j]) * len(words[i])
				}
			}
		}
	}
	return length
}

// leetcode326
func IsPowerOfThree(n int) bool {
	l := 1
	if n == 1 {
		return true
	}
	for n > l {
		l *= 3
	}
	return n == l
}

// leetcode334
func IncreasingTriplet(nums []int) bool {
	// 最长上升子序列
	stack := []int{}
	for i := 0; i < len(nums); i++ {
		if len(stack) == 0 {
			stack = append(stack, nums[i])
			continue
		}
		top := stack[len(stack)-1]
		if top >= nums[i] {
			// 找到第一个比nums[i]小的数进行修改
			cur := len(stack) - 1
			for cur >= 0 && stack[cur] > nums[i] {
				cur--
			}
			stack[cur+1] = nums[i]
		} else {
			stack = append(stack, nums[i])
		}
	}
	fmt.Println(stack)
	return len(stack) >= 3
}

// leetcode335
func IsSelfCrossing(distance []int) bool {
	for i := 3; i < len(distance); i++ {
		// 第 1 类路径交叉的情况
		if distance[i] >= distance[i-2] && distance[i-1] <= distance[i-3] {
			return true
		}

		// 第 2 类路径交叉的情况
		if i == 4 && distance[3] == distance[1] &&
			distance[4] >= distance[2]-distance[0] {
			return true
		}

		// 第 3 类路径交叉的情况
		if i >= 5 && distance[i-3]-distance[i-5] <= distance[i-1] &&
			distance[i-1] <= distance[i-3] &&
			distance[i] >= distance[i-2]-distance[i-4] &&
			distance[i-2] > distance[i-4] {
			return true
		}
	}
	return false
}

// leetcode344
func ReverseString(s []byte) {
	n := len(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		tmp := s[j]
		s[j] = s[i]
		s[i] = tmp
	}
}

// leetcode345
func ReverseVowels(s string) string {
	i, j := 0, len(s)-1
	data := []byte(s)
	for i < j {
		for i <= len(s)-1 && !isAEIOU(data[i]) {
			i++
		}
		for j >= 0 && !isAEIOU(data[j]) {
			j--
		}
		if i < j {
			data[i], data[j] = data[j], data[i]
		}
		i++
		j--
	}
	fmt.Println(string(data))
	return string(data)
}

func isAEIOU(x byte) bool {
	if x == 'a' || x == 'A' || x == 'e' || x == 'E' || x == 'i' || x == 'I' || x == 'o' || x == 'O' || x == 'U' || x == 'u' {
		return true
	}
	return false
}

// leetcode347
type frequency struct {
	val int
	cnt int
}
type freHeap []frequency

func (fh freHeap) Len() int            { return len(fh) }
func (fh freHeap) Less(i, j int) bool  { return fh[i].cnt > fh[j].cnt }
func (fh freHeap) Swap(i, j int)       { fh[i], fh[j] = fh[j], fh[i] }
func (fh *freHeap) Push(x interface{}) { *fh = append(*fh, x.(frequency)) }
func (fh *freHeap) Pop() interface{}   { top := (*fh)[:len(*fh)-1]; *fh = (*fh)[:len(*fh)-1]; return top }
func TopKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}
	h := freHeap{}
	for key, val := range mp {
		h = append(h, frequency{key, val})
	}
	heap.Init(&h)
	ans := []int{}
	for i := 0; i < k; i++ {
		ans = append(ans, h[0].val)
		heap.Pop(&h)
	}
	return ans
}

// leetcode349
func Intersection(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		mp[nums1[i]] = 1
	}
	res := []int{}
	for i := 0; i < len(nums2); i++ {
		if val, ok := mp[nums2[i]]; ok && val > 0 {
			res = append(res, nums2[i])
			mp[nums2[i]]--
		}
	}
	return res
}

func Intersect(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		mp[nums1[i]]++
	}
	res := []int{}
	for i := 0; i < len(nums2); i++ {
		if cnt, ok := mp[nums2[i]]; ok && cnt > 0 {
			res = append(res, nums2[i])
			mp[nums2[i]]--
		}
	}
	return res
}

// leetcode350
func IsHappy(n int) bool {
	mp := make(map[int]int)
	for {
		sum := digitalSum(n)
		if sum == 1 {
			return true
		}
		if _, ok := mp[sum]; !ok {
			mp[sum]++
		} else {
			return false
		}
		n = sum
	}
}

func digitalSum(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
