package main

//goals
//1. validate user imputs
//2. read in tax rate from file
//3. store calculated result on file

import (
	"fmt"
	"os"
	"strconv"
)

const taxratefile = "taxrate.txt"
const profitfile = "profit.txt"

func main() {
	revenue := getRevenue()
	expenses := getExpenses()
	taxrate := readTaxRate()
	ebt := revenue - expenses
	profit := ebt * (1 - taxrate/100)
	ratio := ebt / profit

	fmt.Println("Your earning before tax is : ", ebt)
	fmt.Println("Set tax rate is :", taxrate)
	fmt.Println("Your profit is :", profit)
	fmt.Printf("Your ratio is : %.2f \n", ratio)

	writeProfit(profit)

}

// ///////////
func getRevenue() (rev float64) {
	for {
		fmt.Print("Please input your revenue: ")
		fmt.Scan(&rev)
		if rev <= 0 {
			fmt.Println("Invalid input.")
			continue
		} else {
			return
		}
	} // end for loop for input validation
}

func getExpenses() (exps float64) {
	for {
		fmt.Print("Please input your expenses: ")
		fmt.Scan(&exps)
		if exps <= 0 {
			fmt.Println("Invalid input.")
			continue
		} else {
			return
		}
	} // end for loop for input validation
}

func readTaxRate() (taxrate float64) {
	taxRateData, err := os.ReadFile(taxratefile)
	if err != nil {
		fmt.Println("Tax rate data file not found!")
		return
	} else {
		taxRateText := string(taxRateData)
		taxrate, err = strconv.ParseFloat(taxRateText, 64)
		if err != nil {
			fmt.Println("Can't get tax rate data!")
			return
		}
		return
	}

}

func writeProfit(profit float64) {
	profitText := fmt.Sprint(profit)
	os.WriteFile(profitfile, []byte(profitText), 0644)
}
