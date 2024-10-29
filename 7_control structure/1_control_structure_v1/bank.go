package main

import (
	"fmt"
	"os"
)

var currentBalance float64 = 1000

func main() {

	fmt.Println("Welcome to Go Bank!")
	for {
		choice := choice()

		if choice == "1" {
			checkBalance()
		} else if choice == "2" {
			currentBalance = deposit(currentBalance)
		} else if choice == "3" {
			currentBalance = withdraw(currentBalance)
		} else {
			exit()
		}
		fmt.Println("Your current balance is ", currentBalance)
	}

}

func checkBalance() {
	fmt.Print("Your current balance is ", currentBalance)
}

func choice() (c string) {
	for {
		fmt.Println("Please choose from one of the following choices by entering the choice number:")
		fmt.Println("1. Check balance.")
		fmt.Println("2. Deposit money.")
		fmt.Println("3. Widthdraw money.")
		fmt.Println("4. Exit.")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&c)

		if c == "1" || c == "2" || c == "3" || c == "4" {
			return
		} else {
			fmt.Println("Invalid choice.")
		}
	}

}

func deposit(currentBalance float64) (updatedbalance float64) {
	var depositamt float64

	updatedbalance = currentBalance

	for {
		fmt.Print("Please input an amount to deposit: ")
		fmt.Scan(&depositamt)

		if depositamt <= 0 {
			fmt.Println("Invalid amount.")
		} else {
			updatedbalance = currentBalance + depositamt
			//println("Your current balance is ", updatedbalance)
			return
		}
	}
}

func withdraw(currentBalance float64) (updatedbalance float64) {
	var widthdrawamt float64

	updatedbalance = currentBalance

	for {
		fmt.Print("Please input an amount to widthdraw: ")
		fmt.Scan(&widthdrawamt)
		if widthdrawamt <= 0 {
			fmt.Println("Invalid amount.")
		} else if widthdrawamt > currentBalance {
			fmt.Println("Insufficient funds.")
			return
		} else {
			updatedbalance = currentBalance - widthdrawamt
			//fmt.Println("Your current balance is ", updatedbalance)
			return
		}
	}
}

func exit() {
	fmt.Print("Good bye.")
	os.Exit(0)
}
