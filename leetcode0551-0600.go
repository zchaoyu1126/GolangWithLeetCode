package leetcode

import (
	"fmt"
	"programs/internal/algorithmingo/algorithm"
	"sort"
)

type Node559 struct {
	Val      int
	Children []*Node559
}

// leetcode559
func MaxDepth559(root *Node559) int {
	if root == nil {
		return 0
	}
	res := 0
	for i := 0; i < len(root.Children); i++ {
		tmp := MaxDepth559(root.Children[i])
		if tmp > res {
			res = tmp
		}
	}
	return res + 1
}

// leetcode563
func FindTilt(root *algorithm.TreeNode) int {
	var traverse func(root *algorithm.TreeNode, sum *int) int
	sum := 0
	traverse = func(root *algorithm.TreeNode, sum *int) int {
		if root == nil {
			return 0
		}
		lValue := traverse(root.Left, sum)
		rValue := traverse(root.Right, sum)
		(*sum) += abs_int(rValue - lValue)
		return lValue + rValue + root.Val
	}
	traverse(root, &sum)
	return sum
}

// leetcode573
func Fib(n int) int {
	preOne, preTwo := 0, 1
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	res := 0
	for n-2 >= 0 {
		n--
		res = preOne + preTwo
		preOne = preTwo
		preTwo = res
	}
	return res
}

// leetcode575
func DistributeCandies(candyType []int) int {
	mp := make(map[int]int)
	for i := 0; i < len(candyType); i++ {
		mp[candyType[i]] = i
	}
	res := len(mp)
	if res > len(candyType)/2 {
		return len(candyType) / 2
	}
	return res
}

// leetcode594
func FindLHS(nums []int) int {
	sort.Ints(nums)
	cnt1, cnt2, res := 0, 0, 0
	start, end := nums[0], nums[0]-1
	for i := 0; i < len(nums); i++ {
		if nums[i] == start {
			cnt1++
			//fmt.Println(start, end, nums[i], cnt1, cnt2, res, 1)
		} else if nums[i] == end {
			cnt2++
			//fmt.Println(start, end, nums[i], cnt1, cnt2, res, 2)
		} else if nums[i] == start+1 {
			end = nums[i]
			cnt2++
			//fmt.Println(start, end, nums[i], cnt1, cnt2, res, 3)
		} else if nums[i] == end+1 {
			if cnt1+cnt2 > res {
				res = cnt1 + cnt2
			}
			cnt1 = cnt2
			cnt2 = 1
			start = end
			end = nums[i]
			//fmt.Println(start, end, nums[i], cnt1, cnt2, res, 4)
		} else if nums[i] > end+1 {
			if cnt1+cnt2 > res {
				res = cnt1 + cnt2
			}
			cnt1 = 1
			cnt2 = 0
			start = nums[i]
			end = nums[i] - 1
			//fmt.Println(start, end, nums[i], cnt1, cnt2, res, 5)
		}
	}
	fmt.Println(start, end, res)
	if start == end-1 {
		if cnt1+cnt2 > res {
			res = cnt1 + cnt2
		}
	}
	fmt.Println(res)
	return res
}
