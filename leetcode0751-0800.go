package leetcode

import (
	"container/heap"
)

// leetcode786
type Fraction struct {
	up, down   int
	idxI, idxJ int
}
type FractionHeap []Fraction

func (h FractionHeap) Len() int           { return len(h) }
func (h FractionHeap) Less(i, j int) bool { return h[i].up*h[j].down < h[i].down*h[j].up }
func (h FractionHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *FractionHeap) Push(x interface{}) {
	*h = append(*h, x.(Fraction))
}
func (h *FractionHeap) Pop() interface{} {
	length := len(*h)
	x := (*h)[length-1]
	*h = (*h)[:length-1]
	return x
}

func KthSmallestPrimeFraction(arr []int, k int) []int {
	length := len(arr)
	h := make(FractionHeap, length-1)
	for j := 1; j < length; j++ {
		h[j-1] = Fraction{arr[0], arr[j], 0, j}
	}
	heap.Init(&h)

	for loop := k - 1; loop > 0; loop-- {
		f := heap.Pop(&h).(Fraction)
		if f.idxI+1 < f.idxJ {
			heap.Push(&h, Fraction{arr[f.idxI+1], arr[f.idxJ], f.idxI + 1, f.idxJ})
		}
	}
	return []int{h[0].up, h[0].down}
}

// leetcode794
func ValidTicTacToe(board []string) bool {
	xnum, onum := 0, 0
	res := make([][]byte, 8)
	for i := 0; i < len(res); i++ {
		res[i] = make([]byte, 0)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' {
				onum++
			} else if board[i][j] == 'X' {
				xnum++
			}
			res[j] = append(res[j], board[i][j])
		}
	}
	if !(xnum == onum || xnum == onum+1) {
		return false
	}
	res[3] = []byte(board[0])
	res[4] = []byte(board[1])
	res[5] = []byte(board[2])
	res[6] = append(res[6], board[0][0], board[1][1], board[2][2])
	res[7] = append(res[7], board[0][2], board[1][1], board[2][0])
	for i := 0; i < len(res); i++ {
		if string(res[i]) == "OOO" && onum != xnum {
			return false
		}
		if string(res[i]) == "XXX" && xnum != onum+1 {
			return false
		}
	}
	return true
}
