package main

import (
	"errors"
	"fmt"
)

func validateAge(age int) error {
	if age < 18 {
		return errors.New("l'âge ne peut pas être inférieur à 18 ans")
	}
	return nil
}

func main() {
	age := 16
	err := validateAge(age)
	age2 := 20
	err2 := validateAge(age2)
	if err != nil {
		fmt.Println("Erreur lors de la validation de l'âge :", err)
	} else {
		fmt.Println("Âge valide :", age)
	}
	if err2 != nil {
		fmt.Println("Erreur lors de la validation de l'âge :", err2)
	} else {
		fmt.Println("Âge valide :", age2)
	}
}
