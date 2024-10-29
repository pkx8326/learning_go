package main

import (
	"fmt"
)

type userData struct {
	firstName string
	lastName  string
	dob       string
}

func main() {
	firstName, lastName, dob := getUserData() //get user data
	uData := createStrucVar(firstName, lastName, dob)
	displayUserData(&uData) //passing address instead of value

}

// ///////////
func createStrucVar(firstName, lastName, dob string) (uData userData) {
	uData = userData{
		firstName: firstName,
		lastName:  lastName,
		dob:       dob,
	} //put all user data in one struct
	return
}

func getUserData() (firstName, lastName, dob string) {
	fmt.Print("Input user's first name :")
	fmt.Scan(&firstName)
	fmt.Print("Input user's last name :")
	fmt.Scan(&lastName)
	fmt.Print("Input user's date of birth (YYYY/MM/DD) :")
	fmt.Scan(&dob)
	return
}

func displayUserData(userData *userData) { //de-referencing variable's address to get value
	fmt.Println("User's first name is ", (*userData).firstName)
	fmt.Println("User's last name is ", (*userData).lastName)
	fmt.Println("User's date of birth is ", (*userData).dob)
}
