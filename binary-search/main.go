package main

import "fmt"

func main() {
	fmt.Println("In Binary Search Folder")
	ele := 3
	arr := []int{3, 6, 9, 12, 17, 19, 30, 40, 45, 50, 60}

	fmt.Printf("Element %d on index : %d in array %#v\n", ele, binarySearch(arr, ele), arr)
	fmt.Printf("Recursion - Element %d on index : %d in array %#v\n", ele, binarySearchR(arr, ele, 0, len(arr)-1), arr)
}

func binarySearch(arr []int, ele int) int {

	si, li := 0, len(arr)-1
	mid := 0

	for si <= li {
		mid = (si + li) / 2
		if arr[mid] == ele {
			return mid
		}

		if arr[mid] > ele {
			li = mid - 1
		} else {
			si = mid + 1
		}
	}

	return -1
}

func binarySearchR(arr []int, ele, si, li int) int {
	if si > li {
		return -1
	}
	mid := int((si + li) / 2)

	if arr[mid] == ele {
		return mid
	}

	if arr[mid] > ele {
		return binarySearchR(arr, ele, si, mid-1)
	} else {
		return binarySearchR(arr, ele, mid+1, li)
	}
}
