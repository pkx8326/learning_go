package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("\"hello world\"")
	fmt.Println("สวัสดี")
	fmt.Println(1 + 1)
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
