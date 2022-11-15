package main

import "fmt"

func main() {
	iq := iqueue{}
	iq.Enqueue(2)
	iq.Enqueue(5)
	iq.Enqueue(10)

	for !iq.IsEmpty() {
		fmt.Println(iq.Dequeue())
	}
	iq.Dequeue()

	sq := squeue{}
	sq.Enqueue("Sona")
	sq.Enqueue("Babu")
	sq.Enqueue("Mer Kahi")

	for !sq.IsEmpty() {
		fmt.Println(sq.Dequeue())
	}
	sq.Dequeue()
}
