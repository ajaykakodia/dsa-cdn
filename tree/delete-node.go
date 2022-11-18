package main

func deleteLeafNodes(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.leftChild == nil && root.rightChild == nil {
		return nil
	}
	root.leftChild = deleteLeafNodes(root.leftChild)
	root.rightChild = deleteLeafNodes(root.rightChild)
	return root
}
