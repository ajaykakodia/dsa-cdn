package main

func heapSort(arr []int) {
	size := len(arr)
	if size < 2 {
		return
	}

	//create priority queue in existing array
	//for i, _ := range arr {
	//	hippyfyArray(arr, i)
	//}
	//fmt.Println(arr)

	// create priority queue in existing array using non-leaf logic
	createPriorityQueueInArray(arr)

	// remove array from Heap and storing in last in same array, It gives us sorted array
	for i := 0; i < size; i++ {
		poppedData := arr[0]
		arr[0] = arr[size-1-i]
		arr[size-i-1] = poppedData
		downHippyfyArray(arr, size-i-1)
	}
}

// using concept of using downhippfy for non-leaf elements, and It works and create P queue
func createPriorityQueueInArray(arr []int) {
	size := len(arr)
	nonLeafNodes := size / 2
	for i := nonLeafNodes - 1; i >= 0; i-- {
		downHippyfyArrayFromInBetween(arr, i)
	}
}

func downHippyfyArrayFromInBetween(arr []int, startIndex int) {
	size := len(arr)
	if size == 0 {
		return
	}
	pIndex, minIndex, lcIndex, rcIndex := startIndex, startIndex, startIndex*2+1, startIndex*2+2
	for lcIndex < size {
		minIndex = pIndex
		if arr[minIndex] > arr[lcIndex] {
			minIndex = lcIndex
		}
		if rcIndex < size && arr[minIndex] > arr[rcIndex] {
			minIndex = rcIndex
		}
		if minIndex == pIndex {
			return
		}
		arr[pIndex], arr[minIndex] = arr[minIndex], arr[pIndex]
		pIndex = minIndex
		lcIndex = pIndex*2 + 1
		rcIndex = pIndex*2 + 2
	}
}

func downHippyfyArray(arr []int, lastIndex int) {
	size := len(arr)
	if size == 0 {
		return
	}
	pIndex, minIndex, lcIndex, rcIndex := 0, 0, 1, 2
	for lcIndex < lastIndex {
		minIndex = pIndex
		if arr[minIndex] > arr[lcIndex] {
			minIndex = lcIndex
		}
		if rcIndex < lastIndex && arr[minIndex] > arr[rcIndex] {
			minIndex = rcIndex
		}
		if minIndex == pIndex {
			return
		}
		arr[pIndex], arr[minIndex] = arr[minIndex], arr[pIndex]
		pIndex = minIndex
		lcIndex = pIndex*2 + 1
		rcIndex = pIndex*2 + 2
	}
}

func hippyfyArray(arr []int, cIndex int) {
	pIndex := 0
	for cIndex > 0 {
		pIndex = (cIndex - 1) / 2
		if arr[pIndex] > arr[cIndex] {
			arr[pIndex], arr[cIndex] = arr[cIndex], arr[pIndex]
			cIndex = pIndex
			continue
		}
		return
	}
}
