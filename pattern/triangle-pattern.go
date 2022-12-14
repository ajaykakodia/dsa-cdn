package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("We are in triangular patterns")
	tpattern1(5)
	println()
	tpattern2(5)
	println()
	tpattern3(5)
	println()
	tpattern4(4)

	println()
	tpattern5(4)
	println()
	tpattern6(4)
	println()
	tpattern7(4)
	println()
	tpattern8(4)
	println()
	tpattern9(4)
	println()
	tpattern10(4)
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
func tpattern1(n int) {
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
func tpattern2(n int) {
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
func tpattern3(n int) {
	k := 1
	for r := 1; r <= n; r++ {
		for c := 1; c <= r; c++ {
			fmt.Printf("%-3d", k)
			k++
		}
		fmt.Println()
	}
}

/*
Print triangle of numbers for a given number i.e
# N = 3

* * *
* *
*
*/
func tpattern4(n int) {
	for r := 1; r <= n; r++ {
		for c := 1; c <= n-r+1; c++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
Print triangle of numbers for a given number i.e
# N = 3

		*
	  * *
	* * *
*/
func tpattern5(n int) {
	for r := 1; r <= n; r++ {
		for c := 1; c <= n-r; c++ {
			fmt.Print(" ")
		}
		for c := 1; c <= r; c++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
Print triangle of numbers for a given number i.e
# N = 4

		  1
		 121
		12321
	   1234321
*/
func tpattern6(n int) {
	for r := 1; r <= n; r++ {
		for c := 1; c <= n-r; c++ {
			fmt.Print(" ")
		}
		for c := 1; c <= r; c++ {
			fmt.Print(c)
		}
		for c := 1; c < r; c++ {
			fmt.Print(r - c)
		}
		fmt.Println()
	}
}

/*
Print triangle of numbers for a given number i.e
# N = 4

		  *
		 ***
		*****
	   *******
*/
func tpattern7(n int) {
	for r := 1; r <= n; r++ {
		for c := 1; c <= n-r; c++ {
			fmt.Print(" ")
		}
		for c := 1; c <= r; c++ {
			fmt.Print("*")
		}
		for c := 1; c < r; c++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
Print triangle of numbers for a given number i.e
# N = 4

		  1
		 121
		12321
	   1234321
	   	12321
		 121
		  1
*/
func tpattern8(n int) {
	for r := 1; r <= n; r++ {
		for c := 1; c <= n-r; c++ {
			fmt.Print(" ")
		}
		for c := 1; c <= r; c++ {
			fmt.Print(c)
		}
		for c := 1; c < r; c++ {
			fmt.Print(r - c)
		}
		fmt.Println()
	}
	for r := 1; r <= n-1; r++ {
		for c := 1; c <= r; c++ {
			fmt.Print(" ")
		}
		for c := 1; c <= n-r; c++ {
			fmt.Print(c)
		}
		for c := 1; c <= n-1-r; c++ {
			fmt.Print(n - r - c)
		}
		fmt.Println()
	}
}

/*
Print triangle of numbers for a given number i.e
# N = 4

	   1     1
		2   2
		 3 3
		  4
		 3 3
		2   2
	  1      1
*/
func tpattern9(n int) {
	for r := 1; r <= n; r++ {
		for c := 1; c <= n; c++ {
			if c == r {
				fmt.Print(c)
				continue
			}
			fmt.Print(" ")
		}
		for c := n - 1; c > 0; c-- {
			if c == r {
				fmt.Print(c)
				continue
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
	for r := n - 1; r > 0; r-- {
		for c := 1; c <= n; c++ {
			if c == r {
				fmt.Print(c)
				continue
			}
			fmt.Print(" ")
		}
		for c := n - 1; c > 0; c-- {
			if c == r {
				fmt.Print(c)
				continue
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}

}

/*
Print triangle of numbers for a given number i.e
# N = 4

		  *
		 ***
		*****
	   *******
	   	*****
		 ***
		  *
*/
func tpattern10(n int) {
	for r := 1; r <= n; r++ {
		for c := 1; c <= n-r; c++ {
			fmt.Print(" ")
		}
		for c := 1; c <= r; c++ {
			fmt.Print("*")
		}
		for c := 1; c < r; c++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for r := 1; r < n; r++ {
		for c := 1; c <= r; c++ {
			fmt.Print(" ")
		}
		for c := 1; c <= n-r; c++ {
			fmt.Print("*")
		}
		for c := 1; c <= n-r-1; c++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
