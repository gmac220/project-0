package employees

import (
	"fmt"
	"log"
	"os"

	"github.com/gmac220/project-0/opendb"
)

// EmployeePage prompts the employee what they could do and asks them to pick a choice.
func EmployeePage() {
	var num int
	var acntnumber int
	var username string

	fmt.Println("What do you want to do")
	fmt.Println("1: Approve")
	fmt.Println("2: Deny")
	fmt.Println("3: View Customer Info")
	fmt.Println("4: Show Applications")
	fmt.Println("5: Exit")
	fmt.Printf("Please type in a number: ")
	fmt.Scanln(&num)

	switch num {
	case 1:
		fmt.Printf("Which customer's application do you want to approve?" +
			" (Please input application number): ")
		fmt.Scanln(&acntnumber)
		Approve(acntnumber)
		EmployeePage()
	case 2:
		fmt.Printf("Which customer's application do you want to deny?" +
			" (Please input application number): ")
		fmt.Scanln(&acntnumber)
		if CheckApplication(acntnumber) {
			DeleteApplication(acntnumber)
			fmt.Println("Application Denied :(")
		} else {
			fmt.Println("Application does not exist.")
		}
		EmployeePage()

	case 3:
		fmt.Printf("Which customer's information do you want to look at?" +
			" (Please input customer's username): ")
		fmt.Scanln(&username)
		CustomerInfo(username)
	case 4:
		Applications()
	case 5:
		os.Exit(0)
		// default:
		// 	fmt.Println("Choice does not exist.")
		// 	EmployeePage()
	}
}

// Approve the Customer's application
func Approve(num int) {
	var uname string
	var acntname string
	var joint bool
	var uname2 string

	db := opendb.OpenDB()
	defer db.Close()
	if CheckApplication(num) {
		row := db.QueryRow("SELECT username, acntname, joint, username2 FROM applications WHERE acntnumber = $1", num)
		row.Scan(&uname, &acntname, &joint, &uname2)
		if joint {
			db.Exec("INSERT INTO accounts (acntname, balance, username, username2) VALUES ($1, $2, $3, $4)", "joint"+acntname, 0, uname, uname2)
		} else {
			db.Exec("INSERT INTO accounts (acntname, balance, username) VALUES ($1, $2, $3)", acntname, 0, uname)
		}
		fmt.Println("Application Approved!")
		DeleteApplication(num)
	} else {
		fmt.Println("Application does not exist.")
	}
}

// CheckApplication verifies if application exists
func CheckApplication(num int) bool {
	var acntnumber int
	db := opendb.OpenDB()
	defer db.Close()
	row := db.QueryRow("SELECT acntnumber FROM applications WHERE acntnumber = $1", num)
	row.Scan(&acntnumber)
	return acntnumber == num
}

// DeleteApplication deletes row from applications table
func DeleteApplication(num int) {
	db := opendb.OpenDB()
	defer db.Close()
	db.Exec("DELETE FROM applications WHERE acntnumber = $1", num)
}

// CustomerInfo looks at all of customers account information by passing in their username
func CustomerInfo(username string) {
	var acntnumber int
	var acntname string
	var balance float64
	var uname string
	var uname2 string
	var pw string
	var fname string
	var lname string
	var otheruname string

	db := opendb.OpenDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM accounts FULL OUTER JOIN customers on customers.username = accounts.username OR customers.username = accounts.username2 WHERE customers.username = $1", username)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("---------------------------" + username + "'s INFORMATION------------------------------")
	fmt.Println()
	for rows.Next() {
		rows.Scan(&acntnumber, &acntname, &balance, &uname, &uname2, &otheruname, &pw, &fname, &lname)
		if uname2 == username {
			fmt.Println("Account #:", acntnumber, "|Account Name:", acntname, "|Balance:", balance, "|Other Account Holder:", uname)
		} else if username == uname && uname2 != "" {
			fmt.Println("Account #:", acntnumber, "|Account Name:", acntname, "|Balance:", balance, "|Other Account Holder:", uname2)
		} else if uname != "" || uname2 != "" {
			fmt.Println("Account #:", acntnumber, "|Account Name:", acntname, "|Balance:", balance)
		} else {
			fmt.Println()
		}
	}
	fmt.Println()
	fmt.Println("---------------------------------------------------------------------------------")
	EmployeePage()
}

// Applications loops through applications table listing if they are joint applications or not
func Applications() {
	var acntnum int
	var acntname string
	var uname string
	var fname string
	var lname string
	var joint bool
	var uname2 string
	var fname2 string
	var lname2 string

	db := opendb.OpenDB()
	defer db.Close()
	fmt.Println("----------------------LISTED APPLICATIONS----------------------")
	fmt.Println()
	rows, err := db.Query("SELECT * FROM applications")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		rows.Scan(&acntnum, &uname, &fname, &lname, &acntname, &joint, &uname2, &fname2, &lname2)
		if joint {
			fmt.Println("Account #:", acntnum, "|Type: Joint", "|Username1:", uname,
				"|First name:", fname, "|Last name:", lname, "|Username2:", uname2,
				"|First name:", fname2, "|Last name", lname2)
		} else {
			fmt.Println("Account #:", acntnum, "|Type: Solo ", "|Username:", uname,
				" |First name:", fname, "|Last name:", lname)
		}
	}
	fmt.Println()
	fmt.Println("---------------------------------------------------------------")
	EmployeePage()
}
