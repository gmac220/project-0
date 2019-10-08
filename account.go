package main

import "fmt"

type Account struct {
	accountname   string
	balance       float64
	accountnumber int
}

// Account constructor
// need to do add account number in here
func NewAccount(name string, money float64) *Account {
	return &Account{
		accountname: name,
		balance:     money,
	}
}

// Balance returns the current balance
func (a *Account) Balance() float64 {
	return a.balance
}

// Withdraw money from the bank
func (a *Account) Withdraw(money float64) {
	if a.balance-money >= 0 {
		a.balance -= money
	} else {
		fmt.Println("Not enough money to withdraw")
	}
}

// Deposit money to balance
func (a *Account) Deposit(money float64) {
	if money > 0 {
		a.balance += money
	} else {
		fmt.Println("Please enter a positive integer")
	}
}
