package navigation

import (
	"log"
	"testing"

	"github.com/gmac220/project-0/opendb"
)

// TestEmployeeCreateAccount tests creating an employee.
// Queries the database to see if the values passed in are inside.
func TestEmployeeCreateAccount(t *testing.T) {
	//GOCACHE="/home/godfrey/.cache/go-build"
	var fname, lname, uname, pw string
	firstnameinput := "Ro"
	lastnameinput := "Bert"
	usernameinput := "robert"
	passwordinput := "r"

	CreateAccount(firstnameinput, lastnameinput, usernameinput, passwordinput, "e")
	db := opendb.OpenDB()
	row := db.QueryRow("SELECT * FROM employees WHERE username = $1", usernameinput)
	row.Scan(&uname, &pw, &fname, &lname)
	if firstnameinput == fname && lastnameinput == lname && usernameinput == uname && passwordinput == pw {
		log.Printf("TestEmployeeCreateAccount Passed")
	} else {
		log.Fatal("Inputs do not match. Firstname input:", firstnameinput, " Actual firstname:",
			fname, " Lastname input:", lastnameinput, " Actual lastname:", lname,
			" Username input:", usernameinput, " Actual username:", uname,
			" Password input:", passwordinput, " Actual password:", pw)
	}

}

// TestEmployeeSignIn tests signing in an employee account using the created account inputs.
func TestEmployeeSignIn(t *testing.T) {
	firstnameinput := "Ro"
	lastnameinput := "Bert"
	usernameinput := "robert"
	passwordinput := "r"

	usernamedb, passdb, fname, lname := SignIn(usernameinput, passwordinput, true)
	if firstnameinput == fname && lastnameinput == lname && usernameinput == usernamedb && passwordinput == passdb {
		log.Printf("TestEmployeeSignIn Passed")
	} else {
		log.Fatal("Inputs do not match. Firstname input:", firstnameinput, " Actual firstname:",
			fname, " Lastname input:", lastnameinput, " Actual lastname:", lname,
			" Username input:", usernameinput, " Actual username:", usernamedb,
			" Password input:", passwordinput, " Actual password:", passdb)
	}
}

//TestCustomerCreateAccount tests creating an employee.
// Queries the database to see if the values passed in are inside.
func TestCustomerCreateAccount(t *testing.T) {
	var fname, lname, uname, pw string
	firstnameinput := "Bob"
	lastnameinput := "Truss"
	usernameinput := "bobt"
	passwordinput := "b"

	CreateAccount(firstnameinput, lastnameinput, usernameinput, passwordinput, "c")
	db := opendb.OpenDB()
	row := db.QueryRow("SELECT * FROM customers WHERE username = $1", usernameinput)
	row.Scan(&uname, &pw, &fname, &lname)
	if firstnameinput == fname && lastnameinput == lname && usernameinput == uname && passwordinput == pw {
		log.Printf("TestCustomerCreateAccount Passed")
	} else {
		log.Fatal("Inputs do not match. Firstname input:", firstnameinput, " Actual firstname:",
			fname, " Lastname input:", lastnameinput, " Actual lastname:", lname,
			" Username input:", usernameinput, " Actual username:", uname,
			" Password input:", passwordinput, " Actual password:", pw)
	}
}

// TestCustomerSignIn tests signing in a customer account using the created account inputs.
func TestCustomerSignIn(t *testing.T) {
	firstnameinput := "Bob"
	lastnameinput := "Truss"
	usernameinput := "bobt"
	passwordinput := "b"

	usernamedb, passdb, fname, lname := SignIn(usernameinput, passwordinput, false)
	if firstnameinput == fname && lastnameinput == lname && usernameinput == usernamedb && passwordinput == passdb {
		log.Printf("TestCustomerSignIn Passed")
	} else {
		log.Fatal("Inputs do not match. Firstname input:", firstnameinput, " Actual firstname:",
			fname, " Lastname input:", lastnameinput, " Actual lastname:", lname,
			" Username input:", usernameinput, " Actual username:", usernamedb,
			" Password input:", passwordinput, " Actual password:", passdb)
	}
}

// TestSecondCustomerCreateAccount makes a second account for testing purposes later.
func TestSecondCustomerCreateAccount(t *testing.T) {
	var fname, lname, uname, pw string
	firstnameinput := "Mark"
	lastnameinput := "Bo"
	usernameinput := "markb"
	passwordinput := "b"

	CreateAccount(firstnameinput, lastnameinput, usernameinput, passwordinput, "c")
	db := opendb.OpenDB()
	row := db.QueryRow("SELECT * FROM customers WHERE username = $1", usernameinput)
	row.Scan(&uname, &pw, &fname, &lname)
	if firstnameinput == fname && lastnameinput == lname && usernameinput == uname && passwordinput == pw {
		log.Printf("TestCustomerCreateAccount Passed")
	} else {
		log.Fatal("Inputs do not match. Firstname input:", firstnameinput, " Actual firstname:",
			fname, " Lastname input:", lastnameinput, " Actual lastname:", lname,
			" Username input:", usernameinput, " Actual username:", uname,
			" Password input:", passwordinput, " Actual password:", pw)
	}
}

// TestSecondCustomerSignIn tests signing in a customer account using the created account inputs.
func TestSecondCustomerSignIn(t *testing.T) {
	firstnameinput := "Mark"
	lastnameinput := "Bo"
	usernameinput := "markb"
	passwordinput := "b"

	usernamedb, passdb, fname, lname := SignIn(usernameinput, passwordinput, false)
	if firstnameinput == fname && lastnameinput == lname && usernameinput == usernamedb && passwordinput == passdb {
		log.Printf("TestCustomerSignIn Passed")
	} else {
		log.Fatal("Inputs do not match. Firstname input:", firstnameinput, " Actual firstname:",
			fname, " Lastname input:", lastnameinput, " Actual lastname:", lname,
			" Username input:", usernameinput, " Actual username:", usernamedb,
			" Password input:", passwordinput, " Actual password:", passdb)
	}
}
