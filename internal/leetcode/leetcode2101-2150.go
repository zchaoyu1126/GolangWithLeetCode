package leetcode

// leetcode2132
func PossibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	m, n := len(grid), len(grid[0])
	sum := make([][]int, m+1)
	sum[0] = make([]int, n+1)
	diff := make([][]int, m+1)
	diff[0] = make([]int, n+1)
	blockNum := 0
	for i := 1; i <= m; i++ {
		sum[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			if grid[i-1][j-1] == 1 {
				blockNum++
			}
			sum[i][j] = sum[i][j-1] + sum[i-1][j] - sum[i-1][j-1] + grid[i-1][j-1]
		}
		diff[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				// i, j 为邮票的左上角位置
				// x, y 为邮票的右下角位置
				x, y := i+stampHeight-1, j+stampWidth-1
				// 利用前缀和判断格子是否被占据
				if x < m && y < n && sum[x+1][y+1]-sum[i][y+1]-sum[x+1][j]+sum[i][j] == 0 {
					diff[i][j]++
					diff[i][y+1]--
					diff[x+1][j]--
					diff[x+1][y+1]++
				}
			}
		}
	}
	cnt := make([][]int, m+1)
	cnt[0] = make([]int, n+1)
	emptyNum := 0
	for i := 1; i <= m; i++ {
		cnt[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			cnt[i][j] = cnt[i-1][j] + cnt[i][j-1] - cnt[i-1][j-1] + diff[i-1][j-1]
			if cnt[i][j] == 0 {
				emptyNum++
			}
		}
	}
	return emptyNum == blockNum
}
