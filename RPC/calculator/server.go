package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Calulator struct{}

func (c *Calulator) Add(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

type Args struct {
	A, B int
}

func main() {
	cal := new(Calulator)
	rpc.Register(cal)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 1234...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}

}
