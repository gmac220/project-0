package main

import (
	"fmt"
	"os"
)

// EmployeePage prompts the employee what they could do and asks them to pick a choice.
func EmployeePage() {
	var num int
	var acntnumber int
	var username string

	fmt.Println("What do you want to do")
	fmt.Println("1: Approve")
	fmt.Println("2: Deny")
	fmt.Println("3: View Customer Info")
	fmt.Println("4: Show Applications")
	fmt.Println("5: Logout")
	fmt.Println("6: Exit")
	fmt.Printf("Please type in a number: ")
	fmt.Scanln(&num)

	switch num {
	case 1:
		fmt.Printf("Which customer's application do you want to approve?" +
			" (Please input application number): ")
		fmt.Scanln(&acntnumber)
		Approve(acntnumber)
	case 2:
		fmt.Printf("Which customer's application do you want to deny?" +
			" (Please input application number): ")
		fmt.Scanln(&acntnumber)
		DeleteApplication(acntnumber)

	case 3:
		fmt.Printf("Which customer's information do you want to look at?" +
			" (Please input customer's username): ")
		fmt.Scanln(&username)
		CustomerInfo(username)
	case 4:
		Applications()
	case 5:
		main()
	case 6:
		os.Exit(0)
	}
}

// Approve the Customer's application
func Approve(num int) {
	var uname string
	var appcount int
	var acntname string
	var joint bool
	var uname2 string

	db := OpenDB()
	defer db.Close()
	row := db.QueryRow("SELECT username, acntname, joint, username2 FROM applications WHERE acntnumber = $1", num)
	row.Scan(&uname, &acntname, &joint, &uname2)
	if joint {
		// Updates appcount from both users in joint account. Adds the account in accounts table.
		row = db.QueryRow("SELECT appcount FROM customers WHERE username = $1", uname)
		row.Scan(&appcount)
		db.Exec("UPDATE customers SET appcount = $1 WHERE username = $2", appcount-1, uname)
		row = db.QueryRow("SELECT appcount FROM customers WHERE username = $1", uname2)
		row.Scan(&appcount)
		db.Exec("UPDATE customers SET appcount = $1 WHERE username = $2", appcount-1, uname2)
		db.Exec("INSERT INTO accounts (acntname, balance, username) VALUES ($1, $2, $3)", "joint"+acntname+uname2+uname2, 0, uname)
		db.Exec("INSERT INTO accounts (acntname, balance, username) VALUES ($1, $2, $3)", "joint"+acntname+uname2+uname2, 0, uname2)
	} else {
		row = db.QueryRow("SELECT appcount FROM customers WHERE username = $1", uname)
		row.Scan(&appcount)
		db.Exec("UPDATE customers SET appcount = $1 WHERE username = $2", appcount-1, uname)
		db.Exec("INSERT INTO accounts (acntname, balance, username) VALUES ($1, $2, $3)", acntname, 0, uname)

	}
	DeleteApplication(num)
}

// DeleteApplication deletes row from applications table
func DeleteApplication(num int) {
	db := OpenDB()
	defer db.Close()
	db.Exec("DELETE FROM applications WHERE acntnumber = $1", num)
	EmployeePage()
}

// CustomerInfo looks at all of customers account information by passing in their username
func CustomerInfo(username string) {
	var acntnumber int
	var acntname string
	var balance float64
	var uname string
	var uname2 string
	var pw string
	var fname string
	var lname string
	var appcount int

	db := OpenDB()
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM accounts FULL OUTER JOIN customers on customers.username = accounts.username WHERE customers.username = $1", username)
	fmt.Println("---------------------------" + username + "'s INFORMATION------------------------------")
	fmt.Println()
	for rows.Next() {
		rows.Scan(&acntnumber, &acntname, &balance, &uname, &uname2, &pw, &fname, &lname, &appcount)
		fmt.Println(acntnumber, acntname, balance, uname)
	}
	fmt.Println()
	fmt.Println("---------------------------------------------------------------------------------")
	EmployeePage()
}

// Applications loops through applications table listing if they are joint applications or not
func Applications() {
	var acntnum int
	var acntname string
	var username string
	var fname string
	var lname string
	var joint bool
	var uname2 string
	var fname2 string
	var lname2 string

	db := OpenDB()
	defer db.Close()
	fmt.Println("----------------------LISTED APPLICATIONS----------------------")
	fmt.Println()
	rows, _ := db.Query("SELECT * FROM applications")
	for rows.Next() {
		rows.Scan(&acntnum, &username, &fname, &lname, &acntname, &joint, &uname2, &fname2, &lname2)
		if joint {
			fmt.Println(acntnum, joint, username, fname, lname, uname2, fname2, lname2)
		} else {
			fmt.Println(acntnum, joint, username, fname, lname)
		}
	}
	fmt.Println()
	fmt.Println("---------------------------------------------------------------")
	EmployeePage()
}
