package main

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
