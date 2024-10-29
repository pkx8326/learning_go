package main

import "fmt"

var currentBalance float64

func main() {
	var nextTxn string

	fmt.Println("Welcome to Go Bank v. 3.0")
	for {
		c := menu()
		if c == "1" {
			checkBalance()
			nextTxn = nextTransaction()
			if nextTxn == "Y" || nextTxn == "y" {
				continue
			} else {
				fmt.Println("Thank you for choosing our bank. Good bye.")
				break
			}
		} else if c == "2" {
			currentBalance = deposit(currentBalance)
			nextTxn = nextTransaction()
			if nextTxn == "Y" || nextTxn == "y" {
				continue
			} else {
				fmt.Println("Thank you for choosing our bank. Good bye.")
				break
			}
		} else if c == "3" {
			currentBalance = withdraw(currentBalance)
			nextTxn = nextTransaction()
			if nextTxn == "Y" || nextTxn == "y" {
				continue
			} else {
				fmt.Println("Thank you for choosing our bank. Good bye.")
				break
			}
		} else {
			fmt.Println("Thank you for choosing our bank. Good bye.")
			break
		}

	} //end loop of main program
} //end main function

//////////////////////////////////////////////////////////////////////////////////////////

func menu() (c string) {
	for {
		fmt.Println("Please choose between the following choices by entering the number:")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")
		fmt.Print("Your choice: ")
		fmt.Scan(&c)

		if c == "1" || c == "2" || c == "3" || c == "4" {
			break
		} else {
			fmt.Println("Invalid choice.")
			continue
		}
	}
	return
}

func checkBalance() {
	fmt.Println("Your current balance is ", currentBalance)
}

func deposit(currentBalance float64) (updatedBalance float64) {
	var amt float64
	for {
		fmt.Print("Please input an amount to deposit : ")
		fmt.Scan(&amt)
		if amt <= 0 {
			fmt.Println("Invalid amount.")
			continue
		} else {
			break
		} // for loop for deposit amount validation
	}
	updatedBalance = currentBalance + amt
	fmt.Println("Your current balance is ", updatedBalance)
	return
}

func withdraw(currentBalance float64) (updatedBalance float64) {
	var amt float64
	if currentBalance == 0 {
		fmt.Println("Your current balance is ", currentBalance, ". Please consider making a deposit.")
		return
	} else {
		for {
			fmt.Print("Please input an amount to withdraw :")
			fmt.Scan(&amt)

			if amt <= 0 {
				fmt.Println("Invalid amount.")
				continue
			} else if amt > currentBalance {
				fmt.Println("Insufficient funds.")
				continue
			} else {
				break
			}
		} // for loop for withdraw amount validation
	}
	updatedBalance = currentBalance - amt
	fmt.Println("Your current balance is ", updatedBalance)
	return
}

func nextTransaction() (newTxn string) {
	for {
		fmt.Print("Do you have another transaction? [Y/N] :")
		fmt.Scan(&newTxn)
		if newTxn == "Y" || newTxn == "y" || newTxn == "N" || newTxn == "n" {
			return
		} else {
			fmt.Println("Invalid choice.")
			continue
		}
	}

}
