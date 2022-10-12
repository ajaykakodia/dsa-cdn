package main

import "fmt"

func main() {
	fmt.Println("In replace character in string")
	fmt.Println(replaceChar("afdhrdsedhdd", 'd', 'z'))
	fmt.Println("Replace pi :")
	fmt.Println(replacePi("love you pi"))
}

func replaceChar(str string, a, b rune) string {
	if str == "" {
		return str
	}
	resStr := replaceChar(str[1:], a, b)

	if rune(str[0]) == a {
		return string(b) + resStr
	}

	return string(str[0]) + resStr
}

func replacePi(str string) string {
	if len(str) == 0 || len(str) == 1 {
		return str
	}

	var sol string

	if str[0] == 'p' && str[1] == 'i' {
		sol = replacePi(str[2:])
		return "3.14" + sol
	}

	sol = replacePi(str[1:])
	return string(str[0]) + sol
}
