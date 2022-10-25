package main

import "fmt"

func main() {
	fmt.Println("We are in array folder")

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("Array before reverse: %v\n", arr)
	reverseSlice(arr)
	fmt.Printf("Array after reverse: %v\n", arr)
}

func reverseSlice(arr []int) {
	l := len(arr)
	mid := l / 2
	for i := 0; i < mid; i++ {
		arr[i], arr[l-1-i] = arr[l-1-i], arr[i]
	}
}
