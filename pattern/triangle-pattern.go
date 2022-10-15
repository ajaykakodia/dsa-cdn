package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("We are in triangular patterns")
	pattern1(5)
	println()
	pattern2(5)
	println()
	pattern3(5)
}

/*
Print triangle of numbers for a given number i.e
# N = 5

1
2 2
3 3 3
4 4 4 4
5 5 5 5 5
*/
func pattern1(n int) {
	for r := 1; r <= n; r++ {
		for c := 1; c <= r; c++ {
			fmt.Print(strconv.Itoa(r) + " ")
		}
		fmt.Println()
	}
}

/*
Print triangle of numbers for a given number i.e
# N = 5

1
1 2
1 2 3
1 2 3 4
1 2 3 4 5
*/
func pattern2(n int) {
	for r := 1; r <= n; r++ {
		for c := 1; c <= r; c++ {
			fmt.Print(strconv.Itoa(c) + " ")
		}
		fmt.Println()
	}
}

/*
Print triangle of numbers for a given number i.e
# N = 5

1
2 3
4 5 6
7 8 9 10
11 12 13 14 15
*/
func pattern3(n int) {
	k := 1
	for r := 1; r <= n; r++ {
		for c := 1; c <= r; c++ {
			fmt.Printf("%-3d", k)
			k++
		}
		fmt.Println()
	}
}
