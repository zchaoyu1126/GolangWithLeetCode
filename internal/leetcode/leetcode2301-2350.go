package leetcode

import (
	"programs/internal/algorithmingo/algorithm"
	"sort"
	"strings"
)

// leetcode2303
func CalculateTax(brackets [][]int, income int) float64 {
	res := 0.0
	prev := brackets[0][0]
	if income >= prev {
		res += float64(brackets[0][1]) * float64(prev) / 100.0
	} else {
		res += float64(income) * float64(brackets[0][1]) / 100.0
		return res
	}

	for i := 1; i < len(brackets); i++ {
		if brackets[i][0] > income {
			res += float64(income-prev) * float64(brackets[i][1]) / 100.0
			break
		}
		res += float64(brackets[i][0]-prev) * float64(brackets[i][1]) / 100.0
		prev = brackets[i][0]
	}
	return res
}

// leetcode2305
var res2305 int

func DistributeCookies(cookies []int, k int) int {
	res2305 = 0
	sort.Sort(sort.Reverse(sort.IntSlice(cookies)))
	num := make([]int, 8)
	num[0] += cookies[0]
	traceback(cookies, 1, num)
	return res2305
}
func traceback(cookies []int, pid int, num []int) {
	// 下面要开始分配第pid包
	if pid >= len(cookies) {
		maxNum := 0
		for i := 0; i < len(num); i++ {
			if num[i] > maxNum {
				maxNum = num[i]
			}
		}
		if maxNum < res2305 {
			res2305 = maxNum
		}
		return
	}
	for j := 0; j < len(num); j++ {
		// 将pid包分配给了j号小朋友
		num[j] += cookies[pid]
		traceback(cookies, pid+1, num)
		// 取消pid包的分配
		num[j] -= cookies[pid]
	}
}

// leetcode2306
// func DistinctNames(ideas []string) int64 {
// 	mp := make(map[string]int32)
// 	cnt := make(map[string]int)

// 	for _, idea := range ideas {
// 		head := idea[0]
// 		suffix := idea[1:]
// 		mp[suffix] |= 1 << int(head-'a')
// 		cnt[suffix]++
// 	}

// 	var res int64
// 	arr1 := make([]int32, 0, len(mp))
// 	arr2 := make([]int, 0, len(mp))
// 	for k, v := range mp {
// 		arr1 = append(arr1, v)
// 		arr2 = append(arr2, cnt[k])
// 	}

// 	for i := 0; i < len(arr1); i++ {
// 		for j := i + 1; j < len(arr1); j++ {
// 			num := bits.OnesCount(uint(arr1[i] & arr1[j]))
// 			cnt1 := arr2[i] - num
// 			cnt2 := arr2[j] - num
// 			res += int64(cnt1 * cnt2)
// 		}
// 	}
// 	return res * 2
// }
func DistinctNames(ideas []string) int64 {
	group := make(map[string]int)

	for _, idea := range ideas {
		head := idea[0]
		suffix := idea[1:]
		group[suffix] |= 1 << int(head-'a')
	}

	var res int
	var cnt [26][26]int
	for _, mask := range group {
		for i := 0; i < 26; i++ {
			if mask>>i&1 == 0 {
				// 包含 i 不包含j的数目
				// 含i 不含j
				for j := 0; j < 26; j++ {
					if mask>>j&1 > 0 {
						cnt[i][j]++
					}
				}
			} else {
				// 不包含i
				for j := 0; j < 26; j++ {
					// 含j不含i
					if mask>>j&1 == 0 {
						res += cnt[i][j]
					}
				}
			}
		}
	}

	return int64(res * 2)
}

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	res := 0
	for i := 0; i < len(energy); i++ {
		if initialEnergy <= energy[i] {
			res += energy[i] - initialEnergy + 1
			initialEnergy = energy[i] + 1
		}
		initialEnergy -= energy[i]
	}

	for i := 0; i < len(experience); i++ {
		if initialExperience <= experience[i] {
			res += experience[i] - initialExperience + 1
			initialEnergy = experience[i] + 1
		}
		initialExperience += experience[i]
	}
	return res
}

func largestPalindromic(num string) string {
	mp := make([]int, 10)
	for i := 0; i < len(num); i++ {
		mp[(num[i]-'0')]++
	}
	var sb strings.Builder
	stack := []byte{}
	// 偶数的结束了
	flag := false // 允许填入0
	for i := 9; i >= 0; i-- {
		if i == 0 && !flag {
			break
		}
		for mp[i] >= 2 {
			sb.WriteByte(byte('0' + i))
			stack = append(stack, byte('0'+i))
			mp[i] -= 2
			flag = true
		}
	}
	for i := 9; i >= 0; i-- {
		if mp[i] > 0 {
			sb.WriteByte(byte('0' + i))
			break
		}
	}
	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		sb.WriteByte(top)
	}	
	return sb.String()
}

func amountOfTime(root *algorithm.TreeNode, start int) int {
	if root == nil {
		return 0
	}

	mp := make(map[*algorithm.TreeNode]*algorithm.TreeNode)
	visited := make(map[*algorithm.TreeNode]struct{})
	mp[root] = nil
	var dfs func(root *algorithm.TreeNode)
	var infectionNode *algorithm.TreeNode
	dfs = func(root *algorithm.TreeNode) {
		if root.Val == start {
			infectionNode = root
		}
		if root.Left != nil {
			mp[root.Left] = root
			dfs(root.Left)
		}
		if root.Right != nil {
			mp[root.Right] = root
			dfs(root.Right)
		}
	}
	dfs(root)
	queue := []*algorithm.TreeNode{}
	queue = append(queue, infectionNode)

	res := 0
	for len(queue) != 0 {
		size := len(queue)
		for size != 0 {
			size--
			front := queue[0]
			// fmt.Println(front.Val)
			visited[front] = struct{}{}
			queue = queue[1:]
			if front.Left != nil {
				if _, has := visited[front.Left]; !has {
					queue = append(queue, front.Left)
				}
			}
			if front.Right != nil {
				if _, has := visited[front.Right]; !has {
					queue = append(queue, front.Right)
				}
			}
			if mp[front] != nil {
				if _, has := visited[mp[front]]; !has {
					queue = append(queue, mp[front])
				}
			}
		}
		res++
	}
	return res - 1
}
