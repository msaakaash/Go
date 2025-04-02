package main

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

type KeyValue struct {
	Key   string
	Value string
}

type Response struct {
	Value string
	Found bool
}

type KeyValueStore struct {
	store map[string]string
	mu    sync.Mutex
}

//put method

func (k *KeyValueStore) Put(req KeyValue, res *string) error {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.store[req.Key] = req.Value
	*res = "Stored key successfully"
	for key, value := range k.store {
		fmt.Println(key, ":", value)
	}
	return nil
}

// get method
func (k *KeyValueStore) Get(key string, res *Response) error {
	k.mu.Lock()
	defer k.mu.Unlock()
	value, found := k.store[key]
	res.Value, res.Found = value, found
	return nil
}

func (k *KeyValueStore) Delete(key string, res *string) error {
	k.mu.Lock()
	defer k.mu.Unlock()
	_, found := k.store[key]
	if found {
		delete(k.store, key)
		for key, value := range k.store {
			fmt.Println(key, ":", value)
		}
		*res = "Key deleted successfully"
	} else {
		*res = "Key not found"
	}
	return nil
}

func main() {
	kv := &KeyValueStore{store: make(map[string]string)}
	err := rpc.Register(kv)
	if err != nil {
		fmt.Println("Error")
		return
	}
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error")
	}
	defer listener.Close()

	fmt.Println("Server running on port 1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error")
		}
		go rpc.ServeConn(conn)
	}
}
