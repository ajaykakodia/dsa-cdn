package linkedlist

func (ll *LinkedList) AddNode(num int) {
	newNode := Node{
		data: num,
	}
	if ll.head == nil {
		ll.head = &newNode
		return
	}
	currentNode := ll.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = &newNode
}

func (ll *LinkedList) AddNodeAtIndex(num, index int) {
	if index < 0 {
		fmt.Println("Please provide index greated than or equal to 0")
		return
	}
	newNode := Node{
		data: num,
	}

	var previousNode *Node
	currentNode := ll.head

	for index > 0 && currentNode != nil {
		previousNode = currentNode
		currentNode = currentNode.next
		index--
	}

	if index > 0 && previousNode != nil {
		fmt.Println("Index out of range")
		return
	}
	newNode.next = currentNode
	if previousNode == nil {
		ll.head = &newNode
		return
	}

	previousNode.next = &newNode
}

func AddNodeAtIndexRec(head *Node, num, index int) *Node {
	if index == 0 {
		newNode := Node{
			data: num,
		}
		newNode.next = head
		return &newNode
	}

	if head == nil {
		fmt.Println("Index is out of range to add Node")
		return head
	}

	head.next = AddNodeAtIndexRec(head.next, num, i-1)
	return head
}
