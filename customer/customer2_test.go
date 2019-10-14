package customer

import (
	"log"
	"testing"

	"github.com/gmac220/project-0/opendb"
)

// TestVerifyAccount
func TestVerifyAccount(t *testing.T) {
	num := 1
	balanceexpected := 0.0

	balance, _ := VerifyAccount(num)
	if balanceexpected == balance {
		log.Printf("TestVerifyAccount Passed")
	} else {
		log.Fatal("TestVerifyAccount failed. Balance Expected:", balanceexpected, " Actual Balance:", balance)
	}
}

// TestCheckOwnAccount assumes that the account has been approved and account number is the first one.
func TestCheckOwnAccount(t *testing.T) {
	num := 1
	username := "bobt"
	username2 := "bobt"

	check, uname, uname2 := CheckOwnAccount(num)
	if check && (username == uname || username2 == uname2) {
		log.Printf("TestCheckOwnAccount Passed")
	} else {
		log.Fatal("Input does not match. Username input:", username, " Actual Username:", uname,
			" Username2:", username2, " Actual username2:", uname2, " Check:", check)
	}
}

func TestDeposit(t *testing.T) {
	var balance float64
	num := 1
	expectedinput := 100.0
	balanceinput := 0.0

	Deposit(num, expectedinput, balanceinput)
	db := opendb.OpenDB()
	row := db.QueryRow("SELECT balance FROM accounts WHERE acntnumber = $1", num)
	row.Scan(&balance)
	if balance == expectedinput {
		log.Printf("TestDeposit Passed")
	} else {
		log.Fatal("TestDeposit failed. Amount deposited:", expectedinput, " Actual balance:", balance)
	}
}

// TestShowBalance checks if account number passed in has balance.
func TestShowBalance(t *testing.T) {
	num := 1
	balanceexpected := 100.0

	balance := ShowBalance(num)
	if balanceexpected == balance {
		log.Printf("TestShow Passed")
	} else {
		log.Fatal("TestShowBalance failed. Balance Expected:", balanceexpected, " Actual Balance:", balance)
	}
}

func TestWithdraw(t *testing.T) {
	var balance float64
	num := 1
	expectedinput := 50.0
	balanceinput := 100.0

	Withdraw(num, expectedinput, balanceinput)
	db := opendb.OpenDB()
	row := db.QueryRow("SELECT balance FROM accounts WHERE acntnumber = $1", num)
	row.Scan(&balance)
	if expectedinput == balance {
		log.Printf("TestWithdraw Passed")
	} else {
		log.Fatal("TestWithdraw failed. Amount withdrawn:", expectedinput, " Actual balance:", balance)
	}
}

func TestTransfer(t *testing.T) {
	var withdrawnbalance, depositedbalance float64

	numwithdraw := 1
	numdeposit := 2
	expectedinput := 25.0
	balanceinput := 50.0

	Transfer(numwithdraw, numdeposit, balanceinput, expectedinput)
	db := opendb.OpenDB()
	row := db.QueryRow("SELECT balance FROM accounts WHERE acntnumber = $1", numwithdraw)
	row.Scan(&withdrawnbalance)
	row = db.QueryRow("SELECT balance FROM accounts WHERE acntnumber = $1", numdeposit)
	row.Scan(&depositedbalance)
	if withdrawnbalance == expectedinput && expectedinput == depositedbalance {
		log.Printf("TestTransfer Passed")
	} else {
		log.Fatal("TestTransfer failed. Amount withdrawn:", expectedinput,
			" Actual withdrawn balance:", withdrawnbalance, " Amount deposited:", expectedinput,
			" Actual deposited balance:", depositedbalance)
	}
}
