package main

import "fmt"

func main() {
	fmt.Println(`Here we are shifting elements in an array by D place
		i.e [3,5,2,6,7,8] and D = 2
		Expected Output: [2,6,7,8,3,5]`)

	arr := []int{3, 5, 2, 6, 7, 8, 4, 1}
	fmt.Println(arr)
	shiftDigitsInArrayB(arr, 2)
	fmt.Println(arr)
	shiftDigitsInArrayBySplit(arr, 3)
	fmt.Println(arr)
	shiftDigitsInArrayByReversing(arr, 3)
	fmt.Println(arr)
	shiftDigitsInArrayByGoWay(arr, 3)
	fmt.Println(arr)
}

// Brute force solution
func shiftDigitsInArrayB(arr []int, d int) {
	temp := 0
	l := len(arr)
	for d > 0 {
		temp = arr[0]
		for j := 0; j < l-1; j++ {
			arr[j] = arr[j+1]
		}
		arr[l-1] = temp
		d--
	}
}

// shiftDigitsInArrayBySplit - first take D elements in separate array than move remaining array to d position
// now put separated ary in last remaining places of array
func shiftDigitsInArrayBySplit(arr []int, d int) {
	l := len(arr)
	if l > d {
		arrTemp := make([]int, d)

		for i := 0; i < d; i++ {
			arrTemp[i] = arr[i]
		}

		for i := 0; i < l-d; i++ {
			arr[i] = arr[d+i]
		}

		for j, i := 0, l-d; i < l; i++ {
			arr[i] = arrTemp[j]
			j++
		}
	}
}

// shiftDigitsInArrayByReversing - First Shift complete array than reverse array in 2 parts
// 1st reverse array with len-d length and than remaining array - should provide result
func shiftDigitsInArrayByReversing(arr []int, d int) {
	l := len(arr)
	reverseArray(arr)
	reverseArray(arr[:l-d-1])
	reverseArray(arr[l-d:])
}

// shiftDigitsInArrayByGoWay -
func shiftDigitsInArrayByGoWay(arr []int, d int) {
	arr = append(arr[d:], arr[:d]...)
	fmt.Println(arr)
}

// reverse the array
func reverseArray(arr []int) {
	l := len(arr)
	mid := l / 2
	for i := 0; i <= mid; i++ {
		arr[i], arr[l-i-1] = arr[l-i-1], arr[i]
	}
}
