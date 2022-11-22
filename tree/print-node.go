package main

import "fmt"

func printKthLevel(root *Node, level int) {
	if root == nil {
		return
	}
	if level == 0 {
		fmt.Print(root.data, " ")
	}
	printKthLevel(root.leftChild, level-1)
	printKthLevel(root.rightChild, level-1)
}

func printKthLevel1(root *Node, level, currentHeight int) {
	if root == nil {
		return
	}
	if level == currentHeight {
		fmt.Print(root.data, " ")
	}
	printKthLevel1(root.leftChild, level, currentHeight+1)
	printKthLevel1(root.rightChild, level, currentHeight+1)
}

func printNodeToRootPath(root *Node, data int) []int {
	if root == nil {
		return nil
	}
	if root.data == data {
		return []int{root.data}
	}

	leftTreeMatch := printNodeToRootPath(root.leftChild, data)
	if leftTreeMatch != nil {
		leftTreeMatch = append(leftTreeMatch, root.data)
		return leftTreeMatch
	}

	rightTreeMatch := printNodeToRootPath(root.rightChild, data)
	if rightTreeMatch != nil {
		rightTreeMatch = append(rightTreeMatch, root.data)
		return rightTreeMatch
	}

	return nil
}
