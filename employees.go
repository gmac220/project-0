package main

import (
	"fmt"
)

// Employee struct
type Employee struct {
	username string
	password string
	// figure out how to add customers in approve
	customers    []Customer
	accounts     []Account
	applications []string
}

func NewEmployee(username string, firstname string, lastname string, password string) {
	//db.exec("INSERT INTO employees (username, firstname, lastname, password)
	//  VALUES($1, $2, $3, $4)", username, firstname, lastname, password)
}

// Approve the Customer's application
func (e *Employee) Approve() {
	//
	fmt.Println()
}

// Deny the Customer's application
func (e *Employee) Deny() {

}
