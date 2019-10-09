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
	fmt.Println("What do you want to do")
	fmt.Println("1: Sign In")
	fmt.Println("2: Employee Sign In")
	fmt.Println("3: Create an account")
	fmt.Println("4: Exit")
	fmt.Printf("Please type in a number: ")
	fmt.Scanln(&num)

	switch num {
	case 1:
		SignIn()
	case 2:
		EmployeeSignIn()
	case 3:
		CreateAccount()
	case 4:
		os.Exit(0)
	}
}

// Approve the Customer's application
func Approve(username string) {
	// look at customer table
	// match customer with username
	// db.QueryRow("SELECT * FROM customers WHERE username = $1", username)
	// if choose to approve then add them to customers slice
	//e.customers.append()
	// remove from application
	//DeleteApplication(username)
	db := OpenDB()
	defer db.Close()
	DeleteApplication(username)
	fmt.Println()
}

// Deny the Customer's application
func (e *Employee) Deny() {
	db := OpenDB()
	defer db.Close()
}

// DeleteApplication deletes row from table
func DeleteApplication(username string) {
	db := OpenDB()
	defer db.Close()
	db.Exec("DELETE FROM applications (username) VALUES($1)", username)

}

// Applications loops through database and puts employees in struct
func Applications() {
	db := OpenDB()
	defer db.Close()
	rows, _ := db.Query("SELECT * FROM applications")
	for rows.Next() {
		var username string
		var fname string
		var lname string
		var joint bool
		var uname2 string
		var fname2 string
		var lname2 string

		rows.Scan(&username, &fname, &lname, &joint, &uname2, &fname2, &lname2)
		if joint {
			fmt.Println(username, fname, lname, uname2, fname2, lname2, joint)
		} else {
			fmt.Println(username, fname, lname)
		}

	}
}
