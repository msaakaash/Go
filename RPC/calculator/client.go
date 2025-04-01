package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")

	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer client.Close()

	args := Args{A: 1, B: 9}
	var reply int

	err = client.Call("Calulator.Add", args, &reply)

	if err != nil {
		fmt.Println("Error calling remote method:", err)
		return
	}
	fmt.Println("Result:", reply)

}
