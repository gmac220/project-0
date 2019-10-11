package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	fmt.Println()
	fmt.Println("╦ ╦┌─┐┬  ┌─┐┌─┐┌┬┐┌─┐  ┌┬┐┌─┐  ┌┐ ┬┬  ┬  ┌─┐  ┌─┐┌─┐┬─┐┌─┐┌─┐")
	fmt.Println("║║║├┤ │  │  │ ││││├┤    │ │ │  ├┴┐││  │  └─┐  ├┤ ├─┤├┬┘│ ┬│ │")
	fmt.Println("╚╩╝└─┘┴─┘└─┘└─┘┴ ┴└─┘   ┴ └─┘  └─┘┴┴─┘┴─┘└─┘  └  ┴ ┴┴└─└─┘└─┘")
	Selection()

}

// Selection prompts the user what they would want to do
func Selection() {
	var num int
	fmt.Println("What would you like to do?")
	fmt.Println("1: Sign In")
	fmt.Println("2: Employee Sign In")
	fmt.Println("3: Create an account")
	fmt.Println("4: Exit")
	fmt.Printf("Please type in a number: ")
	fmt.Scanln(&num)

	switch num {
	case 1:
		SignIn(false)
	case 2:
		SignIn(true)
	case 3:
		CreateAccount()
	case 4:
		os.Exit(0)
	}
}

// SignIn verifies if customer or employee credentials match database
func SignIn(employee bool) {
	var username string
	var pass string
	var usernamedb string
	var passdb string
	var fname string
	var lname string
	var row *sql.Row

	fmt.Printf("Enter username: ")
	fmt.Scanln(&username)
	fmt.Printf("Enter password: ")
	SttyCommand("-echo")
	fmt.Scanln(&pass)
	SttyCommand("echo")
	fmt.Println()
	db := OpenDB()
	defer db.Close()
	if employee {
		row = db.QueryRow("SELECT password FROM employees WHERE username = $1", username)
		row.Scan(&passdb)
	} else {
		row = db.QueryRow("SELECT * FROM customers WHERE username = $1", username)
		row.Scan(&usernamedb, &passdb, &fname, &lname)
	}

	if passdb == pass {
		fmt.Println("Login Successful!")
		if employee {
			EmployeePage()
		} else {
			CustomerInit(usernamedb, fname, lname)
		}
	} else {
		fmt.Println("Password does not match")
		Selection()
	}
}

// CreateAccount for either a customer or employee
func CreateAccount() {
	var choice string
	var username string
	var pw string
	var firstname string
	var lastname string

	db := OpenDB()
	defer db.Close()
	fmt.Printf("Enter your firstname: ")
	fmt.Scanln(&firstname)
	fmt.Printf("Enter your lastname: ")
	fmt.Scanln(&lastname)
	fmt.Printf("Enter username: ")
	fmt.Scanln(&username)
	fmt.Printf("Enter password: ")
	SttyCommand("-echo")
	fmt.Scanln(&pw)
	SttyCommand("echo")
	fmt.Println()
	fmt.Printf("Is this Account for a Customer or an Employee? type c or e: ")
	fmt.Scanln(&choice)
	switch choice {
	case "c":
		db.Exec("INSERT INTO customers (username, password, firstname, lastname)"+
			"VALUES ($1, $2, $3, $4)", username, pw, firstname, lastname)
		CustomerInit(username, firstname, lastname)

	case "e":
		db.Exec("INSERT INTO employees (username, password, firstname, lastname)"+
			"VALUES ($1, $2, $3, $4)", username, pw, firstname, lastname)
		EmployeePage()
	}
}

// SttyCommand hides the password so it doesn't show up in the terminal when user types it
func SttyCommand(flag string) {
	cmd := exec.Command("stty", flag)
	cmd.Stdin = os.Stdin
	_, err := cmd.Output()
	cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
}
