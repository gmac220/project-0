package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Welcome to Bills FarGO!")
	Selection()

}

// Selection
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
		SignIn()
	case 2:
		EmployeeSignIn()
	case 3:
		CreateAccount()
	case 4:
		os.Exit(0)
	}
}

// SignIn verifies if user has an account with the bank
func SignIn() {
	var username string
	var pass string

	fmt.Printf("Enter username: ")
	fmt.Scanln(&username)
	fmt.Printf("Enter password: ")
	SttyCommand("-echo")
	fmt.Scanln(&pass)
	SttyCommand("echo")
}

// EmployeeSignIn()
func EmployeeSignIn() {
	//SignIn()
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
		db.Exec("INSERT INTO customers (username, password, firstname, lastname, appcount)"+
			"VALUES ($1, $2, $3, $4, 0)", username, pw, firstname, lastname)

	case "e":
		db.Exec("INSERT INTO employees (username, password, firstname, lastname)"+
			"VALUES ($1, $2, $3, $4)", username, pw, firstname, lastname)
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
