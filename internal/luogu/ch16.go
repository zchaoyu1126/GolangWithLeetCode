package luogu

import "fmt"

type Edge struct {
	next int
	to   int
}

var head []int
var edge []Edge
var tot int
var n, m int

func P3916() {
	fmt.Scan(&n, &m)
	head = make([]int, n)
	for i := 0; i < n; i++ {
		head[i] = -1
	}
	edge = make([]Edge, m)
	tot = 0
	for i := 0; i < m; i++ {
		var from, to int
		fmt.Scan(&from, &to)
		build(from, to)
	}
	for i := 0; i < n; i++ {
		fmt.Println(dfs(i))
	}
}

func dfs(x int) int {
	maxV := x
	for i := head[x]; i != -1; i = edge[i].next {
		to := edge[i].to
		maxV = max(maxV, dfs(to))
	}
	return maxV
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func build(from, to int) {
	edge[tot].to = to
	edge[tot].next = head[from]
	head[from] = tot
	tot++
}
