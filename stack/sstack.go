package main

import "fmt"

type Node struct {
	data string
	next *Node
}

type SStack struct {
	head  *Node
	count uint
}

func (s *SStack) Push(data string) {
	newNode := Node{
		data: data,
	}
	newNode.next = s.head
	s.head = &newNode
}

func (s *SStack) Pop() string {
	if s.IsEmpty() {
		fmt.Println("Stack is Empty! nothing to pop..")
		return ""
	}
	data := s.head.data
	s.head = s.head.next
	return data
}

func (s *SStack) Top() string {
	if s.IsEmpty() {
		fmt.Println("Stack is Empty! nothing to pop..")
		return ""
	}
	return s.head.data
}

func (s *SStack) IsEmpty() bool {
	return s.head == nil
}

func (s *SStack) Size() uint {
	return s.count
}
