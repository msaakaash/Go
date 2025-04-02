package main

import (
	"fmt"
	"net/rpc"
)

type KeyValue struct {
	Key   string
	Value string
}

type Response struct {
	Value string
	Found bool
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer client.Close()

	// PUT operation
	putReq := KeyValue{"username", "aakaash123"}
	var putRes string
	err = client.Call("KeyValueStore.Put", putReq, &putRes)
	if err != nil {
		fmt.Println("Put error:", err)
	} else {
		fmt.Println("Put Response:", putRes)
	}

	putRq := KeyValue{"username2", "aiyyappan123"}
	var putRs string
	err = client.Call("KeyValueStore.Put", putRq, &putRs)
	if err != nil {
		fmt.Println("Put error:", err)
	} else {
		fmt.Println("Put Response:", putRs)
	}

	// GET operation
	var getRes Response
	err = client.Call("KeyValueStore.Get", "username", &getRes)
	if err != nil {
		fmt.Println("Get error:", err)
	} else if getRes.Found {
		fmt.Println("Get Response: Key found with value:", getRes.Value)
	} else {
		fmt.Println("Get Response: Key not found")
	}

	// DELETE operation
	var delRes string
	err = client.Call("KeyValueStore.Delete", "username", &delRes)
	if err != nil {
		fmt.Println("Delete error:", err)
	} else {
		fmt.Println("Delete Response:", delRes)
	}

	// GET operation again to verify deletion
	err = client.Call("KeyValueStore.Get", "username", &getRes)
	if err != nil {
		fmt.Println("Get error:", err)
	} else if getRes.Found {
		fmt.Println("Get Response After Delete: Key found with value:", getRes.Value)
	} else {
		fmt.Println("Get Response After Delete: Key not found")
	}
}
