package main

import "fmt"

func main() {
	fmt.Println("We are in tree folder.. ")

	rootNode := NewNode(5)
	n1 := NewNode(2)
	n2 := NewNode(3)
	bTree := NewBinaryTree(rootNode)
	rootNode.leftChild = n1
	rootNode.rightChild = n2
	bTree.Print()
	n3, n4 := NewNode(10), NewNode(12)
	n1.leftChild = n3
	n1.rightChild = n4
	bTree.Print()
}
