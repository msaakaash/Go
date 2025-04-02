package main

import "fmt"

func merge(arr1 []int, arr2 []int, ch chan []int) {
	i, j := 0, 0
	merged := make([]int, 0, len(arr1)+len(arr2))
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merged = append(merged, arr1[i])
			i++
		} else {
			merged = append(merged, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		merged = append(merged, arr1[i])
		i++
	}
	for j < len(arr2) {
		merged = append(merged, arr2[j])
		j++
	}

	ch <- merged
}

func main() {
	arr1 := []int{1, 3, 5, 7}
	arr2 := []int{2, 4, 6, 8}

	ch := make(chan []int)

	go merge(arr1, arr2, ch)
	mergedArray := <-ch // Receive merged array
	fmt.Println("Merged Sorted Array:", mergedArray)
}
