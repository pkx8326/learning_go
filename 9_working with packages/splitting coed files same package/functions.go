package main

import (
	"fmt"
	"os"
	"strconv"
)

// /////////////////////////////////////////////////////////////////////////////////////////
func readBalance() (updatedBalance float64) {
	balanceData, err := os.ReadFile(BalanceFile)
	if err != nil {
		fmt.Println("No balance data found. Please consider making a deposit with us.")
		return
	} else {
		balanceText := string(balanceData)
		updatedBalance, err = strconv.ParseFloat(balanceText, 64)
		if err != nil {
			fmt.Println("Data corruption detected. THe last valid balance is returned.")
			return
		} else {
			return
		}
	}
}

func writeBalance(updatedBalance float64) {
	balanceText := fmt.Sprint(updatedBalance)
	os.WriteFile(BalanceFile, []byte(balanceText), 0644)
}

func menu() (c string) {
	for {
		fmt.Println("Please choose from the following choices:")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")
		fmt.Scan(&c)
		if c == "1" || c == "2" || c == "3" || c == "4" {
			return
		} else {
			fmt.Println("Invalid choice")
			continue
		}
	}
}

func checkBalance(currentBalance float64) {
	updatedBalance := currentBalance
	fmt.Println("Your current balance is ", updatedBalance)
}

func deposit(currentBalance float64) (updatedBalance float64) {
	var amount float64
	for {
		fmt.Print("Please input an amount to deposit: ")
		fmt.Scan(&amount)
		if amount <= 0 {
			fmt.Println("Invalid amount")
			continue
		} else {
			updatedBalance = currentBalance + amount
			fmt.Println("Your current balance is ", updatedBalance)
			writeBalance(updatedBalance)
			return
		}
	}
}

func withdraw(currentBalance float64) (updatedBalance float64) {
	var amount float64
	for {
		fmt.Print("Please input an amount to withdraw: ")
		fmt.Scan(&amount)
		if amount <= 0 {
			fmt.Println("Invalid amount")
			continue
		} else if amount > currentBalance {
			fmt.Println("Insufficient funds")
			return
		} else {
			updatedBalance = currentBalance - amount
			fmt.Println("Your current balance is", updatedBalance)
			writeBalance(updatedBalance)
			return
		}
	}
}

func exit() (nxttxn string) {
	fmt.Println("Thank you for choosing our bank. Good bye.")
	nxttxn = "n"
	return
}

func nexttxn() (nxttxn string) {
	var ntxn string
	for {
		fmt.Print("Do you have another transaaction? [Y/N]: ")
		fmt.Scan(&ntxn)
		if ntxn == "Y" || ntxn == "y" || ntxn == "N" || ntxn == "n" {
			nxttxn = ntxn
			return
		} else {
			fmt.Println("Invalid input")
			continue
		}
	}
}
