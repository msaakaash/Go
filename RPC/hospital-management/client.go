package main

import (
	"fmt"
	"net/rpc"
)

// Request struct
type Request struct {
	Patient  string
	Doctor   string
	TimeSlot string
}

// Response struct
type Response struct {
	Message string
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")

	if err != nil {
		fmt.Println("Error connecting to server..")
		return
	}

	defer client.Close()

	//Book Appointment
	req := Request{Patient: "Aakaash", Doctor: "Smith", TimeSlot: "10:00 AM"}
	var res Response

	err = client.Call("Hospital.BookAppointment", req, &res)
	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println(res.Message)
	}

	// View appointment
	err = client.Call("Hospital.ViewAppointment", req, &res)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(res.Message)
	}
	// Cancel appointment
	err = client.Call("Hospital.CancelAppointment", req, &res)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(res.Message)
	}
}
