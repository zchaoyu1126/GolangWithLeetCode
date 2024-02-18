package example

import (
	"programs/internal/algorithmingo/algorithm"
	"testing"
)

func TestCrossAt(t *testing.T) {
	commonList := algorithm.NewLinkedList([]int{1, 2, 3}...)
	list1 := algorithm.NewLinkedList([]int{9, 8, 7}...)
	list2 := algorithm.NewLinkedList([]int{10, 11, 12}...)
	head := commonList.DummyHead
	tail1 := list1.DummyHead
	tail2 := list2.DummyHead
	for tail1.Next != nil {
		tail1 = tail1.Next
	}
	for tail2.Next != nil {
		tail2 = tail2.Next
	}
	tail1.Next = head.Next
	tail2.Next = head.Next

	node := CrossAt(&list1, &list2)
	if node.Val != 1 {
		t.Errorf("expected node's val is 1, but get %d", node.Val)
	}
}
