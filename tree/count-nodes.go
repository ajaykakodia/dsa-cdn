package main

import "math"

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

func getHeight(root *Node) int {
	if root == nil {
		return 0
	}
	lcHeight := getHeight(root.leftChild)
	rcHeight := getHeight(root.rightChild)
	return 1 + int(math.Max(float64(lcHeight), float64(rcHeight)))
}

func getMaxNodeInTree(root *Node) int {
	if root == nil {
		return 0
	}
	lcHeight := getMaxNodeInTree(root.leftChild)
	rcHeight := getMaxNodeInTree(root.rightChild)

	return int(math.Max(float64(root.data), math.Max(float64(lcHeight), float64(rcHeight))))
}
