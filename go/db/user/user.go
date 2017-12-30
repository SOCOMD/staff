package user

import (
	"fmt"
)

func GetAll() {
	fmt.Println("Called: Get User")
}

func GetSingle(id int) {
	fmt.Printf("Getting User %d \n", id)
}
