package main

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

type Appointment struct {
	Patient  string
	Doctor   string
	TimeSlot string
}

type Request struct {
	Patient  string
	Doctor   string
	TimeSlot string
}

type Response struct {
	Message string
}

type Hospital struct {
	mu           sync.Mutex
	Appointments map[string]Appointment
}

// Book Appointment
func (h *Hospital) BookAppointment(req Request, res *Response) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, exists := h.Appointments[req.Patient]; exists {
		res.Message = "You already have an appointment"
		return nil
	}

	h.Appointments[req.Patient] = Appointment{
		Patient:  req.Patient,
		Doctor:   req.Doctor,
		TimeSlot: req.TimeSlot,
	}

	res.Message = fmt.Sprintf("Appointment booked for %s with Dr. %s at %s", req.Patient, req.Doctor, req.TimeSlot)
	return nil
}

// View appointment
func (h *Hospital) ViewAppointment(req Request, res *Response) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	appointment, exists := h.Appointments[req.Patient]
	if !exists {
		res.Message = "No appointmet found"
		return nil
	}

	res.Message = fmt.Sprintf("Appointment booked for %s with Dr. %s at %s", appointment.Patient, appointment.Doctor, appointment.TimeSlot)
	return nil
}

//Delete appointment

func (h *Hospital) CancelAppointment(req Request, res *Response) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	_, exists := h.Appointments[req.Patient]

	if !exists {
		res.Message = "No appointment found"
		return nil
	}

	delete(h.Appointments, req.Patient)
	res.Message = "Appointment successfully cancelled"
	return nil
}

func main() {

	hos := &Hospital{
		Appointments: make(map[string]Appointment),
	}

	err := rpc.Register(hos)

	if err != nil {
		fmt.Println("Error registering RPC server")
		return
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server..")
	}

	defer listener.Close()

	fmt.Println("Server running on port 1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error")
			continue
		}
		go rpc.ServeConn(conn)
	}
}
