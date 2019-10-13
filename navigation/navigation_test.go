package navigation

import (
	"log"
	"testing"

	"github.com/gmac220/project-0/opendb"
)

func TestEmployeeCreateAccount(t *testing.T) {
	var fname string
	var lname string
	var uname string
	var pw string
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
		log.Fatal("Inputs do not match. Firstname input:", firstnameinput, "Actual firstname:",
			fname, "Lastname input:", lastnameinput, "Actual lastname:", lname,
			"Username input:", usernameinput, "Actual username:", uname,
			"Password input:", passwordinput, "Actual password:", pw)
	}

}

func TestEmployeeSignIn(t *testing.T) {
	firstnameinput := "Ro"
	lastnameinput := "Bert"
	usernameinput := "robert"
	passwordinput := "r"

	usernamedb, passdb, fname, lname := SignIn("robert", "r", true)
	if firstnameinput == fname && lastnameinput == lname && usernameinput == usernamedb && passwordinput == passdb {
		log.Printf("TestEmployeeSignIn Passed")
	} else {
		log.Fatal("Inputs do not match. Firstname input:", firstnameinput, "Actual firstname:",
			fname, "Lastname input:", lastnameinput, "Actual lastname:", lname,
			"Username input:", usernameinput, "Actual username:", usernamedb,
			"Password input:", passwordinput, "Actual password:", passdb)
	}
}

func TestCustomerCreateAccount(t *testing.T) {
	var fname string
	var lname string
	var uname string
	var pw string
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
		log.Fatal("Inputs do not match. Firstname input:", firstnameinput, "Actual firstname:",
			fname, "Lastname input:", lastnameinput, "Actual lastname:", lname,
			"Username input:", usernameinput, "Actual username:", uname,
			"Password input:", passwordinput, "Actual password:", pw)
	}
}

func TestCustomerSignIn(t *testing.T) {
	firstnameinput := "Bob"
	lastnameinput := "Truss"
	usernameinput := "bobt"
	passwordinput := "b"

	usernamedb, passdb, fname, lname := SignIn("bobt", "b", false)
	if firstnameinput == fname && lastnameinput == lname && usernameinput == usernamedb && passwordinput == passdb {
		log.Printf("TestCustomerSignIn Passed")
	} else {
		log.Fatal("Inputs do not match. Firstname input:", firstnameinput, "Actual firstname:",
			fname, "Lastname input:", lastnameinput, "Actual lastname:", lname,
			"Username input:", usernameinput, "Actual username:", usernamedb,
			"Password input:", passwordinput, "Actual password:", passdb)
	}
}
