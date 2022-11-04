package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(`Find duplicate in a array that starts with 0 and all numbers except one duplicate
		and We have to find that duplicate`)

	findDuplicateB([]int{2, 1, 0, 5, 3, 4, 7, 6, 5})
	findDuplicateSort([]int{2, 1, 0, 5, 3, 4, 7, 6, 5})
	findDuplicateMap([]int{2, 1, 0, 5, 3, 4, 7, 6, 5})
	findDuplicateNaturalSum([]int{2, 1, 0, 5, 3, 4, 7, 6, 5})
}

// Brute force solution
func findDuplicateB(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i != j && arr[i] == arr[j] {
				fmt.Println("Duplicate Element Brute Fores : ", arr[i])
				return
			}
		}
	}
}

// Sorting solution
func findDuplicateSort(arr []int) {
	sort.Ints(arr)
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			fmt.Println("Duplicate Element Sorting Solution : ", arr[i])
			return
		}
	}
}

// Mapping solution
func findDuplicateMap(arr []int) {
	m := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		if _, ok := m[arr[i]]; ok {
			fmt.Println("Duplicate Element Mapping Solution : ", arr[i])
			return
		}
		m[arr[i]] = true
	}
}

// Natural sum solution
// It starts with 0 and one number is duplicate so We are calculating sum for (len-2) numbers
// And then calculate Sum of array and Subtract it from Natural Sum num give us answer
func findDuplicateNaturalSum(arr []int) {
	l := len(arr)
	nSum := (l - 2) * (l - 1) / 2
	arrSum := 0

	for _, v := range arr {
		arrSum = arrSum + v
	}
	fmt.Println("Duplicate Element Natural Sum Solution : ", arrSum-nSum)
}
