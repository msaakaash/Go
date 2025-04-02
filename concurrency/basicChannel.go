package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Vanakam da maple"
	}()

	msg := <-ch
	fmt.Println(msg)
}
