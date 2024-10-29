package main

import (
	"fmt"
	"math"
)

func main() {

	amount, years, returnRate, inflationRate := getValue()
	futureVal, adjustedFutureVal := futureValue(amount, years, returnRate, inflationRate)

	fmt.Printf("The future value is : %.2f\n", futureVal)
	fmt.Printf("The adjusted future value according to the inflation rate is : %.2f", adjustedFutureVal)

}

// func futureValue(amt, yr, rR, ifR float64) (float64, float64) {
// 	fv := amt * math.Pow(1+rR/100, yr)
// 	afv := fv / math.Pow(1+ifR/100, yr)
// 	return fv, afv
// }

func getValue() (float64, float64, float64, float64) {
	var amount, years, returnRate, inflationRate float64
	fmt.Println("Please input values of investment amount, years, expected return rate, and inflation rate respectively:")

	fmt.Print("Please enter the investment amount: ")
	fmt.Scan(&amount)

	fmt.Print("Please enter the number of investment years: ")
	fmt.Scan(&years)

	fmt.Print("Please enter the expected return rate (%): ")
	fmt.Scan(&returnRate)

	fmt.Print("Please enter the inflation rate (&) :")
	fmt.Scan(&inflationRate)

	return amount, years, returnRate, inflationRate
}

// alternative value return style
func futureValue(amt, yr, rR, ifR float64) (fv float64, afv float64) {
	fv = amt * math.Pow(1+rR/100, yr)
	afv = fv / math.Pow(1+ifR/100, yr)
	return
}
