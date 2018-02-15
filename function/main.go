package main

import (
	"fmt"
)

func main() {
	functionA()
	fmt.Println(functionB())
}

func functionA() {
	fmt.Println("Hellow functionA")
}

func functionB() string {
	return "Hellow functionB"
}
