package leetcode

import (
	"fmt"
	"programs/internal/algorithmingo/algorithm"
	"sort"
)

// leetcode851
func LoudAndRich(richer [][]int, quiet []int) []int {
	res := make([]int, len(quiet))
	for i := range res {
		res[i] = i
	}
	mp := make(map[int][]int)
	cnt := make([]int, len(quiet))
	for i := 0; i < len(richer); i++ {
		x, y := richer[i][0], richer[i][1]
		if _, ok := mp[x]; !ok {
			mp[x] = make([]int, 0)
		}
		mp[x] = append(mp[x], y)
		cnt[y]++
	}
	queue := []int{}
	for i := 0; i < len(cnt); i++ {
		if cnt[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) != 0 {
		top := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		for i := 0; i < len(mp[top]); i++ {
			node := mp[top][i]
			if quiet[res[top]] < quiet[res[node]] {
				res[node] = res[top]
			}
			cnt[node]--
			if cnt[node] == 0 {
				queue = append(queue, node)
			}
		}
	}
	return res
}

// leetcode859
func BuddyStrings(s string, goal string) bool {
	len1, len2 := len(s), len(goal)
	if len1 != len2 {
		return false
	}
	first, second := -1, -1
	mp := make(map[byte]int)
	for i := 0; i < len1; i++ {
		mp[s[i]-'a']++
		if s[i] == goal[i] {
			continue
		} else if first == -1 {
			first = i
		} else if second == -1 {
			second = i
		} else {
			return false
		}
	}
	// 完全相同的两个字符串
	if first == -1 && second == -1 {
		for _, v := range mp {
			if v >= 2 {
				return true
			}
		}
		return false
	}
	// 只有一位不同的两个字符串
	if first != -1 && second == -1 {
		return false
	}
	// 恰好有两位相同的两个字符串
	if s[first] == goal[second] && s[second] == goal[first] {
		return true
	}
	return false
}

// leetcode863
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var res []int

func DistanceK(root *algorithm.TreeNode, target *algorithm.TreeNode, k int) []int {
	res = []int{}
	if root.Val == target.Val {
		res = append(res, root.Val)
	}
	parent := findTargetParent(root, target.Val)
	preOrder(root, target.Val, -1, k, false)
	preOrder(root, parent.Val, target.Val, k-1, false)
	fmt.Println(res)
	return res
}

func findTargetParent(root *algorithm.TreeNode, val int) *algorithm.TreeNode {
	if root.Val == val {
		return root
	}

	queue := []*algorithm.TreeNode{}
	queue = append(queue, root)

	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]
		if top.Left != nil {
			if top.Left.Val == val {
				return top
			} else {
				queue = append(queue, top.Left)
			}
		}

		if top.Right != nil {
			if top.Right.Val == val {
				return top
			} else {
				queue = append(queue, top.Right)
			}
		}
	}
	return nil
}

func preOrder(root *algorithm.TreeNode, val, pre, k int, isTargetSon bool) {
	if root == nil {
		return
	}

	if isTargetSon {
		k--
		if k == 0 && root.Val != pre {
			res = append(res, root.Val)
		}
	}

	if root.Val == val {
		isTargetSon = true
	}

	preOrder(root.Left, val, pre, k, isTargetSon)
	preOrder(root.Right, val, pre, k, isTargetSon)
}

// func preOrder2(root *algorithm.TreeNode) {
// 	stack := []*algorithm.TreeNode{}
// 	cur := root
// 	for cur != nil || len(stack) != 0 {
// 		for cur != nil {
// 			fmt.Println(cur.Val)
// 			stack = append(stack, cur)
// 			cur = cur.Left
// 		}

// 		if len(stack) != 0 {
// 			cur = stack[len(stack)-1].Right
// 			stack = stack[:len(stack)-1]
// 		}
// 	}
// }

// leetcode881
func NumRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	for i, j := 0, len(people)-1; i < j; i, j = i+1, j-1 {
		people[i], people[j] = people[j], people[i]
	}
	res := 0
	end := len(people)
	for i := 0; i < end; i++ {
		fmt.Println(i, end-1)
		if people[i]+people[end-1] <= limit {
			end--
		}
		res++
	}
	fmt.Println(people)
	fmt.Println(res)
	return res
}

// leetcode894
func AllPossibleFBT(n int) []*algorithm.TreeNode {
	if n%2 == 0 {
		return []*algorithm.TreeNode{}
	}

	mp := make(map[int][]*algorithm.TreeNode)
	var generate func(n int) []*algorithm.TreeNode
	generate = func(n int) []*algorithm.TreeNode {
		if n == 1 {
			node := &algorithm.TreeNode{}
			return []*algorithm.TreeNode{node}
		}
		res := []*algorithm.TreeNode{}
		for i := 1; i <= n-2; i = i + 2 {
			if _, ok := mp[i]; !ok {
				mp[i] = generate(i)
			}
			if _, ok := mp[n-i-1]; !ok {
				mp[n-i-1] = generate(n - i - 1)
			}
			for _, leftTree := range mp[i] {
				for _, rightTree := range mp[n-i-1] {
					root := &algorithm.TreeNode{}
					root.Left = leftTree
					root.Right = rightTree
					res = append(res, root)
				}
			}
		}
		return res
	}
	return generate(n)
}
