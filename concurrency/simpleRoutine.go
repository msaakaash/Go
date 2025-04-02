package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Vanakam da maple na tha Goroutine!!!")
}

func main() {
	go sayHello()
	time.Sleep(time.Second)
	fmt.Println("Main function finished")

}
