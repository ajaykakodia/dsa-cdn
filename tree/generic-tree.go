package main

import (
	"fmt"
	"math"
)

type GNode struct {
	data     int
	children []*GNode
}

func NewGNode(data int) *GNode {
	return &GNode{data: data, children: []*GNode{}}
}

func (n *GNode) Print() {
	fmt.Print(n.data, ": ")
	for _, val := range n.children {
		fmt.Print(val.data, ", ")
	}
	fmt.Println()
}

type GenericTree struct {
	root *GNode
}

func NewGenericTree() *GenericTree {
	return &GenericTree{}
}

func (gt *GenericTree) Print() {
	if gt.root == nil {
		return
	}
	printGTree(gt.root)
}

func (gt *GenericTree) NumberOfNodes() int {
	if gt.root == nil {
		return 0
	}
	return countGTreeNode(gt.root)
}

func (gt *GenericTree) Height() int {
	if gt.root == nil {
		return 0
	}
	return heightOfGTree(gt.root)
}

func heightOfGTree(root *GNode) int {
	maxHeight := 0
	for _, node := range root.children {
		maxHeight = int(math.Max(float64(maxHeight), float64(heightOfGTree(node))))
	}
	return 1 + maxHeight
}

func countGTreeNode(root *GNode) int {
	numberOfNodes := 0
	for _, node := range root.children {
		nums := countGTreeNode(node)
		numberOfNodes = numberOfNodes + nums
	}
	return 1 + numberOfNodes
}

func printGTree(root *GNode) {
	root.Print()
	for _, child := range root.children {
		printGTree(child)
	}
}
