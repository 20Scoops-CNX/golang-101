package main

import (
	"fmt"

	"github.com/golang-101/package_01/utils"
)

func main() {
	fmt.Println(utils.FirstName)
	fmt.Println(utils.LastName)
	fmt.Println(utils.CompanyName)
	fmt.Println(utils.Age)

	fmt.Println(utils.Plus(10, 35))
	fmt.Println(utils.Minus(2, 100))
	fmt.Println(utils.Multiply(45, 0))
	fmt.Println(utils.Divide(6, 3))
	fmt.Println(utils.Divide(1, 0))
}
