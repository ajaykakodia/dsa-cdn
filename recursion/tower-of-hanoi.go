package main

import "fmt"

func main() {
	fmt.Println("Tower of Hanoi Problem : ")
	towerOfHanoi(4, "src", "helper", "dest")
}

func towerOfHanoi(n int, a, b, c string) {
	if n == 1 {
		fmt.Println("Moving 1st disk from ", a, " to ", c)
		return
	}
	towerOfHanoi(n-1, a, c, b)
	fmt.Printf("Moving %dth disk from %s to %s\n", n, a, c)
	towerOfHanoi(n-1, b, a, c)
}
