package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// BankAccountV2 represents a shared resource with balance and a mutex
type BankAccountV2 struct {
	balance int
	mu      sync.Mutex // Mutex to protect balance modifications
}

// Deposit adds money to the account (thread-safe)
func (a *BankAccountV2) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
	fmt.Printf("Deposited %d, new balance: %d\n", amount, a.balance)
}

// Withdraw removes money from the account (thread-safe)
func (a *BankAccountV2) Withdraw(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.balance >= amount {
		a.balance -= amount
		fmt.Printf("Withdrew %d, new balance: %d\n", amount, a.balance)
	} else {
		fmt.Printf("Failed to withdraw %d, insufficient funds. Current balance: %d\n", amount, a.balance)
	}
}

// GetBalance returns the current balance (thread-safe)
func (a *BankAccountV2) GetBalance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	// Create 10 bank accounts with initial balances
	accounts := make([]*BankAccountV2, 10)
	for i := 0; i < 10; i++ {
		accounts[i] = &BankAccountV2{balance: rand.Intn(1000) + 500} // Random initial balance between 500 and 1500
	}

	for i, v := range accounts {
		fmt.Printf("account[%d] balance: %d\n", i, v.balance)
	}
	
	var wg sync.WaitGroup

	// Simulate 10 goroutines for deposits and withdrawals on each account
	for i, account := range accounts {
		wg.Add(2) // One goroutine for deposit, one for withdrawal

		// Goroutine for deposits
		go func(acc *BankAccountV2, id int) {
			defer wg.Done()
			amount := rand.Intn(500) + 1 // Random deposit amount between 1 and 500
			fmt.Printf("Account %d: Attempting to deposit %d\n", id, amount)
			acc.Deposit(amount)
		}(account, i+1)

		// Goroutine for withdrawals
		go func(acc *BankAccountV2, id int) {
			defer wg.Done()
			amount := rand.Intn(500) + 1 // Random withdrawal amount between 1 and 500
			fmt.Printf("Account %d: Attempting to withdraw %d\n", id, amount)
			acc.Withdraw(amount)
		}(account, i+1)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the final balances of all accounts
	fmt.Println("\nFinal balances of all accounts:")
	for i, account := range accounts {
		fmt.Printf("Account %d: Final balance: %d\n", i+1, account.GetBalance())
	}
}
