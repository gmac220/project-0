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

var cusername string
var cfirstname string
var clastname string

// NewCustomer constructor
func (c *Customer) NewCustomer(name string, pw string) *Customer {

	return &Customer{
		username: name,
		password: pw,
	}
}

// CustomerInit initializes global variables
func CustomerInit(username string, firstname string, lastname string) {
	cusername = username
	cfirstname = firstname
	clastname = lastname
	CustomerPage()
}

// CustomerPage takes user to what the customer would like to do
func CustomerPage() {
	var num int
	fmt.Println("What would you like to do today?")
	fmt.Println("1: Apply")
	fmt.Println("2: Apply to a Joint Account")
	fmt.Println("3: Show Accounts")
	fmt.Println("4: Show Balance")
	fmt.Println("5: Withdraw")
	fmt.Println("6: Deposit")
	fmt.Println("7: Transfer funds")
	fmt.Println("8: Show number of pending applications")
	fmt.Println("9: Logout")
	fmt.Println("10: Exit")
	fmt.Printf("Please type in a number: ")
	fmt.Scanln(&num)

	switch num {
	case 1:
		Apply(cusername, cfirstname, clastname, false)
	case 2:
		JointApp(cusername, cfirstname, clastname, true)
	case 3:
		ShowAccounts(cusername)
	case 4:
		ShowBalance(cusername)
	case 5:
		Withdraw(cusername)
	case 6:
		Deposit(cusername)
	case 7:
		Transfer(cusername)
	case 8:
		ShowPendingApps(cusername)
	case 9:
		main()
	case 10:
		os.Exit(0)
	}
}

// Apply adds to applications table
func Apply(username string, firstname string, lastname string, joint bool) {
	//Adds data to application table on database
	var acntname string
	var appcount int

	db := OpenDB()
	defer db.Close()
	fmt.Printf("What type of account do you want to open? checking or savings: ")
	fmt.Scanln(&acntname)
	// db.Exec("INSERT INTO applications (username, firstname, lastname, acntname, joint)"+
	// "VALUES ($1, $2, $3, $4, $5)", username, firstname, lastname, acntname, joint)
	db.Exec("INSERT INTO applications"+
		"(username, firstname, lastname, acntname, joint, username2, firstname2, lastname2)"+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		username, firstname, lastname, acntname, joint, "N/A", "N/A", "N/A")
	row := db.QueryRow("SELECT appcount FROM customers WHERE username = $1", username)
	row.Scan(&appcount)
	db.Exec("UPDATE customers SET appcount = $1 WHERE username = $2", appcount+1, username)
	CustomerPage()
}

// JointApp adds to applications table
func JointApp(username string, firstname string, lastname string, joint bool) {
	var uname2 string
	var fname2 string
	var lname2 string
	var acntname string

	db := OpenDB()
	defer db.Close()
	fmt.Printf("What type of account do you want to open? checking or savings: ")
	fmt.Scanln(&acntname)
	fmt.Printf("What user do you want to share an account with? Please input username of that user: ")
	fmt.Scanln(&uname2)
	uname2, fname2, lname2 = CheckCustomer(uname2)
	for uname2 == "" {
		fmt.Printf("This user does not exist please input a valid user: ")
		fmt.Scanln(&uname2)
		uname2, fname2, lname2 = CheckCustomer(uname2)
	}

	db.Exec("INSERT INTO applications"+
		"(username, firstname, lastname, acntname, joint, username2, firstname2, lastname2)"+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		username, firstname, lastname, acntname, joint, uname2, fname2, lname2)
	CustomerPage()
}

// CheckCustomer verifies if the customer exists
func CheckCustomer(username string) (string, string, string) {
	var uname2 string
	var fname2 string
	var lname2 string
	var appcnt int
	var pw string

	db := OpenDB()
	defer db.Close()
	row := db.QueryRow("SELECT * FROM customers WHERE username = $1", username)
	row.Scan(&uname2, &pw, &fname2, &lname2, &appcnt)

	return uname2, fname2, lname2
}

// ShowAccounts lists out the accounts the user currently has
func ShowAccounts(username string) {
	db := OpenDB()
	defer db.Close()
	//loop through both accounts and joint accounts
	rows, _ := db.Query("SELECT * FROM accounts WHERE username = ", username)
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
	CustomerPage()
}

// ShowBalance of account
func ShowBalance(username string) {
	var acntnum int
	var balance float64

	fmt.Println("What account number do you want to check the balance for?: ")
	fmt.Scanln(&acntnum)
	db := OpenDB()
	defer db.Close()
	row := db.QueryRow("SELECT balance FROM accounts WHERE acntnumber = $1", acntnum)
	row.Scan(&balance)
	fmt.Println("This is your balance: ", balance)
}

// Withdraw money to bank account balance
func Withdraw(username string) {

}

// Deposit money to bank account balance
func Deposit(username string) {

}

// ShowPendingApps shows the amount of applications the user has applied to
func ShowPendingApps(username string) {

}

// Transfer money from one account to the other
func Transfer(username string) {
	// account1 := c.accounts[acntname]
	// account2 := c.accounts[acntname2]
	// account1.Withdraw(money)
	// account2.Deposit(money)
}
