package datastruct

import (
	"programs/internal/algorithmingo/algorithm"
	"strconv"
	"strings"
)

func GenerateTree(str string) *algorithm.TreeNode {
	seq := strings.Split(str, ",")
	nodes := []*algorithm.TreeNode{}
	for i := 0; i < len(seq); i++ {
		if seq[i] == "null" {
			nodes = append(nodes, nil)
		} else {
			val, _ := strconv.Atoi(seq[i])
			node := &algorithm.TreeNode{Val: val}
			nodes = append(nodes, node)
		}
	}
	root := nodes[0]
	queue := []*algorithm.TreeNode{nodes[0]}
	nodes = nodes[1:]
	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]
		if top == nil {
			continue
		}
		if len(nodes) >= 1 {
			top.Left = nodes[0]
			queue = append(queue, nodes[0])
			nodes = nodes[1:]
		}
		if len(nodes) >= 1 {
			top.Right = nodes[0]
			queue = append(queue, nodes[0])
			nodes = nodes[1:]
		}
	}
	return root
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenerateList(nums []int) *ListNode {
	preHead := &ListNode{}
	cur := preHead
	for _, num := range nums {
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	return preHead.Next
}

type MultiListNode struct {
	Val   int
	Prev  *MultiListNode
	Next  *MultiListNode
	Child *MultiListNode
}

func GenerateMutilList(str string) *MultiListNode {
	seq := strings.Split(str, ",")
	nodes := []*MultiListNode{}
	pos := []int{}
	for i := 0; i < len(seq); i++ {
		if seq[i] == "null" {
			nodes = append(nodes, nil)
			pos = append(pos, i)
		} else {
			val, _ := strconv.Atoi(seq[i])
			node := &MultiListNode{Val: val}
			nodes = append(nodes, node)
		}
	}

	index := 0
	pos = append(pos, len(nodes))

	for len(pos) != 0 {
		levelEnd := pos[0]
		pos = pos[1:]

		//fmt.Println(levelEnd, pos)

		for i := 0; i < levelEnd-1-index; i++ {
			//fmt.Println(index+i, index+i+1)
			nodes[index+i].Next = nodes[index+i+1]
			nodes[index+i+1].Prev = nodes[index+i]
		}

		prev := levelEnd
		sum := 0

		for len(pos) != 0 {
			if pos[0] == prev+1 {
				prev = pos[0]
				pos = pos[1:]
				sum++
			} else {
				break
			}
		}

		if levelEnd+sum+1 > len(nodes) {
			break
		}
		nodes[index+sum].Child = nodes[levelEnd+sum+1]
		index = levelEnd + sum + 1
	}
	return nodes[0]
}
