package main

import "fmt"

func main() {
	bTree := NewBinaryTree(nil)
	bTree.root = createBinaryTreeByInput()
	fmt.Println("PreOrder Traversal")
	PreOrderTraversal(bTree.root)
	fmt.Println("PostOrder Traversal")
	PostOrderTraversal(bTree.root)
	fmt.Println("InOrder Traversal")
	InOrderTraversal(bTree.root)
	fmt.Println("Total Node Count: ", countTotalNode(bTree.root))
	fmt.Println("Total Leaf Node: ", countLeafNode(bTree.root))
}
