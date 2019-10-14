package employees

import (
	"log"
	"testing"

	"github.com/gmac220/project-0/opendb"
)

// TestCheckApplication checks if application testing is correct
func TestCheckApplication(t *testing.T) {
	num := 1
	check, numdb := CheckApplication(num)
	if check && numdb == num {
		log.Printf("TestCheckApplication Passed")
	} else {
		log.Fatal("TestCheckApplication Failed. Num Expected:", num, " Actual Num:", numdb, " Check:", check)
	}
}

// TestDeleteApplication verifies if application is actually deleted
func TestDeleteApplication(t *testing.T) {
	var actualnum int
	num := 2

	DeleteApplication(num)
	db := opendb.OpenDB()
	row := db.QueryRow("SELECT acntnumber FROM applications WHERE acntnumber = $1", num)
	row.Scan(&actualnum)
	if actualnum == 0 {
		log.Printf("TestDeleteApplication Passed")
	} else {
		log.Fatal("TestDeleteApplication Failed. Account Number Expected: 0 Actual Account Number:", actualnum)
	}
}

// TestApprove verifies if account is made
func TestApprove(t *testing.T) {
	var uname, acntname string
	usernameinput := "bobt"
	acntnameinput := "testchecking"
	num := 1

	Approve(num)
	db := opendb.OpenDB()
	row := db.QueryRow("SELECT username, acntname FROM accounts WHERE username = $1 AND acntname = $2", usernameinput, acntnameinput)
	row.Scan(&uname, &acntname)
	if usernameinput == uname && acntnameinput == acntname {
		log.Printf("TestApprove Passed")
	} else {
		log.Fatal("TestApprove Failed. Username Input:", usernameinput, " Actual Username", uname,
			" Account Name Input:", acntnameinput, " Actual Account Name:", acntname)
	}

}

// TestApprove2 approves a joint account and verifies if it is made
func TestApprove2(t *testing.T) {
	var uname, acntname string
	usernameinput := "markb"
	acntnameinput := "jointtestcheckings"
	num := 3

	Approve(num)
	db := opendb.OpenDB()
	row := db.QueryRow("SELECT username, acntname FROM accounts WHERE username = $1 AND acntname = $2", usernameinput, acntnameinput)
	row.Scan(&uname, &acntname)
	if usernameinput == uname && acntnameinput == acntname {
		log.Printf("TestApprove Passed")
	} else {
		log.Fatal("TestApprove Failed. Username Input:", usernameinput, " Actual Username", uname,
			" Account Name Input:", acntnameinput, " Actual Account Name:", acntname)
	}

}
