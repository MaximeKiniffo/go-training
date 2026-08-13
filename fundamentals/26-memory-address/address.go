package main

import "fmt"

func main() {

	langage := "Go"
	yearsOfExperience := 5
	fmt.Println("Langage: " + langage)
	fmt.Println("Adresse de langage: ", &langage)
	fmt.Println("Années d'expérience: ", yearsOfExperience)
	fmt.Println("Adresse de yearsOfExperience: ", &yearsOfExperience)
}
