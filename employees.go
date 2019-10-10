package main

import (
	"fmt"
	"os"
)

// Employee struct
type Employee struct {
	username  string
	password  string
	firstname string
	lastname  string
	// figure out how to add customers in approve
	customers    []Customer
	applications []string
}

func NewEmployee(username string, firstname string, lastname string, password string) {
	db := OpenDB()
	defer db.Close()
	//db.exec("INSERT INTO employees (username, firstname, lastname, password)
	//  VALUES($1, $2, $3, $4)", username, firstname, lastname, password)
}

func EmployeePage() {
	var num int
	var acntnumber int
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
		fmt.Scanln(&acntnumber)
		//implement how to get info
	case 4:
		Applications()
	case 5:
		main()
	case 6:
		os.Exit(0)
	}
}

// Approve the Customer's application
// Approve not working correctly
func Approve(num int) {
	//var row *sql.Row
	var uname string
	var appcount int
	var acntname string
	var joint bool
	var uname2 string

	db := OpenDB()
	defer db.Close()
	row := db.QueryRow("SELECT username, acntname, joint, username2 FROM applications WHERE acntnumber = $1", num)
	row.Scan(&uname, &acntname, &joint, &uname2)
	fmt.Println(uname, acntname, joint, uname2)
	if joint {
		// Updates appcount from both users in joint account. Adds the account in accounts table.
		row = db.QueryRow("SELECT appcount FROM customers WHERE username = $1", uname)
		row.Scan(&appcount)
		db.Exec("UPDATE customers SET appcount = $1 WHERE username = $2", appcount-1, uname)
		row = db.QueryRow("SELECT appcount FROM customers WHERE username = $1", uname2)
		row.Scan(&appcount)
		db.Exec("UPDATE customers SET appcount = $1 WHERE username = $2", appcount-1, uname2)
		db.Exec("INSERT INTO accounts (acntname, balance, username) VALUES ($1, $2, $3)", "joint", 0, uname)
		db.Exec("INSERT INTO accounts (acntname, balance, username) VALUES ($1, $2, $3)", "joint", 0, uname2)
	} else {
		row = db.QueryRow("SELECT appcount FROM customers WHERE username = $1", uname)
		row.Scan(&appcount)
		db.Exec("UPDATE customers SET appcount = $1 WHERE username = $2", appcount-1, uname)
		db.Exec("INSERT INTO accounts (acntname, balance, username) VALUES ($1, $2, $3)", acntname, 0, uname)

	}
	DeleteApplication(num)
}

// DeleteApplication deletes row from table
func DeleteApplication(num int) {
	db := OpenDB()
	defer db.Close()
	db.Exec("DELETE FROM applications WHERE acntnumber = $1", num)
	EmployeePage()
}

// CustomerInfo looks at all of customers account information
func CustomerInfo() {
	//make a search join where you use username as id
	db := OpenDB()
	defer db.Close()
	EmployeePage()
}

// Applications loops through database and puts employees in struct
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
	rows, _ := db.Query("SELECT * FROM applications")
	for rows.Next() {
		rows.Scan(&acntnum, &username, &fname, &lname, &acntname, &joint, &uname2, &fname2, &lname2)
		if joint {
			fmt.Println(acntnum, joint, username, fname, lname, uname2, fname2, lname2)
		} else {
			fmt.Println(acntnum, joint, username, fname, lname)
		}
	}
	fmt.Println("---------------------------------------------------------------")
	EmployeePage()
}
