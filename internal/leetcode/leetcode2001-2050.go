package leetcode

func Construct2DArray(original []int, m int, n int) [][]int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := 0; i < len(original); i++ {
		res[i/m][i%n] = original[i]
	}
	return res
}
