package main

import (
	"fmt"

	"github.com/golang-101/calculator/model"
)

func main() {
	fmt.Print("Please select operator\n\t+ = 1\n\t- = 2\n\t* = 3\n\t/ = 4")
	fmt.Print("\nEnter number operator: ")
	var operator string
	fmt.Scanln(&operator)
	choice, err := model.GetNumberOperator(operator)
	fmt.Print("\nEnter number 1: ")
	var numberOne string
	fmt.Scanln(&numberOne)
	fmt.Print("\nEnter number 2: ")
	var numberTwo string
	fmt.Scanln(&numberTwo)
	n1, n2, err := model.ConvertStringToInt(numberOne, numberTwo)
	if err != nil {
		fmt.Printf("\n" + err.Error())
	}
	result := model.GetResult(choice, n1, n2)
	fmt.Printf("\nResult: %v\n", result)
}
