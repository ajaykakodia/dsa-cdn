package main

import "fmt"

func main() {
	fmt.Println(`Merge Sort: 
		It depends on recursion, first we find mid point of list and reach to its last node 
		after that we combine/ merge sorted small arrays and it returns us sorted array
		It totally depend on concept of how to merge two already sorted arrays`)

	arr := mergeSortedArray([]int{5, 8, 10, 15, 19}, []int{2, 6, 7, 11, 14})
	fmt.Printf("Merged sorted array: %v\n", arr)
	fmt.Println("Merge Sort")
	arr1 := mergeSort([]int{5, 3, 7, 6, 10, 15, 13, 45, 22})
	fmt.Printf("Sorted array: %v\n", arr1)
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := (len(arr)) / 2

	arr1 := mergeSort(arr[:mid])
	arr2 := mergeSort(arr[mid:])
	return mergeSortedArray(arr1, arr2)
}

func mergeSortedArray(arr1, arr2 []int) []int {
	l1 := len(arr1)
	l2 := len(arr2)
	arr := []int{}
	var i, j int

	for i < l1 && j < l2 {
		if arr1[i] > arr2[j] {
			arr = append(arr, arr2[j])
			j++
			continue
		}
		arr = append(arr, arr1[i])
		i++
	}

	arr = append(arr, arr1[i:]...)
	arr = append(arr, arr2[j:]...)

	return arr
}
