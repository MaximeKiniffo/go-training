package main

import "fmt"

type Developer struct {
	Name              string
	YearsOfExperience int
}

func (d *Developer) GainExperience() {
	d.YearsOfExperience++
}

func main() {
	developer := Developer{Name: "Maxime", YearsOfExperience: 2}
	developer.GainExperience()
	fmt.Printf("%s a %d années d'expérience.\n", developer.Name, developer.YearsOfExperience)
	fmt.Println("Le receiver par pointeur modifie l'instance originale de Developer.")
}
