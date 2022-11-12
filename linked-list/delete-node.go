package main

import "fmt"

// deleteNode => Delete ode Iterative
func (ll *LinkedList) deleteNode(index int) {
	if index == 0 {
		ll.head = ll.head.next
		return
	}
	i := 0
	cn := ll.head
	for i < index-1 && cn != nil {
		i++
		cn = cn.next
	}
	if cn == nil {
		fmt.Println("No node to delete on index ", index)
		return
	}
	cn.next = cn.next.next
}

// deleteNodeRec => Delete Node by Recursion wa
func deleteNodeRec(head *Node, index int) *Node {
	if head == nil {
		return nil
	}
	if index == 0 {
		return head.next
	}
	head.next = deleteNodeRec(head.next, index-1)
	return head
}
