package main

// mergeTwoLinkedList => merge two sorted linked list
func mergeTwoLinkedList(head1, head2 *Node) *Node {
	if head1 == nil && head2 == nil {
		return nil
	}

	var (
		head        *Node
		currentNode *Node
	)

	currentNode1 := head1
	currentNode2 := head2

	for currentNode1 != nil && currentNode2 != nil {
		if currentNode1.data < currentNode2.data {
			if head == nil {
				head = currentNode1
				currentNode = currentNode1
				currentNode1 = currentNode1.next
				continue
			}
			currentNode.next = currentNode1
			currentNode = currentNode.next
			currentNode1 = currentNode1.next
		} else {
			if head == nil {
				head = currentNode2
				currentNode = currentNode2
				currentNode2 = currentNode2.next
				continue
			}
			currentNode.next = currentNode2
			currentNode = currentNode.next
			currentNode2 = currentNode2.next
		}
	}

	if currentNode1 != nil {
		if head == nil {
			head = currentNode1
			return head
		}
		currentNode.next = currentNode1
	}

	if currentNode2 != nil {
		if head == nil {
			head = currentNode2
			return head
		}
		currentNode.next = currentNode2
	}

	return head
}
