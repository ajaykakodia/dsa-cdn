package main

func Fibonacci(num int) int {
	arr := make([]int, num+1)
	return fibb(num, arr)
}

func fibb(num int, arr []int) int {
	if num == 0 || num == 1 {
		return num
	}
	if arr[num-1] == 0 {
		arr[num-1] = fibb(num-1, arr)
	}
	if arr[num-2] == 0 {
		arr[num-2] = fibb(num-2, arr)
	}

	return arr[num-1] + arr[num-2]
}
