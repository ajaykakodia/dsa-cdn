package main

import (
	"strconv"
	"strings"
)

func createBinarySearchTree() *Node {
	var nodeData string = "10 5 2 1 -1 -1 3 -1 -1 6 -1 7 -1 -1 25 15 11 -1 -1 17 -1 -1 30 28 -1 -1 33 -1 35 -1 -1"
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

func createBinarySearchTree2() *Node {
	var nodeData string = "5 2 1 -1 -1 6 -1 -1 10 -1 -1"
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
