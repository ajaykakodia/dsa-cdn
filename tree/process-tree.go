package main

func processInOrderPreOrderForBinaryTree(inOrder, preOrder []int) *Node {
	if len(preOrder) == 0 {
		return nil
	}
	rootData := preOrder[0]
	rootNode := NewNode(rootData)
	rootIndexInOrder := -1

	for i, v := range inOrder {
		if v == rootData {
			rootIndexInOrder = i
		}
	}
	if rootIndexInOrder == -1 {
		return nil
	}

	leftChildTreeInOrder := inOrder[:rootIndexInOrder]
	rightChildTreeInOrder := inOrder[rootIndexInOrder+1:]

	leftChildTreePreOrder := preOrder[1 : len(leftChildTreeInOrder)+1]
	rightChildTreePreOrder := preOrder[len(leftChildTreeInOrder)+1:]

	rootNode.leftChild = processInOrderPreOrderForBinaryTree(leftChildTreeInOrder, leftChildTreePreOrder)
	rootNode.rightChild = processInOrderPreOrderForBinaryTree(rightChildTreeInOrder, rightChildTreePreOrder)

	return rootNode
}

func processInOrderPostOrderForBinaryTree(inOrder, postOrder []int) *Node {
	if len(postOrder) == 0 {
		return nil
	}
	rootData := postOrder[len(postOrder)-1]
	rootNode := NewNode(rootData)
	rootIndexInOrder := -1

	for i, v := range inOrder {
		if v == rootData {
			rootIndexInOrder = i
		}
	}
	if rootIndexInOrder == -1 {
		return nil
	}

	leftChildTreeInOrder := inOrder[:rootIndexInOrder]
	rightChildTreeInOrder := inOrder[rootIndexInOrder+1:]

	leftChildTreePostOrder := postOrder[:len(leftChildTreeInOrder)]
	rightChildTreePostOrder := postOrder[len(leftChildTreeInOrder) : len(postOrder)-1]

	rootNode.leftChild = processInOrderPostOrderForBinaryTree(leftChildTreeInOrder, leftChildTreePostOrder)
	rootNode.rightChild = processInOrderPostOrderForBinaryTree(rightChildTreeInOrder, rightChildTreePostOrder)

	return rootNode
}
