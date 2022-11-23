package main

import (
	"fmt"
	"math"
)

func bstSearch(root *Node, data int) bool {
	if root == nil {
		return false
	}

	if root.data == data {
		return true
	}
	if root.data > data {
		return bstSearch(root.leftChild, data)
	} else {
		return bstSearch(root.rightChild, data)
	}
}

func bstPrintElementInRange(root *Node, start, end int) {
	if root == nil {
		return
	}

	if root.data >= start && root.data <= end {
		fmt.Print(root.data, " ")
	}
	if root.data <= start {
		bstPrintElementInRange(root.rightChild, start, end)
	} else if root.data > end {
		bstPrintElementInRange(root.leftChild, start, end)
	} else {
		bstPrintElementInRange(root.rightChild, start, end)
		bstPrintElementInRange(root.leftChild, start, end)
	}
}

func getMinData(root *Node) int {
	if root == nil {
		return math.MaxInt
	}
	leftMin := getMinData(root.leftChild)
	rightMin := getMinData(root.rightChild)
	return int(math.Min(float64(root.data), math.Min(float64(leftMin), float64(rightMin))))
}

func getMaxData(root *Node) int {
	if root == nil {
		return math.MinInt
	}
	leftMax := getMaxData(root.leftChild)
	rightMax := getMaxData(root.rightChild)
	return int(math.Max(float64(root.data), math.Max(float64(leftMax), float64(rightMax))))
}

func isTreeBST(root *Node) bool {
	if root == nil {
		return true
	}
	isLeftTreeBST := isTreeBST(root.leftChild)
	isRightTreeBST := isTreeBST(root.rightChild)
	if !(isRightTreeBST && isLeftTreeBST) {
		return false
	}
	maxDataLeftTree := getMaxData(root.leftChild)
	minDataRightTree := getMinData(root.rightChild)
	return root.data > maxDataLeftTree && root.data <= minDataRightTree
}

func isTreeBST2(root *Node) (int, int, bool) {
	if root == nil {
		return math.MaxInt, math.MinInt, true
	}
	minLeft, maxLeft, isLeftTreeBST := isTreeBST2(root.leftChild)
	minRight, maxRight, isRightTreeBST := isTreeBST2(root.rightChild)
	if !(isRightTreeBST && isLeftTreeBST) {
		return 0, 0, false
	}
	minData := int(math.Min(float64(root.data), math.Min(float64(minLeft), float64(minRight))))
	maxData := int(math.Max(float64(root.data), math.Max(float64(maxLeft), float64(maxRight))))
	return minData, maxData, root.data > maxLeft && root.data <= minRight
}

func isTreeBST3(root *Node, min, max int) bool {
	if root == nil {
		return true
	}
	isLeftTreeBST := isTreeBST3(root.leftChild, min, root.data-1)
	isRightTreeBST := isTreeBST3(root.rightChild, root.data, max)
	if !(isRightTreeBST && isLeftTreeBST) {
		return false
	}
	return root.data >= min && root.data <= max
}
