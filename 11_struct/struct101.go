package main

import (
	"fmt"
	"time"
)

// defining a struct template
type userNames struct {
	firstName  string
	middleName string
	lastName   string
	createdAt  time.Time
}

func main() {
	var appUser userNames //declaring a variable of a custom struct type

	userFirstName := "Phisan"
	userMiddleName := "Peter"
	userLastName := "Kane"

	//group the data user data together into one struct variable
	appUser = userNames{
		firstName:  userFirstName, //can omit the field names but the order must be the same as the struct declared
		middleName: userMiddleName,
		lastName:   userLastName,
		createdAt:  time.Now(), //The last field in a struct variable must still be followed by a comma
	}

	fmt.Println(appUser)
}
