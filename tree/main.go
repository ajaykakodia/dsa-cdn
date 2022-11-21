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

	//bTree.root = createBinaryTreeLevelWiseInput()
	//bTree.LevelWisePrint()
	inOrder := []int{8, 1, 10, 2, 15, 5, 6, 19, 3, 4}
	preOrder := []int{3, 2, 1, 8, 10, 5, 15, 6, 19, 4}
	fmt.Println("In Order & Pre Order Process")
	bTree.root = processInOrderPreOrderForBinaryTree(inOrder, preOrder)
	bTree.LevelWisePrint()

	postOrder := []int{8, 10, 1, 15, 19, 6, 5, 2, 4, 3}
	fmt.Println("In Order & Post Order Process")
	bTree.root = processInOrderPostOrderForBinaryTree(inOrder, postOrder)
	bTree.LevelWisePrint()
}
