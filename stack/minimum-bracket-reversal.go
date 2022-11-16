package main

// minimumBracketReversal => One way to solve
func minimumBracketReversal(str string) int {
	count := 0
	if len(str)%2 != 0 {
		return -1
	}
	stack := SStack{}

	for _, v := range str {
		c := string(v)
		if c == "{" {
			stack.Push(c)
			continue
		}
		if stack.IsEmpty() || stack.Top() == "}" {
			stack.Push(c)
			continue
		}
		stack.Pop()
	}
	for !stack.IsEmpty() {
		c1 := stack.Pop()
		c2 := stack.Pop()
		if c1 == c2 {
			count++
			continue
		}
		count = count + 2
	}
	return count
}

// minimumBracketReversal1 => 2nd Method '}}{{{}{{{{'
func minimumBracketReversal1(str string) int {
	count := 0
	if len(str)%2 != 0 {
		return -1
	}
	stack := SStack{}

	for _, v := range str {
		c := string(v)
		if c == "{" {
			stack.Push("{")
			continue
		}
		if !stack.IsEmpty() {
			stack.Pop()
			continue
		}
		stack.Push("{")
		count++
	}
	count = count + stack.Size()/2
	return count
}
