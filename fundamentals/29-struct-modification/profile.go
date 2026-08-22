package main

import "fmt"

type Developer struct {
	Name              string
	YearsOfExperience int
}

func promotion(developer *Developer) {
	developer.YearsOfExperience++
}

func main() {
	developer := Developer{
		Name:              "Alice",
		YearsOfExperience: 5,
	}
	fmt.Printf("Avant: %s possède %d années d'expérience \n", developer.Name, developer.YearsOfExperience)
	promotion(&developer)
	fmt.Printf("Après: %s possède %d années d'expérience \n", developer.Name, developer.YearsOfExperience)
}
