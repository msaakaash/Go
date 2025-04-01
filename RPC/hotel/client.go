package main

import (
	"fmt"
	"net/rpc"
)

type Request struct {
	Roomtype int
}

type Response struct {
	Answer    string
	Timestamp string
}

func main() {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		fmt.Println("Error occurred while connecting to server:", err)
		return
	}
	defer client.Close()
	req := Request{Roomtype: 1}
	var res Response

	err = client.Call("Hotel.Bookroom", req, &res)
	if err != nil {
		fmt.Println("Error occured.")
		return
	} else {
		fmt.Println("Response:", res.Answer, "Timestamp:", res.Timestamp)
	}
}
