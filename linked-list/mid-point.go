package main

func (ll *LinkedList) midPoint() *Node {
	if ll.head == nil {
		return nil
	}
	currentNode := ll.head
	fastNode := ll.head

	for fastNode.next != nil && fastNode.next.next != nil {
		currentNode = currentNode.next
		fastNode = fastNode.next.next
	}
	return currentNode
}

func midPointOfLinkedlist(head *Node) *Node {
	if head == nil {
		return nil
	}
	currentNode := head
	fastNode := head

	for fastNode.next != nil && fastNode.next.next != nil {
		currentNode = currentNode.next
		fastNode = fastNode.next.next
	}
	return currentNode
}
