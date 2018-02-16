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

	// while
	var index int
	for {
		if index == len(fruits) {
			break
		}
		fmt.Printf("index: %d, value: %s\n", index, fruits[index])
		index++
	}

	// do-while
	var position = 0
	for flag := true; flag; flag = (position != 100) {
		fmt.Printf("position: %d\n", (position + 1))
		position++
	}
}
