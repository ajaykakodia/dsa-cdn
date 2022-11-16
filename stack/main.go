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
	fmt.Println("Queue with two stack....")
	iq := queue{}
	iq.Enqueue(2)
	iq.Enqueue(5)
	iq.Enqueue(10)

	for !iq.IsEmpty() {
		fmt.Println(iq.Dequeue())
	}
	iq.Dequeue()

	minBracketReversal := minimumBracketReversal("}}{{{}{{{{")
	if minBracketReversal == -1 {
		fmt.Println("Brackets are not reversible")
	} else {
		fmt.Println(minBracketReversal, " reversal required")
	}
	fmt.Println("reversible method 2")

	minBracketReversal = minimumBracketReversal1("}}{{{}{{{{")
	if minBracketReversal == -1 {
		fmt.Println("Brackets are not reversible")
	} else {
		fmt.Println(minBracketReversal, " reversal required")
	}

	fmt.Println("Stock Span Problem... ")

	stockSpanB([]int{5, 3, 8, 7, 7, 4, 10, 12, 16})
	stockSpan([]int{5, 3, 8, 7, 7, 4, 10, 12, 16})
	stockSpan([]int{5, 3, 8, 7, 7, 20, 10, 12, 16})

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
