package main

import "fmt"

type queue struct {
	s1 IStack
	s2 IStack
}

// Enqueue => Choose Order(n) Enqueue operation to save Front order(n) operation
func (q *queue) Enqueue(data int) {
	for !q.s1.IsEmpty() {
		q.s2.Push(q.s1.Pop())
	}
	q.s1.Push(data)
	for !q.s2.IsEmpty() {
		q.s1.Push(q.s2.Pop())
	}
}

// Dequeue => Choose Order(1) Enqueue operation
func (q *queue) Dequeue() int {
	if q.IsEmpty() {
		fmt.Println("Nothing to dequeue!!")
		return -1
	}
	return q.s1.Pop()
}

func (q *queue) Front() int {
	if q.IsEmpty() {
		fmt.Println("Queue is empty..")
		return -1
	}
	return q.s1.Top()
}

func (q *queue) Size() int {
	return q.s1.Size()
}

func (q *queue) IsEmpty() bool {
	return q.s1.IsEmpty()
}
