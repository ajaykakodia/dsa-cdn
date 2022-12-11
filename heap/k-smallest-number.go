package main

import "github.com/ajaykakodia/lib/heap"

func kSmallestElements(arr []int, k int) []int {
	if len(arr) == 0 {
		return nil
	}
	if len(arr) <= k {
		return arr
	}
	kArr := arr[:k]
	heap.HeapifyMax(kArr)
	for i := k; i < len(arr); i++ {
		if kArr[0] > arr[i] {
			heap.HeapMaxReplace(kArr, arr[i])
		}
	}
	return kArr
}
