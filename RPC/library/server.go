package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// defining request
type Request struct {
	Title string
}

// defining response
type Response struct {
	Message string
}

// defining Library system
type Library struct {
	Books map[string]int
}

//checking availability of the book

func (l *Library) CheckAvailability(req Request, res *Response) error {
	if count, exists := l.Books[req.Title]; exists && count > 0 {
		res.Message = req.Title + " book is available"
	} else {
		res.Message = "Book not available"
	}
	return nil
}

//borrowing book

func (l *Library) BorrowBook(req Request, res *Response) error {
	if count, exists := l.Books[req.Title]; exists && count > 0 {
		l.Books[req.Title]--
		res.Message = "Borrowed " + req.Title + " " + "book successfully"
	} else {
		res.Message = "Book not available"
	}
	return nil
}

// returning book
func (l *Library) ReturnBook(req Request, res *Response) error {
	if count, exists := l.Books[req.Title]; exists && count > 0 {
		l.Books[req.Title]++
		res.Message = "Returned " + req.Title + " " + "book successfully"
	} else {
		res.Message = "Book not available"
	}
	return nil
}

func main() {
	lib := &Library{
		Books: map[string]int{
			"go": 10,
			"ml": 15,
			"dl": 0,
		}}

	err := rpc.Register(lib)
	if err != nil {
		fmt.Println("error occured")
		return
	}

	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		fmt.Println("erro occured")
	}

	defer listener.Close()
	fmt.Println("Server running on the port 1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error connecting")
		}

		go rpc.ServeConn(conn)
	}

}
