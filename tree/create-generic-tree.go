package main

import (
	"fmt"
	queue "github.com/ajaykakodia/lib/queue"
)

func createGenericTree(child *int) *GNode {
	if child == nil {
		fmt.Println("Enter root data: ")
	} else {
		fmt.Println("Enter data for child: ", *child)
	}

	var data, numChild int
	_, err := fmt.Scanln(&data)
	if err != nil || data == -1 {
		return nil
	}
	root := NewGNode(data)
	fmt.Println("Please enter no of child for ", data)
	_, err = fmt.Scanln(&numChild)
	if err == nil {
		for numChild > 0 {
			childNode := createGenericTree(&data)
			root.children = append(root.children, childNode)
			numChild--
		}
	}
	return root
}

func createGenericTreeLevelWise() *GNode {
	fmt.Println("Please enter root data: ")

	var data, numChild int
	var que queue.GQueue

	_, err := fmt.Scanln(&data)
	if err != nil || data == -1 {
		return nil
	}
	root := NewGNode(data)
	que.Enqueue(root)

	for !que.IsEmpty() {
		node := que.Dequeue().(*GNode)
		fmt.Println("Please enter no of child for ", node.data)
		_, err = fmt.Scanln(&numChild)
		if err == nil {
			for numChild > 0 {
				fmt.Println("Please enter child for ", node.data)
				_, err := fmt.Scanln(&data)
				if err == nil {
					childNode := NewGNode(data)
					node.children = append(node.children, childNode)
					que.Enqueue(childNode)
				}
				numChild--
			}
		}
	}

	return root
}
