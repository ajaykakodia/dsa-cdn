package main

import "fmt"

func main() {
	stack := new(IStack)
	stack.Push(2)
	stack.Push(5)
	stack.Push(7)

	for !stack.IsEmpty() {
		fmt.Println(stack.Pop())
	}
	fmt.Println("Size of stack - ", stack.Size())
	stack.Push(10)

	fmt.Println(stack.Pop())

	s := new(SStack)
	s.Push("Ajay")
	s.Push("Rekha")
	s.Push("Tanmay")

	for !s.IsEmpty() {
		fmt.Println(s.Pop())
	}

	fmt.Println(s.Pop())
	s.Push("Yashika")
	fmt.Println(s.Pop())
}
