package main

import "fmt"

func main() {
	type DeveloperProfile struct {
		Name              string `json:"name"`
		MainLanguage      string `json:"main_language"`
		YearsOfExperience int    `json:"years_of_experience"`
	}
	developer := DeveloperProfile{
		Name:              "Maxim",
		MainLanguage:      "Go",
		YearsOfExperience: 4,
	}
	fmt.Printf("%v programme principalement en %v et a %v ans d'expérience.\n", developer.Name, developer.MainLanguage, developer.YearsOfExperience)
}
