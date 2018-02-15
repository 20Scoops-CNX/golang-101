package main

import (
	"fmt"
)

func main() {
	var fruits = []string{"Banana", "Melon", "Potato", "Coconut"}

	for index := 0; index < len(fruits); index++ {
		fmt.Printf("index: %d, value: %s\n", index, fruits[index])
	}

	for index, value := range fruits {
		fmt.Printf("index: %d, value: %s\n", index, value)
	}

	var index int
	for {
		if index == len(fruits) {
			break
		}
		fmt.Printf("index: %d, value: %s\n", index, fruits[index])
		index++
	}
}
