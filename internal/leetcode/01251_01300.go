package leetcode

// leetcode1260
func ShiftGrid(grid [][]int, k int) [][]int {
	n, m := len(grid), len(grid[0])
	arr := make([]int, n*m)
	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			arr[cnt] = grid[i][j]
			cnt++
		}
	}
	k = k % (n * m)
	for i := k; i < n*m; i++ {
		x, y := i/m, i%m
		grid[x][y] = arr[i-k]
	}
	for i := 0; i < k; i++ {
		x, y := i/m, i%m
		grid[x][y] = arr[n*m-k+i]
	}
	return grid
}

// leetcode1289
func MinFallingPathSum(grid [][]int) int {
	n := len(grid)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[i][j] = grid[i][j]
				continue
			}
			minN := 0xFFFFFFFF
			for k := 0; k < n; k++ {
				if k == j {
					continue
				}
				if minN > dp[i-1][k] {
					minN = dp[i-1][k]
				}
			}
			dp[i][j] = minN + grid[i][j]
		}
	}
	res := 0xFFFFFFFF
	for k := 0; k < n; k++ {
		if res > dp[n-1][k] {
			res = dp[n-1][k]
		}
	}
	return res
}
