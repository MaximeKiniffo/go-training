package main

import (
	"fmt"
	"strconv"
)

func lireScore(score string) (int, error) {
	scoreToString, err := strconv.Atoi(score)

	return scoreToString, err
}

func chargerScore(score string) (int, error) {
	nbToStringLu, err := lireScore(score)
	if err != nil {
		return 0, err
	}

	return nbToStringLu, err
}

func handleScore(score string) {
	scoreTexte, err := chargerScore(score)
	if err != nil {
		fmt.Println("Impossible de charger le score '", score, "'")
		return
	}
	fmt.Println("Score chargé :", scoreTexte)

}

func main() {
	handleScore("abc")
	handleScore("87")

}
