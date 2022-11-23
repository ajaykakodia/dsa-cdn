package main

import (
	"fmt"
	"github.com/ajaykakodia/lib/queue"
)

type BinaryTree struct {
	root *Node
}

func NewBinaryTree(root *Node) *BinaryTree {
	return &BinaryTree{root}
}

type Node struct {
	data       int
	leftChild  *Node
	rightChild *Node
}

func NewNode(data int) *Node {
	return &Node{data: data}
}

func (n *Node) Print() {
	fmt.Print(n.data, ": ")
	if n.leftChild != nil {
		fmt.Print("L", n.leftChild.data, ", ")
	}
	if n.rightChild != nil {
		fmt.Print("R", n.rightChild.data)
	}
	fmt.Println()
}

func (n *Node) IsLeaf() bool {
	return n == nil || (n.leftChild == nil && n.rightChild == nil)
}

func (bn *BinaryTree) Print() {
	if bn.root == nil {
		return
	}
	bn.root.Print()
	lbTree := NewBinaryTree(bn.root.leftChild)
	rbTree := NewBinaryTree(bn.root.rightChild)
	lbTree.Print()
	rbTree.Print()
}

func (bn *BinaryTree) PostOrderPrint() {
	if bn.root == nil {
		return
	}
	lbTree := NewBinaryTree(bn.root.leftChild)
	rbTree := NewBinaryTree(bn.root.rightChild)
	lbTree.Print()
	rbTree.Print()
	bn.root.Print()
}

func (bn *BinaryTree) InOrderPrint() {
	if bn.root == nil {
		return
	}
	lbTree := NewBinaryTree(bn.root.leftChild)
	rbTree := NewBinaryTree(bn.root.rightChild)
	lbTree.Print()
	bn.root.Print()
	rbTree.Print()
}

func (bn *BinaryTree) LevelWisePrint() {
	if bn.root == nil {
		return
	}
	que := queue.GQueue{}
	que.Enqueue(bn.root)
	for !que.IsEmpty() {
		ele := que.Dequeue().(*Node)
		ele.Print()
		if ele.leftChild != nil {
			que.Enqueue(ele.leftChild)
		}
		if ele.rightChild != nil {
			que.Enqueue(ele.rightChild)
		}
	}
}
