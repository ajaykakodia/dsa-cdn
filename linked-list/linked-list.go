package linkedlist

import "fmt"

type LinkedList struct {
	head *Node
}

type Node struct {
	data int
	next *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Length() int {
	count := 0
	currentNode := ll.head
	for currentNode != nil {
		currentNode = currentNode.next
		count++
	}
	return count
}

func (ll *LinkedList) Print() {
	if ll.head == nil {
		fmt.Println("No node for print")
		return
	}
	currentNode := ll.head
	for currentNode != nil {
		fmt.Printf("%d -> ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println("none")
}
