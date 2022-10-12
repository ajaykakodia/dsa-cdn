package main

import "fmt"

func main() {
	fmt.Println(`Bubble Sort : It checks two consecutive element and replace/ bubbles
	if one is greater than other and moved to next index, 
	by this it set the largest remaining element in last at every iteration`)

	arr := []int{5, 3, 7, 6, 10, 15, 11, 8, 4, 20, 9}
	fmt.Printf("unsorted array : %v\n", arr)
	bubbleSort(arr)
	fmt.Printf("sorted array : %v\n", arr)
}

func bubbleSort(arr []int) {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
