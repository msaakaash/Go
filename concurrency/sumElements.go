package main

import (
	"fmt"
	"sync"
)

func sumElements(nums []int, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, i := range nums {
		sum += i
	}
	result <- sum
}

func main() {
	nums := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := make(chan int, 2)
	numbers := nums[:]

	mid := len(nums) / 2

	var wg sync.WaitGroup

	//First part
	wg.Add(1)
	go sumElements(numbers[:mid], res, &wg)
	wg.Wait()

	//Second part
	wg.Add(1)
	go sumElements(numbers[mid:], res, &wg)
	wg.Wait()

	close(res)

	sum1 := <-res
	sum2 := <-res
	fmt.Println("Array:", numbers)
	fmt.Println(sum1)
	fmt.Println(sum2)
	fmt.Println("Total Sum:", (sum1 + sum2))
}
