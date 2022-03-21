package leetcode

import (
	"math"
	"programs/internal/algorithmingo/algorithm"
	"strconv"
	"strings"
)

// leetcode162
func FindPeakElement(nums []int) int {
	n := len(nums)

	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}

	l, r := 0, n-1
	for {
		mid := (l + r) / 2
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
}

// leetcode165
func CompareVersion(version1 string, version2 string) int {
	res1 := strings.Split(version1, ".")
	res2 := strings.Split(version2, ".")
	i, j := 0, 0
	for ; i < len(res1) && j < len(res2); i, j = i+1, j+1 {
		num1, _ := strconv.Atoi(res1[i])
		num2, _ := strconv.Atoi(res2[j])
		if num1 < num2 {
			return -1
		} else if num1 > num2 {
			return 1
		}
	}
	for ; i < len(res1); i++ {
		num, _ := strconv.Atoi(res1[i])
		if num > 0 {
			return 1
		}
	}
	for ; j < len(res2); j++ {
		num, _ := strconv.Atoi(res2[j])
		if num > 0 {
			return -1
		}
	}
	return 0
}

// leetcode168
func ConvertToTitle(columnNumber int) string {
	var single byte
	var res string
	if columnNumber == 0 {
		return ""
	}
	for columnNumber != 0 {
		columnNumber--
		single = byte('A' + columnNumber%26)
		res = string(single) + res
		columnNumber /= 26
	}
	return res
}

// leetcode187
func FindRepeatedDnaSequences(s string) []string {
	root := algorithm.NewRuneTrie()
	res := []string{}
	for i := 0; i+10 <= len(s); i++ {
		if !root.Find(s[i : i+10]) {
			root.Insert(s[i : i+10])
		} else {
			res = append(res, s[i:i+10])
		}
	}
	return RemoveRepByMap(res)
}

func RemoveRepByMap(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

// leetcode199
func RightSideView(root *algorithm.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := [][]int{}

	queue := []*algorithm.TreeNode{root}
	for len(queue) != 0 {
		size := len(queue)
		cur := []int{}
		for i := 0; i < size; i++ {
			front := queue[0]
			queue = queue[1:]
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
			cur = append(cur, front.Val)
		}
		res = append(res, cur)
	}
	ans := []int{}
	for _, cur := range res {
		ans = append(ans, cur[len(cur)-1])
	}
	return ans
}
