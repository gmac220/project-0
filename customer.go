package main

import "fmt"

// Customer struct
type Customer struct {
	username string
	password string
	appcount int
	accounts map[string]Account
}

// NewCustomer constructor
func (c *Customer) NewCustomer(name string, pw string) *Customer {
	return &Customer{
		username: name,
		password: pw,
	}
}

// Apply would request employee to open an account
func (c *Customer) Apply() {
	//Adds data to application table on database
	c.appcount++
}

// JointApp requests employee to open an account with another customer
func (c *Customer) JointApp(othercust Customer) {
	//Adds data to application table on database
	c.appcount++
}

// ShowAccounts lists out the accounts the user currently has
func (c *Customer) ShowAccounts() {
	for _, v := range c.accounts {
		fmt.Println(v)
	}
}

// NumApplications shows the amount of applications the user has applied to
func (c *Customer) NumApplications() int {
	return c.appcount
}

// Transfer money from one account to the other
func (c *Customer) Transfer(acntname string, acntname2 string, money float64) {
	account1 := c.accounts[acntname]
	account2 := c.accounts[acntname2]
	account1.Withdraw(money)
	account2.Deposit(money)
}
