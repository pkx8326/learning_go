package main

import "fmt"

type user struct {
	firstName string
	lastName  string
	dob       string
}

func main() {

	userData := getUserData()
	userData.displayUserData()
	userData.clearUserData()

}

/////////////////////////////////////////////////
func getUserData() (userData user) {
	var firstName, lastName, dob string
	fmt.Print("Input the user's firstname :")
	fmt.Scan(&firstName)
	fmt.Print("Input the user's lastname :")
	fmt.Scan(&lastName)
	fmt.Print("Input the user's date of birth (YYYY/MM/DD) :")
	fmt.Scan(&dob)

	userData = user{
		firstName: firstName,
		lastName:  lastName,
		dob:       dob,
	}
	return
}

func (userData user) displayUserData() {
	fmt.Println("The user's firstname is ", userData.firstName)
	fmt.Println("The user's lastname is ", userData.lastName)
	fmt.Println("The user's date of birth is ", userData.dob)
}

func (userData user) clearUserData() {
	userData.firstName = "-"
	fmt.Println("Clearing user data...")
	userData.lastName = "-"
	userData.dob = "-"

	fmt.Println("The user's firstname is ", userData.firstName)
	fmt.Println("The user's lastname is ", userData.lastName)
	fmt.Println("The user's date of birth is ", userData.dob)

}
