package main

import "fmt"

func main() {
	fmt.Println("Equilibrium Index of an array: Its an index from there sum of left side array is equal to sum of right side array")

	//findEquilibriumIndexB([]int{1, 3, 4, 2, 2, 5, 4, 1})

	findEquilibriumIndexOptimized([]int{1, 3, 4, 2, 2, 5, 4, 1})
}

// findEquilibriumIndexB - Brute-force solution
func findEquilibriumIndexB(arr []int) {
	lsum, rsum := 0, 0

	for i := range arr {
		lsum, rsum = 0, 0
		for j := i - 1; j >= 0; j-- {
			lsum = lsum + arr[j]
		}
		for j := i + 1; j < len(arr); j++ {
			rsum = rsum + arr[j]
		}
		if lsum == rsum {
			fmt.Printf("Equilibrium Index is %d and sum is %d\n", i, lsum)
		}
	}
}

// findEquilibriumIndexOptimized - Optimized solution by taking sum of all array and then move lsum and rsum
func findEquilibriumIndexOptimized(arr []int) {
	tsum, lsum, rsum := 0, 0, 0
	for _, ele := range arr {
		tsum = tsum + ele
	}
	lsum, rsum = 0, tsum
	if len(arr) > 0 {
		rsum = rsum - arr[0]
	}

	for i := 1; i < len(arr); i++ {
		// avoiding check here for every iteration, put outside and manage rsum there
		// if i == 0 {
		// 	lsum = lsum + arr[i-1]
		// }

		lsum = lsum + arr[i-1]
		rsum = rsum - arr[i]

		if lsum == rsum {
			fmt.Printf("Equilibrium Index is %d and sum is %d\n", i, lsum)
			return
		}
	}
	fmt.Println("No Equilibrium Point Found")
}
