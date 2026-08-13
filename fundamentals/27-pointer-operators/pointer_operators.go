package main

import "fmt"

func main() {
	yearsOfExperience := 1
	yearsPointer := &yearsOfExperience
	fmt.Println("Valeur initiale d'expérience: ", yearsOfExperience)
	fmt.Println("Valeur via le pointeur: ", *yearsPointer)
	*yearsPointer = 2
	fmt.Println("Nouvelle valeur d'expérience : ", yearsOfExperience)
}
