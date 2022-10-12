package main

import "fmt"

func main() {
	fmt.Println(`Insertion Sort: 
		Here we start from index 1 and assume all array left to current index is sorted and we have to place current element in right place 
		in already sorted lift side array by moving larger elements to one place and puss element in right place`)

	arr := []int{5, 3, 7, 6, 10, 15, 11, 8, 4, 20, 9}
	fmt.Printf("unsorted array : %v\n", arr)
	insertionSort(arr)
	fmt.Printf("sorted array : %v\n", arr)
}

func insertionSort(arr []int) {
	ele := 0

	for i := 1; i < len(arr); i++ {
		ele = arr[i]
		j := i - 1
		for j = i - 1; j >= 0; j-- {
			if arr[j] > ele {
				arr[j+1] = arr[j]
				continue
			}
			break
		}
		arr[j+1] = ele
	}
}
