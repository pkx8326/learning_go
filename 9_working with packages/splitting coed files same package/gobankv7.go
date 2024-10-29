package main

import (
	"fmt"
)

const BalanceFile = "Balance.txt"

var currentBalance float64

func main() {
	var nxttxn string
	for {
		fmt.Println("Welcome to Go Bank.")
		currentBalance = readBalance()
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
