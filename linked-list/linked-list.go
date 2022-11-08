package main

import "fmt"

func main() {
	ll := NewLinkedList()
	ll.Print()
	ll.AddNode(5)
	ll.AddNode(3)
	ll.AddNode(8)
	ll.Print()
	ll.AddNode(10)
	ll.AddNode(2)
	ll.AddNode(15)
	ll.Print()
	ll.AddNodeAtIndex(1, 0)
	ll.Print()
	ll.AddNodeAtIndex(4, 7)
	ll.Print()
	ll.AddNodeAtIndex(9, 9)
	ll.Print()
	ll.AddNodeAtIndex(9, 2)
	ll.Print()
}

type linkedList struct {
	head *Node
}

type Node struct {
	data int
	next *Node
}

func NewLinkedList() *linkedList {
	return &linkedList{}
}

func (ll *linkedList) AddNode(num int) {
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

func (ll *linkedList) Length() int {
	count := 0
	currentNode := ll.head
	for currentNode != nil {
		currentNode = currentNode.next
		count++
	}
	return count
}

func (ll *linkedList) AddNodeAtIndex(num, index int) {
	if index < 0 {
		fmt.Println("Please provide index greated than or equal to 0")
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

func (ll *linkedList) Print() {
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
