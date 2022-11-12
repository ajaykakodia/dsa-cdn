package main

import "fmt"

func (ll *LinkedList) Print() {
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

func (ll *LinkedList) PrintParticularNode(index int) {
	i := 0
	cn := ll.head
	for cn != nil {
		if i == index {
			fmt.Printf("Node at Index %d is %v\n", index, cn)
			return
		}
		i++
		cn = cn.next
	}
	fmt.Printf("No node to print at index %d\n", index)
}
