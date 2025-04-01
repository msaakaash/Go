package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"sort"
	"sync"
)

type Patient struct {
	ID   int
	Name string
	Age  int
}

var (
	patients = map[int]Patient{
		1: {ID: 1, Name: "Aakaash", Age: 20},
		2: {ID: 2, Name: "Aiyyappan", Age: 21},
	}
	mutex sync.Mutex
)

type Health struct{}

// Get a patient by ID
func (h *Health) GetPatient(id int, reply *Patient) error {
	mutex.Lock()
	defer mutex.Unlock()

	if p, found := patients[id]; found {
		*reply = p
		return nil
	}
	return errors.New("patient not found")
}

// Add a new patient
func (h *Health) AddPatient(p Patient, reply *string) error {
	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := patients[p.ID]; exists {
		return errors.New("patient ID already exists")
	}

	patients[p.ID] = p
	*reply = "Patient added successfully"
	return nil
}

// Remove a patient

func (h *Health) RemovePatient(id int, reply *string) error {
	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := patients[id]; !exists {
		return errors.New("Patient not found")
	}

	delete(patients, id)

	*reply = "Patient removed successfully"
	return nil
}

// sorting patients by ID

func (h *Health) Sort(_ string, reply *[]Patient) error {
	mutex.Lock()
	defer mutex.Unlock()

	var sorted []Patient

	for _, p := range patients {
		sorted = append(sorted, p)
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].ID < sorted[j].ID
	})

	*reply = sorted
	return nil
}

func main() {
	rpc.Register(&Health{})
	listener, _ := net.Listen("tcp", ":1234")
	defer listener.Close()

	fmt.Println("Server Running on port 1234")

	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}
}
