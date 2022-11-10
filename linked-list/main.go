package main

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
	ll.head = AddNodeAtIndexRec(ll.head, 12, 4)
	ll.Print()
}
