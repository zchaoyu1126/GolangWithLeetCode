package example

import (
	"fmt"
	"programs/internal/algorithmingo/algorithm"
)

func EntryPoint(list *algorithm.LinkedList) *algorithm.ListNode {
	// 考虑特殊情况
	if list.Size == 1 {
		return nil
	}

	slow, fast := list.DummyHead.Next, list.DummyHead.Next.Next
	for fast != nil && fast.Next != nil && slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if slow != fast {
		fmt.Println("no circle")
		return nil
	}
	cur := list.DummyHead
	for cur != slow {
		cur = cur.Next
		slow = slow.Next
	}
	return slow
}
