package main

import "strings"

var openingParenthesis = "[{("
var closingParenthesis = ")}]"

func isOpeningParenthesis(char string) bool {
	return strings.Index(openingParenthesis, string(char)) > -1
}

func isClosingParenthesis(char string) bool {
	return strings.Index(closingParenthesis, string(char)) > -1
}

func getOpeningParenthesis(char string) string {
	if char == ")" {
		return "("
	}
	if char == "}" {
		return "{"
	}
	if char == "]" {
		return "["
	}
	return ""
}

func BalancedParenthesis(str string) bool {
	stack := new(SStack)

	for _, char := range str {
		c := string(char)
		if isOpeningParenthesis(c) {
			stack.Push(c)
			continue
		}
		if isClosingParenthesis(c) {
			if stack.Top() != getOpeningParenthesis(c) {
				return false
			}
			stack.Pop()
			continue
		}
	}

	if !stack.IsEmpty() {
		return false
	}

	return true
}
