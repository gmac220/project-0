package main

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// Customer struct
type Customer struct {
	username  string
	password  string
	firstname string
	lastname  string
	appcount  int
	accounts  map[string]Account
}

// NewCustomer constructor
func (c *Customer) NewCustomer(name string, pw string) *Customer {

	return &Customer{
		username: name,
		password: pw,
	}
}

// CustomerPage takes user to what the customer would like to do
func CustomerPage(username string, firstname string, lastname string) {
	var num int
	fmt.Println("1: Apply")
	fmt.Println("2: Apply to a Joint Account")
	fmt.Println("3: Show Accounts")
	fmt.Println("4: Show Balance")
	fmt.Println("5: Withdraw")
	fmt.Println("6: Deposit")
	fmt.Println("7: Transfer funds")
	fmt.Println("8: Show number of pending applications")
	fmt.Println("9: Exit")
	fmt.Printf("Please type in a number: ")
	fmt.Scanln(&num)

	switch num {
	case 1:
		Apply()
	case 2:
		JointApp()
	case 3:
		ShowAccounts()
	case 4:
		ShowBalance()
	case 5:
		Withdraw()
	case 6:
		Deposit()
	case 7:
		Transfer()
	case 8:
		ShowPendingApps()
	case 9:
		os.Exit(0)
	}
}

// Apply adds to applications table
func Apply() {
	//Adds data to application table on database

}

// JointApp adds to applications table
func JointApp() {
	//Adds data to application table on database

}

// ShowAccounts lists out the accounts the user currently has
func ShowAccounts() {

}

func ShowBalance() {

}

func Withdraw() {

}

func Deposit() {

}

// ShowPendingApps shows the amount of applications the user has applied to
func ShowPendingApps() {

}

// Transfer money from one account to the other
func Transfer() {
	// account1 := c.accounts[acntname]
	// account2 := c.accounts[acntname2]
	// account1.Withdraw(money)
	// account2.Deposit(money)
}
