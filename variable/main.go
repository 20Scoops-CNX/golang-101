package main

import (
	"fmt"
	"os"
)

const firstName = "Jedsada"

const (
	lastName     = "Tiwongvorakul"
	companyName  = "20scoops CNX"
	sex          = "male"
	privateKeyS3 = os.Getenv("S3_PRIVATE_KEY") // compile error because is not constant
)

var (
	httpPort       = os.Getenv("PORT")
	messageError   string
	messageSuccess = ""
	// messageAlret // complie error because syntax error
)

func main() {
	hostName := getHostName()

	hostName1 = getHostName() // complie error because undefined

	var hostName2 = getHostName()

	var hostName3 string = getHostName() // warning because inferred type from value

	fmt.Println(hostName)
	fmt.Println(hostName2)
	fmt.Println(hostName3)
}

func getHostName() string {
	return "20scoops"
}
