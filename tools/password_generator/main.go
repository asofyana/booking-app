package main

import (
	"booking-app/internal/utils"
	"fmt"
	"os"
)

func main() {
	pwd := os.Args[1]
	fmt.Println(pwd)
	hashedPwd, _ := utils.HashPassword(pwd)
	fmt.Println(hashedPwd)
	verified := utils.VerifyPassword(pwd, hashedPwd)
	fmt.Println(verified)

}
