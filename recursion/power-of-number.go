package main

import "fmt"

func main() {
	fmt.Println("Getting power of an number, (2)3 -> 8")

	fmt.Printf("Power %d of Num %d: %d\n", 7, 8, power(7, 8))
	fmt.Printf("Power %d of Num %d: %d\n", 7, 0, power(7, 0))
	fmt.Printf("Power %d of Num %d: %d\n", 2, 5, power(2, 5))

	fmt.Printf("Recursion: Power %d of Num %d: %d\n", 7, 8, powerRec(7, 8))
	fmt.Printf("Recursion: Power %d of Num %d: %d\n", 7, 0, powerRec(7, 0))
	fmt.Printf("Recursion: Power %d of Num %d: %d\n", 2, 5, powerRec(2, 5))

	fmt.Printf("Optimized Solution Recursion: Power %d of Num %d: %d\n", 7, 8, powerRecOptimize(7, 8))
	fmt.Printf("Optimized Solution Recursion: Power %d of Num %d: %d\n", 7, 0, powerRecOptimize(7, 0))
	fmt.Printf("Optimized Solution Recursion: Power %d of Num %d: %d\n", 7, 1, powerRecOptimize(7, 1))
	fmt.Printf("Optimized Solution Recursion: Power %d of Num %d: %d\n", 7, 2, powerRecOptimize(7, 2))
	fmt.Printf("Optimized Solution Recursion: Power %d of Num %d: %d\n", 7, 3, powerRecOptimize(7, 3))
	fmt.Printf("Optimized Solution Recursion: Power %d of Num %d: %d\n", 2, 5, powerRecOptimize(2, 5))
}

// power - Brute-force solution
func power(num, pow int) int {
	ans := 1
	for i := 0; i < pow; i++ {
		ans = ans * num
	}
	return ans
}

// powerRec - Recursive solution of Power
func powerRec(num, pow int) int {
	if pow == 0 {
		return 1
	}
	return num * powerRec(num, pow-1)
}

// powerRecOptimize - Optimize Solution of Power
// for even power - p/2*p/2
// for odd power - num* p/2*p/2
func powerRecOptimize(num, pow int) int {
	if pow == 0 {
		return 1
	}
	halfPow := powerRecOptimize(num, pow/2)
	if pow%2 == 0 {
		return halfPow * halfPow
	}

	return num * halfPow * halfPow
}
