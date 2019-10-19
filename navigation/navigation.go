package navigation

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
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
	var username, pass, firstname, lastname, choice string

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
		if username == "" {
			fmt.Printf("Can't have empty username \n")
			Selection()
		}
		fmt.Printf("Enter password: ")
		SttyCommand("-echo")
		fmt.Scanln(&pass)
		SttyCommand("echo")
		fmt.Println()
		usernamedb, passdb, fname, lname := SignIn(username, pass, false)
		if passdb == pass {
			fmt.Println("Login Successful!")
			customer.SetCustomerVars(usernamedb, fname, lname)
			customer.ShowCustomerPrompts()
		} else {
			fmt.Println("Password does not match")
			Selection()
		}

	case 2:
		fmt.Printf("Enter username: ")
		fmt.Scanln(&username)
		if username == "" {
			fmt.Printf("Can't have empty username \n")
			Selection()
		}
		fmt.Printf("Enter password: ")
		SttyCommand("-echo")
		fmt.Scanln(&pass)
		SttyCommand("echo")
		fmt.Println()
		_, passdb, _, _ := SignIn(username, pass, true)
		if passdb == pass {
			fmt.Println("Login Successful!")
			employees.EmployeePage()
		} else {
			fmt.Println("Password does not match")
			Selection()
		}
	case 3:
		fmt.Printf("Enter your firstname: ")
		fmt.Scanln(&firstname)
		fmt.Printf("Enter your lastname: ")
		fmt.Scanln(&lastname)
		fmt.Printf("Enter username: ")
		fmt.Scanln(&username)
		if username == "" {
			fmt.Printf("Can't have empty username \n")
			Selection()
		}
		fmt.Printf("Enter password: ")
		SttyCommand("-echo")
		fmt.Scanln(&pass)
		SttyCommand("echo")
		fmt.Println()
		fmt.Printf("Is this Account for a Customer or an Employee? type c or e: ")
		fmt.Scanln(&choice)
		CreateAccount(firstname, lastname, username, pass, choice)
		switch choice {
		case "c":
			customer.SetCustomerVars(username, firstname, lastname)
			customer.ShowCustomerPrompts()
		case "e":
			employees.EmployeePage()
		}
	case 4:
		os.Exit(0)
	default:
		fmt.Println("Choice does not exist.")
		Selection()
	}
}

// SignIn verifies if customer or employee credentials match database
func SignIn(username string, password string, employee bool) (string, string, string, string) {
	var usernamedb, passdb, fname, lname string
	var row *sql.Row

	db := opendb.OpenDB()
	defer db.Close()
	if employee {
		row = db.QueryRow("SELECT * FROM employees WHERE username = $1", username)
		row.Scan(&usernamedb, &passdb, &fname, &lname)
	} else {
		row = db.QueryRow("SELECT * FROM customers WHERE username = $1", username)
		row.Scan(&usernamedb, &passdb, &fname, &lname)
	}
	return usernamedb, passdb, fname, lname
}

// CreateAccount for either a customer or employee
func CreateAccount(firstname string, lastname string, username string, password string, choice string) {

	db := opendb.OpenDB()
	defer db.Close()
	switch choice {
	case "c":
		db.Exec("INSERT INTO customers (username, password, firstname, lastname)"+
			"VALUES ($1, $2, $3, $4)", username, password, firstname, lastname)
	case "e":
		db.Exec("INSERT INTO employees (username, password, firstname, lastname)"+
			"VALUES ($1, $2, $3, $4)", username, password, firstname, lastname)
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

func Formsubmit(response http.ResponseWriter, request *http.Request) {
	var username = request.FormValue("username")
	var pass = request.FormValue("pass")
	var firstname = request.FormValue("firstname")
	var lastname = request.FormValue("lastname")

	fmt.Println("Username", username)
	fmt.Println("Password", pass)
	fmt.Println("First name", firstname)
	fmt.Println("Last name", lastname)

	db := opendb.OpenDB()
	defer db.Close()
	result, err := db.Exec("INSERT INTO customers (username, password, firstname, lastname) VALUES ($1, $2, $3, $4)", username, pass, firstname, lastname)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
