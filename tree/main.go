package main

import (
	"fmt"
	"math"
)

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

	paths := printNodeToRootPath(bTree.root, 5)
	fmt.Println("Path to leaf 5: ", paths)

	paths = printNodeToRootPath(bTree.root, 19)
	fmt.Println("Path to leaf 19: ", paths)

	bTree.root = createBinarySearchTree()
	// bTree.LevelWisePrint()
	fmt.Println("28 exits in BST: ", bstSearch(bTree.root, 28))
	fmt.Println("18 exits in BST: ", bstSearch(bTree.root, 18))
	fmt.Println("7 exits in BST: ", bstSearch(bTree.root, 7))
	fmt.Println("Print in range => 3 & 30")
	bstPrintElementInRange(bTree.root, 3, 30)
	fmt.Println()
	fmt.Println("IsTreeBST method")
	fmt.Println("Is current Tree is Binary Search Tree: ", isTreeBST(bTree.root))
	bTree.root = createBinarySearchTree2()
	fmt.Println("Is current Tree is Binary Search Tree: ", isTreeBST(bTree.root))

	fmt.Println("IsTreeBST2 method")
	bTree.root = createBinarySearchTree()
	_, _, result := isTreeBST2(bTree.root)
	fmt.Println("Is current Tree is Binary Search Tree: ", result)
	bTree.root = createBinarySearchTree2()
	_, _, result = isTreeBST2(bTree.root)
	fmt.Println("Is current Tree is Binary Search Tree: ", result)

	fmt.Println("IsTreeBST3 method")
	bTree.root = createBinarySearchTree()
	fmt.Println("Is current Tree is Binary Search Tree: ", isTreeBST3(bTree.root, math.MinInt, math.MaxInt))
	bTree.root = createBinarySearchTree2()
	fmt.Println("Is current Tree is Binary Search Tree: ", isTreeBST3(bTree.root, math.MinInt, math.MaxInt))

	bsTree := &BinarySearchTree{}
	bsTree.Insert(10)
	bsTree.Insert(5)
	bsTree.Insert(2)
	bsTree.Insert(6)
	bsTree.Insert(1)
	bsTree.Insert(3)
	bsTree.Insert(7)
	bsTree.Insert(25)
	bsTree.Insert(15)
	bsTree.Insert(30)
	bsTree.Insert(11)
	bsTree.Insert(17)
	bsTree.Insert(28)
	bsTree.Insert(33)
	bsTree.Insert(35)
	bsTree.Insert(32)
	bsTree.Print()
	fmt.Println("Deleted 25 from BST: ", bsTree.Delete(25))
	bsTree.Print()
	fmt.Println("Deleted 1 from BST: ", bsTree.Delete(1))
	bsTree.Print()
	fmt.Println("Deleted 6 from BST: ", bsTree.Delete(6))
	bsTree.Print()
	fmt.Println("Deleted 10 from BST: ", bsTree.Delete(10))
	bsTree.Print()

	fmt.Println("================================= Generic Tree =====================================")
	gTree := NewGenericTree()
	gTree.root = createGenericTree(nil)
	gTree.Print()
	fmt.Println("Number of Nodes in current generic tree: ", gTree.NumberOfNodes())
	fmt.Println("Height of current generic tree: ", gTree.Height())

	gTree.root = createGenericTreeLevelWise()
	gTree.Print()
	fmt.Println("Number of Nodes in current generic tree: ", gTree.NumberOfNodes())
	fmt.Println("Height of current generic tree: ", gTree.Height())
	fmt.Println("Print generic tree level wise")
	gTree.PrintLevelWise()
}
