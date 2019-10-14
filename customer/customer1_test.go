package customer

import (
	"log"
	"testing"

	"github.com/gmac220/project-0/opendb"
)

// TestSetCustomerVars passes in values that check to see if global variables in customers get assigned.
func TestSetCustomerVars(t *testing.T) {
	usernameinput := "bobt"
	firstnameinput := "Bob"
	lastnameinput := "Truss"

	uname, fname, lname := SetCustomerVars(usernameinput, firstnameinput, lastnameinput)
	if usernameinput == uname && firstnameinput == fname && lastnameinput == lname {
		log.Printf("TestSetCustomerVars Passed")
	} else {
		log.Fatal("Inputs returned are incorrect. Username Input:", usernameinput,
			" Actual Username:", uname, " Firstname Input:", firstnameinput, " Actual Firstname:",
			fname, " Lastname Input:", lastnameinput, " Actual lastname:", lname)
	}

}

// TestApply passes in inputs that would mimic applying to an account.
func TestApply(t *testing.T) {
	var uname, fname, lname, acntname string

	usernameinput := "bobt"
	firstnameinput := "Bob"
	lastnameinput := "Truss"
	// acntnameinput := "checking"
	acntnameinput := "testchecking"

	Apply(usernameinput, firstnameinput, lastnameinput, acntnameinput)
	db := opendb.OpenDB()
	row := db.QueryRow("SELECT username, firstname, lastname, acntname FROM applications WHERE username = $1", usernameinput)
	row.Scan(&uname, &fname, &lname, &acntname)
	if firstnameinput == fname && lastnameinput == lname && usernameinput == uname && acntnameinput == acntname {
		log.Printf("TestApply Passed")
	} else {
		log.Fatal("Inputs do not match. Firstname input:", firstnameinput, " Actual firstname:",
			fname, " Lastname input:", lastnameinput, " Actual lastname:", lname,
			" Username input:", usernameinput, " Actual username:", uname,
			" Accountname input:", acntnameinput, " Actual accountname:", acntname)
	}
}

// TestJointApply passes in inputs that would mimic applying to a joint account.
func TestJointApp(t *testing.T) {
	var uname, uname2, fname, lname, acntname string

	usernameinput := "bobt"
	firstnameinput := "Bob"
	lastnameinput := "Truss"
	// acntnameinput := "savings"
	acntnameinput := "testsavings"
	usernameinput2 := "markb"

	JointApp(usernameinput, firstnameinput, lastnameinput, acntnameinput, usernameinput2)
	db := opendb.OpenDB()
	row := db.QueryRow("SELECT username, firstname, lastname, acntname, username2 FROM applications WHERE username = $1 AND username2 = $2", usernameinput, usernameinput2)
	row.Scan(&uname, &fname, &lname, &acntname, &uname2)
	if firstnameinput == fname && lastnameinput == lname && usernameinput == uname && acntnameinput == acntname {
		log.Printf("TestJointApp Passed")
	} else {
		log.Fatal("Inputs do not match. Firstname input:", firstnameinput, " Actual firstname:",
			fname, " Lastname input:", lastnameinput, " Actual lastname:", lname,
			" Username input:", usernameinput, " Actual username:", uname,
			" Accountname input:", acntnameinput, " Actual accountname:", acntname,
			" Username2 input:", usernameinput2, " Actual Username2:", uname2)
	}
}

// TestJointApp2 tests if username2 is the same as username2 input
func TestJointApp2(t *testing.T) {
	var uname2 string

	firstnameinput := "Mark"
	lastnameinput := "Bo"
	usernameinput := "markb"
	// acntnameinput := "checkings"
	acntnameinput := "testcheckings"
	usernameinput2 := "bobt"

	JointApp(usernameinput, firstnameinput, lastnameinput, acntnameinput, usernameinput2)
	db := opendb.OpenDB()
	row := db.QueryRow("SELECT username2 FROM applications WHERE username2 = $1", usernameinput2)
	row.Scan(&uname2)
	if usernameinput2 == uname2 {
		log.Printf("TestJointApp2 Passed")
	} else {
		log.Fatal("Input does not match. Username2 input:", usernameinput2, " Actual Username2:", uname2)
	}
}

// TestCheckCustomer verifies if that customer exists on not
func TestCheckCustomer(t *testing.T) {

	usernameinput := "markb"
	firstname := "Mark"
	lastname := "Bo"

	uname, fname, lname := CheckCustomer(usernameinput)
	if usernameinput == uname && firstname == fname && lastname == lname {
		log.Printf("TestCheckCustomer Passed")
	} else {
		log.Fatal("Input does not match. Username input:", usernameinput, " Actual Username:", uname,
			" Firstname:", firstname, " Actual firstname:", fname, " Lastname:", lastname,
			" Actual lastname:", lname)
	}
}
