package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	Balance int
	mu      sync.Mutex
}

func (ba *BankAccount) Deposit(amount int) {
	ba.mu.Lock()
	defer ba.mu.Unlock()
	if amount != 0 {
		ba.Balance += amount
	} else {
		return
	}
}

func (ba *BankAccount) WithDraw(amount int) bool {
	ba.mu.Lock()
	defer ba.mu.Unlock()
	if ba.Balance > amount && amount > 0 {
		ba.Balance -= amount
		return true
	}
	return false
}

func main() {
	bankAccount := BankAccount{
		Balance: 3000,
	}

	var wg sync.WaitGroup

	wg.Add(7)

	go func() {
		wg.Done()
		bankAccount.Deposit(3000)
	}()

	go func() {
		wg.Done()
		bankAccount.WithDraw(2000)
	}()

	go func() {
		wg.Done()
		bankAccount.WithDraw(3000)
	}()

	go func() {
		wg.Done()
		bankAccount.Deposit(1000)
	}()

	go func() {
		wg.Done()
		bankAccount.WithDraw(2000)
	}()

	go func() {
		wg.Done()
		bankAccount.WithDraw(1000)
	}()

	go func() {
		wg.Done()
		bankAccount.Deposit(500)
	}()

	fmt.Println("Main goroutine: Awaiting all transactions!...")
	wg.Wait()
	fmt.Printf("Main goroutine: Finished all transactions with balance %v. Exited!\n", bankAccount.Balance)
}
