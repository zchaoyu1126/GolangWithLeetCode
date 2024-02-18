package leetcode

import "programs/kit/utils"

// leetcode904
func TotalFruit1(fruits []int) int {
	mp, res := make(map[int]int), 0
	for l, r := 0, 0; r < len(fruits); r++ {
		if _, ok := mp[fruits[r]]; !ok {
			// fruits[r] 不存在有两种情况
			// 初始化窗口的时候
			if len(mp) < 2 {
				mp[fruits[r]] = r
				if res < r-l+1 {
					res = r - l + 1
				}
			} else {
				// 结算的时候
				if r-l > res {
					res = r - l
				}
				keys, vals := []int{}, []int{}
				for key, val := range mp {
					keys = append(keys, key)
					vals = append(vals, val)
				}
				if vals[0] < vals[1] {
					delete(mp, keys[0])
					l = vals[0] + 1
				} else {
					delete(mp, keys[1])
					l = vals[1] + 1
				}
				mp[fruits[r]] = r
			}
		} else {
			// 更新最后出现的下标位置
			mp[fruits[r]] = r
			if r-l+1 > res {
				res = r - l + 1
			}
		}
	}
	return res
}

func TotalFruit2(fruits []int) int {
	mp, res := make(map[int]int), 0
	for l, r := 0, 0; r < len(fruits); r++ {
		mp[fruits[r]]++
		for len(mp) > 2 {
			mp[fruits[l]]--
			if mp[fruits[l]] == 0 {
				delete(mp, fruits[l])
			}
			l++
		}

		if res < r-l+1 {
			res = r - l + 1
		}
	}
	return res
}

// leetcode908
func SmallestRangeI(nums []int, k int) int {
	var minV int = 1e5
	var maxV int = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxV {
			maxV = nums[i]
		}
		if nums[i] < minV {
			minV = nums[i]
		}
	}
	return utils.MaxNum(0, maxV-minV-2*k)
}

// leetcode913
const (
	draw     = 0
	mouseWin = 1
	catWin   = 2
)

// 自顶向下的方法
func CatMouseGame1(graph [][]int) int {
	n := len(graph)
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n*(n-1)*2)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	var getResult, getNextResult func(int, int, int) int
	getResult = func(mouse, cat, turns int) int {
		if turns == 2*n*(n-1) {
			return draw
		}
		if dp[mouse][cat][turns] != -1 {
			return dp[mouse][cat][turns]
		}
		if mouse == 0 {
			return mouseWin
		} else if mouse == cat {
			return catWin
		} else {
			res := getNextResult(mouse, cat, turns)
			dp[mouse][cat][turns] = res
			return res
		}
	}
	getNextResult = func(mouse, cat, turns int) int {
		curMove := mouse
		defaultRes := catWin
		if turns%2 == 1 {
			curMove = cat
			defaultRes = mouseWin
		}

		res := defaultRes
		for _, next := range graph[curMove] {
			if curMove == cat && next == 0 {
				continue
			}
			nextMouse, nextCat := mouse, cat
			if curMove == mouse {
				nextMouse = next
			} else {
				nextCat = next
			}
			nextRes := getResult(nextMouse, nextCat, turns+1)
			if nextRes != defaultRes {
				// 老鼠移动是，默认是猫赢，如果找到一个非必输的状态
				// 就更新res，如果res一直都是draw，那就继续
				// 如果找到了一个必胜的状态，那就直接结束
				res = nextRes
				if res != draw {
					break
				}
			}
		}
		return res

	}

	// 如果当前玩家存在一种移动方法到达非必败状态，则用该状态更新游戏结果。

	// 如果该移动方法到达必胜状态，则将当前状态（移动前的状态）设为必胜状态，结束遍历其他可能的移动。

	// 如果该移动方法到达必和状态，则将当前状态（移动前的状态）设为必和状态，继续遍历其他可能的移动，因为可能存在到达必胜状态的移动方法。

	// 如果当前玩家的任何移动方法都到达必败状态，则将当前状态（移动前的状态）设为必败状态。

	return getResult(1, 2, 0)
}

// 采用自底向上的方法，消除轮数的影响
type GraphNode struct {
	mouse int
	cat   int
	turn  int
}

func CatMouseGame2(graph [][]int) int {
	win, lose := 1, 2
	n := len(graph)

	degrees := make([][][2]int, n)
	results := make([][][2]int, n)
	for i := range degrees {
		degrees[i] = make([][2]int, n)
		results[i] = make([][2]int, n)
	}
	// 计算入度
	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			// 0:mouseTurn 1:catTurn
			// mouse在i处时，cat在j处，cat不能在0的位置
			// mouse是从graph[i]转移过来的, cat是从graph[j]转移过来的
			degrees[i][j][0] = len(graph[i])
			degrees[i][j][1] = len(graph[j])
		}
	}

	queue := make([]GraphNode, 0)
	for i := 1; i < n; i++ {
		queue = append(queue, GraphNode{0, i, 0})
		queue = append(queue, GraphNode{0, i, 1})
	}

	for i := 1; i < n; i++ {
		results[i][i][0] = catWin
		results[i][i][1] = catWin
		queue = append(queue, GraphNode{i, i, 0})
		queue = append(queue, GraphNode{i, i, 1})
	}
	mp := make(map[int]int)
	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]
		mouse, cat, curMove := top.mouse, top.cat, top.turn

		prevMouse, prevCat, prevMove, passby := mouse, cat, (curMove+1)%2, mouse
		if curMove == 0 {
			passby = cat
		}
		curRes := results[mouse][cat][curMove]
		for _, next := range graph[passby] {
			prevMouse = next
			if curMove == 0 {
				prevCat = next
			}
			if curRes == lose {
				// 当前是一个必负态
				results[prevMouse][prevCat][prevMove] = win
			} else if curRes == win {
				// 当前是一个必胜态
				// 用一个数组维护B连向的非"必胜态"的最大个数, 当推导过程中,
				// 某状态该个数为0时, 就知道这个状态是必负态了
				// prevMouse, prevCat prevMove 连了一个非必胜态
				idx := n*(n-1)*prevMove + prevMouse*(n-1) + prevCat
				mp[idx]++
				if curMove == 0 && mp[idx] == len(graph[prevCat]) || curMove == 1 && mp[idx] == len(graph[prevMouse]) {
					results[prevMouse][prevCat][prevMove] = lose
				}
				degrees[prevMouse][prevCat][prevMove]--
				if degrees[prevMouse][prevCat][prevMove] == 0 {
					queue = append(queue, GraphNode{prevMouse, prevCat, prevMove})
				}
			}
		}
	}
	return results[1][2][0]
}

// leetcode933
type RecentCounter struct {
	// 我基本的思路是类似滑动窗口，每次Ping的时候，对当前的窗口进行移动
	// 移动到max(0,t-3000)
	// 这么写会有内存性能的问题吗？
	// 预先分配好3000的容量
	cnt   int
	left  int
	times []int
}

func NewRecentCounter() RecentCounter {
	return RecentCounter{0, 0, make([]int, 0, 3000)}
}

func (r *RecentCounter) Ping(t int) int {
	r.cnt++
	r.times = append(r.times, t)
	v := utils.MaxNum(0, t-3000)
	for i := 0; i < len(r.times); i++ {
		if r.times[i] < v {
			r.cnt--
		} else {
			r.left = i
			break
		}
	}
	r.times = r.times[r.left:]
	return r.cnt
}

// leetcode942
func DiStringMatch(s string) []int {
	n := len(s)
	res := make([]int, n+1)
	l, r := 0, n

	for i := 0; i < n; i++ {
		if s[i] == 'I' {
			// res[i] < res[i+1]
			res[i] = l
			l++
		} else if s[i] == 'D' {
			// res[i] > res[i+1]
			res[i] = r
			r--
		}
	}
	res[n] = l
	return res
}

// leetcode944
func MinDeletionSize(strs []string) int {
	cnt := 0
	for j := 0; j < len(strs[0]); j++ {
		for i := 1; i < len(strs); i++ {
			if strs[i][j] < strs[i-1][j] {
				cnt++
				break
			}
		}
	}
	return cnt
}
