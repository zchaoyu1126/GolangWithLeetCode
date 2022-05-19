package main

import (
	"fmt"
	"programs/internal/algorithmingo/algorithm"
)

func main() {
	nums := []int{2, 1, 3, 5, 4, 8}
	algorithm.MergeSort(nums)
	fmt.Println(nums)
}

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// var m, n, x int
// var skill1 [][]int //面试官会的语言
// var skill2 []int   //面试者会的语言
// var conn [][]int   //面试者i是否可以被面试官j面试

// var match [][]int //记录面试官j面了那几个人
// var visited []int //记录面试官j是否被询问过能否安排面试

// func read() {
// 	// 注：面试官，面试者的编号均从1开始
// 	fmt.Scanf("%d %d %d\n", &m, &n, &x)
// 	skill1 = make([][]int, m+5)
// 	skill2 = make([]int, n+5)
// 	reader := bufio.NewReader(os.Stdin)

// 	for i := 1; i <= m; i++ {
// 		skill1[i] = make([]int, 0)
// 		bytes, _, _ := reader.ReadLine()
// 		str := string(bytes)
// 		skills := strings.Split(str, " ")
// 		for _, skill := range skills {
// 			if skill == "C++" {
// 				skill1[i] = append(skill1[i], 0)
// 			} else if skill == "Python" {
// 				skill1[i] = append(skill1[i], 1)
// 			} else if skill == "Java" {
// 				skill1[i] = append(skill1[i], 2)
// 			}
// 		}
// 	}

// 	for i := 1; i <= n; i++ {
// 		bytes, _, _ := reader.ReadLine()
// 		skill := string(bytes)
// 		if skill == "C++" {
// 			skill2[i] = 0
// 		} else if skill == "Python" {
// 			skill2[i] = 1
// 		} else if skill == "Java" {
// 			skill2[i] = 2
// 		}
// 	}

// 	// 建图 conn = connection
// 	// 使用邻接矩阵存储
// 	// 由于面试者要被面试两次，所以直接将面试者加倍，等价于有2n个面试者要进行面试
// 	conn = make([][]int, 2*n+5)
// 	for i := range conn {
// 		conn[i] = make([]int, m+5)
// 	}
// 	for i := 1; i <= 2*n; i++ {
// 		for j := 1; j <= m; j++ {
// 			for k := 0; k < len(skill1[j]); k++ {
// 				if skill1[j][k] == skill2[i%n] {
// 					conn[i][j] = 1
// 				}
// 			}
// 		}
// 	}

// }

// func search(i int) bool {
// 	for j := 1; j <= m; j++ {
// 		// 面试者i可以被面试官j面试，且j没有被询问过是否可以安排面试
// 		if conn[i][j] == 1 && visited[j] == 0 {
// 			visited[j] = 1
// 			// 判断match[j]是否含有i-n，即一面是否面过面试者i
// 			if i-n > 0 {
// 				for p := 0; p < len(match[j]); p++ {
// 					// 若含有，看i-n号面试者是否能被其他面试官面
// 					if match[j][p] == i-n {
// 						if search(i - n) {
// 							match[j][p] = i
// 							return true
// 						} else {
// 							return false
// 						}
// 					}
// 				}
// 			}

// 			if len(match[j]) < x {
// 				match[j] = append(match[j], i)
// 				return true
// 			} else {
// 				for k := 0; k < len(match[j]); k++ {
// 					if search(match[j][k]) {
// 						match[j][k] = i
// 						return true
// 					}
// 				}
// 				return false
// 			}
// 		}
// 	}
// 	return false
// }

// func main() {
// 	read()
// 	match = make([][]int, m+5)
// 	visited = make([]int, m+5)
// 	cnt := 0
// 	for i := 1; i <= 2*n; i++ {
// 		for i := 1; i <= m; i++ {
// 			visited[i] = 0
// 		}
// 		if search(i) {
// 			cnt++
// 		}
// 	}

// 	if cnt == 2*n {
// 		fmt.Println(true)
// 		for i := 1; i <= m; i++ {
// 			for j := 0; j < len(match[i]); j++ {
// 				if match[i][j] > n {
// 					fmt.Printf("%d ", match[i][j]-n)
// 				} else {
// 					fmt.Printf("%d ", match[i][j])
// 				}

// 			}
// 			fmt.Println()
// 		}
// 		return
// 	}

// 	fmt.Println(false)
// }
