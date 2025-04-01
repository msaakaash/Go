package main

import (
	"fmt"
	"net/rpc"
)

type Patient struct {
	ID   int
	Name string
	Age  int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting")
		return
	}
	// Add a new patient
	newPatient := Patient{ID: 0, Name: "Anna Liza", Age: 20}
	var addReply string
	client.Call("Health.AddPatient", newPatient, &addReply)
	fmt.Println(addReply)

	// Remove a patient
	var removeReply string
	client.Call("Health.RemovePatient", 1, &removeReply)
	fmt.Println(removeReply)

	// Get patient details
	var patient Patient
	client.Call("Health.GetPatient", 2, &patient)
	fmt.Println("Patient Details:", patient)

	// Sort and display patients
	var sortedPatients []Patient
	client.Call("Health.Sort", "", &sortedPatients)
	fmt.Println("Sorted Patients:")
	for _, p := range sortedPatients {
		fmt.Println(p)
	}

}
