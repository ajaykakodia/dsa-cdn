package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(`We need to find pair in an array for a given sum
		i.e - [2, -6, 6, 4, -4, -2] & pairSum = 0
		expected output - [-6, 6] [4, -4] [2, -2]
		i.e - [2, 3, 3, 4, 5, 8, 9, 9, 9] & pairSum = 12
		expected output - [3, 9] [3, 9] [3, 9]
						  [3, 9] [3, 9] [3, 9]
						  [4, 8]
		i.e - [2, 2, 2, 2, 2] & pairSum = 4
		expected output - [2, 2] [2, 2] [2, 2] [2, 2]
						  [2, 2] [2, 2] [2, 2]
						  [2, 2] [2, 2]
						  [2, 2] `)

	pairSumB([]int{2, -6, 4, 6, -4, 5, 7, -2}, 0)
	pairSumB([]int{2, 3, 3, 4, 5, 8, 9, 9, 9}, 12)
	pairSumB([]int{2, 2, 2, 2, 2}, 4)
	fmt.Println("Pair Sum Sorting Solution")
	pairSumSortingSolution([]int{2, -6, 4, 6, -4, 5, 7, -2}, 0)
	fmt.Println()
	pairSumSortingSolution([]int{2, 3, 3, 4, 5, 8, 9, 9, 9}, 12)
	fmt.Println()
	pairSumSortingSolution([]int{2, 2, 2, 2, 2}, 4)

}

// Brute Force solution
func pairSumB(arr []int, pairSum int) {
	requiredEle := 0
	for i := 0; i < len(arr); i++ {
		requiredEle = pairSum - arr[i]
		for j := i + 1; j < len(arr); j++ {
			if arr[j] == requiredEle {
				fmt.Printf("[%d, %d] ", arr[i], arr[j])
			}
		}
	}
	fmt.Println()
}

// Sorting solution
func pairSumSortingSolution(arr []int, pairSum int) {
	sort.Ints(arr)
	i, j := 0, len(arr)-1
	sum, k := 0, 0
	for i < j {
		sum = arr[i] + arr[j]
		if sum < pairSum {
			i++
			continue
		}
		if sum == pairSum {
			fmt.Printf("[%d, %d] ", arr[i], arr[j])
			k = j - 1
			for k > i && arr[j] == arr[k] {
				fmt.Printf("[%d, %d] ", arr[i], arr[j])
				k--
			}
			i++
			continue
		}
		j--
	}
}
