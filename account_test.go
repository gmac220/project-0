package main

import (
	"log"
	"strconv"
	"testing"
)

func Test_NewAccount(t *testing.T) {
	account := NewAccount("Savings", 100.0)

	if account.accountname == "Savings" && account.balance == 100.0 {
		log.Printf("Successfully made" + account.accountname + "account with balance of" + strconv.FormatFloat(account.balance, 'f', 2, 64))
	} else {
		t.Errorf("Expected 100.0 but got: " + strconv.FormatFloat(account.balance, 'f', 2, 64))
	}

}

func Test_Balance(t *testing.T) {
	account := NewAccount("Savings", 300.0)

	if account.Balance() == 300.0 {
		log.Printf("Account has balance of " + strconv.FormatFloat(account.balance, 'f', 2, 64))
	} else {
		t.Errorf("Expected 300.0 but got: " + strconv.FormatFloat(account.balance, 'f', 2, 64))
	}
}

func Test_Withdraw(t *testing.T) {
	account := NewAccount("Savings", 300.0)

	account.Withdraw(144.0)
	if account.balance == 156.0 {
		log.Printf("Successfully withdrew 144.0")
	} else {
		t.Errorf("Expected 156.0 but got: " + strconv.FormatFloat(account.balance, 'f', 2, 64))
	}
}

func Test_Deposit(t *testing.T) {
	account := NewAccount("Savings", 0.0)

	account.Deposit(222.0)
	if account.balance == 222.0 {
		log.Printf("Successfully deposited 222.0")
	} else {
		t.Errorf("Expected 222.0 but got: " + strconv.FormatFloat(account.balance, 'f', 2, 64))
	}
}
