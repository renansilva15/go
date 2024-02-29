package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	err := checkmail.ValidateFormat("example@example.com")
	fmt.Println(err)
}
