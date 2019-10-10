package main

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var cusername string
var cfirstname string
var clastname string

// CustomerInit initializes global variables
func CustomerInit(username string, firstname string, lastname string) {
	cusername = username
	cfirstname = firstname
	clastname = lastname
	CustomerPage()
}

// CustomerPage prompts the customer what they could do and asks them to pick a choice.
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
	var acntname string
	var appcount int

	db := OpenDB()
	defer db.Close()
	fmt.Printf("What type of account do you want to open? checking, savings, other...: ")
	fmt.Scanln(&acntname)
	db.Exec("INSERT INTO applications"+
		"(username, firstname, lastname, acntname, joint, username2, firstname2, lastname2)"+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		username, firstname, lastname, acntname, joint, "N/A", "N/A", "N/A")
	row := db.QueryRow("SELECT appcount FROM customers WHERE username = $1", username)
	row.Scan(&appcount)
	db.Exec("UPDATE customers SET appcount = $1 WHERE username = $2", appcount+1, username)
	fmt.Println("Application Successful!")
	CustomerPage()
}

// JointApp adds to applications table
func JointApp(username string, firstname string, lastname string, joint bool) {
	var uname2 string
	var fname2 string
	var lname2 string
	var acntname string
	var appcount int

	db := OpenDB()
	defer db.Close()
	fmt.Printf("What type of account do you want to open? checking, savings, other...: ")
	fmt.Scanln(&acntname)
	fmt.Printf("What user do you want to share an account with? Please input username of that user: ")
	fmt.Scanln(&uname2)
	uname2, fname2, lname2 = CheckCustomer(uname2)
	for uname2 == "" {
		fmt.Printf("This user does not exist please input a valid user: ")
		fmt.Scanln(&uname2)
		uname2, fname2, lname2 = CheckCustomer(uname2)
	}

	row := db.QueryRow("SELECT appcount FROM customers WHERE username = $1", username)
	row.Scan(&appcount)
	db.Exec("UPDATE customers SET appcount = $1 WHERE username = $2", appcount+1, username)
	row = db.QueryRow("SELECT appcount FROM customers WHERE username = $1", uname2)
	row.Scan(&appcount)
	db.Exec("UPDATE customers SET appcount = $1 WHERE username = $2", appcount+1, uname2)
	db.Exec("INSERT INTO applications"+
		"(username, firstname, lastname, acntname, joint, username2, firstname2, lastname2)"+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		username, firstname, lastname, acntname, joint, uname2, fname2, lname2)
	fmt.Println("Joint Application with", uname2, "Successful!")
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
	var username string

	db := OpenDB()
	defer db.Close()
	//loop through both accounts and joint accounts
	fmt.Println("-------------------------YOUR ACCOUNTS-------------------------")
	fmt.Println()
	rows, _ := db.Query("SELECT acntnumber, balance, acntname, username FROM accounts WHERE username = $1 OR username2 = $2", cusername, cusername)
	for rows.Next() {

		rows.Scan(&acntnumber, &balance, &acntname, &username)
		if username != cusername {
			fmt.Println("Account #:", acntnumber, "|Account name:", acntname, "|Balance:", balance, "|Other Account Holder:", username)
		} else {
			fmt.Println("Account #:", acntnumber, "|Account name:", acntname, "|Balance:", balance)
		}

	}
	fmt.Println()
	fmt.Println("---------------------------------------------------------------")
	CustomerPage()
}

// ShowBalance shows the balance of an account the user chooses
func ShowBalance() {
	var acntnum int

	fmt.Printf("What account number do you want to check the balance for?: ")
	fmt.Scanln(&acntnum)
	balance, acntname := VerifyAccount(acntnum)
	if acntname == "" {
		fmt.Println("That account does not exist.")
	} else {
		fmt.Println("This is your balance: ", balance)
	}
	CustomerPage()
}

// VerifyAccount checks if there is an account with number specified
func VerifyAccount(accountnumber int) (float64, string) {
	var balance float64
	var acntname string
	db := OpenDB()
	defer db.Close()
	row := db.QueryRow("SELECT balance, acntname FROM accounts WHERE acntnumber = $1", accountnumber)
	row.Scan(&balance, &acntname)
	if acntname == "" {
		fmt.Println("That account does not exist.")
		CustomerPage()
	}
	return balance, acntname
}

// Withdraw takes out money from an account the user chooses
func Withdraw() {
	var acntnum int
	var withdrawal float64

	fmt.Printf("What account number would you like to withdraw from: ")
	fmt.Scanln(&acntnum)
	balance, _ := VerifyAccount(acntnum)
	fmt.Printf("How much money would you like to withdraw? Ex. 20.02: ")
	fmt.Scanln(&withdrawal)
	db := OpenDB()
	defer db.Close()
	for withdrawal > balance {
		fmt.Printf("Not enough money in balance please enter another amount: ")
		fmt.Scanln(&withdrawal)
	}
	db.Exec("UPDATE accounts SET balance = $1 WHERE acntnumber = $2", balance-withdrawal, acntnum)
	fmt.Println("Withdraw Successful!")
	CustomerPage()
}

// Deposit adds money to an account the user chooses
func Deposit() {
	var acntnum int
	var deposit float64

	fmt.Printf("What account number would you like to deposit into: ")
	fmt.Scanln(&acntnum)
	balance, _ := VerifyAccount(acntnum)
	fmt.Printf("How much money would you like to deposit? Ex. 20.02: ")
	fmt.Scanln(&deposit)
	db := OpenDB()
	defer db.Close()
	db.Exec("UPDATE accounts SET balance = $1 WHERE acntnumber = $2", balance+deposit, acntnum)
	fmt.Println("Deposit Successful!")
	CustomerPage()
}

// ShowPendingApps shows the amount of applications the user has applied to
func ShowPendingApps() {
	var uname string
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
	fmt.Println()
	rows, _ := db.Query("SELECT * FROM applications WHERE username = $1", cusername)
	for rows.Next() {
		rows.Scan(&acntnum, &uname, &fname, &lname, &acntname, &joint, &uname2, &fname2, &lname2)
		if joint {
			fmt.Println("Account #:", acntnum, "|Type: Joint", "|Username1:", uname,
				"|First name:", fname, "|Last name:", lname, "|Username2:", uname2,
				"|First name:", fname2, "|Last name:", lname2)
		} else {
			fmt.Println("Account #:", acntnum, "|Type: Solo ", "|Username:", uname,
				" |First name:", fname, "|Last name:", lname)
		}
	}
	fmt.Println()
	fmt.Println("---------------------------------------------------------------")
	CustomerPage()
}

// Transfer takes money from one account to another account available in the database
func Transfer() {
	var acntnumwithdraw int
	var acntnumdeposit int
	var funds float64

	fmt.Printf("What account number would you like to take money out of?: ")
	fmt.Scanln(&acntnumwithdraw)
	balance, _ := VerifyAccount(acntnumwithdraw)
	fmt.Printf("What account number would you like to transfer into?: ")
	fmt.Scanln(&acntnumdeposit)
	VerifyAccount(acntnumdeposit)
	fmt.Printf("How much money would you like to transfer? Ex. 20.02: ")
	fmt.Scanln(&funds)
	db := OpenDB()
	defer db.Close()
	for funds > balance {
		fmt.Printf("Not enough money in balance please enter another amount: ")
		fmt.Scanln(&funds)
	}
	db.Exec("UPDATE accounts SET balance = $1 WHERE acntnumber = $2", balance-funds, acntnumwithdraw)
	balance, _ = VerifyAccount(acntnumdeposit)
	db.Exec("UPDATE accounts SET balance = $1 WHERE acntnumber = $2", balance+funds, acntnumdeposit)
	fmt.Println("Transfer Successful!")
	CustomerPage()
}
