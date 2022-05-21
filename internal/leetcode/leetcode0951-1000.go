package leetcode

import (
	"fmt"
	"programs/internal/algorithmingo/algorithm"
	"sort"
)

// leetcode913
func IsAlienSorted(words []string, order string) bool {
	mp := make([]int, 26)
	for i := 0; i < len(order); i++ {
		mp[int(order[i]-'a')] = i
	}
	cmp := func(x, y string) bool {
		// 小于返回true，大于返回false
		for i := 0; i < len(x) && i < len(y); i++ {
			if mp[int(x[i]-'a')] > mp[int(y[i]-'a')] {
				return false
			} else if mp[int(x[i]-'a')] < mp[int(y[i]-'a')] {
				return true
			}
		}
		return len(x) <= len(y)
	}
	for i := 1; i < len(words); i++ {
		// 小于返回true，大于返回false
		if !cmp(words[i-1], words[i]) {
			return false
		}
	}
	return true
}

// leetcode961
func RepeatedNTimes(nums []int) int {
	n := len(nums)
	mp := make(map[int]struct{}, n)
	for _, val := range nums {
		if _, has := mp[val]; has {
			return val
		}
		mp[val] = struct{}{}
	}
	return -1
}

// leetcode977
func SortedSquares1(nums []int) []int {
	i, j := 0, len(nums)-1
	res, k := make([]int, len(nums)), len(nums)-1
	for i <= j {
		if nums[i]*nums[i] < nums[j]*nums[j] {
			res[k] = nums[j] * nums[j]
			j--
			k--
		} else {
			res[k] = nums[i] * nums[i]
			i++
			k--
		}
	}
	return res
}

// leetcode981
type pair struct {
	timestamp int
	value     string
}

type TimeMap struct {
	m map[string][]pair
}

func NewTimeMap() TimeMap {
	return TimeMap{map[string][]pair{}}
}

func (timeMap *TimeMap) Set(key string, value string, timestamp int) {
	timeMap.m[key] = append(timeMap.m[key], pair{value: value, timestamp: timestamp})
}

func (timeMap *TimeMap) Get(key string, timestamp int) string {
	pairs := timeMap.m[key]
	i := sort.Search(len(pairs), func(i int) bool { return pairs[i].timestamp > timestamp })
	if i > 0 {
		return pairs[i-1].value
	}
	return ""
}

var mp map[int]map[int][]int
var left int
var right int

func VerticalTraversal(root *algorithm.TreeNode) [][]int {
	mp = make(map[int]map[int][]int)
	res := make([][]int, 0)
	left, right = 0, 0
	preorderTraversal(root, 0, 0)
	cur := 0
	for i := left; i <= right; i++ {
		res = append(res, []int{})
		if mpRow, has := mp[i]; has {
			keys := []int{}
			for key := range mpRow {
				keys = append(keys, key)
			}
			sort.Ints(keys)
			for j := range keys {
				data := mpRow[keys[j]]
				sort.Ints(data)
				res[cur] = append(res[cur], data...)
			}
		}
		cur++
	}
	fmt.Println(res)
	return res
}

// leetcode987
func preorderTraversal(root *algorithm.TreeNode, row, col int) {
	if root == nil {
		return
	}
	if col < left {
		left = col
	}
	if col > right {
		right = col
	}
	if mp[col] == nil {
		mp[col] = make(map[int][]int)
	}
	mp[col][row] = append(mp[col][row], root.Val)

	preorderTraversal(root.Left, row+1, col-1)
	preorderTraversal(root.Right, row+1, col+1)
}
