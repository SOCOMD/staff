package user

import (
	"fmt"
)

func GetAll() {
	fmt.Println("Called: Get User")
}

func GetSingle(id int) {
	fmt.Println("Getting User %d", id)
}
