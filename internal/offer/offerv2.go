package offer

import (
	"programs/internal/algorithmingo/algorithm"
)

// 03
func FindRepeatNumber(nums []int) int {
	n := len(nums)
	mp := make(map[int]struct{}, n)
	for i := 0; i < n; i++ {
		if _, has := mp[nums[i]]; has {
			return nums[i]
		}
		mp[nums[i]] = struct{}{}
	}
	return -1
}

// 06
func ReversePrint(head *algorithm.ListNode) []int {
	cur := head
	res := make([]int, 0)
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	for i, j := 0, len(res)-1; i <= j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

// 09 用两个栈实现队列
type CQueue struct {
	in  []int
	out []int
}

func NewCQueue() CQueue {
	return CQueue{make([]int, 0, 10), make([]int, 0, 10)}
}

func (c *CQueue) AppendTail(value int) {
	c.in = append(c.in, value)
}

func (c *CQueue) DeleteHead() int {
	if len(c.in) == 0 && len(c.out) == 0 {
		return -1
	}
	if len(c.out) == 0 {
		for len(c.in) != 0 {
			top := c.in[len(c.in)-1]
			c.in = c.in[:len(c.in)-1]
			c.out = append(c.out, top)
		}
	}
	head := c.out[len(c.out)-1]
	c.out = c.out[:len(c.out)-1]
	return head
}

// 18 删除列表中的节点
func DeleteNode(head *algorithm.ListNode, val int) *algorithm.ListNode {
	dummyHead := &algorithm.ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
			break
		}
		cur = cur.Next
	}
	return dummyHead.Next
}

// 21
func Exchange(nums []int) []int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		for l < n && nums[l]%2 == 1 {
			l++
		}
		for r >= 0 && nums[r]%2 == 0 {
			r--
		}
		if l < n && r >= 0 && l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
			//fmt.Println(nums, l, r)
		}

	}
	return nums
}

// 24 反转链表
func ReverseList(head *algorithm.ListNode) *algorithm.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	remain := head.Next
	head.Next = nil

	// 返回的是新的头节点，其末尾节点是remain:head.Next
	newHead := ReverseList(remain)
	remain.Next = head
	return newHead
}

// 114
func AlienOrder(words []string) string {
	mp := make(map[int]map[int]struct{}, 26)
	degree := make(map[int]int, 26)
	for _, word := range words {
		for _, ch := range word {
			degree[int(ch-'a')] = 0
		}
	}
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			n, m := len(words[i]), len(words[j])
			for k := 0; k < n; k++ {
				if k >= m {
					return ""
				}
				x, y := int(words[i][k]-'a'), int(words[j][k]-'a')
				if x != y {
					if _, exist := mp[x]; !exist {
						mp[x] = make(map[int]struct{}, 26)
					}
					if _, exist := mp[x][y]; !exist {
						mp[x][y] = struct{}{}
						degree[y]++
					}
					break
				}
			}
		}
	}

	queue := []int{}
	for key, val := range degree {
		if val == 0 {
			queue = append(queue, key)
		}
	}

	cnt, res := 0, []byte{}
	for len(queue) != 0 {
		front := queue[0]
		queue = queue[1:]
		cnt++
		res = append(res, byte(front+'a'))
		for key, _ := range mp[front] {
			degree[key]--
			if degree[key] == 0 {
				queue = append(queue, key)
			}
		}
	}

	if cnt != len(degree) {
		return ""
	}
	return string(res)
}
