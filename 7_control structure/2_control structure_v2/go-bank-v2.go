package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	var currentBalance float64 = 1000
	var amt float64
	var c, newTransaction string

	for {
		//printing menu choices
		for {

			ClearScreen()

			fmt.Println("Welcom to Go Bank V2. Please choose from the following choices by input the number.")
			fmt.Println("1. Check balance.")
			fmt.Println("2. Deposit.")
			fmt.Println("3. Withdraw.")
			fmt.Println("4. Exit.")
			fmt.Print("Your choice: ")
			fmt.Scan(&c)
			if c == "1" || c == "2" || c == "3" || c == "4" {
				break
			} else {
				fmt.Println("Invalid choice.")
				continue
			}
		} //end for loop for main choices

		if c == "1" {
			fmt.Println("Your current balance is: ", currentBalance)
			for {
				fmt.Print("New transaction? [Y/N]: ")
				fmt.Scan(&newTransaction)
				if newTransaction == "Y" || newTransaction == "N" || newTransaction == "y" || newTransaction == "n" {
					break
				} else {
					continue
				}
			} // end for loop for new transaction?
			if newTransaction == "Y" || newTransaction == "y" {
				continue
			} else {
				fmt.Println("Thank you for choosing our bank. Good bye.")
				break
			}
		} else if c == "2" {
			for {
				fmt.Print("Please input an amount to deposit: ")
				fmt.Scan(&amt)
				if amt <= 0 {
					println("Invalid amount.")
					continue
				} else {
					break
				}
			} // end for loop for entering deposit amount
			currentBalance += amt
			fmt.Println("Your current balance is: ", currentBalance)
			for {
				fmt.Print("New transaction? [Y/N]: ")
				fmt.Scan(&newTransaction)
				if newTransaction == "Y" || newTransaction == "N" || newTransaction == "y" || newTransaction == "n" {
					break
				} else {
					continue
				}
			} // end for loop for new transaction?
			if newTransaction == "Y" || newTransaction == "y" {
				continue
			} else {
				fmt.Println("Thank you for choosing our bank. Good bye.")
				break
			}
		} else if c == "3" {
			for {
				fmt.Print("Pleaes input an amount to withdraw: ")
				fmt.Scan(&amt)
				if amt <= 0 {
					fmt.Println("Invalid amount.")
					continue
				} else if amt > currentBalance {
					fmt.Println("Insuffient funds.")
					continue
				} else {
					break
				}
			} // end for loop for entering withdrawal amount
			currentBalance -= amt
			fmt.Println("Your current balance is: ", currentBalance)
			for {
				fmt.Print("New transaction? [Y/N]: ")
				fmt.Scan(&newTransaction)
				if newTransaction == "Y" || newTransaction == "N" || newTransaction == "y" || newTransaction == "n" {
					break
				} else {
					continue
				}
			} // end for loop for new transaction?
			if newTransaction == "Y" || newTransaction == "y" {
				continue
			} else {
				fmt.Println("Thank you for choosing our bank. Good bye.")
				break
			}
		} else {
			fmt.Print("Thank you for choosing our bank. Good bye.")
			break
		}
	} //main for loop
} //func main

// This function clears the terminal screen.
func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
