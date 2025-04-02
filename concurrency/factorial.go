package main

import "fmt"

func factorial(n int, res chan int) {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}

	res <- result
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	channel := make(chan int, len(arr))

	for _, i := range arr {
		go factorial(i, channel)
	}

	for range arr {
		fmt.Println(<-channel)
	}
}
