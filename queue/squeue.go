package main

import "fmt"

type Node struct {
	data string
	next *Node
}

type squeue struct {
	head  *Node
	tail  *Node
	count int
}

func (q *squeue) Enqueue(data string) {
	newNode := &Node{data: data}
	q.count++
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
		return
	}
	q.tail.next = newNode
	q.tail = newNode
}

func (q *squeue) Dequeue() string {
	if q.IsEmpty() {
		fmt.Print("Nothing to Deqeue!!")
		return ""
	}
	ele := q.head
	q.head = q.head.next
	return ele.data
}

func (q *squeue) Front() string {
	if q.IsEmpty() {
		fmt.Println("Queue is empty...")
		return ""
	}
	return q.head.data
}

func (q *squeue) Size() int {
	return q.count
}

func (q *squeue) IsEmpty() bool {
	return q.head == nil
}
