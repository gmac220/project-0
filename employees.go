package main

import (
	"database/sql"
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
	fmt.Println("5: Exit")
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
		fmt.Printf("Which customer's application do you want to look at?" +
			" (Please input application number): ")
		fmt.Scanln(&acntnumber)
		//implement how to get info
	case 4:
		Applications()
	case 5:
		os.Exit(0)
	}
}

// Approve the Customer's application
// Approve not working correctly
func Approve(num int) {
	var row *sql.Row
	var uname string
	var pw string
	var fname string
	var lname string
	var appcount int
	var acntname string
	var acntnum int
	// look at customer table
	// match customer with username
	// if choose to approve then add them to customers slice
	//e.customers.append()
	// remove from application
	//DeleteApplication(username)
	db := OpenDB()
	defer db.Close()
	row = db.QueryRow("SELECT * FROM applications WHERE acntnumber = $1", num)
	row.Scan(&acntnum, &uname, &pw, &fname, &lname, &appcount, &acntname)
	db.Exec("UPDATE applications SET appcount = $1 WHERE username = $2", appcount-1, uname)
	db.Exec("INSERT INTO accounts (acntname, balance, username) VALUES($1, $2, $3)")
	DeleteApplication(acntnum)
}

// DeleteApplication deletes row from table
func DeleteApplication(num int) {
	db := OpenDB()
	defer db.Close()
	db.Exec("DELETE FROM applications WHERE acntnumber = $1", num)
}

// Applications loops through database and puts employees in struct
func Applications() {
	db := OpenDB()
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM applications")
	for rows.Next() {
		var acntnum int
		var username string
		var fname string
		var lname string
		var joint bool
		var uname2 string
		var fname2 string
		var lname2 string

		rows.Scan(&acntnum, &username, &fname, &lname, &joint, &uname2, &fname2, &lname2)
		if joint {
			fmt.Println(acntnum, username, fname, lname, uname2, fname2, lname2, joint)
		} else {
			fmt.Println(acntnum, username, fname, lname)
		}
	}
}
