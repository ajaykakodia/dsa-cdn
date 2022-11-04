package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(`Find unique in array, array have duplicate element for other and there is a single element that is unique.
		i.e array [6,1,6,4,1,3,3] have 4 unique and all other have duplicate`)

	findUniqueEleB([]int{6, 1, 6, 4, 3, 1, 3})
	findUniqueEleSort([]int{6, 1, 6, 3, 1, 3, 4})
	findUniqueEleMap([]int{6, 1, 6, 4, 3, 1, 3})
	findUniqueEleXOR([]int{6, 1, 6, 4, 3, 1, 3})
}

// Brute force solution
func findUniqueEleB(arr []int) {
	isUnique := true
	for i := 0; i < len(arr); i++ {
		isUnique = true
		for j := 0; j < len(arr); j++ {
			if i != j && arr[i] == arr[j] {
				isUnique = false
			}
		}
		if isUnique {
			fmt.Println("Unique Element in Array Brute Force Solution: ", arr[i])
			break
		}
	}
}

// Solution by Sort
func findUniqueEleSort(arr []int) {
	sort.Ints(arr)

	for i := 0; i < len(arr); {
		if i == len(arr)-1 {
			fmt.Println("Unique Element in Array Sort Solution: ", arr[i])
			break
		}
		if arr[i] != arr[i+1] {
			fmt.Println("Unique Element in Array Sort Solution : ", arr[i])
			break
		}
		i = i + 2
	}
}

// Solution by Map
func findUniqueEleMap(arr []int) {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = m[arr[i]] + 1
	}

	for k, v := range m {
		if v == 1 {
			fmt.Println("Unique Element in Array Map Solution : ", k)
			break
		}
	}
}

// Solution by XOR
func findUniqueEleXOR(arr []int) {

	x := 0

	for i := 0; i < len(arr); i++ {
		x = x ^ arr[i]
	}

	fmt.Println("Unique Element in Array XOR Solution: ", x)
}
