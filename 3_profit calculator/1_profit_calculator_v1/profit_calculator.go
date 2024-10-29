package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var taxrate float64
	var expenses float64

	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter the tax rate (%): ")
	fmt.Scan(&taxrate)

	fmt.Print("Enter the expenses: ")
	fmt.Scan(&expenses)

	fmt.Println("Calculating profit...")
	fmt.Print("The profit is ")

	earningBeforeTax := revenue - expenses //declare the inferred earningBeforeTax variable
	fmt.Printf("Earning before tax is: %.2f\n", earningBeforeTax)

	profit := earningBeforeTax * (1 - taxrate/100)
	fmt.Printf("Earning after tax (Profit) is: %.2f\n", profit)

	ratio := earningBeforeTax / profit
	fmt.Printf("The ratio (EBT/Profit) is: %.2f", ratio)
}
