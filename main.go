package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	//"golang.org/x/crypto/ssh/terminal"
)

func main() {
	var pass string
	fmt.Println("Welcome to Bills FarGO")
	fmt.Printf("Enter password: ")

	// using stty to hide the password so it doesn't show up when the user types
	cmd := exec.Command("stty", "-echo")
	cmd.Stdin = os.Stdin
	_, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Scanln(&pass)
	cmd = exec.Command("stty", "echo")
	cmd.Stdin = os.Stdin
	_, err = cmd.Output()
	cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()

}
