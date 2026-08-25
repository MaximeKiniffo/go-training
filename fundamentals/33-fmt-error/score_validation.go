package main

import (
	"fmt"
)

func validateScore(score int) error {
	if score < 0 || score > 100 {
		return fmt.Errorf("Erreur : score invalide : %d, la valeur doit être comprise entre 0 et 100", score)
	}
	return nil
}

func main() {
	score := 99
	err := validateScore(score)
	score2 := -85
	err2 := validateScore(score2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Score valide :", score)
	}
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Score valide :", score2)
	}
}
