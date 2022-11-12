package main

// reverseLinkedList => Reverse a linked list by Iterative way
func (ll *LinkedList) reverseLinkedList() {
	if ll.head == nil || ll.head.next == nil {
		return
	}
	var (
		previousNode *Node
		tempNode     *Node
	)
	currentNode := ll.head
	for currentNode != nil {
		tempNode = currentNode.next
		currentNode.next = previousNode
		previousNode = currentNode
		currentNode = tempNode
	}
	ll.head = previousNode
}

// reverseLinkedListRec => reverse list by looping smaller linked list
func reverseLinkedListRec(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	smlh := reverseLinkedListRec(head.next)
	currentNode := smlh
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = head
	head.next = nil
	return smlh
}

// reverseLinkedListRec2 => reverse linked list by maintaining tail
func reverseLinkedListRec2(head *Node) (*Node, *Node) {
	if head == nil || head.next == nil {
		return head, head
	}
	smlh, tail := reverseLinkedListRec2(head.next)
	tail.next = head
	head.next = nil
	return smlh, head
}

// reverseLinkedListRec3 => reverse linked list by using head.next as tail
func reverseLinkedListRec3(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	smlh := reverseLinkedListRec3(head.next)
	tail := head.next
	tail.next = head
	head.next = nil
	return smlh
}
