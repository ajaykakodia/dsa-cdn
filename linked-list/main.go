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

	fmt.Println("Please enter list of numbers to enter data in linked list space separated - 1 2 3 4 5 6")
	in := bufio.NewReader(os.Stdin)
	llItems, _ := in.ReadString('\n')
	for _, val := range strings.Fields(llItems) {
		data, err := strconv.Atoi(val)
		if err == nil {
			ll.addNode(data)
		}
	}
	ll.print()
	ll.addNodeAtIndex(4, 7)
	ll.print()
	ll.addNodeAtIndex(9, 2)
	ll.print()
	ll.head = addNodeAtIndexRec(ll.head, 12, 4)
	ll.print()
	ll.printParticularNode(4)
	ll.head = deleteNodeRec(ll.head, 4)
	ll.print()
	ll.head = deleteNodeRec(ll.head, 6)
	ll.print()
	ll.deleteNode(2)
	ll.print()
	fmt.Println("Reverse 1")
	ll.reverseLinkedList()
	ll.print()
	fmt.Println("Reverse 2")
	ll.head = reverseLinkedListRec(ll.head)
	ll.print()
	fmt.Println("Reverse 3")
	ll.head, _ = reverseLinkedListRec2(ll.head)
	ll.print()
	fmt.Println("Reverse 4")
	ll.head = reverseLinkedListRec3(ll.head)
	ll.print()
	midNode := ll.midPoint()
	fmt.Printf("Node at mid %v\n", midNode)
	ll.deleteNode(2)
	ll.print()
	midNode = ll.midPoint()
	fmt.Printf("Node at mid %v\n", midNode)

	fmt.Println("Merge Sort")

	ll.head = mergeSort(ll.head)
	ll.print()

	fmt.Println("========================================================")

	// Uncomment below line to take sorted inputs and merge linked list

	// fmt.Println("Entered first sorted list elements")
	// ll1 := NewLinkedList()
	// llItems, _ = in.ReadString('\n')
	// for _, val := range strings.Fields(llItems) {
	// 	data, err := strconv.Atoi(val)
	// 	if err == nil {
	// 		ll1.addNode(data)
	// 	}
	// }
	// ll1.print()

	// fmt.Println("Entered first sorted list elements")
	// ll2 := NewLinkedList()
	// llItems, _ = in.ReadString('\n')
	// for _, val := range strings.Fields(llItems) {
	// 	data, err := strconv.Atoi(val)
	// 	if err == nil {
	// 		ll2.addNode(data)
	// 	}
	// }
	// ll2.print()

	// ll.head = mergeTwoLinkedList(ll1.head, ll2.head)
	// ll.print()
}
