package main

import (
	"fmt"
	"net/rpc"
)

type Transaction struct {
	AccountName string
	Amount      int
}

type Transfer struct {
	From   string
	To     string
	Amount int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	defer client.Close()
	var response string
	var balance int

	// Create Accounts
	client.Call("Bank.CreateAccount", "Alice", &response)
	fmt.Println(response)
	client.Call("Bank.CreateAccount", "Bob", &response)
	fmt.Println(response)
	// Deposit Money
	client.Call("Bank.Deposit", Transaction{"Alice", 5000}, &response)
	fmt.Println(response)
	client.Call("Bank.Deposit", Transaction{"Bob", 2000}, &response)
	fmt.Println(response)

	// Check Balance
	client.Call("Bank.CheckBalance", "Alice", &balance)
	fmt.Println("Alice's Balance:", balance)
	client.Call("Bank.CheckBalance", "Bob", &balance)
	fmt.Println("Bob's Balance:", balance)

	// Withdraw Money
	client.Call("Bank.Withdraw", Transaction{"Alice", 3000}, &response)
	fmt.Println(response)
	client.Call("Bank.Withdraw", Transaction{"Bob", 5000}, &response)
	fmt.Println(response) // Should show "Insufficient funds"

	// Transfer Money
	client.Call("Bank.TransferMoney", Transfer{"Alice", "Bob", 1000}, &response)
	fmt.Println(response)

	// Check Balance After Transfer
	client.Call("Bank.CheckBalance", "Alice", &balance)
	fmt.Println("Alice's Balance After Transfer:", balance)
	client.Call("Bank.CheckBalance", "Bob", &balance)
	fmt.Println("Bob's Balance After Transfer:", balance)

	// Close an Account
	client.Call("Bank.CloseAccount", "Bob", &response)
	fmt.Println(response)

	// Check Balance of Closed Account
	client.Call("Bank.CheckBalance", "Bob", &balance)
	fmt.Println("Bob's Balance After Closing Account:", balance)
}
