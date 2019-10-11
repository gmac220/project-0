package customer

import (
	"fmt"
	"log"
	"os"

	"github.com/gmac220/project-0/opendb"
)

var cusername string
var cfirstname string
var clastname string

// SetCustomerVars initializes global variables
func SetCustomerVars(username string, firstname string, lastname string) {
	cusername = username
	cfirstname = firstname
	clastname = lastname
	ShowCustomerPrompts()
}

// ShowCustomerPrompts asks the customer what they could do and asks them to pick a choice.
func ShowCustomerPrompts() {
	var num int
	var acntname string
	var username2 string
	var acntnum int
	var withdrawal float64
	var deposit float64
	var acntnumwithdraw int
	var acntnumdeposit int
	var funds float64

	fmt.Println("What would you like to do today?")
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
		fmt.Printf("What type of account do you want to open? checking, savings, other...: ")
		fmt.Scanln(&acntname)
		Apply(cusername, cfirstname, clastname, acntname)
	case 2:
		fmt.Printf("What type of account do you want to open? checking, savings, other...: ")
		fmt.Scanln(&acntname)
		fmt.Printf("What user do you want to share an account with? Please input username of that user: ")
		fmt.Scanln(&username2)
		JointApp(cusername, cfirstname, clastname, acntname, username2)
	case 3:
		ShowAccounts(cusername)
	case 4:
		fmt.Printf("What account number do you want to check the balance for?: ")
		fmt.Scanln(&acntnum)
		ShowBalance(acntnum)
	case 5:
		fmt.Printf("What account number would you like to withdraw from: ")
		fmt.Scanln(&acntnum)
		balance, _ := VerifyAccount(acntnum)
		if CheckOwnAccount(acntnum) {
			fmt.Printf("How much money would you like to withdraw? Ex. 20.02: ")
			fmt.Scanln(&withdrawal)
			Withdraw(acntnum, withdrawal, balance)
		} else {
			fmt.Println("You do not have access to this account.")
		}
		ShowCustomerPrompts()
	case 6:
		fmt.Printf("What account number would you like to deposit into: ")
		fmt.Scanln(&acntnum)
		balance, _ := VerifyAccount(acntnum)
		if CheckOwnAccount(acntnum) {
			fmt.Printf("How much money would you like to deposit? Ex. 20.02: ")
			fmt.Scanln(&deposit)
			Deposit(acntnum, deposit, balance)
		} else {
			fmt.Println("You do not have access to this account.")
		}
		ShowCustomerPrompts()
	case 7:
		fmt.Printf("What account number would you like to take money out of?: ")
		fmt.Scanln(&acntnumwithdraw)
		balance, _ := VerifyAccount(acntnumwithdraw)
		if CheckOwnAccount(acntnumwithdraw) {
			fmt.Printf("What account number would you like to transfer into?: ")
			fmt.Scanln(&acntnumdeposit)
			VerifyAccount(acntnumdeposit)
			fmt.Printf("How much money would you like to transfer? Ex. 20.02: ")
			fmt.Scanln(&funds)
			Transfer(acntnumwithdraw, acntnumdeposit, balance, funds)
		} else {
			fmt.Println("You do not have access to this account.")
		}
		ShowCustomerPrompts()
	case 8:
		ShowPendingApps(cusername)
	case 9:
		os.Exit(0)
	}
}

// Apply adds to applications table
func Apply(username string, firstname string, lastname string, acntname string) {
	//var acntname string

	db := opendb.OpenDB()
	defer db.Close()
	db.Exec("INSERT INTO applications"+
		"(username, firstname, lastname, acntname, joint, username2, firstname2, lastname2)"+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		username, firstname, lastname, acntname, false, "N/A", "N/A", "N/A")
	fmt.Println("Application Successful!")
	ShowCustomerPrompts()
}

// JointApp adds to applications table
func JointApp(username string, firstname string, lastname string, acntname string, username2 string) {

	db := opendb.OpenDB()
	defer db.Close()
	uname2, fname2, lname2 := CheckCustomer(username2)
	for uname2 == "" {
		fmt.Printf("This user does not exist please input a valid user: ")
		fmt.Scanln(&uname2)
		uname2, fname2, lname2 = CheckCustomer(uname2)
	}
	db.Exec("INSERT INTO applications"+
		"(username, firstname, lastname, acntname, joint, username2, firstname2, lastname2)"+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		username, firstname, lastname, acntname, true, uname2, fname2, lname2)
	fmt.Println("Joint Application with", uname2, "Successful!")
	ShowCustomerPrompts()
}

// CheckCustomer verifies if the customer exists
func CheckCustomer(username string) (string, string, string) {
	var uname string
	var fname string
	var lname string

	db := opendb.OpenDB()
	defer db.Close()
	row := db.QueryRow("SELECT username, firstname, lastname FROM customers WHERE username = $1", username)
	row.Scan(&uname, &fname, &lname)

	return uname, fname, lname
}

// ShowAccounts lists out the accounts the user currently has
func ShowAccounts(username string) {
	var acntnumber int
	var acntname string
	var balance float64
	var uname string
	var username2 string

	db := opendb.OpenDB()
	defer db.Close()
	//loop through both accounts and joint accounts
	fmt.Println("-------------------------YOUR ACCOUNTS-------------------------")
	fmt.Println()
	rows, err := db.Query("SELECT acntnumber, balance, acntname, username, username2 FROM accounts WHERE username = $1 OR username2 = $2", username, username)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		username = ""
		username2 = ""
		rows.Scan(&acntnumber, &balance, &acntname, &uname, &username2)
		if uname != cusername && username2 != "" {
			fmt.Println("Account #:", acntnumber, "|Account name:", acntname, "|Balance:", balance, "|Other Account Holder:", uname)
		} else if username2 != cusername && username2 != "" {
			fmt.Println("Account #:", acntnumber, "|Account name:", acntname, "|Balance:", balance, "|Other Account Holder:", username2)
		} else {
			fmt.Println("Account #:", acntnumber, "|Account name:", acntname, "|Balance:", balance)
		}

	}
	fmt.Println()
	fmt.Println("---------------------------------------------------------------")
	ShowCustomerPrompts()
}

// CheckOwnAccount verifies if the customer owns the account
func CheckOwnAccount(num int) bool {
	var username string
	var username2 string

	db := opendb.OpenDB()
	defer db.Close()
	row := db.QueryRow("SELECT username, username2 FROM accounts WHERE acntnumber = $1", num)
	row.Scan(&username, &username2)
	return cusername == username || cusername == username2
}

// ShowBalance shows the balance of an account the user chooses
func ShowBalance(acntnum int) {
	balance, _ := VerifyAccount(acntnum)
	if CheckOwnAccount(acntnum) {
		fmt.Println("This is your balance: ", balance)
	} else {
		fmt.Println("You do not have access to this account.")
	}

	ShowCustomerPrompts()
}

// VerifyAccount checks if there is an account with number specified
func VerifyAccount(accountnumber int) (float64, string) {
	var balance float64
	var acntname string

	db := opendb.OpenDB()
	defer db.Close()
	row := db.QueryRow("SELECT balance, acntname FROM accounts WHERE acntnumber = $1", accountnumber)
	row.Scan(&balance, &acntname)
	if acntname == "" {
		fmt.Println("That account does not exist.")
		ShowCustomerPrompts()
	}
	return balance, acntname
}

// Withdraw takes out money from an account the user chooses
func Withdraw(acntnum int, withdrawal float64, balance float64) {
	db := opendb.OpenDB()
	defer db.Close()
	for withdrawal > balance {
		fmt.Printf("Not enough money in balance please enter another amount: ")
		fmt.Scanln(&withdrawal)
	}
	db.Exec("UPDATE accounts SET balance = $1 WHERE acntnumber = $2", balance-withdrawal, acntnum)
	fmt.Println("Withdraw Successful!")
}

// Deposit adds money to an account the user chooses
func Deposit(acntnum int, deposit float64, balance float64) {
	db := opendb.OpenDB()
	defer db.Close()
	db.Exec("UPDATE accounts SET balance = $1 WHERE acntnumber = $2", balance+deposit, acntnum)
	fmt.Println("Deposit Successful!")
}

// ShowPendingApps shows the amount of applications the user has applied to
func ShowPendingApps(username string) {
	var uname string
	var acntnum int
	var acntname string
	var fname string
	var lname string
	var joint bool
	var uname2 string
	var fname2 string
	var lname2 string

	db := opendb.OpenDB()
	defer db.Close()
	fmt.Println("----------------------PENDING APPLICATIONS----------------------")
	fmt.Println()
	rows, _ := db.Query("SELECT * FROM applications WHERE username = $1", username)
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
	ShowCustomerPrompts()
}

// Transfer takes money from one account to another account available in the database
func Transfer(acntnumwithdraw int, acntnumdeposit int, balanceinput float64, fundsinput float64) {
	var funds float64 = fundsinput

	db := opendb.OpenDB()
	defer db.Close()
	for funds > balanceinput {
		fmt.Printf("Not enough money in balance please enter another amount: ")
		fmt.Scanln(&funds)
	}
	db.Exec("UPDATE accounts SET balance = $1 WHERE acntnumber = $2", balanceinput-funds, acntnumwithdraw)
	balancedeposit, _ := VerifyAccount(acntnumdeposit)
	db.Exec("UPDATE accounts SET balance = $1 WHERE acntnumber = $2", balancedeposit+funds, acntnumdeposit)
	fmt.Println("Transfer Successful!")
}
