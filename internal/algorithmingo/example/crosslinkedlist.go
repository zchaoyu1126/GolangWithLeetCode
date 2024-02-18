package example

import "programs/internal/algorithmingo/algorithm"

// 求两个链表相交
// 首先求出两个链表的长度，记较长的链表长度为n，较短的链表长度为m
// 长链表先移动n-m步，此时两个链表剩余的还未遍历的长度是相同的
// 两个链表同时向后移动，直到node1 == node2
// node1 == node2有两种情况，一种是遍历完了，都为nil
// 另一种情况是遇到了两个链表相交的节点
func CrossAt(list1, list2 *algorithm.LinkedList) *algorithm.ListNode {
	n, m := list1.Size, list2.Size
	gap := n - m
	if n < m {
		list1, list2 = list2, list1
		gap = m - n
	}
	cur1 := list1.DummyHead
	cur2 := list2.DummyHead
	for gap > 0 {
		gap--
		cur1 = cur1.Next
	}
	for cur1 != cur2 {
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	return cur1
}
