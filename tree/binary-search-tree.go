package main

import "math"

type BinarySearchTree struct {
	root     *Node
	numCount int
}

func (bst *BinarySearchTree) NodeCount() int {
	return bst.numCount
}

func (bst *BinarySearchTree) Print() {
	printBST(bst.root)
}

func (bst *BinarySearchTree) IsPresent(data int) bool {
	return bstSearch(bst.root, data)
}

func (bst *BinarySearchTree) Insert(data int) {
	bst.root = insertNodeBST(bst.root, data)
}

func (bst *BinarySearchTree) Delete(data int) bool {
	isDeleted, newRoot := deleteNodeBST(bst.root, data)
	if isDeleted {
		bst.root = newRoot
	}
	return isDeleted
}

func deleteNodeBST(root *Node, data int) (bool, *Node) {
	if root == nil {
		return false, nil
	}
	var isDeleted bool
	if root.data > data {
		isDeleted, root.leftChild = deleteNodeBST(root.leftChild, data)
		return isDeleted, root
	}
	if root.data < data {
		isDeleted, root.rightChild = deleteNodeBST(root.rightChild, data)
		return isDeleted, root
	}

	if root.IsLeaf() {
		return true, nil
	}
	if root.leftChild == nil {
		return true, root.rightChild
	}
	if root.rightChild == nil {
		return true, root.leftChild
	}

	minRightData := getMinData(root.rightChild)

	root.data = minRightData

	_, root.rightChild = deleteNodeBST(root.rightChild, minRightData)

	return true, root
}

func getMinBST(root *Node) int {
	if root == nil {
		return math.MinInt
	}
	if root.leftChild == nil {
		return root.data
	}
	return getMinBST(root.leftChild)
}

func insertNodeBST(root *Node, data int) *Node {
	if root == nil {
		newNode := NewNode(data)
		return newNode
	}
	if root.data > data {
		root.leftChild = insertNodeBST(root.leftChild, data)
	} else {
		root.rightChild = insertNodeBST(root.rightChild, data)
	}
	return root
}

func printBST(root *Node) {
	if root == nil {
		return
	}
	root.Print()
	printBST(root.leftChild)
	printBST(root.rightChild)
}
