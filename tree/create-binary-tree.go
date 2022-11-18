package main

import (
	"fmt"
	"strconv"
	"strings"
)

var countPopulatedData int = 0

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
	var nodeData string = "5 2 6 8 -1 -1 -1 10 -1 -1 3 9 -1 -1 11 -1 -1"
	rootNode := &Node{}
	dataToPopulate := []int{}
	for _, val := range strings.Fields(nodeData) {
		data, _ := strconv.Atoi(val)
		dataToPopulate = append(dataToPopulate, data)
	}
	countPopulatedData = 0
	rootNode = populateData(dataToPopulate)
	return rootNode
}

func populateData(arr []int) *Node {
	if countPopulatedData >= len(arr) {
		return nil
	}

	if arr[countPopulatedData] == -1 {
		return nil
	}

	node := NewNode(arr[countPopulatedData])
	countPopulatedData++
	node.leftChild = populateData(arr)
	countPopulatedData++
	node.rightChild = populateData(arr)
	return node
}
