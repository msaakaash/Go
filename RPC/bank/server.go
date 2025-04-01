package main

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

type Account struct {
	Name    string
	Balance int
}

type Transaction struct {
	AccountName string
	Amount      int
}

type Transfer struct {
	From   string
	To     string
	Amount int
}

type Bank struct {
	mu       sync.Mutex
	Accounts map[string]*Account
}

func (b *Bank) CreateAccount(name string, res *string) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, exists := b.Accounts[name]; exists {
		*res = "Account already exists"
		return nil
	}

	b.Accounts[name] = &Account{Name: name, Balance: 0}
	*res = "Account created successfully"
	return nil

}

func (b *Bank) Deposit(txt Transaction, res *string) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	if acc, exists := b.Accounts[txt.AccountName]; exists {
		acc.Balance += txt.Amount
		*res = fmt.Sprintf("Deposited %d to %s. New balance: %d", txt.Amount, acc.Name, acc.Balance)
		return nil
	} else {
		*res = "Account not found"
	}
	return nil
}

func (b *Bank) Withdraw(txt Transaction, res *string) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if acc, exists := b.Accounts[txt.AccountName]; exists {
		if acc.Balance >= txt.Amount {
			acc.Balance -= txt.Amount
			*res = fmt.Sprintf("withdrawn %d successully.Remaining Balance:%d", txt.Amount, acc.Balance)
		} else {
			*res = "Insufficient funds"
		}
	} else {
		*res = "Account not found"
	}
	return nil
}

func (b *Bank) CheckBalance(name string, res *int) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if acc, exists := b.Accounts[name]; exists {
		*res = acc.Balance
	} else {
		*res = -1
	}
	return nil
}

func (b *Bank) TransferMoney(trans Transfer, res *string) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	fromAcc, fromExists := b.Accounts[trans.From]
	toAcc, toExists := b.Accounts[trans.To]

	if !fromExists || !toExists {
		*res = "One or both accounts not found"
		return nil
	}

	if fromAcc.Balance < trans.Amount {
		*res = "Insufficient funds"
		return nil
	}

	fromAcc.Balance -= trans.Amount
	toAcc.Balance += trans.Amount
	*res = fmt.Sprintf("Transferred %d from %s to %s successfully", trans.Amount, fromAcc.Name, toAcc.Name)
	return nil
}

func (b *Bank) CloseAccount(name string, res *string) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, exists := b.Accounts[name]; exists {
		delete(b.Accounts, name)
		*res = "Account closed successfully"
	} else {
		*res = "Account not found"
	}
	return nil
}

func main() {
	bank := &Bank{
		Accounts: make(map[string]*Account),
	}
	rpc.Register(bank)
	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	defer listener.Close()

	fmt.Println("Banking System running on port 1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
