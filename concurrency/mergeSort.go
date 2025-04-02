package main

import (
	"fmt"
)

func mergeSort(arr []int, ch chan []int) {
	if len(arr) <= 1 {
		ch <- arr
		return
	}

	mid := len(arr) / 2
	leftChan, rightChan := make(chan []int), make(chan []int)

	go mergeSort(arr[:mid], leftChan)
	go mergeSort(arr[mid:], rightChan)

	left, right := <-leftChan, <-rightChan

	result := make([]int, 0, len(arr))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	ch <- result
}

func main() {
	arr := []int{9, 2, 4, 12, -2, -4, 0}
	ch := make(chan []int)

	go mergeSort(arr, ch)
	sortedArray := <-ch
	fmt.Println("Array:", sortedArray)
}
