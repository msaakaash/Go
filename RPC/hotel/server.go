package main

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
	"time"
)

type Room struct {
	Available int
	Price     int
}

type Request struct {
	Roomtype int
}

type Response struct {
	Answer    string
	Timestamp string
}

type Hotel struct {
	mu    sync.Mutex
	Rooms map[int]*Room
}

func (h *Hotel) Bookroom(req Request, res *Response) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	room, exists := h.Rooms[req.Roomtype]
	if !exists {
		res.Answer = "Invalid room type"
	} else if room.Available > 0 {
		room.Available--
		res.Answer = "room booked successfully"
	} else {
		res.Answer = "room not available"
	}
	res.Timestamp = time.Now().Format(time.RFC3339)
	return nil

}

func main() {
	hotel := &Hotel{
		Rooms: map[int]*Room{
			0: {Available: 10, Price: 1000},
			1: {Available: 20, Price: 1500},
			2: {Available: 5, Price: 2000},
			3: {Available: 3, Price: 3000},
			4: {Available: 2, Price: 5000},
		},
	}
	err := rpc.Register(hotel)
	if err != nil {
		fmt.Println("error occured")
		return
	}
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("error occured")
		return
	}
	defer listener.Close()
	fmt.Println("server listening")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error occured")
			return
		}
		go rpc.ServeConn(conn)
	}

}
