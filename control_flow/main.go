package main

import (
	"fmt"
)

func main() {
	max := getMax(3, 5)
	fmt.Printf("Max value: %v\n", max)
	printStringNumber(2)
}

func getMax(number1 int, number2 int) int {
	if number1 > number2 {
		return number1
	}
	return number2
}

func printStringNumber(value int) {
	switch value {
	case 1:
		fmt.Println("Number one")
	case 2:
		fmt.Println("Number two")
	case 3:
		fmt.Println("Number three")
	case 4:
		fmt.Println("Number four")
	case 5:
		fmt.Println("Number five")
	default:
		fmt.Println("Is Number")
	}
}
