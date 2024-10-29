package main

import "fmt"

type user struct {
	firstname string
	lastname  string
	age       int
}

func main() {
	var firstname, lastname string
	var age int
	fmt.Print("User first name: ")
	fmt.Scan(&firstname)
	fmt.Print("User last name: ")
	fmt.Scan(&lastname)
	fmt.Print("User age: ")
	fmt.Scan(&age)

	var appUser *user = newUser(firstname, lastname, age) //The *user is the same type as that of the return value's
	appUser.displaydata()
}

/////////////////////////////

func newUser(userFirstname string, userLastname string, userAge int) (newUser *user) {
	newUser = &user{
		firstname: userFirstname,
		lastname:  userLastname,
		age:       userAge,
	}
	return
}

// func (u *user) insertdata() {
// 	(*u).firstname = "Peter"
// 	(*u).lastname = "Kane"
// 	(*u).age = 40
// }

func (u user) displaydata() {
	fmt.Println("First name :", u.firstname)
	fmt.Println("Last name :", u.lastname)
	fmt.Println("Age :", u.age)
}
