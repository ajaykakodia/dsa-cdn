package main

import "fmt"

func main() {
	fmt.Println(`Selection Sort: 
		Here we select smallest number from array and place it on its correct place starts with index 0 
		with assumption that its smallest and compare with
		rest of element and find actual smallest number and replace it with current index`)

	arr := []int{5, 3, 7, 6, 10, 15, 11, 8, 4, 20, 9}
	fmt.Printf("unsorted array : %v\n", arr)
	selectionSort(arr)
	fmt.Printf("sorted array : %v\n", arr)
}

func selectionSort(arr []int) {
	si := 0
	l := len(arr)
	for i := 0; i < l; i++ {
		si = i
		for j := i; j < l; j++ {
			if arr[si] > arr[j] {
				si = j
			}
		}
		arr[si], arr[i] = arr[i], arr[si]
	}
}
