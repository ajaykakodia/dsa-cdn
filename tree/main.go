package main

import "fmt"

func main() {
	bTree := NewBinaryTree(nil)
	bTree.root = createBinaryTree()
	fmt.Println("PreOrder Traversal")
	PreOrderTraversal(bTree.root)
	fmt.Println("PostOrder Traversal")
	PostOrderTraversal(bTree.root)
	fmt.Println("InOrder Traversal")
	InOrderTraversal(bTree.root)
	fmt.Println("Total Node Count: ", countTotalNode(bTree.root))
	fmt.Println("Total Leaf Node: ", countLeafNode(bTree.root))
	fmt.Println("Height of Tree: ", getHeight(bTree.root))
	fmt.Println("Max Node of Tree: ", getMaxNodeInTree(bTree.root))
	fmt.Print("Nodes at level 2: ")
	printKthLevel(bTree.root, 2)
	fmt.Println()
	fmt.Print("Nodes at level 3: ")
	printKthLevel1(bTree.root, 3, 0)
	fmt.Println()
	bTree.root = deleteLeafNodes(bTree.root)
	fmt.Println("Tree after delete leaf nodes")
	bTree.Print()
	bTree.root = createBinaryTree1()
	bTree.Print()
	balanced, _ := isTreeBalanced(bTree.root)
	fmt.Println("Is current tree is balanced: ", balanced)
	bTree.root = createBinaryTree2()
	bTree.Print()
	balanced, _ = isTreeBalanced(bTree.root)
	fmt.Println("Is current tree is balanced: ", balanced)

	bTree.root = createBinaryTree3()
	bTree.Print()
	height, dia := getDiameterOfTree(bTree.root)
	fmt.Println("Height of tree ", height, " & Diameter of tree :", dia)

	bTree.root = createBinaryTreeLevelWiseInput()
	bTree.LevelWisePrint()
}
