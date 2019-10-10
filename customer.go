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
func ShowAccounts() {
	var acntnumber int
	var acntname string
	var balance float64

	db := OpenDB()
	defer db.Close()
	//loop through both accounts and joint accounts
	fmt.Println("-------------------------SHOW ACCOUNTS-------------------------")
	fmt.Println()
	rows, _ := db.Query("SELECT acntnumber, balance, acntname FROM accounts WHERE username = $1", cusername)
	for rows.Next() {

		rows.Scan(&acntnumber, &balance, &acntname)
		fmt.Println(acntnumber, acntname, balance)
	}
	fmt.Println()
	fmt.Println("---------------------------------------------------------------")
	CustomerPage()
}

// ShowBalance of account
func ShowBalance() {
	var acntnum int
	var balance float64

	fmt.Printf("What account number do you want to check the balance for?: ")
	fmt.Scanln(&acntnum)
	db := OpenDB()
	defer db.Close()
	row := db.QueryRow("SELECT balance FROM accounts WHERE acntnumber = $1", acntnum)
	row.Scan(&balance)
	fmt.Println("This is your balance: ", balance)
	CustomerPage()
}

// Withdraw money to bank account balance
func Withdraw() {
	var acntnum int
	var withdrawal float64
	var balance float64

	fmt.Printf("What account number would you like to withdraw from: ")
	fmt.Scanln(&acntnum)
	fmt.Printf("How much money would you like to withdraw? Ex. 20.02:")
	fmt.Scanln(&withdrawal)
	db := OpenDB()
	defer db.Close()
	row := db.QueryRow("SELECT balance FROM accounts WHERE acntnumber = $1", acntnum)
	row.Scan(&balance)
	for withdrawal > balance {
		fmt.Printf("Not enough money in balance please enter another amount: ")
		fmt.Scanln(&withdrawal)
	}
	db.Exec("UPDATE accounts SET balance = $1 WHERE acntnumber = $2", balance-withdrawal, acntnum)
	CustomerPage()
}

// Deposit money to bank account balance
func Deposit() {
	var acntnum int
	var deposit float64
	var balance float64

	fmt.Printf("What account number would you like to deposit into: ")
	fmt.Scanln(&acntnum)
	fmt.Printf("How much money would you like to deposit? Ex. 20.02:")
	fmt.Scanln(&deposit)
	db := OpenDB()
	defer db.Close()
	row := db.QueryRow("SELECT balance FROM accounts WHERE acntnumber = $1", acntnum)
	row.Scan(&balance)
	db.Exec("UPDATE accounts SET balance = $1 WHERE acntnumber = $2", balance+deposit, acntnum)
	CustomerPage()
}

// ShowPendingApps shows the amount of applications the user has applied to
func ShowPendingApps() {
	var username string
	var acntnum int
	var acntname string
	var fname string
	var lname string
	var joint bool
	var uname2 string
	var fname2 string
	var lname2 string

	db := OpenDB()
	defer db.Close()
	fmt.Println("----------------------PENDING APPLICATIONS----------------------")
	rows, _ := db.Query("SELECT * FROM applications WHERE username = $1", cusername)
	for rows.Next() {
		rows.Scan(&acntnum, &username, &fname, &lname, &acntname, &joint, &uname2, &fname2, &lname2)
		if joint {
			fmt.Println(acntnum, username, fname, lname, uname2, fname2, lname2)
		} else {
			fmt.Println(acntnum, username, fname, lname)
		}
	}
	fmt.Println("---------------------------------------------------------------")
	CustomerPage()
}

// Transfer money from one account to the other
func Transfer() {
	var acntnumwithdraw int
	var acntnumdeposit int
	var funds float64
	var balance float64

	fmt.Printf("What account number would you like to take money out of?: ")
	fmt.Scanln(&acntnumwithdraw)
	fmt.Printf("What account number would you like to transfer into?: ")
	fmt.Scanln(&acntnumdeposit)
	fmt.Printf("How much money would you like to transfer? Ex. 20.02: ")
	fmt.Scanln(&funds)
	db := OpenDB()
	defer db.Close()
	row := db.QueryRow("SELECT balance FROM accounts WHERE acntnumber = $1", acntnumwithdraw)
	row.Scan(&balance)
	for funds > balance {
		fmt.Printf("Not enough money in balance please enter another amount: ")
		fmt.Scanln(&funds)
	}
	db.Exec("UPDATE accounts SET balance = $1 WHERE acntnumber = $2", balance-funds, acntnumwithdraw)
	row = db.QueryRow("SELECT balance FROM accounts WHERE acntnumber = $1", acntnumdeposit)
	row.Scan(&balance)
	db.Exec("UPDATE accounts SET balance = $1 WHERE acntnumber = $2", balance+funds, acntnumdeposit)
	CustomerPage()
}
