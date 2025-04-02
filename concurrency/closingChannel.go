package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 100

	close(ch)

	value, ok := <-ch
	fmt.Println(value, ok)

	value, ok = <-ch
	fmt.Println(value, ok)
}
