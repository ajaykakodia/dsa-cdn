package main

func reverseStack(s1, s2 *IStack) {
	if s1.IsEmpty() || s1.Size() == 1 {
		return
	}
	for s1.Size() != 1 {
		s2.Push(s1.Pop())
	}
	temp := s1.Pop()
	for !s2.IsEmpty() {
		s1.Push(s2.Pop())
	}
	reverseStack(s1, s2)
	s1.Push(temp)
}
