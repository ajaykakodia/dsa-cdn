package main

import "fmt"

func main() {
	fmt.Println("we are in character pattern")
	pattern1(4)
	println()
	pattern1(5)
	println()
	pattern2(4)
	println()
	pattern3(4)
	println()
	pattern4(4)
	println()
	pattern5(5)
}

/*
Print character pattern for a given number i.e
# N = 4

ABCD
ABCD
ABCD
ABCD

*/
func pattern1(n int) {
	for r := 1; r <= n; r++ {
		for c := 1; c <= n; c++ {
			fmt.Print(string('A' + c - 1))
		}
		fmt.Println()
	}
}

/*
Print character pattern for a given number i.e
# N = 4

ABCD
BCDE
CDEF
DEFG
*/
func pattern2(n int) {
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			fmt.Print(string('A' + r + c))
		}
		fmt.Println()
	}
}

/*
Print character pattern for a given number i.e
# N = 4

A
AB
ABC
ABCD
*/
func pattern3(n int) {
	for r := 0; r < n; r++ {
		for c := 0; c < r; c++ {
			fmt.Print(string('A' + c))
		}
		fmt.Println()
	}
}

/*
Print character pattern for a given number i.e
# N = 4

ABCD
ABC
AB
A
*/
func pattern4(n int) {
	for r := 0; r < n; r++ {
		for c := 0; c < n-r; c++ {
			fmt.Print(string('A' + c))
		}
		fmt.Println()
	}
}

/*
Print character pattern for a given number i.e
# N = 4

A
BB
CCC
DDDD
*/
func pattern5(n int) {
	for r := 0; r < n; r++ {
		for c := 0; c <= r; c++ {
			fmt.Print(string('A' + r))
		}
		fmt.Println()
	}
}
