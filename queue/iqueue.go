package main

import "fmt"

type iqueue struct {
	arr []int
}

func (q *iqueue) Enqueue(data int) {
	q.arr = append(q.arr, data)
}

func (q *iqueue) Dequeue() int {
	if len(q.arr) == 0 {
		fmt.Println("Nothing to dequeue!!")
		return -1
	}
	ele := q.arr[0]
	q.arr = q.arr[1:]
	return ele
}

func (q *iqueue) Front() int {
	if len(q.arr) == 0 {
		fmt.Println("Queue is empty..")
		return -1
	}
	return q.arr[0]
}

func (q *iqueue) Size() int {
	return len(q.arr)
}

func (q *iqueue) IsEmpty() bool {
	return q.Size() == 0
}
