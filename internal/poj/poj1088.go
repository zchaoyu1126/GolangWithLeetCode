package poj

import "programs/kit/utils"

func MaxLength(hight [][]int) int {
	row := len(hight)
	col := len(hight[0])
	length := make([][]int, row)
	maxLength := 0
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	for i := 0; i < row; i++ {
		length[i] = make([]int, col)
	}
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if length[x][y] != 0 {
			return length[x][y]
		}
		for i := 0; i < 4; i++ {
			nx := x + dx[i]
			ny := y + dy[i]
			if nx >= 0 && nx < row && ny >= 0 && ny < col {
				if hight[x][y] > hight[nx][ny] {
					length[x][y] = utils.MaxNum(length[x][y], dfs(nx, ny)+1)
				}
			}
		}
		if maxLength < length[x][y] {
			maxLength = length[x][y]
		}
		return length[x][y]
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			dfs(i, j)
		}
	}
	return maxLength
}
