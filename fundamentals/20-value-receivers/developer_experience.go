package main

import "fmt"

type Developer struct {
	YearsOfExperience int
}

func (developer Developer) GainExperience() {
	// Le receiver par valeur crée une copie de la struct Developer.
	// Cette incrémentation ne modifie donc que cette copie locale.
	// Cet exemple sert à observer ce comportement ; ici, le nom de la
	// méthode suggère volontairement une modification qui n'a pas lieu.
	developer.YearsOfExperience++
}

func main() {
	developer := Developer{YearsOfExperience: 5}
	fmt.Println("avant l'expérience :", developer.YearsOfExperience)
	developer.GainExperience()
	fmt.Println("après l'expérience :", developer.YearsOfExperience)
	fmt.Println("Le receiver par valeur ne modifie pas l'instance originale de Developer.")
}
