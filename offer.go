package leetcode

import (
	"fmt"
	"math"
	"programs/internal/algorithmingo/algorithm"
	"programs/kit/common"
	"sort"
)

//offer04.二维数组中的查找
func FindNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	i, j := 0, n-1
	for i <= m-1 && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target { // 往小的方向移动
			j--
		} else if matrix[i][j] < target { // 往大的方向移动
			i++
		}
	}
	return false
}

//offer07.重建二叉树
func BuildTree(preorder []int, inorder []int) *algorithm.TreeNode {
	length := len(preorder)
	if length == 0 {
		return nil
	}
	root := &algorithm.TreeNode{Val: preorder[0], Left: nil, Right: nil}
	rootIndex := -1
	for i := 0; i < length; i++ {
		if inorder[i] == preorder[0] { // 找到了树的根，左侧是左子树，右侧是右子树
			rootIndex = i
			break
		}
	}
	root.Left = BuildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = BuildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
	return root
}

//offer42 连续子数组的最大和
func MaxSubArray5(nums []int) int {
	// dp[i] i代表以nums[i]结尾的连续字数组的最大和
	dp := make([]int, len(nums))
	var res int
	dp[0], res = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = common.LargerNumber(dp[i-1]+nums[i], nums[i])
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

// ****************************
//offerII03 前n个数字二进制中1的个数
func CountBits(n int) []int {
	dp := make([]int, n)
	base := 2
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i] = 0
		} else if i == 1 {
			dp[i] = 1
		} else if i == 2 {
			dp[i] = 1
		} else {
			if i-base == base {
				dp[i] = 1
				base *= 2
			} else {
				dp[i] = dp[base] + dp[i-base]
			}
		}
	}
	return dp
}

//offerII06
// 双指针的解法
func TwoSumInOffer(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		if numbers[i]+numbers[j] < target {
			i++
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			return []int{i, j}
		}
	}
	return []int{}
}

//offerII07
func ThreeSumInOffer(nums []int) [][]int {
	if len(nums) < 2 {
		return [][]int{}
	}
	sort.Ints(nums)

	ans := [][]int{}
	prev := nums[0] - 1

	twoSum := func(arr []int, target int) {
		i, j := 0, len(arr)-1
		for i < j {
			if arr[i]+arr[j] < target {
				i++
			} else if arr[i]+arr[j] > target {
				j--
			} else {
				ans = append(ans, []int{-target, arr[i], arr[j]})
				cur := arr[i]
				for arr[i] == cur {
					i++
				}
			}
		}
	}
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == prev {
			continue
		}
		twoSum(nums[i+1:], -nums[i])
		prev = nums[i]
	}
	fmt.Println(ans)
	return ans
}

//offerII08
func MinSubArrayLen2(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i, j, sum, res := 0, 0, nums[0], math.MaxInt32
	if nums[0] >= target {
		res = 1
	}

	for j = 1; j < len(nums); j++ {
		if sum < target {
			sum += nums[j]
		}
		for sum >= target {
			sum -= nums[i]
			res = common.SmallerNumber(res, j-i+1)
			i++
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func NumSubarrayProductLessThanK(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	i, j, mul, res := 0, 0, 1, 0
	for j = 0; j < len(nums); j++ {
		mul *= nums[j]

		for i <= j && mul >= k {
			mul /= nums[i]
			i++
		}
		if i <= j {
			res += (j - i + 1)
		} else {
			res += 0
		}
	}
	return res
}

//offerII058
func ReverseLeftWords(s string, n int) string {
	bytes := []byte(s)
	ReverseStr(bytes)
	ReverseStr(bytes[:len(bytes)-n])
	ReverseStr(bytes[len(bytes)-n:])
	return string(bytes)
}

func ReverseStr(bytes []byte) {
	n := len(bytes)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
}

//offerII088
func MinCostClimbingStairs2(cost []int) int {
	dp := make([]int, len(cost)+5)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= len(cost); i++ {
		dp[i] = common.SmallerNumber(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}

func DeleteText(article string, index int) string {
	if article[index] == ' ' {
		return article
	}
	bytes := []byte(article)
	l, r := index, index
	for l >= 0 && article[l] != ' ' {
		l--
	}
	for r < len(bytes) && article[r] != ' ' {
		r++
	}
	ans := string(bytes[:l]) + string(bytes[r:])
	return ans
}

func NumFlowers(roads [][]int) int {
	return 2
}
