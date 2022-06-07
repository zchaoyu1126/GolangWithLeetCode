package leetcode

import (
	"fmt"
	"math"
	"programs/internal/algorithmingo/algorithm"
	"programs/kit/utils"
	"sort"
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

// leetcode668
func FindKthNumber(m int, n int, k int) int {
	l, r := 0, m*n
	lessThanCnt := func(x int) int {
		ans := 0
		for i := 0; i < m; i++ {
			ans += utils.MinNum(x/i, n)
		}
		return ans
	}
	for l <= r {
		m := (l + r) / 2
		cnt := lessThanCnt(m)
		if cnt >= k {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
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

// leetcode675
type treeInfo struct {
	Height int
	x, y   int
}

type treeInfoSlice []treeInfo

func (t treeInfoSlice) Len() int           { return len(t) }
func (t treeInfoSlice) Less(i, j int) bool { return t[i].Height < t[j].Height }
func (t treeInfoSlice) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

func CutOffTree(forest [][]int) int {
	n, m := len(forest), len(forest[0])
	tree := make([]treeInfo, 0, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if forest[i][j] >= 2 {
				tree = append(tree, treeInfo{forest[i][j], i, j})
			}
		}
	}
	sort.Sort(treeInfoSlice(tree))

	prevX, prevY, cnt := 0, 0, 0
	for _, t := range tree {
		curX, curY := t.x, t.y
		steps := bfs(forest, prevX, prevY, curX, curY)
		if steps == -1 {
			return -1
		}
		cnt += steps
		prevX, prevY = curX, curY
	}
	return cnt
}

func bfs(arr [][]int, startX, startY, endX, endY int) int {
	n, m := len(arr), len(arr[0])
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	steps := 0
	visited := make([][]int, n)
	for i := range visited {
		visited[i] = make([]int, m)
	}
	queue := [][2]int{{startX, startY}}
	visited[startX][startY] = 1
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			top := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			x, y := top[0], top[1]
			if x == endX && y == endY {
				return steps
			}
			for i := 0; i < 4; i++ {
				nx, ny := x+dx[i], y+dy[i]
				// 检查是否合法
				if nx < n && nx >= 0 && ny < m && ny >= 0 && visited[nx][ny] != 1 && arr[nx][ny] != 0 {
					queue = append(queue, [2]int{nx, ny})
					visited[nx][ny] = 1
				}
			}
		}
		steps++
	}
	return -1
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

// leetcode691
// 错误的写法
// func MinStickers(stickers []string, target string) int {
// 	// 一眼贪心，优先选匹配度高的
// 	// 求target需要各个字符各多少个
// 	need := make(map[byte]int)
// 	for i := 0; i < len(target); i++ {
// 		need[target[i]]++
// 	}
// 	res := 0
// 	// 求stickers[i]与target有多少个匹配的字符， 有的话是什么？
// 	has := make([]map[byte]int, len(stickers)) //有哪些字符，有多少个
// 	for i := 0; i < len(stickers); i++ {
// 		has[i] = make(map[byte]int)
// 	}
// 	total := make([]int, len(stickers))
// 	for k, str := range stickers {
// 		for i := 0; i < len(str); i++ {
// 			for j := 0; j < len(target); j++ {
// 				if str[i] == target[j] {
// 					has[k][str[i]]++
// 					total[k]++
// 				}
// 			}
// 		}
// 	}
// 	// 已经求得了每个字符串有多少个匹配的字符，
// 	// 优先选择最多的
// 	max := func(a, b int) int {
// 		if a < b {
// 			return b
// 		}
// 		return a
// 	}
// 	for len(need) != 0 {
// 		maxIdx := -1
// 		maxNum := 0
// 		for i := 0; i < len(stickers); i++ {
// 			if total[i] > maxNum {
// 				maxNum = total[i]
// 				maxIdx = i
// 			}
// 		}
// 		// 最多的maxIdx
// 		cnt := 0
// 		for k, v := range has[maxIdx] {
// 			if need[k] != 0 {
// 				need[k] = max(need[k]-v, 0)
// 				if need[k] == 0 {
// 					delete(need, k)
// 				}
// 			} else {
// 				cnt++
// 			}
// 		}
// 		if cnt == len(has[maxIdx]) {
// 			total[maxIdx] = 0
// 		}
// 		res++
// 	}
// 	return res
// }

func MinStickers(stickers []string, target string) int {
	// bfs写法，将target视为起始态，每次用stickers去删除
	// 判断最后能否得到空
	targetSet := map[rune]bool{} // target字符串中的所有字符
	for _, r := range target {
		targetSet[r] = true
	}
	availables := []map[rune]int{} // stickers中存在的字符的数量
	getCounter := func(str string, set map[rune]bool) map[rune]int {
		res := map[rune]int{}
		for _, ch := range str {
			if set[ch] {
				res[ch]++
			}
		}
		if len(res) == 0 {
			return nil
		}
		return res
	}
	transfer := func(s string, mp map[rune]int) string {
		for k, v := range mp {
			s = strings.Replace(s, string(k), "", v)
		}
		//fmt.Println("hi", s)
		return s
	}

	for _, s := range stickers {
		if c := getCounter(s, targetSet); c != nil {
			availables = append(availables, c)
		}
	}
	//fmt.Println(availables)
	queue := []string{target}
	explored := map[string]int{target: 0}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		//fmt.Println(cur)
		for _, avl := range availables {
			if avl[rune(cur[0])] > 0 {
				nxt := transfer(cur, avl)
				if len(nxt) == 0 {
					return explored[cur] + 1
				}
				if _, ok := explored[nxt]; !ok {
					queue = append(queue, nxt)
					explored[nxt] = explored[cur] + 1
				}
			}
		}
	}
	return -1
}

// leetcode699

func FallingSquares(positions [][]int) []int {
	// 1 2 2
	// 2 4 3
	// 6 6 1
	return []int{}
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
