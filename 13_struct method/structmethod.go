package main

import "fmt"

type user struct {
	firstName string
	lastName  string
	dob       string
}

func main() {
	uData := getUserData()
	uData.displayUserData() //The struct's method

}

//////////////////////////////////

func getUserData() (userData user) {
	var firstName, lastName, dob string
	fmt.Print("Enter the user's firstname :")
	fmt.Scan(&firstName)
	fmt.Print("Enter the user's lastname :")
	fmt.Scan(&lastName)
	fmt.Print("Enter the user's date of birth (YYYY/MM/DD) :")
	fmt.Scan(&dob)

	userData = user{
		firstName: firstName,
		lastName:  lastName,
		dob:       dob,
	}
	return
}

//declare that this function is a method of the userData struct (type user)
func (userData user) displayUserData() {
	fmt.Println("The user's first name is ", userData.firstName)
	fmt.Println("The user's last name is ", userData.lastName)
	fmt.Println("The user's date of birth is ", userData.dob)
}
