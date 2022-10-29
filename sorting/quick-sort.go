package main

import "fmt"

func main() {
	fmt.Println(`We are in Quick Sort folder and Quick sort depends on Pivoting - 
		Where we first find any random number proper place and then move
		smaller number on left side and greater number on right side
		and return pivot index - index of random number
		and then call Quicksort for remaining both side of array parts`)

	arr := []int{5, 3, 7, 6, 10, 15, 13, 45, 22}
	fmt.Printf("UnSorted array : %v\n", arr)
	quickSort(arr)
	fmt.Printf("Sorted array : %v\n", arr)
}

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pvi := pivotIndex(arr)
	quickSort(arr[:pvi])
	quickSort(arr[pvi+1:])
}

// pivotIndex -> get pivot index and manage elements around pivot elements
func pivotIndex(arr []int) int {
	// checking for length
	if len(arr) == 1 {
		return 0
	}
	// we can use any random index but for simplicity we are taking 1st element
	eleToPivot := arr[0]
	pi := 0
	// calculating pivot index
	for _, v := range arr {
		if eleToPivot > v {
			pi++
		}
	}
	// replace pvi element with selected element
	arr[0], arr[pi] = arr[pi], arr[0]

	i, j := 0, len(arr)-1

	// run loop until all element is in place
	for i < pi || j > pi {
		// check for left side elements to swap
		for arr[i] < eleToPivot && i < pi {
			i++
		}
		// check for right side elements to swap
		for arr[j] >= eleToPivot && j > pi {
			j--
		}
		// swap elements and change i and j values
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}

	return pi
}
