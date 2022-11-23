package main

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

func printBST(root *Node) {
	if root == nil {
		return
	}
	root.Print()
	printBST(root.leftChild)
	printBST(root.rightChild)
}
