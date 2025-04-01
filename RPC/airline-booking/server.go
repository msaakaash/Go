package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Request struct {
	Class string
	Count int
}

type Response struct {
	Message string
}

type Airline struct {
	Ticket map[string]int
	Price  map[string]int
}

func (a *Airline) Bookticket(req Request, res *Response) error {
	count, exists := a.Ticket[req.Class]
	if !exists {
		fmt.Println("Invalid ticket class")
		fmt.Println(count)
	} else if a.Ticket[req.Class] >= req.Count {
		a.Ticket[req.Class] -= req.Count
		res.Message = fmt.Sprintf("%d Tickets booked successfully \n Total cost of tickets is %d", req.Count, (a.Price[req.Class])*req.Count)
	} else {
		res.Message = "Ticket sold out"
	}
	return nil
}

func main() {
	airline := &Airline{
		Ticket: map[string]int{
			"E": 50,
			"B": 30,
			"F": 10,
		},
		Price: map[string]int{
			"E": 5000,
			"B": 30000,
			"F": 100000,
		},
	}

	err := rpc.Register(airline)
	if err != nil {
		fmt.Println("Error occurred during RPC registration")
		return
	}

	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error occurred during TCP listening")
		return
	}
	defer listener.Close()
	fmt.Println("Server running on port 1234")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error occurred during connection acceptance")
			continue
		}
		go rpc.ServeConn(conn)
	}
}
