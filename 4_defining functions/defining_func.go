package main

import "fmt"

func main() {
	outputText()
	plusOperation()

}

func outputText() {
	var text string
	fmt.Print("Please type any word:")
	fmt.Scan(&text)
	fmt.Print(text)
}

func plusOperation() {
	var a, b int
	fmt.Print("Please input an integer for A: ")
	fmt.Scan(&a)

	fmt.Print("Please intput an integer for B: ")
	fmt.Scan(&b)

	fmt.Print("The result of A + B is ", (a + b))
}
