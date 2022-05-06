package leetcode

import (
	"fmt"
	"math"
	"programs/internal/algorithmingo/algorithm"
	"strings"
)

// leetcode654
func ConstructMaximumBinaryTree(nums []int) *algorithm.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[idx] {
			idx = i
		}
	}
	node := &algorithm.TreeNode{Val: nums[idx]}
	node.Left = ConstructMaximumBinaryTree(nums[:idx])
	node.Right = ConstructMaximumBinaryTree(nums[idx+1:])
	return node
}

// leetcode669
func TrimBST(root *algorithm.TreeNode, low int, high int) *algorithm.TreeNode {
	for root != nil {
		if root.Val <= high && root.Val >= low {
			break
		} else if root.Val > high {
			root = root.Left
		} else if root.Val < low {
			root = root.Right
		}
	}
	if root == nil {
		return nil
	}

	root.Left = TrimBST(root.Left, low, high)
	root.Right = TrimBST(root.Right, low, high)

	return root
}

// leetcode674 最长的递增子数组
func FindLengthOfLCIS(nums []int) int {
	// dp[i] 表示以i结尾时，能够达到的最长的递增子数组长度
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	res := 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
			if res < dp[i] {
				res = dp[i]
			}
		}
	}
	return res
}

func FindLengthOfLCISII(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	cur, res := 1, 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			cur++
		} else {
			if cur > res {
				res = cur
			}
			cur = 1
		}
	}
	if cur > res {
		res = cur
	}
	return res
}

// leetcode686
func RepeatedStringMatch(a string, b string) int {
	lenB := len(b)
	lenA := len(a)
	cnt := int(math.Ceil(float64(lenB) / float64(lenA)))
	repulicateStr := ""
	for i := 0; i < cnt; i++ {
		repulicateStr += a
	}
	fmt.Println(cnt)
	if strings.Contains(repulicateStr, b) {
		return cnt
	} else if strings.Contains(repulicateStr+a, b) {
		return cnt + 1
	}
	return -1
}

// leetcode689
func MaxSumOfThreeSubarrays(nums []int, k int) []int {
	if 3*k > len(nums) || nums == nil {
		return nil
	}
	sum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	leftMax, rightMax := sum[k]-sum[0], sum[len(nums)]-sum[len(nums)-k]

	// 求出以i为右边界时的最大值leftMax
	leftV := make([]int, len(nums))
	leftV[k-1] = 0
	for i := k; i < len(nums); i++ {
		if sum[i+1]-sum[i+1-k] > leftMax {
			leftMax = sum[i+1] - sum[i+1-k]
			leftV[i] = i - k + 1
		} else {
			leftV[i] = leftV[i-1]
		}

	}

	// 求出以i为左边界的最大值rightMax
	rightV := make([]int, len(nums))
	rightV[len(nums)-k] = len(nums) - k
	for i := len(nums) - k - 1; i >= 0; i-- {
		if sum[i+k]-sum[i] >= rightMax {
			rightMax = sum[i+k] - sum[i]
			rightV[i] = i
		} else {
			rightV[i] = rightV[i+1]
		}

	}

	totalSum := 0
	p1, p2, p3 := 0, 0, 0
	for m := k; m <= len(nums)-2*k; m++ {
		l, r := leftV[m-1], rightV[m+k]
		if sum[m+k]-sum[m]+sum[l+k]-sum[l]+sum[r+k]-sum[r] > totalSum {
			totalSum = sum[m+k] - sum[m] + sum[l+k] - sum[l] + sum[r+k] - sum[r]
			p1, p2, p3 = l, m, r
		}
	}
	ans := []int{p1, p2, p3}
	return ans
}

// leetcode700
func SearchBST700(root *algorithm.TreeNode, val int) *algorithm.TreeNode {
	cur := root
	for cur != nil {
		if cur.Val == val {
			return cur
		} else if cur.Val < val {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	return nil
}
