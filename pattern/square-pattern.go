package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("We are in Simple Patterns")
	// SquarePattern
	SquarePattern1(4)
	fmt.Println()
	SquarePattern1(5)
	fmt.Println()
	SquarePattern2(5)
	fmt.Println()
	SquarePattern3(5)
	fmt.Println()
	SquarePattern3(3)
	fmt.Println()
	SquarePattern4(5)
	fmt.Println()
	SquarePattern5(5)
	fmt.Println()
	SquarePattern6(4)
	fmt.Println()
	SquarePattern6(6)
	fmt.Println()
	SquarePattern7(4)
}

/*
Print Square of star of provide number i.e

****
****
****
****
*/
func SquarePattern1(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
Print Square of row of provide number i.e
# N = 5

1 1 1 1 1
2 2 2 2 2
3 3 3 3 3
4 4 4 4 4
5 5 5 5 5
*/
func SquarePattern2(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(strconv.Itoa(i+1) + " ")
		}
		fmt.Println()
	}
}

/*
Print Square of row of provide number i.e
# N = 5

1 2 3 4 5
1 2 3 4 5
1 2 3 4 5
1 2 3 4 5
1 2 3 4 5
*/
func SquarePattern3(n int) {
	for r := 1; r <= n; r++ {
		for c := 1; c <= n; c++ {
			print(strconv.Itoa(c) + " ")
		}
		println()
	}
}

/*
Print Square of row of provide number i.e
# N = 5

5 4 3 2 1
5 4 3 2 1
5 4 3 2 1
5 4 3 2 1
5 4 3 2 1
*/
func SquarePattern4(n int) {
	for r := 1; r <= n; r++ {
		for c := 1; c <= n; c++ {
			print(strconv.Itoa(n-c+1) + " ")
		}
		println()
	}
}

/*
Print Square of row of provide number i.e
# N = 5

1 2 3 4 5
2 3 4 5 6
3 4 5 6 7
4 5 6 7 8
5 6 7 8 9
*/
func SquarePattern5(n int) {
	for r := 1; r <= n; r++ {
		for c := 1; c <= n; c++ {
			fmt.Print(strconv.Itoa(r+c-1) + " ")
		}
		fmt.Println()
	}
}

/*
Print Square of row of provide number i.e
# N = 4

4555
3455
2345
1234
*/
func SquarePattern6(n int) {
	k := n + 1
	for r := 1; r <= n; r++ {
		for c := 1; c <= r; c++ {
			fmt.Print(strconv.Itoa(k-r+c-1) + " ")
		}
		for c := 1; c <= n-r; c++ {
			fmt.Print(strconv.Itoa(k) + " ")
		}
		fmt.Println()
	}
}

/*
Print Square of row of provide number i.e
# N = 4

12344321
123**321
12****21
1******1
*/
func SquarePattern7(n int) {
	for r := 1; r <= n; r++ {
		for c := 1; c <= n-r+1; c++ {
			fmt.Print(strconv.Itoa(c) + " ")
		}
		for c := 1; c <= 2*(r-1); c++ {
			fmt.Print("* ")
		}
		for c := 1; c <= n-r+1; c++ {
			fmt.Print(strconv.Itoa(n-c+1) + " ")
		}
		fmt.Println()
	}
}
