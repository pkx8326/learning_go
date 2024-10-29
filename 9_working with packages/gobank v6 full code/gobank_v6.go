package main

import (
	"fmt"
	"os"
	"strconv"
)

const BalanceFile = "Balance.text"

var currentBalance float64

func main() {
	var nxttxn string
	fmt.Println("Welcome to Go Bank!")
	for {
		currentBalance = readBalance()
		c := menu()
		switch c {
		case "1":
			checkBalance()
			nxttxn = nextTxn()
		case "2":
			currentBalance = deposit(currentBalance)
			nxttxn = nextTxn()
		case "3":
			currentBalance = withdraw(currentBalance)
			nxttxn = nextTxn()
		case "4":
			nxttxn = exit()
		} //end swithc - case
		if nxttxn == "Y" || nxttxn == "y" {
			continue
		} else {
			break
		}
	} // end for loop
}

// ////////////////////////////
func menu() (c string) {
	for {
		fmt.Println("Please choose one of the following choices: ")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")
		fmt.Print("Your choice: ")
		fmt.Scan(&c)
		if c == "1" || c == "2" || c == "3" || c == "4" {
			return
		} else {
			fmt.Println("Invalid input.")
			continue
		}
	}
}

func readBalance() (updatedBalance float64) {
	balanceData, err := os.ReadFile(BalanceFile)
	if err != nil {
		fmt.Println("No balance data found. Please consider making a deposit.")
		return
	} else {
		balanceText := string(balanceData)
		updatedBalance, err = strconv.ParseFloat(balanceText, 64)
		if err != nil {
			fmt.Println("Corrupted balance data found.")
			return
		} else {
			return
		}

	}
} // end function readBalance

func checkBalance() {
	fmt.Println("Your current balance is ", currentBalance)
}

func deposit(currentBalance float64) (updatedBalance float64) {
	var amount float64
	for {
		fmt.Print("Please input an amount to deposit :")
		fmt.Scan(&amount)
		if amount <= 0 {
			fmt.Println("Invalid amount")
			continue
		} else {
			updatedBalance = currentBalance + amount
			fmt.Println("Your curretn balance is ", updatedBalance)
			writeBalance(updatedBalance)
			return
		}
	}
}

func withdraw(currentBalance float64) (updatedBalance float64) {
	var amount float64
	for {
		fmt.Print("Please inut an amount to withdraw :")
		fmt.Scan(&amount)
		if amount <= 0 {
			fmt.Println("Invalid amount")
			continue
		} else if amount > currentBalance {
			fmt.Println("Insufficient funds")
			updatedBalance = currentBalance
			fmt.Println("Your current balance is ", updatedBalance)
			return
		} else {
			updatedBalance = currentBalance - amount
			fmt.Println("Your current balance is ", updatedBalance)
			writeBalance(updatedBalance)
			return
		}
	}
}

func exit() (nxtxn string) {
	fmt.Println("Thank you for choosing our bank. Good bye.")
	nxtxn = "n"
	return
}

func writeBalance(currentBalance float64) {
	balanceText := fmt.Sprint(currentBalance)
	os.WriteFile(BalanceFile, []byte(balanceText), 0644)
}

func nextTxn() (nttxn string) {
	for {
		fmt.Print("Do you have another transaction? [Y/N]: ")
		fmt.Scan(&nttxn)
		if nttxn == "Y" || nttxn == "y" || nttxn == "N" || nttxn == "n" {
			return
		} else {
			fmt.Println("Invalid input")
			continue
		}
	}

}
