package user

import (
	"fmt"
)

//GetAll returns a list of all users from the database
func GetAll() {
	fmt.Printf("Get all users\n")
}

//GetSingle returns a single user from the database
func GetSingle(id int) {
	fmt.Printf("Get user with id:%d \n", id)
}
