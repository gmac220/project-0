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

// Approve the Customer's application
func (e *Employee) Approve() {
	fmt.Println()
}

// Deny the Customer's application
func (e *Employee) Deny() {

}
