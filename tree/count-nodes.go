package main

func countTotalNode(root *Node) int {
	if root == nil {
		return 0
	}
	lcCount := countTotalNode(root.leftChild)
	rcCount := countTotalNode(root.rightChild)
	return lcCount + rcCount + 1 // Add 1 for current node count
}

func countLeafNode(root *Node) int {
	if root == nil {
		return 0
	}
	if root.rightChild == nil && root.leftChild == nil {
		return 1
	}

	lcLeafCount := countLeafNode(root.leftChild)
	rcLeafCount := countLeafNode(root.rightChild)

	return lcLeafCount + rcLeafCount
}
