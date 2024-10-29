package main

import (
	"fmt"

	"example.com/gobankv7-packages/fileops"
)

const BalanceFile = "Balance.txt"

var currentBalance float64

func main() {
	var nxttxn string
	for {
		fmt.Println("Welcome to Go Bank.")
		currentBalance = fileops.ReadFloat(BalanceFile)
		c := menu()
		switch c {
		case "1":
			checkBalance(currentBalance)
			nxttxn = nexttxn()
		case "2":
			deposit(currentBalance)
			nxttxn = nexttxn()
		case "3":
			withdraw(currentBalance)
			nxttxn = nexttxn()
		case "4":
			nxttxn = exit()
		}
		if nxttxn == "Y" || nxttxn == "y" {
			continue
		} else {
			break
		}
	}
}
