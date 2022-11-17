package main

import (
	"fmt"
	"strconv"
	"strings"
)

func createBinaryTreeByInput() *Node {
	var nodeData string
	fmt.Scanln(&nodeData)
	val, err := strconv.Atoi(nodeData)
	if err != nil || val == -1 {
		return nil
	}
	node := NewNode(val)
	node.leftChild = createBinaryTreeByInput()
	node.rightChild = createBinaryTreeByInput()
	return node
}

func createBinaryTree() *Node {
	var nodeData string = "5 2 6 8 -1 -1 -1 7 -1 -1 3 -1 -1"
	rootNode := &Node{}
	dataToPopulate := []int{}
	for _, val := range strings.Fields(nodeData) {
		data, _ := strconv.Atoi(val)
		dataToPopulate = append(dataToPopulate, data)
	}
	rootNode = populateData1(dataToPopulate)
	return rootNode
}

// TODO fixing required
func populateData1(arr []int) *Node {
	if len(arr) == 0 || arr[0] == -1 {
		return nil
	}

	node := NewNode(arr[0])
	node.leftChild = populateData1(arr[1:])
	if node.leftChild == nil {

	}
	node.rightChild = populateData1(arr[2:])
	return node
}

// TODO fixing required
func populateData(arr []int, root *Node) []int {
	if len(arr) == 0 {
		root = nil
		return nil
	}
	if arr[0] == -1 {
		root = nil
		return arr[1:]
	}
	root.data = arr[0]
	root.leftChild = &Node{}
	root.rightChild = &Node{}
	if len(arr) > 1 {
		arr = populateData(arr[1:], root.leftChild)
	}
	if len(arr) > 1 {
		arr = populateData(arr[1:], root.rightChild)
	}
	return arr
}
