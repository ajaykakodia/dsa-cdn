package main

import "fmt"

func main() {
	fmt.Println("In recursion -1 - recursion folder")
	print1ToN(5)
	fmt.Println()
	printNT1(5)
	fmt.Println()
	fmt.Println(fib(5))
	fmt.Println(fibLoop(5))
	fmt.Println(firstIndexInList([]int{3, 5, 7, 9, 11, 11, 12, 6, 5, 7, 8, 11, 15}, 11))
	fmt.Println(lastIndexInList([]int{3, 5, 7, 9, 10, 11, 12, 6, 5, 7, 8, 12, 11, 15}, 11))

	fmt.Println(isListSorted([]int{3, 5, 7, 8, 10, 11, 12}))
}

func print1ToN(n int) {
	if n == 0 {
		return
	}
	print1ToN(n - 1)
	fmt.Print(n)
}

func printNT1(n int) {
	if n == 0 {
		return
	}
	fmt.Print(n)
	print1ToN(n - 1)
}

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fibLoop(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	a, b := 1, 1
	for i := 0; i < n-2; i++ {
		a, b = b, a+b
	}
	return b
}

func firstIndexInList(arr []int, x int) int {
	if len(arr) == 0 {
		return -1
	}
	if arr[0] == x {
		return 0
	}
	smallListIndex := firstIndexInList(arr[1:], x)
	if smallListIndex != -1 {
		return smallListIndex + 1
	}
	return smallListIndex
}

func lastIndexInList(arr []int, x int) int {
	if len(arr) == 0 {
		return -1
	}

	smallListIndex := lastIndexInList(arr[1:], x)
	if smallListIndex != -1 {
		return smallListIndex + 1
	}

	if arr[0] == x {
		return 0
	}

	return -1
}

func isListSorted(arr []int) bool {
	l := len(arr)
	if l == 0 || l == 1 {
		return true
	}
	if arr[0] > arr[1] {
		return false
	}

	isSorted := isListSorted(arr[1:])
	return isSorted
}
