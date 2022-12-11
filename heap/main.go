package main

import "fmt"

func main() {
	priorityQueue := NewPriorityQueue()

	priorityQueue.Insert("V1", 25)
	priorityQueue.Insert("V2", 35)
	priorityQueue.Insert("V3", 50)
	fmt.Println("Top Heap Element: ", priorityQueue.Get())
	priorityQueue.Insert("V4", 60)
	priorityQueue.Insert("V5", 100)
	priorityQueue.Insert("V6", 15)
	fmt.Println("Top Heap Element: ", priorityQueue.Get())
	priorityQueue.Insert("V7", 45)
	priorityQueue.Insert("V8", 5)

	fmt.Println("Top Heap Element: ", priorityQueue.Get())
	fmt.Println("Get Heap Element: ", priorityQueue.Remove())
	fmt.Println("Get Heap Element: ", priorityQueue.Remove())
	fmt.Println("Get Heap Element: ", priorityQueue.Remove())
	fmt.Println("Get Heap Element: ", priorityQueue.Remove())
	fmt.Println("Get Heap Element: ", priorityQueue.Remove())
	fmt.Println("Get Heap Element: ", priorityQueue.Remove())
	fmt.Println("Get Heap Element: ", priorityQueue.Remove())
	fmt.Println("Get Heap Element: ", priorityQueue.Remove())
	fmt.Println("Get Heap Element: ", priorityQueue.Remove())

	arr := []int{20, 13, 44, 11, 5, 8, 32, 1, 7, 9}
	heapSort(arr)
	fmt.Println("Sorted Array: ", arr)
	arr = []int{20, 13, 44, 11, 5, 8, 32, 1, 7, 9}
	sm := kSmallestElements(arr, 4)
	fmt.Println("Smallest 4 Elements: ", sm)
}
