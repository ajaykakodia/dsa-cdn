package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ll := NewLinkedList()

	fmt.Println("Please enter list of numbers to enter data in linked list space seperated - 1 2 3 4 5 6")
	in := bufio.NewReader(os.Stdin)
	llItems, _ := in.ReadString('\n')
	for _, val := range strings.Fields(llItems) {
		data, err := strconv.Atoi(val)
		if err == nil {
			ll.AddNode(data)
		}
	}
	ll.Print()
	ll.AddNodeAtIndex(4, 7)
	ll.Print()
	ll.AddNodeAtIndex(9, 2)
	ll.Print()
	ll.head = AddNodeAtIndexRec(ll.head, 12, 4)
	ll.Print()
	ll.PrintParticularNode(4)
	ll.head = deleteNodeRec(ll.head, 4)
	ll.Print()
	ll.head = deleteNodeRec(ll.head, 6)
	ll.Print()
	ll.head = deleteNodeRec(ll.head, 0)
	ll.Print()
	ll.DeleteNode(2)
	ll.Print()
}
