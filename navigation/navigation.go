package navigation

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/gmac220/project-0/customer"
	"github.com/gmac220/project-0/employees"
	"github.com/gmac220/project-0/opendb"
)

// Welcome prints a welcome screen to the user
func Welcome() {
	fmt.Println()
	fmt.Println("╦ ╦┌─┐┬  ┌─┐┌─┐┌┬┐┌─┐  ┌┬┐┌─┐  ┌┐ ┬┬  ┬  ┌─┐  ┌─┐┌─┐┬─┐┌─┐┌─┐")
	fmt.Println("║║║├┤ │  │  │ ││││├┤    │ │ │  ├┴┐││  │  └─┐  ├┤ ├─┤├┬┘│ ┬│ │")
	fmt.Println("╚╩╝└─┘┴─┘└─┘└─┘┴ ┴└─┘   ┴ └─┘  └─┘┴┴─┘┴─┘└─┘  └  ┴ ┴┴└─└─┘└─┘")
	Selection()
}

// Selection prompts the user what they would want to do
func Selection() {
	var num int
	var username string
	var pass string
	var firstname string
	var lastname string
	var choice string

	fmt.Println("What would you like to do?")
	fmt.Println("1: Sign In")
	fmt.Println("2: Employee Sign In")
	fmt.Println("3: Create an account")
	fmt.Println("4: Exit")
	fmt.Printf("Please type in a number: ")
	fmt.Scanln(&num)

	switch num {
	case 1:
		fmt.Printf("Enter username: ")
		fmt.Scanln(&username)
		fmt.Printf("Enter password: ")
		SttyCommand("-echo")
		fmt.Scanln(&pass)
		SttyCommand("echo")
		fmt.Println()
		SignIn(username, pass, false)
	case 2:
		fmt.Printf("Enter username: ")
		fmt.Scanln(&username)
		fmt.Printf("Enter password: ")
		SttyCommand("-echo")
		fmt.Scanln(&pass)
		SttyCommand("echo")
		fmt.Println()
		SignIn(username, pass, true)
	case 3:
		fmt.Printf("Enter your firstname: ")
		fmt.Scanln(&firstname)
		fmt.Printf("Enter your lastname: ")
		fmt.Scanln(&lastname)
		fmt.Printf("Enter username: ")
		fmt.Scanln(&username)
		fmt.Printf("Enter password: ")
		SttyCommand("-echo")
		fmt.Scanln(&pass)
		SttyCommand("echo")
		fmt.Println()
		fmt.Printf("Is this Account for a Customer or an Employee? type c or e: ")
		fmt.Scanln(&choice)
		CreateAccount(firstname, lastname, username, pass, choice)
	case 4:
		os.Exit(0)
	}
}

// SignIn verifies if customer or employee credentials match database
func SignIn(username string, password string, employee bool) {
	var usernamedb string
	var passdb string
	var fname string
	var lname string
	var row *sql.Row
	db := opendb.OpenDB()
	defer db.Close()
	if employee {
		row = db.QueryRow("SELECT password FROM employees WHERE username = $1", username)
		row.Scan(&passdb)
	} else {
		row = db.QueryRow("SELECT * FROM customers WHERE username = $1", username)
		row.Scan(&usernamedb, &passdb, &fname, &lname)
	}

	if passdb == password {
		fmt.Println("Login Successful!")
		if employee {
			employees.EmployeePage()
		} else {
			customer.SetCustomerVars(usernamedb, fname, lname)
		}
	} else {
		fmt.Println("Password does not match")
		Selection()
	}
}

// CreateAccount for either a customer or employee
func CreateAccount(firstname string, lastname string, username string, password string, choice string) {
	db := opendb.OpenDB()
	defer db.Close()
	switch choice {
	case "c":
		db.Exec("INSERT INTO customers (username, password, firstname, lastname)"+
			"VALUES ($1, $2, $3, $4)", username, password, firstname, lastname)
		customer.SetCustomerVars(username, firstname, lastname)

	case "e":
		db.Exec("INSERT INTO employees (username, password, firstname, lastname)"+
			"VALUES ($1, $2, $3, $4)", username, password, firstname, lastname)
		employees.EmployeePage()
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
