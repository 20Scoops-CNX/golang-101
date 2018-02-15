package main

import (
	"fmt"
)

// AnimalModel is model for create obj
type AnimalModel struct {
	Name  string
	Age   int
	Sex   bool
	Onwer OnwerDetail
}

// OnwerDetail is model for create obj
type OnwerDetail struct {
	Name        string
	Address     string
	Email       string
	PhoneNumber string
}

func main() {
	var onwerModel OnwerDetail
	onwerModel.Name = "Pat"
	onwerModel.Address = "Chaing Mai"
	onwerModel.PhoneNumber = "0978307474"
	onwerModel.Email = "pat@gmail.com"

	var animalModel = AnimalModel{
		Name:  "Peak",
		Age:   5,
		Sex:   true,
		Onwer: onwerModel,
	}

	fmt.Println(animalModel)
}
