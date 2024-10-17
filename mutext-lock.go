package main

import (
	"fmt"
	"sync"
)

// BankAccount represents a shared resource that multiple goroutines will access
type BankAccount struct {
	balance int
	mu      sync.Mutex // Mutex to protect access to the balance
}

// Deposit adds money to the account
func (a *BankAccount) Deposit(amount int) {
	a.mu.Lock()         // Acquire the lock to ensure exclusive access
	defer a.mu.Unlock() // Release the lock when the function exits

	a.balance += amount // Modify the shared resource
	fmt.Printf("Deposited %d, new balance: %d\n", amount, a.balance)
}

// Withdraw removes money from the account
func (a *BankAccount) Withdraw(amount int) {
	a.mu.Lock()         // Acquire the lock to ensure exclusive access
	defer a.mu.Unlock() // Release the lock when the function exits

	if a.balance >= amount {
		a.balance -= amount // Modify the shared resource
		fmt.Printf("Withdrew %d, new balance: %d\n", amount, a.balance)
	} else {
		fmt.Printf("Failed to withdraw %d, insufficient funds. Balance: %d\n", amount, a.balance)
	}
}

// GetBalance returns the current account balance (protected by mutex)
func (a *BankAccount) GetBalance() int {
	a.mu.Lock()         // Acquire the lock to ensure exclusive access
	defer a.mu.Unlock() // Release the lock when the function exits

	return a.balance
}

func main() {
	// must use pointer because pass into many function go routine
	// not using pointer it create and copy
	// efficiency (avoid copy)
	account := &BankAccount{balance: 1000} // Initial balance of 1000

	var wg sync.WaitGroup

	// Simulate multiple goroutines performing deposits and withdrawals
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			account.Deposit(amount)
		}(100 * i) // Deposit amounts: 0, 100, 200, 300, 400
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(amount int) {
			defer wg.Done()
			account.Withdraw(amount)
		}(50 * i) // Withdraw amounts: 0, 50, 100, 150, 200
	}

	wg.Wait() // Wait for all goroutines to finish

	// Print final account balance
	fmt.Printf("Final balance: %d\n", account.GetBalance())
}
