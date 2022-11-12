package main

import "fmt"

// mergeSort => sort linked list
func mergeSort(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	mid := midPointOfLinkedlist(head)
	head1 := head
	head2 := mid.next
	mid.next = nil
	smlh1 := mergeSort(head1)
	smlh2 := mergeSort(head2)

	node := mergeTwoLinkedList(smlh1, smlh2)
	return node
}
