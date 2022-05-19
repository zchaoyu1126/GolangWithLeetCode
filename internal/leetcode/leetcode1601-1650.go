package leetcode

import (
	"fmt"
	"programs/kit/utils"
)

// leetcode1614
func MaxDepth(s string) int {
	// "8*((1*(5+6))*(8/6))"
	// "(1+(2*3)+((8)/4))+1"
	stack := []byte{}
	deep := []int{}
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] != '(' && s[i] != ')' {
			continue
		}
		if s[i] == '(' {
			stack = append(stack, '(')
			deep = append(deep, 1)
			//cur = 0
		} else if s[i] == ')' {
			// 肯定是满足括号匹配的
			num := len(stack)
			prev := num - 2

			stack = stack[:num-1]
			d, deep := deep[num-1], deep[:num-1]
			fmt.Println(stack, deep)
			if prev >= 0 {
				deep[prev] = utils.MaxNum(deep[prev], d+1)
				res = utils.MaxNum(res, deep[prev])
			} else {
				res = utils.MaxNum(res, d)
			}
		}
	}
	fmt.Println(res)
	return res
}

// leetcode1631
func abs_int(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func MinimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	mp := make(map[int]map[int]int)
	for i := 0; i < m*n; i++ {
		mp[i] = make(map[int]int)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			right, down := j+1, i+1
			if right < n {
				v := abs_int(heights[i][right] - heights[i][j])
				mp[i*n+j][i*n+right] = v
				mp[i*n+right][i*n+j] = v
			}
			if down < m {
				v := abs_int(heights[down][j] - heights[i][j])
				mp[i*n+j][down*n+j] = v
				mp[down*n+j][i*n+j] = v
			}
		}
	}

	dis := make([]int, m*n)
	vis := make([]int, m*n)
	for i := range dis {
		dis[i] = 0xFFFFFFFF
	}
	for k, v := range mp[0] {
		dis[k] = v
	}
	dis[0] = 0
	vis[0] = 1

	for t := 1; t < m*n; t++ {
		minnum := 0xFFFFFFFF
		nextNode := -1
		for i := 0; i < m*n; i++ {
			if vis[i] != 1 && dis[i] < minnum {
				minnum = dis[i]
				nextNode = i
			}
		}
		fmt.Printf("find %d \n", nextNode)
		fmt.Println(vis)
		fmt.Println(dis)
		fmt.Println()

		vis[nextNode] = 1
		for k, v := range mp[nextNode] {
			if vis[k] == 1 {
				continue
			}
			fmt.Println(k, v, nextNode)
			if dis[k] == 0xFFFFFFFF {
				dis[k] = utils.MaxNum(dis[nextNode], v)
			} else {
				dis[k] = utils.MinNum(dis[k], utils.MaxNum(dis[nextNode], v))
			}
		}

		fmt.Printf("after %d \n", nextNode)
		fmt.Println(vis)
		fmt.Println(dis)
		fmt.Println()
	}
	fmt.Println(dis[m*n-1])
	return dis[m*n-1]
}

// leetcode1646
func GetMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	res := make([]int, n+1)
	res[0] = 0
	res[1] = 1
	ans := -1
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			res[i] = res[i/2]
		} else {
			res[i] = res[i/2] + res[i/2+1]
		}
		if ans < res[i] {
			res[i] = ans
		}
	}
	return ans
}
