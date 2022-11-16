package main

import "fmt"

/*
In stock span problem, we have a series of n daily price quotes for a stock
we need to calculate the span of the stock’s price for all n days.
The span Si of the stock’s price on a given day i is defined as the maximum number of consecutive days just before the given day,
for which the price of the stock on the current day is less than its price on the given day.
Input: N = 7, price[] = [100 80 60 70 60 75 85]
Output: 1 1 1 2 1 4 6
*/
func stockSpanB(arr []int) {
	spArr := make([]int, len(arr))

	for i := len(arr) - 1; i > 0; i-- {
		span := 0
		for j := i - 1; j >= 0; j-- {
			if arr[j] >= arr[i] {
				break
			}
			span++
		}
		spArr[i] = span + 1
	}
	spArr[0] = 1
	fmt.Println(spArr)
}

// stockSpan => Optimized solution with stack
func stockSpan(arr []int) {
	spArr := make([]int, len(arr))
	stack := IStack{}

	for i := 0; i < len(arr); i++ {
		span := 1
		if stack.IsEmpty() || arr[stack.Top()] >= arr[i] {
			spArr[i] = span
			stack.Push(i)
			continue
		}
		for !stack.IsEmpty() && arr[stack.Top()] < arr[i] {
			popIndex := stack.Pop()
			span = i - popIndex + spArr[popIndex]
		}
		stack.Push(i)
		spArr[i] = span
	}
	fmt.Println(spArr)
}
