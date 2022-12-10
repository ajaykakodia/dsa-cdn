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
}