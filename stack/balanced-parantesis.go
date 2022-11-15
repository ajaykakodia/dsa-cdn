package main

import "strings"

func BalancedParathesis(str string) bool {
	openParanthesis := "[{("
	stack := new(SStack)

	for _, char := range str {
		if strings.Index(openParanthesis, string(char)) > -1 {
			stack.Push(string(char))
			continue
		}
		if string(char) == ")" {
			if stack.Top() != "(" {
				return false
			}
			stack.Pop()
			continue
		}
		if string(char) == "}" {
			if stack.Top() != "{" {
				return false
			}
			stack.Pop()
			continue
		}
		if string(char) == "]" {
			if stack.Top() != "[" {
				return false
			}
			stack.Pop()
		}
	}

	if !stack.IsEmpty() {
		return false
	}

	return true
}
