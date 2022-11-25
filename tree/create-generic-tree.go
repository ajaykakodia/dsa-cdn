package main

import "fmt"

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
