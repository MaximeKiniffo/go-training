package main

import (
	"errors"
	"fmt"
)

func calculMoyenneDesTemps(temps []int) (float64, error) {
	if len(temps) == 0 {
		return 0, errors.New("Mesure impossible, aucun temps de réponse")
	}
	somme := 0
	for _, mesure := range temps {
		somme += mesure
	}
	return float64(somme) / float64(len(temps)), nil
}

func chargerMoyenneDesTemps(temps []int) (float64, error) {
	moyenneLu, err := calculMoyenneDesTemps(temps)
	if err != nil {
		return 0, err
	}
	return moyenneLu, err
}

func handleMoyenneDesTemps(temps []int) {
	moyenneTemps, err := chargerMoyenneDesTemps(temps)
	if err != nil {
		fmt.Println("Impossible de produire le rapport :", err)
		return
	}
	fmt.Println("Temps de réponse moyen : ", moyenneTemps, "ms")
}

func main() {
	handleMoyenneDesTemps([]int{12, 150, 45, 654, 98})
	handleMoyenneDesTemps([]int{})
}
