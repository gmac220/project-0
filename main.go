package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Welcome to Bills FarGO")
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

// SignIn
func SignIn() {
	var username string
	var pass string
	fmt.Printf("Enter username: ")
	fmt.Scanln(&username)
	fmt.Printf("Enter password: ")
	SttyCommand("-echo")
	fmt.Scanln(&pass)
	SttyCommand("echo")
	fmt.Println(username, pass)
}

// EmployeeSignIn()
func EmployeeSignIn() {

}

// CreateAccount
func CreateAccount() {

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
