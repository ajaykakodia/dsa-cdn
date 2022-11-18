package main

import "math"

/*
A tree is balanced when difference of height of Its children is not more than 1
*/
// isTreeBalancedB => Not an optimized method as we are calculation height at every step, and it causes traversing every node again and again
func isTreeBalancedB(root *Node) bool {
	if root == nil {
		return true
	}
	isLCTreeBalanced := isTreeBalancedB(root.leftChild)
	isRCTreeBalanced := isTreeBalancedB(root.rightChild)
	if !(isRCTreeBalanced && isLCTreeBalanced) {
		return false
	}
	lcHeight := getHeight(root.leftChild)
	rcHeight := getHeight(root.rightChild)
	return !(lcHeight-rcHeight > 1 || rcHeight-lcHeight > 1)
}

// isTreeBalanced
func isTreeBalanced(root *Node) (bool, int) {
	if root == nil {
		return true, 0
	}
	isLCTreeBalanced, lcHeight := isTreeBalanced(root.leftChild)
	isRCTreeBalanced, rcHeight := isTreeBalanced(root.rightChild)
	if !(isRCTreeBalanced && isLCTreeBalanced) {
		return false, 0
	}

	return !(lcHeight-rcHeight > 1 || rcHeight-lcHeight > 1), 1 + int(math.Max(float64(lcHeight), float64(rcHeight)))
}
