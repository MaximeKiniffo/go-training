package main

import "fmt"

func main() {
	type Developer struct {
		Name              string
		MainLanguage      string
		YearsOfExperience int
	}

	developer := Developer{
		Name:              "Maxim",
		MainLanguage:      "Go",
		YearsOfExperience: 4,
	}
	fmt.Printf("%v programme principalement en %v et a %v ans d'expérience.\n", developer.Name, developer.MainLanguage, developer.YearsOfExperience)
}
