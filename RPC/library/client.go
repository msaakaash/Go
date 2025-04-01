package main

import (
	"fmt"
	"net/rpc"
)

type BookRequest struct {
	Title string
}
type Response struct {
	Message string
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting")
		return
	}
	defer client.Close()
	fmt.Println("Connected to the server on port 1234")
	var res Response

	client.Call("Library.CheckAvailability", BookRequest{"go"}, &res)
	fmt.Println(res.Message)

	client.Call("Library.BorrowBook", BookRequest{"go"}, &res)
	fmt.Println(res.Message)

	client.Call("Library.BorrowBook", BookRequest{"ml"}, &res)
	fmt.Println(res.Message)

	client.Call("Library.ReturnBook", BookRequest{"go"}, &res)
	fmt.Println(res.Message)
}
