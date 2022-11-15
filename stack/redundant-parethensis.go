package main

func RedundantParenthesis(str string) bool {

	stack := new(SStack)

	for _, val := range str {
		c := string(val)
		if !isClosingParenthesis(c) {
			stack.Push(c)
			continue
		}

		if isClosingParenthesis(c) {
			otherChar := 0
			cp := getOpeningParenthesis(c)
			for !stack.IsEmpty() && stack.Top() != cp {
				stack.Pop()
				otherChar++
			}
			if stack.IsEmpty() || otherChar == 0 {
				return true
			}
			stack.Pop()
		}
	}

	if !stack.IsEmpty() {
		return true
	}

	return false
}
