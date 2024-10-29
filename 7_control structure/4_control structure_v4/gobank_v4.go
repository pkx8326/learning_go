package main

import "fmt"

var currentBalance float64

func main() {
	var choice, nxttxn string
	for {
		choice = menu()
		switch choice {
		case "1":
			checkBalance()
		case "2":
			currentBalance = deposit(currentBalance)
		case "3":
			currentBalance = withdraw(currentBalance)
		case "4":
			exit()
		} //end switch
		nxttxn = nextTxn()
		if nxttxn == "N" || nxttxn == "n" {
			exit()
			break
		} else {
			continue
		}
	} //end for

} //end main func

////////////////////////////////////

func menu() (c string) {
	for {
		fmt.Println("Please choose from the following choices by inputting the number :")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")
		fmt.Print("Your choice :")
		fmt.Scan(&c)
		if c == "1" || c == "2" || c == "3" || c == "4" {
			return
		} else {
			fmt.Println("Invalid choice.")
			continue
		}
	}
}

func checkBalance() {
	fmt.Println("Your current balance is ", currentBalance)
}

func deposit(currentBalance float64) (updatedBalance float64) {
	var amt float64
	for {
		fmt.Print("Please input an amout to deposit: ")
		fmt.Scan(&amt)
		if amt <= 0 {
			fmt.Println("Invalid amount.")
			continue
		} else {
			break
		}
	} // end for loop for the inpput validation for deposit amount validation
	updatedBalance = currentBalance + amt
	fmt.Println("Your current balance is ", updatedBalance)
	return
}

func withdraw(currentBalance float64) (updatedBalance float64) {
	var amt float64
	for {
		if currentBalance == 0 {
			fmt.Println("Your current balance is 0. Please consider depositing some balances with our bank.")
			return
		} else {
			fmt.Print("Please input an amout to withdraw: ")
			fmt.Scan(&amt)

			if amt > currentBalance {
				fmt.Println("Insufficient funds.")
				updatedBalance = currentBalance
				return
			} else if amt <= 0 {
				fmt.Println("Invalid amount.")
				continue
			} else {
				updatedBalance = currentBalance - amt
				fmt.Println("Your current balance is ", updatedBalance)
				return
			}
		} // end main else
	} // end for loop for input validation
} // end function

func nextTxn() (ntx string) {
	for {
		fmt.Print("Do you have a next transaction? [Y/N]: ")
		fmt.Scan(&ntx)
		if ntx == "Y" || ntx == "y" || ntx == "N" || ntx == "n" {
			return
		} else {
			fmt.Println("Invalid choice.")
			continue
		}
	}
}

func exit() {
	fmt.Println("Thank you for choosing our bank. Good bye.")
}
