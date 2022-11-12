package main

import "fmt"

// addNode => Add new node at last Node of list
func (ll *LinkedList) addNode(num int) {
	newNode := Node{
		data: num,
	}
	if ll.head == nil {
		ll.head = &newNode
		return
	}
	currentNode := ll.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = &newNode
}

// addNodeAtIndex=> add node at particular index
func (ll *LinkedList) addNodeAtIndex(num, index int) {
	if index < 0 {
		fmt.Println("Please provide index greater than or equal to 0")
		return
	}
	newNode := Node{
		data: num,
	}

	var previousNode *Node
	currentNode := ll.head

	for index > 0 && currentNode != nil {
		previousNode = currentNode
		currentNode = currentNode.next
		index--
	}

	if index > 0 && previousNode != nil {
		fmt.Println("Index out of range")
		return
	}
	newNode.next = currentNode
	if previousNode == nil {
		ll.head = &newNode
		return
	}

	previousNode.next = &newNode
}

// addNodeAtIndexRec=> add node at particular index
func addNodeAtIndexRec(head *Node, num, index int) *Node {
	if index == 0 {
		newNode := Node{
			data: num,
		}
		newNode.next = head
		return &newNode
	}

	if head == nil {
		fmt.Println("Index is out of range to add Node")
		return head
	}

	head.next = addNodeAtIndexRec(head.next, num, index-1)
	return head
}
