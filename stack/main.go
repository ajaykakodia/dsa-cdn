package main

import "fmt"

func main() {
	fmt.Println("Parenthesis Balanced ", BalancedParenthesis("[a+b +(a+j)]"))
	fmt.Println("Parenthesis Balanced ", BalancedParenthesis("[a+b +(a+j])"))
	fmt.Println("Parenthesis Balanced ", BalancedParenthesis("[a+b +(a+j)"))

	fmt.Println("Redundant Balanced ", RedundantParenthesis("[a+b +(a+j)]"))
	fmt.Println("Redundant Balanced ", RedundantParenthesis("((a+b))"))
	fmt.Println("Redundant Balanced ", RedundantParenthesis("[a+{b +(a+j)}]"))

	stack := new(IStack)
	stack.Push(2)
	stack.Push(5)
	stack.Push(7)
	stack.Push(10)
	fmt.Println(stack.arr)
	reverseStack(stack, new(IStack))
	fmt.Println(stack.arr)

	for !stack.IsEmpty() {
		fmt.Println(stack.Pop())
	}

	//fmt.Println("Size of stack - ", stack.Size())
	//stack.Push(10)
	//
	//fmt.Println(stack.Pop())
	//
	//s := new(SStack)
	//s.Push("Ajay")
	//s.Push("Rekha")
	//s.Push("Tanmay")
	//
	//for !s.IsEmpty() {
	//	fmt.Println(s.Pop())
	//}
	//
	//fmt.Println(s.Pop())
	//s.Push("Yashika")
	//fmt.Println(s.Pop())
}
