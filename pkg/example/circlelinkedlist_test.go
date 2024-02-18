package example

import (
	"programs/internal/algorithmingo/algorithm"
	"testing"
)

func TestEntryPointWithCircle(t *testing.T) {
	// 构造有环链表
	list := algorithm.NewLinkedList([]int{1, 2, 3}...)
	circle := algorithm.NewLinkedList([]int{4, 5, 6, 7, 8, 9}...)
	tail1 := list.DummyHead
	for tail1.Next != nil {
		tail1 = tail1.Next
	}

	head, tail2 := circle.DummyHead, circle.DummyHead
	for tail2.Next != nil {
		tail2 = tail2.Next
	}
	tail2.Next = head.Next
	tail1.Next = head.Next

	output := EntryPoint(&list)
	if output.Val != 4 {
		t.Errorf("expected node's val:%d, but get %d", 4, output.Val)
	}
}

func TestEntryPointWithOutCircle(t *testing.T) {
	list := algorithm.NewLinkedList([]int{1, 2, 3, 4}...)
	output := EntryPoint(&list)
	if output != nil {
		t.Errorf("expected nil, but get %v", output)
	}
}
