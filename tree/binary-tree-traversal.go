package main

// postOrderTraversal
func PostOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	PostOrderTraversal(root.leftChild)
	PostOrderTraversal(root.rightChild)
	root.Print()
}

// InOrderTraversal
func InOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	InOrderTraversal(root.leftChild)
	root.Print()
	InOrderTraversal(root.rightChild)
}

// PreOrderTraversal
func PreOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	root.Print()
	PreOrderTraversal(root.leftChild)
	PreOrderTraversal(root.rightChild)
}
