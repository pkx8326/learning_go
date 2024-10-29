package main

import "fmt"

func main() {
	age := float64(16)
	fmt.Println("age = ", age)
	fmt.Println("age's memory address = ", &age)
	add(&age)

}

func add(age *float64) {
	var yearsToLegalAge float64
	fmt.Println("In-memory age's address in function = ", &age)
	if *age >= 18 {
		fmt.Println("Legal age reached.")
	} else {
		yearsToLegalAge = 18 - *age
		fmt.Println("There are still ", yearsToLegalAge, " years to legal age.")
	}
}
