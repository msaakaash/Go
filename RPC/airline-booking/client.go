package main

import (
	"fmt"
	"net/rpc"
)

type Request struct {
	Class string
	Count int
}

type Response struct {
	Message string
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")

	if err != nil {
		fmt.Println("Error")
	}

	defer client.Close()

	var res Response
	req := Request{Class: "E", Count: 40}

	err = client.Call("Airline.Bookticket", req, &res)
	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println(res.Message)
	}
}
