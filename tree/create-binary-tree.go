package main

import (
	"fmt"
	"github.com/ajaykakodia/lib/queue"
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

// createBinaryTree1 => create a balanced tree
func createBinaryTree1() *Node {
	var nodeData string = "5 2 6 8 -1 -1 -1 3 -1 -1 7 -1 10 -1 -1"
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

// createBinaryTree1 => create an unbalanced tree
func createBinaryTree2() *Node {
	var nodeData string = "5 2 6 8 9 -1 -1 -1 -1 3 -1 -1 7 -1 10 -1 -1"
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

// createBinaryTree1 => create an unbalanced tree
func createBinaryTree3() *Node {
	var nodeData string = "3 2 1 8 11 -1 -1 -1 10 -1 -1 5 15 -1 -1 6 -1 19 -1 -1 4 -1 -1"
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

func createBinaryTreeLevelWiseInput() *Node {
	que := queue.GQueue{}
	fmt.Println("Please enter interger values and if want nil for any data enter -1")
	fmt.Println("Enter Root Element: ")
	var (
		sData   string
		newNode *Node
	)

	fmt.Scanln(&sData)

	data, err := strconv.Atoi(sData)
	if err != nil || data == -1 {
		return nil
	}
	rootNode := NewNode(data)
	que.Enqueue(rootNode)

	for !que.IsEmpty() {
		ele := (que.Dequeue().(*Node))
		fmt.Println("Enter left child of ", ele.data)
		fmt.Scanln(&sData)
		data, err = strconv.Atoi(sData)
		if err == nil && data != -1 {
			newNode = NewNode(data)
			ele.leftChild = newNode
			que.Enqueue(newNode)
		}
		fmt.Println("Enter right child of ", ele.data)
		fmt.Scanln(&sData)
		data, err = strconv.Atoi(sData)
		if err == nil && data != -1 {
			newNode = NewNode(data)
			ele.rightChild = newNode
			que.Enqueue(newNode)
		}
	}

	return rootNode
}
