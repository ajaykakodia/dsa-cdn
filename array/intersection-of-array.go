package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Getting Intersection of two array: ")
	arr1 := []int{7, 8, 5, 6, 3, 10, 11}
	arr2 := []int{5, 8, 15, 12, 13, 7}

	intersectionOfArrayB(arr1, arr2)
	fmt.Println()

	intersectionOfArrayMerge([]int{7, 8, 5, 6, 3, 10, 11}, []int{5, 8, 15, 12, 13, 7})
	fmt.Println()

	intersectionOfArrayBinarySearch([]int{7, 8, 5, 6, 3, 10, 11, 15, 20, 25, 40}, []int{5, 8, 15, 12, 13, 7, 3})
	fmt.Println()

	intersectionOfArrayByMap([]int{7, 8, 5, 6, 3, 10, 11, 15, 20, 25, 40}, []int{5, 8, 15, 12, 13, 7, 3})
}

// intersectionOfArrayB - Brute-force solution
func intersectionOfArrayB(arr1, arr2 []int) {
	fmt.Println("Brute-fore Solution for Intersection of array: ")
	for _, val := range arr1 {
		for _, ele := range arr2 {
			if val == ele {
				fmt.Printf("%4d", ele)
			}
		}
	}
}

// intersectionOfArrayMerge - solution by sorting two array and then use merging of sorted array solution
func intersectionOfArrayMerge(arr1, arr2 []int) {
	fmt.Println("Merging of Sorted Array Solution for Intersection of array: ")
	// sort array by inbuilt methods
	sort.Ints(arr1)
	sort.Ints(arr2)
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			i++
		} else if arr1[i] > arr2[j] {
			j++
		} else {
			fmt.Printf("%4d", arr1[i])
			i++
			j++
		}
	}
}

// intersectionOfArrayBinarySearch - solution by sorting small array and search element of bigger array in sorted array by binary search - mlogn
func intersectionOfArrayBinarySearch(arr1, arr2 []int) {
	fmt.Println("Binary Search Solution for Intersection of array: ")
	// sort array by inbuilt methods
	var searchingArray, sortedArray []int

	if len(arr1) < len(arr2) {
		sortedArray = arr1
		searchingArray = arr2
	} else {
		sortedArray = arr2
		searchingArray = arr1
	}
	sort.Ints(sortedArray)
	for _, ele := range searchingArray {
		if binarySearch(sortedArray, ele) != -1 {
			fmt.Printf("%4d", ele)
		}
	}
}

// intersectionOfArrayByMap - store data of one in map and then check in other with loop - m+n
func intersectionOfArrayByMap(arr1, arr2 []int) {
	fmt.Println("Optimized Solution by map for Intersection of array: ")

	m := make(map[int]bool)

	for _, ele := range arr1 {
		m[ele] = true
	}

	for _, ele := range arr2 {
		if _, ok := m[ele]; ok {
			fmt.Printf("%4d", ele)
		}
	}
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
