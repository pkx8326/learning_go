package main

import "fmt"

func main() {
	var revenue, expenses, taxRate, ebt, profit, ratio float64
	fmt.Print("Please enter the revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Please enter the expense: ")
	fmt.Scan(&expenses)

	fmt.Print("Please enter the tax rate (%): ")
	fmt.Scan(&taxRate)

	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)

	ratio = ebt / profit

	//using `` to print multi-line strings
	//using fmt.Sprint() to generate the said multi-line string
	result := fmt.Sprintf(`
Earning before tax (EBT) is : %.2f
The profit is %.2f
EBT/profit ratio is %.2f
	`, ebt, profit, ratio)

	fmt.Print(result)
}
